// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListDelegatedServicesForAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
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

type ListDelegatedServicesForAccountResponse struct {
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DelegatedServices *ListDelegatedServicesForAccountResponseDelegatedServices `json:"DelegatedServices,omitempty" xml:"DelegatedServices,omitempty" require:"true" type:"Struct"`
}

func (s ListDelegatedServicesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponse) SetRequestId(v string) *ListDelegatedServicesForAccountResponse {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetDelegatedServices(v *ListDelegatedServicesForAccountResponseDelegatedServices) *ListDelegatedServicesForAccountResponse {
	s.DelegatedServices = v
	return s
}

type ListDelegatedServicesForAccountResponseDelegatedServices struct {
	DelegatedService []*ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService `json:"DelegatedService,omitempty" xml:"DelegatedService,omitempty" require:"true" type:"Repeated"`
}

func (s ListDelegatedServicesForAccountResponseDelegatedServices) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseDelegatedServices) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseDelegatedServices) SetDelegatedService(v []*ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService) *ListDelegatedServicesForAccountResponseDelegatedServices {
	s.DelegatedService = v
	return s
}

type ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService struct {
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty" require:"true"`
	ServicePrincipal      *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty" require:"true"`
}

func (s ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService) SetDelegationEnabledTime(v string) *ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService) SetServicePrincipal(v string) *ListDelegatedServicesForAccountResponseDelegatedServicesDelegatedService {
	s.ServicePrincipal = &v
	return s
}

type RegisterDelegatedAdministratorRequest struct {
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty" require:"true"`
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

type RegisterDelegatedAdministratorResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RegisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponse) SetRequestId(v string) *RegisterDelegatedAdministratorResponse {
	s.RequestId = &v
	return s
}

type ListDelegatedAdministratorsRequest struct {
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsRequest) SetServicePrincipal(v string) *ListDelegatedAdministratorsRequest {
	s.ServicePrincipal = &v
	return s
}

type ListDelegatedAdministratorsResponse struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Accounts  *ListDelegatedAdministratorsResponseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" require:"true" type:"Struct"`
}

func (s ListDelegatedAdministratorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponse) SetRequestId(v string) *ListDelegatedAdministratorsResponse {
	s.RequestId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetAccounts(v *ListDelegatedAdministratorsResponseAccounts) *ListDelegatedAdministratorsResponse {
	s.Accounts = v
	return s
}

type ListDelegatedAdministratorsResponseAccounts struct {
	Account []*ListDelegatedAdministratorsResponseAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Repeated"`
}

func (s ListDelegatedAdministratorsResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseAccounts) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseAccounts) SetAccount(v []*ListDelegatedAdministratorsResponseAccountsAccount) *ListDelegatedAdministratorsResponseAccounts {
	s.Account = v
	return s
}

type ListDelegatedAdministratorsResponseAccountsAccount struct {
	DelegationEnabledTime *string `json:"DelegationEnabledTime,omitempty" xml:"DelegationEnabledTime,omitempty" require:"true"`
	JoinMethod            *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	AccountId             *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ServicePrincipal      *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty" require:"true"`
	DisplayName           *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
}

func (s ListDelegatedAdministratorsResponseAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListDelegatedAdministratorsResponseAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponseAccountsAccount) SetDelegationEnabledTime(v string) *ListDelegatedAdministratorsResponseAccountsAccount {
	s.DelegationEnabledTime = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseAccountsAccount) SetJoinMethod(v string) *ListDelegatedAdministratorsResponseAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseAccountsAccount) SetAccountId(v string) *ListDelegatedAdministratorsResponseAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseAccountsAccount) SetServicePrincipal(v string) *ListDelegatedAdministratorsResponseAccountsAccount {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedAdministratorsResponseAccountsAccount) SetDisplayName(v string) *ListDelegatedAdministratorsResponseAccountsAccount {
	s.DisplayName = &v
	return s
}

type DeregisterDelegatedAdministratorRequest struct {
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty" require:"true"`
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

type DeregisterDelegatedAdministratorResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeregisterDelegatedAdministratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponse) SetRequestId(v string) *DeregisterDelegatedAdministratorResponse {
	s.RequestId = &v
	return s
}

type ListControlPoliciesRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListControlPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesRequest) SetPolicyType(v string) *ListControlPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageNumber(v int) *ListControlPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesRequest) SetPageSize(v int) *ListControlPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesRequest) SetLanguage(v string) *ListControlPoliciesRequest {
	s.Language = &v
	return s
}

type ListControlPoliciesResponse struct {
	TotalCount      *int                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize        *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber      *int                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	ControlPolicies *ListControlPoliciesResponseControlPolicies `json:"ControlPolicies,omitempty" xml:"ControlPolicies,omitempty" require:"true" type:"Struct"`
}

func (s ListControlPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponse) SetTotalCount(v int) *ListControlPoliciesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListControlPoliciesResponse) SetPageSize(v int) *ListControlPoliciesResponse {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesResponse) SetRequestId(v string) *ListControlPoliciesResponse {
	s.RequestId = &v
	return s
}

func (s *ListControlPoliciesResponse) SetPageNumber(v int) *ListControlPoliciesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesResponse) SetControlPolicies(v *ListControlPoliciesResponseControlPolicies) *ListControlPoliciesResponse {
	s.ControlPolicies = v
	return s
}

type ListControlPoliciesResponseControlPolicies struct {
	ControlPolicy []*ListControlPoliciesResponseControlPoliciesControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" require:"true" type:"Repeated"`
}

func (s ListControlPoliciesResponseControlPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseControlPolicies) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseControlPolicies) SetControlPolicy(v []*ListControlPoliciesResponseControlPoliciesControlPolicy) *ListControlPoliciesResponseControlPolicies {
	s.ControlPolicy = v
	return s
}

type ListControlPoliciesResponseControlPoliciesControlPolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
}

func (s ListControlPoliciesResponseControlPoliciesControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListControlPoliciesResponseControlPoliciesControlPolicy) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetPolicyType(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetUpdateDate(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetDescription(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.Description = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetAttachmentCount(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetPolicyName(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetCreateDate(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetPolicyId(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *ListControlPoliciesResponseControlPoliciesControlPolicy) SetEffectScope(v string) *ListControlPoliciesResponseControlPoliciesControlPolicy {
	s.EffectScope = &v
	return s
}

type ListControlPolicyAttachmentsForTargetRequest struct {
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty" require:"true"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetTargetId(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetLanguage(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.Language = &v
	return s
}

type ListControlPolicyAttachmentsForTargetResponse struct {
	RequestId                *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ControlPolicyAttachments *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments `json:"ControlPolicyAttachments,omitempty" xml:"ControlPolicyAttachments,omitempty" require:"true" type:"Struct"`
}

func (s ListControlPolicyAttachmentsForTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetRequestId(v string) *ListControlPolicyAttachmentsForTargetResponse {
	s.RequestId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetControlPolicyAttachments(v *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments) *ListControlPolicyAttachmentsForTargetResponse {
	s.ControlPolicyAttachments = v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments struct {
	ControlPolicyAttachment []*ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment `json:"ControlPolicyAttachment,omitempty" xml:"ControlPolicyAttachment,omitempty" require:"true" type:"Repeated"`
}

func (s ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments) SetControlPolicyAttachment(v []*ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachments {
	s.ControlPolicyAttachment = v
	return s
}

type ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment struct {
	PolicyType  *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachDate  *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty" require:"true"`
	PolicyName  *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PolicyId    *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
}

func (s ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetPolicyType(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyType = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetDescription(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetAttachDate(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetPolicyName(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetPolicyId(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment) SetEffectScope(v string) *ListControlPolicyAttachmentsForTargetResponseControlPolicyAttachmentsControlPolicyAttachment {
	s.EffectScope = &v
	return s
}

type CreateControlPolicyRequest struct {
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectScope    *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
}

func (s CreateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyRequest) SetPolicyName(v string) *CreateControlPolicyRequest {
	s.PolicyName = &v
	return s
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

type CreateControlPolicyResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ControlPolicy *CreateControlPolicyResponseControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" require:"true" type:"Struct"`
}

func (s CreateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponse) SetRequestId(v string) *CreateControlPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *CreateControlPolicyResponse) SetControlPolicy(v *CreateControlPolicyResponseControlPolicy) *CreateControlPolicyResponse {
	s.ControlPolicy = v
	return s
}

type CreateControlPolicyResponseControlPolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
}

func (s CreateControlPolicyResponseControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateControlPolicyResponseControlPolicy) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseControlPolicy) SetPolicyType(v string) *CreateControlPolicyResponseControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetUpdateDate(v string) *CreateControlPolicyResponseControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetDescription(v string) *CreateControlPolicyResponseControlPolicy {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetAttachmentCount(v string) *CreateControlPolicyResponseControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetPolicyName(v string) *CreateControlPolicyResponseControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetCreateDate(v string) *CreateControlPolicyResponseControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetPolicyId(v string) *CreateControlPolicyResponseControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateControlPolicyResponseControlPolicy) SetEffectScope(v string) *CreateControlPolicyResponseControlPolicy {
	s.EffectScope = &v
	return s
}

type ListTargetAttachmentsForControlPolicyRequest struct {
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPolicyId(v string) *ListTargetAttachmentsForControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageNumber(v int) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyRequest) SetPageSize(v int) *ListTargetAttachmentsForControlPolicyRequest {
	s.PageSize = &v
	return s
}

type ListTargetAttachmentsForControlPolicyResponse struct {
	TotalCount        *int                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize          *int                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId         *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber        *int                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TargetAttachments *ListTargetAttachmentsForControlPolicyResponseTargetAttachments `json:"TargetAttachments,omitempty" xml:"TargetAttachments,omitempty" require:"true" type:"Struct"`
}

func (s ListTargetAttachmentsForControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetTotalCount(v int) *ListTargetAttachmentsForControlPolicyResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetPageSize(v int) *ListTargetAttachmentsForControlPolicyResponse {
	s.PageSize = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetRequestId(v string) *ListTargetAttachmentsForControlPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetPageNumber(v int) *ListTargetAttachmentsForControlPolicyResponse {
	s.PageNumber = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetTargetAttachments(v *ListTargetAttachmentsForControlPolicyResponseTargetAttachments) *ListTargetAttachmentsForControlPolicyResponse {
	s.TargetAttachments = v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseTargetAttachments struct {
	TargetAttachment []*ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment `json:"TargetAttachment,omitempty" xml:"TargetAttachment,omitempty" require:"true" type:"Repeated"`
}

func (s ListTargetAttachmentsForControlPolicyResponseTargetAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseTargetAttachments) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseTargetAttachments) SetTargetAttachment(v []*ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) *ListTargetAttachmentsForControlPolicyResponseTargetAttachments {
	s.TargetAttachment = v
	return s
}

type ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment struct {
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty" require:"true"`
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty" require:"true"`
	TargetId   *string `json:"TargetId,omitempty" xml:"TargetId,omitempty" require:"true"`
}

func (s ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) SetAttachDate(v string) *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) SetTargetType(v string) *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment {
	s.TargetType = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) SetTargetName(v string) *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment {
	s.TargetName = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment) SetTargetId(v string) *ListTargetAttachmentsForControlPolicyResponseTargetAttachmentsTargetAttachment {
	s.TargetId = &v
	return s
}

type GetControlPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetControlPolicyRequest) SetPolicyId(v string) *GetControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyRequest) SetLanguage(v string) *GetControlPolicyRequest {
	s.Language = &v
	return s
}

type GetControlPolicyResponse struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ControlPolicy *GetControlPolicyResponseControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" require:"true" type:"Struct"`
}

func (s GetControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponse) SetRequestId(v string) *GetControlPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *GetControlPolicyResponse) SetControlPolicy(v *GetControlPolicyResponseControlPolicy) *GetControlPolicyResponse {
	s.ControlPolicy = v
	return s
}

type GetControlPolicyResponseControlPolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PolicyDocument  *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
}

func (s GetControlPolicyResponseControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyResponseControlPolicy) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseControlPolicy) SetPolicyType(v string) *GetControlPolicyResponseControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetUpdateDate(v string) *GetControlPolicyResponseControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetDescription(v string) *GetControlPolicyResponseControlPolicy {
	s.Description = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetAttachmentCount(v string) *GetControlPolicyResponseControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetPolicyName(v string) *GetControlPolicyResponseControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetPolicyDocument(v string) *GetControlPolicyResponseControlPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetCreateDate(v string) *GetControlPolicyResponseControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetPolicyId(v string) *GetControlPolicyResponseControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyResponseControlPolicy) SetEffectScope(v string) *GetControlPolicyResponseControlPolicy {
	s.EffectScope = &v
	return s
}

type EnableControlPolicyRequest struct {
}

func (s EnableControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyRequest) GoString() string {
	return s.String()
}

type EnableControlPolicyResponse struct {
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty" require:"true"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *EnableControlPolicyResponse) SetEnablementStatus(v string) *EnableControlPolicyResponse {
	s.EnablementStatus = &v
	return s
}

func (s *EnableControlPolicyResponse) SetRequestId(v string) *EnableControlPolicyResponse {
	s.RequestId = &v
	return s
}

type GetControlPolicyEnablementStatusRequest struct {
}

func (s GetControlPolicyEnablementStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusRequest) GoString() string {
	return s.String()
}

type GetControlPolicyEnablementStatusResponse struct {
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty" require:"true"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetControlPolicyEnablementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponse) SetEnablementStatus(v string) *GetControlPolicyEnablementStatusResponse {
	s.EnablementStatus = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetRequestId(v string) *GetControlPolicyEnablementStatusResponse {
	s.RequestId = &v
	return s
}

type DisableControlPolicyRequest struct {
}

func (s DisableControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyRequest) GoString() string {
	return s.String()
}

type DisableControlPolicyResponse struct {
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty" require:"true"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponse) SetEnablementStatus(v string) *DisableControlPolicyResponse {
	s.EnablementStatus = &v
	return s
}

func (s *DisableControlPolicyResponse) SetRequestId(v string) *DisableControlPolicyResponse {
	s.RequestId = &v
	return s
}

type DetachControlPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty" require:"true"`
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

type DetachControlPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DetachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponse) SetRequestId(v string) *DetachControlPolicyResponse {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
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

type DeleteControlPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) SetRequestId(v string) *DeleteControlPolicyResponse {
	s.RequestId = &v
	return s
}

type UpdateControlPolicyRequest struct {
	PolicyId          *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	NewPolicyName     *string `json:"NewPolicyName,omitempty" xml:"NewPolicyName,omitempty"`
	NewDescription    *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewPolicyDocument *string `json:"NewPolicyDocument,omitempty" xml:"NewPolicyDocument,omitempty"`
}

func (s UpdateControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyRequest) SetPolicyId(v string) *UpdateControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyName(v string) *UpdateControlPolicyRequest {
	s.NewPolicyName = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewDescription(v string) *UpdateControlPolicyRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateControlPolicyRequest) SetNewPolicyDocument(v string) *UpdateControlPolicyRequest {
	s.NewPolicyDocument = &v
	return s
}

type UpdateControlPolicyResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ControlPolicy *UpdateControlPolicyResponseControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" require:"true" type:"Struct"`
}

func (s UpdateControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponse) SetRequestId(v string) *UpdateControlPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateControlPolicyResponse) SetControlPolicy(v *UpdateControlPolicyResponseControlPolicy) *UpdateControlPolicyResponse {
	s.ControlPolicy = v
	return s
}

type UpdateControlPolicyResponseControlPolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	EffectScope     *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty" require:"true"`
}

func (s UpdateControlPolicyResponseControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateControlPolicyResponseControlPolicy) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseControlPolicy) SetPolicyType(v string) *UpdateControlPolicyResponseControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetUpdateDate(v string) *UpdateControlPolicyResponseControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetDescription(v string) *UpdateControlPolicyResponseControlPolicy {
	s.Description = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetAttachmentCount(v string) *UpdateControlPolicyResponseControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetPolicyName(v string) *UpdateControlPolicyResponseControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetCreateDate(v string) *UpdateControlPolicyResponseControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetPolicyId(v string) *UpdateControlPolicyResponseControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyResponseControlPolicy) SetEffectScope(v string) *UpdateControlPolicyResponseControlPolicy {
	s.EffectScope = &v
	return s
}

type AttachControlPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" require:"true"`
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty" require:"true"`
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

type AttachControlPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AttachControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyResponse) SetRequestId(v string) *AttachControlPolicyResponse {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
	CustomSuffix *string `json:"CustomSuffix,omitempty" xml:"CustomSuffix,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetServiceName(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetCustomSuffix(v string) *CreateServiceLinkedRoleRequest {
	s.CustomSuffix = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetDescription(v string) *CreateServiceLinkedRoleRequest {
	s.Description = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Role      *CreateServiceLinkedRoleResponseRole `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Struct"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetRequestId(v string) *CreateServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetRole(v *CreateServiceLinkedRoleResponseRole) *CreateServiceLinkedRoleResponse {
	s.Role = v
	return s
}

type CreateServiceLinkedRoleResponseRole struct {
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty" require:"true"`
	IsServiceLinkedRole      *bool   `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty" require:"true"`
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponseRole) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseRole) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseRole) SetRoleName(v string) *CreateServiceLinkedRoleResponseRole {
	s.RoleName = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetDescription(v string) *CreateServiceLinkedRoleResponseRole {
	s.Description = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetAssumeRolePolicyDocument(v string) *CreateServiceLinkedRoleResponseRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetIsServiceLinkedRole(v bool) *CreateServiceLinkedRoleResponseRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetArn(v string) *CreateServiceLinkedRoleResponseRole {
	s.Arn = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetRoleId(v string) *CreateServiceLinkedRoleResponseRole {
	s.RoleId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetCreateDate(v string) *CreateServiceLinkedRoleResponseRole {
	s.CreateDate = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseRole) SetRolePrincipalName(v string) *CreateServiceLinkedRoleResponseRole {
	s.RolePrincipalName = &v
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

type GetServiceLinkedRoleDeletionStatusResponse struct {
	Status    *string                                           `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Reason    *GetServiceLinkedRoleDeletionStatusResponseReason `json:"Reason,omitempty" xml:"Reason,omitempty" require:"true" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetStatus(v string) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Status = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetRequestId(v string) *GetServiceLinkedRoleDeletionStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetReason(v *GetServiceLinkedRoleDeletionStatusResponseReason) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Reason = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseReason struct {
	Message    *string                                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RoleUsages *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" require:"true" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseReason) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseReason) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReason) SetMessage(v string) *GetServiceLinkedRoleDeletionStatusResponseReason {
	s.Message = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReason) SetRoleUsages(v *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages) *GetServiceLinkedRoleDeletionStatusResponseReason {
	s.RoleUsages = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages struct {
	RoleUsage []*GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage `json:"RoleUsage,omitempty" xml:"RoleUsage,omitempty" require:"true" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages) SetRoleUsage(v []*GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage) *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsages {
	s.RoleUsage = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage struct {
	Region    *string                                                                       `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	Resources *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources `json:"Resources,omitempty" xml:"Resources,omitempty" require:"true" type:"Struct"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage) SetRegion(v string) *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage {
	s.Region = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage) SetResources(v *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources) *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsage {
	s.Resources = v
	return s
}

type GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources struct {
	// Resource
	Resource []*string `json:"Resource,omitempty" xml:"Resource,omitempty" require:"true" type:"Repeated"`
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources) SetResource(v []*string) *GetServiceLinkedRoleDeletionStatusResponseReasonRoleUsagesRoleUsageResources {
	s.Resource = v
	return s
}

type ListTrustedServiceStatusRequest struct {
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AdminAccountId *string `json:"AdminAccountId,omitempty" xml:"AdminAccountId,omitempty"`
}

func (s ListTrustedServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusRequest) SetPageNumber(v int) *ListTrustedServiceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetPageSize(v int) *ListTrustedServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusRequest) SetAdminAccountId(v string) *ListTrustedServiceStatusRequest {
	s.AdminAccountId = &v
	return s
}

type ListTrustedServiceStatusResponse struct {
	TotalCount               *int                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	RequestId                *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageSize                 *int                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber               *int                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	EnabledServicePrincipals *ListTrustedServiceStatusResponseEnabledServicePrincipals `json:"EnabledServicePrincipals,omitempty" xml:"EnabledServicePrincipals,omitempty" require:"true" type:"Struct"`
}

func (s ListTrustedServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponse) SetTotalCount(v int) *ListTrustedServiceStatusResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetRequestId(v string) *ListTrustedServiceStatusResponse {
	s.RequestId = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetPageSize(v int) *ListTrustedServiceStatusResponse {
	s.PageSize = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetPageNumber(v int) *ListTrustedServiceStatusResponse {
	s.PageNumber = &v
	return s
}

func (s *ListTrustedServiceStatusResponse) SetEnabledServicePrincipals(v *ListTrustedServiceStatusResponseEnabledServicePrincipals) *ListTrustedServiceStatusResponse {
	s.EnabledServicePrincipals = v
	return s
}

type ListTrustedServiceStatusResponseEnabledServicePrincipals struct {
	EnabledServicePrincipal []*ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal `json:"EnabledServicePrincipal,omitempty" xml:"EnabledServicePrincipal,omitempty" require:"true" type:"Repeated"`
}

func (s ListTrustedServiceStatusResponseEnabledServicePrincipals) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseEnabledServicePrincipals) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseEnabledServicePrincipals) SetEnabledServicePrincipal(v []*ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal) *ListTrustedServiceStatusResponseEnabledServicePrincipals {
	s.EnabledServicePrincipal = v
	return s
}

type ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal struct {
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty" require:"true"`
	EnableTime       *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty" require:"true"`
}

func (s ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal) String() string {
	return tea.Prettify(s)
}

func (s ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal) GoString() string {
	return s.String()
}

func (s *ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal) SetServicePrincipal(v string) *ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal {
	s.ServicePrincipal = &v
	return s
}

func (s *ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal) SetEnableTime(v string) *ListTrustedServiceStatusResponseEnabledServicePrincipalsEnabledServicePrincipal {
	s.EnableTime = &v
	return s
}

type DeleteServiceLinkedRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
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

type DeleteServiceLinkedRoleResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty" require:"true"`
}

func (s DeleteServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleResponse) SetRequestId(v string) *DeleteServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) SetDeletionTaskId(v string) *DeleteServiceLinkedRoleResponse {
	s.DeletionTaskId = &v
	return s
}

type UpdateRoleRequest struct {
	RoleName                    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	NewAssumeRolePolicyDocument *string `json:"NewAssumeRolePolicyDocument,omitempty" xml:"NewAssumeRolePolicyDocument,omitempty"`
	NewMaxSessionDuration       *int64  `json:"NewMaxSessionDuration,omitempty" xml:"NewMaxSessionDuration,omitempty"`
	NewDescription              *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleRequest) SetNewAssumeRolePolicyDocument(v string) *UpdateRoleRequest {
	s.NewAssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleRequest) SetNewMaxSessionDuration(v int64) *UpdateRoleRequest {
	s.NewMaxSessionDuration = &v
	return s
}

func (s *UpdateRoleRequest) SetNewDescription(v string) *UpdateRoleRequest {
	s.NewDescription = &v
	return s
}

type UpdateRoleResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Role      *UpdateRoleResponseRole `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Struct"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetRequestId(v string) *UpdateRoleResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponse) SetRole(v *UpdateRoleResponseRole) *UpdateRoleResponse {
	s.Role = v
	return s
}

type UpdateRoleResponseRole struct {
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty" require:"true"`
	UpdateDate               *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty" require:"true"`
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty" require:"true"`
}

func (s UpdateRoleResponseRole) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseRole) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseRole) SetMaxSessionDuration(v int64) *UpdateRoleResponseRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *UpdateRoleResponseRole) SetUpdateDate(v string) *UpdateRoleResponseRole {
	s.UpdateDate = &v
	return s
}

func (s *UpdateRoleResponseRole) SetRoleName(v string) *UpdateRoleResponseRole {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleResponseRole) SetDescription(v string) *UpdateRoleResponseRole {
	s.Description = &v
	return s
}

func (s *UpdateRoleResponseRole) SetAssumeRolePolicyDocument(v string) *UpdateRoleResponseRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleResponseRole) SetArn(v string) *UpdateRoleResponseRole {
	s.Arn = &v
	return s
}

func (s *UpdateRoleResponseRole) SetRoleId(v string) *UpdateRoleResponseRole {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleResponseRole) SetCreateDate(v string) *UpdateRoleResponseRole {
	s.CreateDate = &v
	return s
}

func (s *UpdateRoleResponseRole) SetRolePrincipalName(v string) *UpdateRoleResponseRole {
	s.RolePrincipalName = &v
	return s
}

type ListResourcesRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceIds     *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetResourceGroupId(v string) *ListResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesRequest) SetService(v string) *ListResourcesRequest {
	s.Service = &v
	return s
}

func (s *ListResourcesRequest) SetRegion(v string) *ListResourcesRequest {
	s.Region = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetResourceIds(v string) *ListResourcesRequest {
	s.ResourceIds = &v
	return s
}

type ListResourcesResponse struct {
	TotalCount *int                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Resources  *ListResourcesResponseResources `json:"Resources,omitempty" xml:"Resources,omitempty" require:"true" type:"Struct"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetTotalCount(v int) *ListResourcesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesResponse) SetPageSize(v int) *ListResourcesResponse {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponse) SetRequestId(v string) *ListResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponse) SetPageNumber(v int) *ListResourcesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponse) SetResources(v *ListResourcesResponseResources) *ListResourcesResponse {
	s.Resources = v
	return s
}

type ListResourcesResponseResources struct {
	Resource []*ListResourcesResponseResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourcesResponseResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseResources) SetResource(v []*ListResourcesResponseResourcesResource) *ListResourcesResponseResources {
	s.Resource = v
	return s
}

type ListResourcesResponseResourcesResource struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	Service         *string `json:"Service,omitempty" xml:"Service,omitempty" require:"true"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListResourcesResponseResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseResourcesResource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseResourcesResource) SetResourceGroupId(v string) *ListResourcesResponseResourcesResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesResponseResourcesResource) SetResourceId(v string) *ListResourcesResponseResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseResourcesResource) SetService(v string) *ListResourcesResponseResourcesResource {
	s.Service = &v
	return s
}

func (s *ListResourcesResponseResourcesResource) SetResourceType(v string) *ListResourcesResponseResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseResourcesResource) SetRegionId(v string) *ListResourcesResponseResourcesResource {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseResourcesResource) SetCreateDate(v string) *ListResourcesResponseResourcesResource {
	s.CreateDate = &v
	return s
}

type CreateCloudAccountRequest struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
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

func (s *CreateCloudAccountRequest) SetParentFolderId(v string) *CreateCloudAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateCloudAccountRequest) SetEmail(v string) *CreateCloudAccountRequest {
	s.Email = &v
	return s
}

func (s *CreateCloudAccountRequest) SetPayerAccountId(v string) *CreateCloudAccountRequest {
	s.PayerAccountId = &v
	return s
}

type CreateCloudAccountResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *CreateCloudAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s CreateCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponse) SetRequestId(v string) *CreateCloudAccountResponse {
	s.RequestId = &v
	return s
}

func (s *CreateCloudAccountResponse) SetAccount(v *CreateCloudAccountResponseAccount) *CreateCloudAccountResponse {
	s.Account = v
	return s
}

type CreateCloudAccountResponseAccount struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
}

func (s CreateCloudAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountResponseAccount) SetResourceDirectoryId(v string) *CreateCloudAccountResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetAccountId(v string) *CreateCloudAccountResponseAccount {
	s.AccountId = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetDisplayName(v string) *CreateCloudAccountResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetAccountName(v string) *CreateCloudAccountResponseAccount {
	s.AccountName = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetFolderId(v string) *CreateCloudAccountResponseAccount {
	s.FolderId = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetJoinMethod(v string) *CreateCloudAccountResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetModifyTime(v string) *CreateCloudAccountResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetType(v string) *CreateCloudAccountResponseAccount {
	s.Type = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetStatus(v string) *CreateCloudAccountResponseAccount {
	s.Status = &v
	return s
}

func (s *CreateCloudAccountResponseAccount) SetRecordId(v string) *CreateCloudAccountResponseAccount {
	s.RecordId = &v
	return s
}

type DeleteRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
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

type DeleteRoleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponse) SetRequestId(v string) *DeleteRoleResponse {
	s.RequestId = &v
	return s
}

type GetRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s *GetRoleRequest) SetLanguage(v string) *GetRoleRequest {
	s.Language = &v
	return s
}

type GetRoleResponse struct {
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Role      *GetRoleResponseRole `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Struct"`
}

func (s GetRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponse) GoString() string {
	return s.String()
}

func (s *GetRoleResponse) SetRequestId(v string) *GetRoleResponse {
	s.RequestId = &v
	return s
}

func (s *GetRoleResponse) SetRole(v *GetRoleResponseRole) *GetRoleResponse {
	s.Role = v
	return s
}

type GetRoleResponseRole struct {
	MaxSessionDuration       *int64                                 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty" require:"true"`
	UpdateDate               *string                                `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	RoleName                 *string                                `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description              *string                                `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AssumeRolePolicyDocument *string                                `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty" require:"true"`
	IsServiceLinkedRole      *bool                                  `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty" require:"true"`
	Arn                      *string                                `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	RoleId                   *string                                `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	CreateDate               *string                                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	RolePrincipalName        *string                                `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty" require:"true"`
	LatestDeletionTask       *GetRoleResponseRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" require:"true" type:"Struct"`
}

func (s GetRoleResponseRole) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseRole) GoString() string {
	return s.String()
}

func (s *GetRoleResponseRole) SetMaxSessionDuration(v int64) *GetRoleResponseRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *GetRoleResponseRole) SetUpdateDate(v string) *GetRoleResponseRole {
	s.UpdateDate = &v
	return s
}

func (s *GetRoleResponseRole) SetRoleName(v string) *GetRoleResponseRole {
	s.RoleName = &v
	return s
}

func (s *GetRoleResponseRole) SetDescription(v string) *GetRoleResponseRole {
	s.Description = &v
	return s
}

func (s *GetRoleResponseRole) SetAssumeRolePolicyDocument(v string) *GetRoleResponseRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *GetRoleResponseRole) SetIsServiceLinkedRole(v bool) *GetRoleResponseRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *GetRoleResponseRole) SetArn(v string) *GetRoleResponseRole {
	s.Arn = &v
	return s
}

func (s *GetRoleResponseRole) SetRoleId(v string) *GetRoleResponseRole {
	s.RoleId = &v
	return s
}

func (s *GetRoleResponseRole) SetCreateDate(v string) *GetRoleResponseRole {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseRole) SetRolePrincipalName(v string) *GetRoleResponseRole {
	s.RolePrincipalName = &v
	return s
}

func (s *GetRoleResponseRole) SetLatestDeletionTask(v *GetRoleResponseRoleLatestDeletionTask) *GetRoleResponseRole {
	s.LatestDeletionTask = v
	return s
}

type GetRoleResponseRoleLatestDeletionTask struct {
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty" require:"true"`
}

func (s GetRoleResponseRoleLatestDeletionTask) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *GetRoleResponseRoleLatestDeletionTask) SetCreateDate(v string) *GetRoleResponseRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *GetRoleResponseRoleLatestDeletionTask) SetDeletionTaskId(v string) *GetRoleResponseRoleLatestDeletionTask {
	s.DeletionTaskId = &v
	return s
}

type ListRolesRequest struct {
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetPageNumber(v int) *ListRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRolesRequest) SetPageSize(v int) *ListRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRolesRequest) SetLanguage(v string) *ListRolesRequest {
	s.Language = &v
	return s
}

type ListRolesResponse struct {
	TotalCount *int                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Roles      *ListRolesResponseRoles `json:"Roles,omitempty" xml:"Roles,omitempty" require:"true" type:"Struct"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetTotalCount(v int) *ListRolesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListRolesResponse) SetPageSize(v int) *ListRolesResponse {
	s.PageSize = &v
	return s
}

func (s *ListRolesResponse) SetRequestId(v string) *ListRolesResponse {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponse) SetPageNumber(v int) *ListRolesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListRolesResponse) SetRoles(v *ListRolesResponseRoles) *ListRolesResponse {
	s.Roles = v
	return s
}

type ListRolesResponseRoles struct {
	Role []*ListRolesResponseRolesRole `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Repeated"`
}

func (s ListRolesResponseRoles) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseRoles) SetRole(v []*ListRolesResponseRolesRole) *ListRolesResponseRoles {
	s.Role = v
	return s
}

type ListRolesResponseRolesRole struct {
	MaxSessionDuration  *int64                                        `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty" require:"true"`
	UpdateDate          *string                                       `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	RoleName            *string                                       `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description         *string                                       `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	IsServiceLinkedRole *bool                                         `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty" require:"true"`
	Arn                 *string                                       `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	RoleId              *string                                       `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	CreateDate          *string                                       `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	RolePrincipalName   *string                                       `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty" require:"true"`
	LatestDeletionTask  *ListRolesResponseRolesRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" require:"true" type:"Struct"`
}

func (s ListRolesResponseRolesRole) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseRolesRole) GoString() string {
	return s.String()
}

func (s *ListRolesResponseRolesRole) SetMaxSessionDuration(v int64) *ListRolesResponseRolesRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetUpdateDate(v string) *ListRolesResponseRolesRole {
	s.UpdateDate = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetRoleName(v string) *ListRolesResponseRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetDescription(v string) *ListRolesResponseRolesRole {
	s.Description = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetIsServiceLinkedRole(v bool) *ListRolesResponseRolesRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetArn(v string) *ListRolesResponseRolesRole {
	s.Arn = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetRoleId(v string) *ListRolesResponseRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetCreateDate(v string) *ListRolesResponseRolesRole {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetRolePrincipalName(v string) *ListRolesResponseRolesRole {
	s.RolePrincipalName = &v
	return s
}

func (s *ListRolesResponseRolesRole) SetLatestDeletionTask(v *ListRolesResponseRolesRoleLatestDeletionTask) *ListRolesResponseRolesRole {
	s.LatestDeletionTask = v
	return s
}

type ListRolesResponseRolesRoleLatestDeletionTask struct {
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty" require:"true"`
}

func (s ListRolesResponseRolesRoleLatestDeletionTask) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseRolesRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *ListRolesResponseRolesRoleLatestDeletionTask) SetCreateDate(v string) *ListRolesResponseRolesRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseRolesRoleLatestDeletionTask) SetDeletionTaskId(v string) *ListRolesResponseRolesRoleLatestDeletionTask {
	s.DeletionTaskId = &v
	return s
}

type CreateRoleRequest struct {
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty" require:"true"`
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetAssumeRolePolicyDocument(v string) *CreateRoleRequest {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleRequest) SetMaxSessionDuration(v int64) *CreateRoleRequest {
	s.MaxSessionDuration = &v
	return s
}

type CreateRoleResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Role      *CreateRoleResponseRole `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Struct"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetRequestId(v string) *CreateRoleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponse) SetRole(v *CreateRoleResponseRole) *CreateRoleResponse {
	s.Role = v
	return s
}

type CreateRoleResponseRole struct {
	MaxSessionDuration       *int64  `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty" require:"true"`
	RoleName                 *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty" require:"true"`
	Arn                      *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	RoleId                   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	RolePrincipalName        *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty" require:"true"`
}

func (s CreateRoleResponseRole) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseRole) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseRole) SetMaxSessionDuration(v int64) *CreateRoleResponseRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleResponseRole) SetRoleName(v string) *CreateRoleResponseRole {
	s.RoleName = &v
	return s
}

func (s *CreateRoleResponseRole) SetDescription(v string) *CreateRoleResponseRole {
	s.Description = &v
	return s
}

func (s *CreateRoleResponseRole) SetAssumeRolePolicyDocument(v string) *CreateRoleResponseRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleResponseRole) SetArn(v string) *CreateRoleResponseRole {
	s.Arn = &v
	return s
}

func (s *CreateRoleResponseRole) SetRoleId(v string) *CreateRoleResponseRole {
	s.RoleId = &v
	return s
}

func (s *CreateRoleResponseRole) SetCreateDate(v string) *CreateRoleResponseRole {
	s.CreateDate = &v
	return s
}

func (s *CreateRoleResponseRole) SetRolePrincipalName(v string) *CreateRoleResponseRole {
	s.RolePrincipalName = &v
	return s
}

type ListPolicyAttachmentsRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListPolicyAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsRequest) SetResourceGroupId(v string) *ListPolicyAttachmentsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPolicyType(v string) *ListPolicyAttachmentsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPolicyName(v string) *ListPolicyAttachmentsRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPrincipalType(v string) *ListPolicyAttachmentsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPrincipalName(v string) *ListPolicyAttachmentsRequest {
	s.PrincipalName = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPageNumber(v int) *ListPolicyAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetPageSize(v int) *ListPolicyAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentsRequest) SetLanguage(v string) *ListPolicyAttachmentsRequest {
	s.Language = &v
	return s
}

type ListPolicyAttachmentsResponse struct {
	TotalCount        *int                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize          *int                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber        *int                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PolicyAttachments *ListPolicyAttachmentsResponsePolicyAttachments `json:"PolicyAttachments,omitempty" xml:"PolicyAttachments,omitempty" require:"true" type:"Struct"`
}

func (s ListPolicyAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponse) SetTotalCount(v int) *ListPolicyAttachmentsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetPageSize(v int) *ListPolicyAttachmentsResponse {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetRequestId(v string) *ListPolicyAttachmentsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetPageNumber(v int) *ListPolicyAttachmentsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetPolicyAttachments(v *ListPolicyAttachmentsResponsePolicyAttachments) *ListPolicyAttachmentsResponse {
	s.PolicyAttachments = v
	return s
}

type ListPolicyAttachmentsResponsePolicyAttachments struct {
	PolicyAttachment []*ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment `json:"PolicyAttachment,omitempty" xml:"PolicyAttachment,omitempty" require:"true" type:"Repeated"`
}

func (s ListPolicyAttachmentsResponsePolicyAttachments) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponsePolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponsePolicyAttachments) SetPolicyAttachment(v []*ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) *ListPolicyAttachmentsResponsePolicyAttachments {
	s.PolicyAttachment = v
	return s
}

type ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	AttachDate      *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty" require:"true"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty" require:"true"`
}

func (s ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetPolicyType(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetDescription(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetResourceGroupId(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetAttachDate(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetPolicyName(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetPrincipalName(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.PrincipalName = &v
	return s
}

func (s *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment) SetPrincipalType(v string) *ListPolicyAttachmentsResponsePolicyAttachmentsPolicyAttachment {
	s.PrincipalType = &v
	return s
}

type DetachPolicyRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty" require:"true"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty" require:"true"`
}

func (s DetachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyRequest) SetResourceGroupId(v string) *DetachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyType(v string) *DetachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyName(v string) *DetachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalType(v string) *DetachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalName(v string) *DetachPolicyRequest {
	s.PrincipalName = &v
	return s
}

type DetachPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DetachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponse) SetRequestId(v string) *DetachPolicyResponse {
	s.RequestId = &v
	return s
}

type AttachPolicyRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PrincipalType   *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty" require:"true"`
	PrincipalName   *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty" require:"true"`
}

func (s AttachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) SetResourceGroupId(v string) *AttachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyType(v string) *AttachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyName(v string) *AttachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalType(v string) *AttachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalName(v string) *AttachPolicyRequest {
	s.PrincipalName = &v
	return s
}

type AttachPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AttachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponse) SetRequestId(v string) *AttachPolicyResponse {
	s.RequestId = &v
	return s
}

type GetPolicyVersionRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
}

func (s GetPolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionRequest) SetPolicyType(v string) *GetPolicyVersionRequest {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyVersionRequest) SetPolicyName(v string) *GetPolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyVersionRequest) SetVersionId(v string) *GetPolicyVersionRequest {
	s.VersionId = &v
	return s
}

type GetPolicyVersionResponse struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyVersion *GetPolicyVersionResponsePolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" require:"true" type:"Struct"`
}

func (s GetPolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponse) SetRequestId(v string) *GetPolicyVersionResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyVersionResponse) SetPolicyVersion(v *GetPolicyVersionResponsePolicyVersion) *GetPolicyVersionResponse {
	s.PolicyVersion = v
	return s
}

type GetPolicyVersionResponsePolicyVersion struct {
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty" require:"true"`
	PolicyDocument   *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s GetPolicyVersionResponsePolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyVersionResponsePolicyVersion) GoString() string {
	return s.String()
}

func (s *GetPolicyVersionResponsePolicyVersion) SetVersionId(v string) *GetPolicyVersionResponsePolicyVersion {
	s.VersionId = &v
	return s
}

func (s *GetPolicyVersionResponsePolicyVersion) SetIsDefaultVersion(v bool) *GetPolicyVersionResponsePolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetPolicyVersionResponsePolicyVersion) SetPolicyDocument(v string) *GetPolicyVersionResponsePolicyVersion {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyVersionResponsePolicyVersion) SetCreateDate(v string) *GetPolicyVersionResponsePolicyVersion {
	s.CreateDate = &v
	return s
}

type SetDefaultPolicyVersionRequest struct {
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
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

type SetDefaultPolicyVersionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetDefaultPolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionResponse) SetRequestId(v string) *SetDefaultPolicyVersionResponse {
	s.RequestId = &v
	return s
}

type DeleteResourceGroupRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
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

type DeleteResourceGroupResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceGroup *DeleteResourceGroupResponseResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" require:"true" type:"Struct"`
}

func (s DeleteResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) SetRequestId(v string) *DeleteResourceGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetResourceGroup(v *DeleteResourceGroupResponseResourceGroup) *DeleteResourceGroupResponse {
	s.ResourceGroup = v
	return s
}

type DeleteResourceGroupResponseResourceGroup struct {
	Status         *string                                                 `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	AccountId      *string                                                 `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName    *string                                                 `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	Id             *string                                                 `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	CreateDate     *string                                                 `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	Name           *string                                                 `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionStatuses *DeleteResourceGroupResponseResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" require:"true" type:"Struct"`
}

func (s DeleteResourceGroupResponseResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseResourceGroup) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseResourceGroup) SetStatus(v string) *DeleteResourceGroupResponseResourceGroup {
	s.Status = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetAccountId(v string) *DeleteResourceGroupResponseResourceGroup {
	s.AccountId = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetDisplayName(v string) *DeleteResourceGroupResponseResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetId(v string) *DeleteResourceGroupResponseResourceGroup {
	s.Id = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetCreateDate(v string) *DeleteResourceGroupResponseResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetName(v string) *DeleteResourceGroupResponseResourceGroup {
	s.Name = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroup) SetRegionStatuses(v *DeleteResourceGroupResponseResourceGroupRegionStatuses) *DeleteResourceGroupResponseResourceGroup {
	s.RegionStatuses = v
	return s
}

type DeleteResourceGroupResponseResourceGroupRegionStatuses struct {
	RegionStatus []*DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteResourceGroupResponseResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseResourceGroupRegionStatuses) SetRegionStatus(v []*DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus) *DeleteResourceGroupResponseResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *DeleteResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

type GetPolicyRequest struct {
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s *GetPolicyRequest) SetLanguage(v string) *GetPolicyRequest {
	s.Language = &v
	return s
}

type GetPolicyResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Policy    *GetPolicyResponsePolicy `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true" type:"Struct"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetRequestId(v string) *GetPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponse) SetPolicy(v *GetPolicyResponsePolicy) *GetPolicyResponse {
	s.Policy = v
	return s
}

type GetPolicyResponsePolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *int    `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	DefaultVersion  *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty" require:"true"`
	PolicyDocument  *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s GetPolicyResponsePolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponsePolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponsePolicy) SetPolicyType(v string) *GetPolicyResponsePolicy {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetUpdateDate(v string) *GetPolicyResponsePolicy {
	s.UpdateDate = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetDescription(v string) *GetPolicyResponsePolicy {
	s.Description = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetAttachmentCount(v int) *GetPolicyResponsePolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetPolicyName(v string) *GetPolicyResponsePolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetDefaultVersion(v string) *GetPolicyResponsePolicy {
	s.DefaultVersion = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetPolicyDocument(v string) *GetPolicyResponsePolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetPolicyResponsePolicy) SetCreateDate(v string) *GetPolicyResponsePolicy {
	s.CreateDate = &v
	return s
}

type UpdateResourceGroupRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
	NewDisplayName  *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty" require:"true"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) SetResourceGroupId(v string) *UpdateResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetNewDisplayName(v string) *UpdateResourceGroupRequest {
	s.NewDisplayName = &v
	return s
}

type UpdateResourceGroupResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceGroup *UpdateResourceGroupResponseResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" require:"true" type:"Struct"`
}

func (s UpdateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponse) SetRequestId(v string) *UpdateResourceGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponse) SetResourceGroup(v *UpdateResourceGroupResponseResourceGroup) *UpdateResourceGroupResponse {
	s.ResourceGroup = v
	return s
}

type UpdateResourceGroupResponseResourceGroup struct {
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s UpdateResourceGroupResponseResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponseResourceGroup) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseResourceGroup) SetAccountId(v string) *UpdateResourceGroupResponseResourceGroup {
	s.AccountId = &v
	return s
}

func (s *UpdateResourceGroupResponseResourceGroup) SetDisplayName(v string) *UpdateResourceGroupResponseResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateResourceGroupResponseResourceGroup) SetId(v string) *UpdateResourceGroupResponseResourceGroup {
	s.Id = &v
	return s
}

func (s *UpdateResourceGroupResponseResourceGroup) SetCreateDate(v string) *UpdateResourceGroupResponseResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateResourceGroupResponseResourceGroup) SetName(v string) *UpdateResourceGroupResponseResourceGroup {
	s.Name = &v
	return s
}

type ListResourceGroupsRequest struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

type ListResourceGroupsResponse struct {
	TotalCount     *int                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize       *int                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber     *int                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	ResourceGroups *ListResourceGroupsResponseResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" require:"true" type:"Struct"`
}

func (s ListResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) SetTotalCount(v int) *ListResourceGroupsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponse) SetPageSize(v int) *ListResourceGroupsResponse {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsResponse) SetRequestId(v string) *ListResourceGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponse) SetPageNumber(v int) *ListResourceGroupsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsResponse) SetResourceGroups(v *ListResourceGroupsResponseResourceGroups) *ListResourceGroupsResponse {
	s.ResourceGroups = v
	return s
}

type ListResourceGroupsResponseResourceGroups struct {
	ResourceGroup []*ListResourceGroupsResponseResourceGroupsResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceGroupsResponseResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseResourceGroups) SetResourceGroup(v []*ListResourceGroupsResponseResourceGroupsResourceGroup) *ListResourceGroupsResponseResourceGroups {
	s.ResourceGroup = v
	return s
}

type ListResourceGroupsResponseResourceGroupsResourceGroup struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListResourceGroupsResponseResourceGroupsResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseResourceGroupsResourceGroup) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetStatus(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetAccountId(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.AccountId = &v
	return s
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetDisplayName(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetId(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetCreateDate(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *ListResourceGroupsResponseResourceGroupsResourceGroup) SetName(v string) *ListResourceGroupsResponseResourceGroupsResourceGroup {
	s.Name = &v
	return s
}

type ListPoliciesRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetPolicyType(v string) *ListPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesRequest) SetPageNumber(v int) *ListPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesRequest) SetPageSize(v int) *ListPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesRequest) SetLanguage(v string) *ListPoliciesRequest {
	s.Language = &v
	return s
}

type ListPoliciesResponse struct {
	TotalCount *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Policies   *ListPoliciesResponsePolicies `json:"Policies,omitempty" xml:"Policies,omitempty" require:"true" type:"Struct"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetTotalCount(v int) *ListPoliciesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPoliciesResponse) SetPageSize(v int) *ListPoliciesResponse {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesResponse) SetRequestId(v string) *ListPoliciesResponse {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponse) SetPageNumber(v int) *ListPoliciesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesResponse) SetPolicies(v *ListPoliciesResponsePolicies) *ListPoliciesResponse {
	s.Policies = v
	return s
}

type ListPoliciesResponsePolicies struct {
	Policy []*ListPoliciesResponsePoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true" type:"Repeated"`
}

func (s ListPoliciesResponsePolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponsePolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponsePolicies) SetPolicy(v []*ListPoliciesResponsePoliciesPolicy) *ListPoliciesResponsePolicies {
	s.Policy = v
	return s
}

type ListPoliciesResponsePoliciesPolicy struct {
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AttachmentCount *int    `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty" require:"true"`
	PolicyName      *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	DefaultVersion  *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListPoliciesResponsePoliciesPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponsePoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponsePoliciesPolicy) SetPolicyType(v string) *ListPoliciesResponsePoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetUpdateDate(v string) *ListPoliciesResponsePoliciesPolicy {
	s.UpdateDate = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetDescription(v string) *ListPoliciesResponsePoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetAttachmentCount(v int) *ListPoliciesResponsePoliciesPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetPolicyName(v string) *ListPoliciesResponsePoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesResponsePoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesResponsePoliciesPolicy) SetCreateDate(v string) *ListPoliciesResponsePoliciesPolicy {
	s.CreateDate = &v
	return s
}

type ListPolicyVersionsRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
}

func (s ListPolicyVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsRequest) SetPolicyType(v string) *ListPolicyVersionsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyVersionsRequest) SetPolicyName(v string) *ListPolicyVersionsRequest {
	s.PolicyName = &v
	return s
}

type ListPolicyVersionsResponse struct {
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyVersions *ListPolicyVersionsResponsePolicyVersions `json:"PolicyVersions,omitempty" xml:"PolicyVersions,omitempty" require:"true" type:"Struct"`
}

func (s ListPolicyVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponse) SetRequestId(v string) *ListPolicyVersionsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPolicyVersionsResponse) SetPolicyVersions(v *ListPolicyVersionsResponsePolicyVersions) *ListPolicyVersionsResponse {
	s.PolicyVersions = v
	return s
}

type ListPolicyVersionsResponsePolicyVersions struct {
	PolicyVersion []*ListPolicyVersionsResponsePolicyVersionsPolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" require:"true" type:"Repeated"`
}

func (s ListPolicyVersionsResponsePolicyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponsePolicyVersions) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponsePolicyVersions) SetPolicyVersion(v []*ListPolicyVersionsResponsePolicyVersionsPolicyVersion) *ListPolicyVersionsResponsePolicyVersions {
	s.PolicyVersion = v
	return s
}

type ListPolicyVersionsResponsePolicyVersionsPolicyVersion struct {
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListPolicyVersionsResponsePolicyVersionsPolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyVersionsResponsePolicyVersionsPolicyVersion) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsResponsePolicyVersionsPolicyVersion) SetVersionId(v string) *ListPolicyVersionsResponsePolicyVersionsPolicyVersion {
	s.VersionId = &v
	return s
}

func (s *ListPolicyVersionsResponsePolicyVersionsPolicyVersion) SetIsDefaultVersion(v bool) *ListPolicyVersionsResponsePolicyVersionsPolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *ListPolicyVersionsResponsePolicyVersionsPolicyVersion) SetCreateDate(v string) *ListPolicyVersionsResponsePolicyVersionsPolicyVersion {
	s.CreateDate = &v
	return s
}

type CreateResourceAccountRequest struct {
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	ParentFolderId    *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	PayerAccountId    *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
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

func (s *CreateResourceAccountRequest) SetParentFolderId(v string) *CreateResourceAccountRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetPayerAccountId(v string) *CreateResourceAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *CreateResourceAccountRequest) SetAccountNamePrefix(v string) *CreateResourceAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

type CreateResourceAccountResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *CreateResourceAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s CreateResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponse) SetRequestId(v string) *CreateResourceAccountResponse {
	s.RequestId = &v
	return s
}

func (s *CreateResourceAccountResponse) SetAccount(v *CreateResourceAccountResponseAccount) *CreateResourceAccountResponse {
	s.Account = v
	return s
}

type CreateResourceAccountResponseAccount struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s CreateResourceAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *CreateResourceAccountResponseAccount) SetStatus(v string) *CreateResourceAccountResponseAccount {
	s.Status = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetModifyTime(v string) *CreateResourceAccountResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetJoinMethod(v string) *CreateResourceAccountResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetResourceDirectoryId(v string) *CreateResourceAccountResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetType(v string) *CreateResourceAccountResponseAccount {
	s.Type = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetAccountId(v string) *CreateResourceAccountResponseAccount {
	s.AccountId = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetDisplayName(v string) *CreateResourceAccountResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetJoinTime(v string) *CreateResourceAccountResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetFolderId(v string) *CreateResourceAccountResponseAccount {
	s.FolderId = &v
	return s
}

func (s *CreateResourceAccountResponseAccount) SetAccountName(v string) *CreateResourceAccountResponseAccount {
	s.AccountName = &v
	return s
}

type ListHandshakesForResourceDirectoryRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageNumber(v int) *ListHandshakesForResourceDirectoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryRequest) SetPageSize(v int) *ListHandshakesForResourceDirectoryRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForResourceDirectoryResponse struct {
	TotalCount *int                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Handshakes *ListHandshakesForResourceDirectoryResponseHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" require:"true" type:"Struct"`
}

func (s ListHandshakesForResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponse) SetTotalCount(v int) *ListHandshakesForResourceDirectoryResponse {
	s.TotalCount = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetPageSize(v int) *ListHandshakesForResourceDirectoryResponse {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetRequestId(v string) *ListHandshakesForResourceDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetPageNumber(v int) *ListHandshakesForResourceDirectoryResponse {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponse) SetHandshakes(v *ListHandshakesForResourceDirectoryResponseHandshakes) *ListHandshakesForResourceDirectoryResponse {
	s.Handshakes = v
	return s
}

type ListHandshakesForResourceDirectoryResponseHandshakes struct {
	Handshake []*ListHandshakesForResourceDirectoryResponseHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Repeated"`
}

func (s ListHandshakesForResourceDirectoryResponseHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakes) SetHandshake(v []*ListHandshakesForResourceDirectoryResponseHandshakesHandshake) *ListHandshakesForResourceDirectoryResponseHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForResourceDirectoryResponseHandshakesHandshake struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
}

func (s ListHandshakesForResourceDirectoryResponseHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForResourceDirectoryResponseHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetStatus(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetNote(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetTargetType(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.TargetType = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForResourceDirectoryResponseHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForResourceDirectoryResponseHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

type DestroyResourceDirectoryRequest struct {
}

func (s DestroyResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryRequest) GoString() string {
	return s.String()
}

type DestroyResourceDirectoryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DestroyResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DestroyResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponse) SetRequestId(v string) *DestroyResourceDirectoryResponse {
	s.RequestId = &v
	return s
}

type CreatePolicyVersionRequest struct {
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
	SetAsDefault   *bool   `json:"SetAsDefault,omitempty" xml:"SetAsDefault,omitempty"`
}

func (s CreatePolicyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionRequest) SetPolicyName(v string) *CreatePolicyVersionRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetPolicyDocument(v string) *CreatePolicyVersionRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyVersionRequest) SetSetAsDefault(v bool) *CreatePolicyVersionRequest {
	s.SetAsDefault = &v
	return s
}

type CreatePolicyVersionResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyVersion *CreatePolicyVersionResponsePolicyVersion `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty" require:"true" type:"Struct"`
}

func (s CreatePolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponse) SetRequestId(v string) *CreatePolicyVersionResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyVersionResponse) SetPolicyVersion(v *CreatePolicyVersionResponsePolicyVersion) *CreatePolicyVersionResponse {
	s.PolicyVersion = v
	return s
}

type CreatePolicyVersionResponsePolicyVersion struct {
	VersionId        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	IsDefaultVersion *bool   `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s CreatePolicyVersionResponsePolicyVersion) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyVersionResponsePolicyVersion) GoString() string {
	return s.String()
}

func (s *CreatePolicyVersionResponsePolicyVersion) SetVersionId(v string) *CreatePolicyVersionResponsePolicyVersion {
	s.VersionId = &v
	return s
}

func (s *CreatePolicyVersionResponsePolicyVersion) SetIsDefaultVersion(v bool) *CreatePolicyVersionResponsePolicyVersion {
	s.IsDefaultVersion = &v
	return s
}

func (s *CreatePolicyVersionResponsePolicyVersion) SetCreateDate(v string) *CreatePolicyVersionResponsePolicyVersion {
	s.CreateDate = &v
	return s
}

type DeletePolicyVersionRequest struct {
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
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

type DeletePolicyVersionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeletePolicyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionResponse) SetRequestId(v string) *DeletePolicyVersionResponse {
	s.RequestId = &v
	return s
}

type GetResourceGroupRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
}

func (s GetResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) SetResourceGroupId(v string) *GetResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type GetResourceGroupResponse struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceGroup *GetResourceGroupResponseResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" require:"true" type:"Struct"`
}

func (s GetResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponse) SetRequestId(v string) *GetResourceGroupResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponse) SetResourceGroup(v *GetResourceGroupResponseResourceGroup) *GetResourceGroupResponse {
	s.ResourceGroup = v
	return s
}

type GetResourceGroupResponseResourceGroup struct {
	Status         *string                                              `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	AccountId      *string                                              `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName    *string                                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	Id             *string                                              `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	CreateDate     *string                                              `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	Name           *string                                              `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionStatuses *GetResourceGroupResponseResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" require:"true" type:"Struct"`
}

func (s GetResourceGroupResponseResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseResourceGroup) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseResourceGroup) SetStatus(v string) *GetResourceGroupResponseResourceGroup {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetAccountId(v string) *GetResourceGroupResponseResourceGroup {
	s.AccountId = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetDisplayName(v string) *GetResourceGroupResponseResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetId(v string) *GetResourceGroupResponseResourceGroup {
	s.Id = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetCreateDate(v string) *GetResourceGroupResponseResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetName(v string) *GetResourceGroupResponseResourceGroup {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroup) SetRegionStatuses(v *GetResourceGroupResponseResourceGroupRegionStatuses) *GetResourceGroupResponseResourceGroup {
	s.RegionStatuses = v
	return s
}

type GetResourceGroupResponseResourceGroupRegionStatuses struct {
	RegionStatus []*GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" require:"true" type:"Repeated"`
}

func (s GetResourceGroupResponseResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseResourceGroupRegionStatuses) SetRegionStatus(v []*GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus) *GetResourceGroupResponseResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *GetResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

type InitResourceDirectoryRequest struct {
}

func (s InitResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryRequest) GoString() string {
	return s.String()
}

type InitResourceDirectoryResponse struct {
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceDirectory *InitResourceDirectoryResponseResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" require:"true" type:"Struct"`
}

func (s InitResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponse) SetRequestId(v string) *InitResourceDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *InitResourceDirectoryResponse) SetResourceDirectory(v *InitResourceDirectoryResponseResourceDirectory) *InitResourceDirectoryResponse {
	s.ResourceDirectory = v
	return s
}

type InitResourceDirectoryResponseResourceDirectory struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	RootFolderId        *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s InitResourceDirectoryResponseResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s InitResourceDirectoryResponseResourceDirectory) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponseResourceDirectory) SetResourceDirectoryId(v string) *InitResourceDirectoryResponseResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InitResourceDirectoryResponseResourceDirectory) SetMasterAccountId(v string) *InitResourceDirectoryResponseResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *InitResourceDirectoryResponseResourceDirectory) SetMasterAccountName(v string) *InitResourceDirectoryResponseResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *InitResourceDirectoryResponseResourceDirectory) SetRootFolderId(v string) *InitResourceDirectoryResponseResourceDirectory {
	s.RootFolderId = &v
	return s
}

func (s *InitResourceDirectoryResponseResourceDirectory) SetCreateTime(v string) *InitResourceDirectoryResponseResourceDirectory {
	s.CreateTime = &v
	return s
}

type GetHandshakeRequest struct {
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
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

type GetHandshakeResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Handshake *GetHandshakeResponseHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Struct"`
}

func (s GetHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponse) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponse) SetRequestId(v string) *GetHandshakeResponse {
	s.RequestId = &v
	return s
}

func (s *GetHandshakeResponse) SetHandshake(v *GetHandshakeResponseHandshake) *GetHandshakeResponse {
	s.Handshake = v
	return s
}

type GetHandshakeResponseHandshake struct {
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime             *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ResourceDirectoryId    *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	HandshakeId            *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	MasterAccountName      *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	Note                   *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	TargetType             *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	MasterAccountId        *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	MasterAccountRealName  *string `json:"MasterAccountRealName,omitempty" xml:"MasterAccountRealName,omitempty" require:"true"`
	ExpireTime             *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	InvitedAccountRealName *string `json:"InvitedAccountRealName,omitempty" xml:"InvitedAccountRealName,omitempty" require:"true"`
	TargetEntity           *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
}

func (s GetHandshakeResponseHandshake) String() string {
	return tea.Prettify(s)
}

func (s GetHandshakeResponseHandshake) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponseHandshake) SetStatus(v string) *GetHandshakeResponseHandshake {
	s.Status = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetModifyTime(v string) *GetHandshakeResponseHandshake {
	s.ModifyTime = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetResourceDirectoryId(v string) *GetHandshakeResponseHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetHandshakeId(v string) *GetHandshakeResponseHandshake {
	s.HandshakeId = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetMasterAccountName(v string) *GetHandshakeResponseHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetNote(v string) *GetHandshakeResponseHandshake {
	s.Note = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetCreateTime(v string) *GetHandshakeResponseHandshake {
	s.CreateTime = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetTargetType(v string) *GetHandshakeResponseHandshake {
	s.TargetType = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetMasterAccountId(v string) *GetHandshakeResponseHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetMasterAccountRealName(v string) *GetHandshakeResponseHandshake {
	s.MasterAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetExpireTime(v string) *GetHandshakeResponseHandshake {
	s.ExpireTime = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetInvitedAccountRealName(v string) *GetHandshakeResponseHandshake {
	s.InvitedAccountRealName = &v
	return s
}

func (s *GetHandshakeResponseHandshake) SetTargetEntity(v string) *GetHandshakeResponseHandshake {
	s.TargetEntity = &v
	return s
}

type CancelHandshakeRequest struct {
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
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

type CancelHandshakeResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Handshake *CancelHandshakeResponseHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Struct"`
}

func (s CancelHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponse) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponse) SetRequestId(v string) *CancelHandshakeResponse {
	s.RequestId = &v
	return s
}

func (s *CancelHandshakeResponse) SetHandshake(v *CancelHandshakeResponseHandshake) *CancelHandshakeResponse {
	s.Handshake = v
	return s
}

type CancelHandshakeResponseHandshake struct {
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
}

func (s CancelHandshakeResponseHandshake) String() string {
	return tea.Prettify(s)
}

func (s CancelHandshakeResponseHandshake) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponseHandshake) SetHandshakeId(v string) *CancelHandshakeResponseHandshake {
	s.HandshakeId = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetResourceDirectoryId(v string) *CancelHandshakeResponseHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetMasterAccountId(v string) *CancelHandshakeResponseHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetMasterAccountName(v string) *CancelHandshakeResponseHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetTargetEntity(v string) *CancelHandshakeResponseHandshake {
	s.TargetEntity = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetTargetType(v string) *CancelHandshakeResponseHandshake {
	s.TargetType = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetNote(v string) *CancelHandshakeResponseHandshake {
	s.Note = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetStatus(v string) *CancelHandshakeResponseHandshake {
	s.Status = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetCreateTime(v string) *CancelHandshakeResponseHandshake {
	s.CreateTime = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetModifyTime(v string) *CancelHandshakeResponseHandshake {
	s.ModifyTime = &v
	return s
}

func (s *CancelHandshakeResponseHandshake) SetExpireTime(v string) *CancelHandshakeResponseHandshake {
	s.ExpireTime = &v
	return s
}

type CreatePolicyRequest struct {
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty" require:"true"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDocument(v string) *CreatePolicyRequest {
	s.PolicyDocument = &v
	return s
}

type CreatePolicyResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Policy    *CreatePolicyResponsePolicy `json:"Policy,omitempty" xml:"Policy,omitempty" require:"true" type:"Struct"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetRequestId(v string) *CreatePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponse) SetPolicy(v *CreatePolicyResponsePolicy) *CreatePolicyResponse {
	s.Policy = v
	return s
}

type CreatePolicyResponsePolicy struct {
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty" require:"true"`
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s CreatePolicyResponsePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponsePolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponsePolicy) SetPolicyType(v string) *CreatePolicyResponsePolicy {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyResponsePolicy) SetDescription(v string) *CreatePolicyResponsePolicy {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponsePolicy) SetPolicyName(v string) *CreatePolicyResponsePolicy {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponsePolicy) SetDefaultVersion(v string) *CreatePolicyResponsePolicy {
	s.DefaultVersion = &v
	return s
}

func (s *CreatePolicyResponsePolicy) SetCreateDate(v string) *CreatePolicyResponsePolicy {
	s.CreateDate = &v
	return s
}

type DeclineHandshakeRequest struct {
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
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

type DeclineHandshakeResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Handshake *DeclineHandshakeResponseHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Struct"`
}

func (s DeclineHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponse) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponse) SetRequestId(v string) *DeclineHandshakeResponse {
	s.RequestId = &v
	return s
}

func (s *DeclineHandshakeResponse) SetHandshake(v *DeclineHandshakeResponseHandshake) *DeclineHandshakeResponse {
	s.Handshake = v
	return s
}

type DeclineHandshakeResponseHandshake struct {
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
}

func (s DeclineHandshakeResponseHandshake) String() string {
	return tea.Prettify(s)
}

func (s DeclineHandshakeResponseHandshake) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponseHandshake) SetHandshakeId(v string) *DeclineHandshakeResponseHandshake {
	s.HandshakeId = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetResourceDirectoryId(v string) *DeclineHandshakeResponseHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetMasterAccountId(v string) *DeclineHandshakeResponseHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetMasterAccountName(v string) *DeclineHandshakeResponseHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetTargetEntity(v string) *DeclineHandshakeResponseHandshake {
	s.TargetEntity = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetTargetType(v string) *DeclineHandshakeResponseHandshake {
	s.TargetType = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetNote(v string) *DeclineHandshakeResponseHandshake {
	s.Note = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetStatus(v string) *DeclineHandshakeResponseHandshake {
	s.Status = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetCreateTime(v string) *DeclineHandshakeResponseHandshake {
	s.CreateTime = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetModifyTime(v string) *DeclineHandshakeResponseHandshake {
	s.ModifyTime = &v
	return s
}

func (s *DeclineHandshakeResponseHandshake) SetExpireTime(v string) *DeclineHandshakeResponseHandshake {
	s.ExpireTime = &v
	return s
}

type DeletePolicyRequest struct {
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" require:"true"`
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

type DeletePolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetRequestId(v string) *DeletePolicyResponse {
	s.RequestId = &v
	return s
}

type ListHandshakesForAccountRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountRequest) SetPageNumber(v int) *ListHandshakesForAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountRequest) SetPageSize(v int) *ListHandshakesForAccountRequest {
	s.PageSize = &v
	return s
}

type ListHandshakesForAccountResponse struct {
	TotalCount *int                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Handshakes *ListHandshakesForAccountResponseHandshakes `json:"Handshakes,omitempty" xml:"Handshakes,omitempty" require:"true" type:"Struct"`
}

func (s ListHandshakesForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponse) SetTotalCount(v int) *ListHandshakesForAccountResponse {
	s.TotalCount = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetPageSize(v int) *ListHandshakesForAccountResponse {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetRequestId(v string) *ListHandshakesForAccountResponse {
	s.RequestId = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetPageNumber(v int) *ListHandshakesForAccountResponse {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountResponse) SetHandshakes(v *ListHandshakesForAccountResponseHandshakes) *ListHandshakesForAccountResponse {
	s.Handshakes = v
	return s
}

type ListHandshakesForAccountResponseHandshakes struct {
	Handshake []*ListHandshakesForAccountResponseHandshakesHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Repeated"`
}

func (s ListHandshakesForAccountResponseHandshakes) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseHandshakes) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseHandshakes) SetHandshake(v []*ListHandshakesForAccountResponseHandshakesHandshake) *ListHandshakesForAccountResponseHandshakes {
	s.Handshake = v
	return s
}

type ListHandshakesForAccountResponseHandshakesHandshake struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
}

func (s ListHandshakesForAccountResponseHandshakesHandshake) String() string {
	return tea.Prettify(s)
}

func (s ListHandshakesForAccountResponseHandshakesHandshake) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetStatus(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.Status = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetModifyTime(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.ModifyTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetResourceDirectoryId(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetHandshakeId(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.HandshakeId = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetMasterAccountName(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetNote(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.Note = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetCreateTime(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.CreateTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetTargetType(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.TargetType = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetMasterAccountId(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetExpireTime(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.ExpireTime = &v
	return s
}

func (s *ListHandshakesForAccountResponseHandshakesHandshake) SetTargetEntity(v string) *ListHandshakesForAccountResponseHandshakesHandshake {
	s.TargetEntity = &v
	return s
}

type InviteAccountToResourceDirectoryRequest struct {
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
	TargetType   *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	Note         *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s InviteAccountToResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetEntity(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetEntity = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetTargetType(v string) *InviteAccountToResourceDirectoryRequest {
	s.TargetType = &v
	return s
}

func (s *InviteAccountToResourceDirectoryRequest) SetNote(v string) *InviteAccountToResourceDirectoryRequest {
	s.Note = &v
	return s
}

type InviteAccountToResourceDirectoryResponse struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Handshake *InviteAccountToResourceDirectoryResponseHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Struct"`
}

func (s InviteAccountToResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponse) SetRequestId(v string) *InviteAccountToResourceDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponse) SetHandshake(v *InviteAccountToResourceDirectoryResponseHandshake) *InviteAccountToResourceDirectoryResponse {
	s.Handshake = v
	return s
}

type InviteAccountToResourceDirectoryResponseHandshake struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
}

func (s InviteAccountToResourceDirectoryResponseHandshake) String() string {
	return tea.Prettify(s)
}

func (s InviteAccountToResourceDirectoryResponseHandshake) GoString() string {
	return s.String()
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetStatus(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.Status = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetModifyTime(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.ModifyTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetResourceDirectoryId(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetHandshakeId(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.HandshakeId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetMasterAccountName(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetNote(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.Note = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetCreateTime(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.CreateTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetTargetType(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.TargetType = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetMasterAccountId(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetExpireTime(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.ExpireTime = &v
	return s
}

func (s *InviteAccountToResourceDirectoryResponseHandshake) SetTargetEntity(v string) *InviteAccountToResourceDirectoryResponseHandshake {
	s.TargetEntity = &v
	return s
}

type AcceptHandshakeRequest struct {
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
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

type AcceptHandshakeResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Handshake *AcceptHandshakeResponseHandshake `json:"Handshake,omitempty" xml:"Handshake,omitempty" require:"true" type:"Struct"`
}

func (s AcceptHandshakeResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponse) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponse) SetRequestId(v string) *AcceptHandshakeResponse {
	s.RequestId = &v
	return s
}

func (s *AcceptHandshakeResponse) SetHandshake(v *AcceptHandshakeResponseHandshake) *AcceptHandshakeResponse {
	s.Handshake = v
	return s
}

type AcceptHandshakeResponseHandshake struct {
	HandshakeId         *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	TargetEntity        *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty" require:"true"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty" require:"true"`
	Note                *string `json:"Note,omitempty" xml:"Note,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ExpireTime          *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
}

func (s AcceptHandshakeResponseHandshake) String() string {
	return tea.Prettify(s)
}

func (s AcceptHandshakeResponseHandshake) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponseHandshake) SetHandshakeId(v string) *AcceptHandshakeResponseHandshake {
	s.HandshakeId = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetResourceDirectoryId(v string) *AcceptHandshakeResponseHandshake {
	s.ResourceDirectoryId = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetMasterAccountId(v string) *AcceptHandshakeResponseHandshake {
	s.MasterAccountId = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetMasterAccountName(v string) *AcceptHandshakeResponseHandshake {
	s.MasterAccountName = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetTargetEntity(v string) *AcceptHandshakeResponseHandshake {
	s.TargetEntity = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetTargetType(v string) *AcceptHandshakeResponseHandshake {
	s.TargetType = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetNote(v string) *AcceptHandshakeResponseHandshake {
	s.Note = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetStatus(v string) *AcceptHandshakeResponseHandshake {
	s.Status = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetCreateTime(v string) *AcceptHandshakeResponseHandshake {
	s.CreateTime = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetModifyTime(v string) *AcceptHandshakeResponseHandshake {
	s.ModifyTime = &v
	return s
}

func (s *AcceptHandshakeResponseHandshake) SetExpireTime(v string) *AcceptHandshakeResponseHandshake {
	s.ExpireTime = &v
	return s
}

type UpdateAccountRequest struct {
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
}

func (s UpdateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountRequest) SetNewDisplayName(v string) *UpdateAccountRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateAccountRequest) SetNewAccountType(v string) *UpdateAccountRequest {
	s.NewAccountType = &v
	return s
}

func (s *UpdateAccountRequest) SetAccountId(v string) *UpdateAccountRequest {
	s.AccountId = &v
	return s
}

type UpdateAccountResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *UpdateAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s UpdateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponse) SetRequestId(v string) *UpdateAccountResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountResponse) SetAccount(v *UpdateAccountResponseAccount) *UpdateAccountResponse {
	s.Account = v
	return s
}

type UpdateAccountResponseAccount struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s UpdateAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountResponseAccount) SetStatus(v string) *UpdateAccountResponseAccount {
	s.Status = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetModifyTime(v string) *UpdateAccountResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetJoinMethod(v string) *UpdateAccountResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetResourceDirectoryId(v string) *UpdateAccountResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetType(v string) *UpdateAccountResponseAccount {
	s.Type = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetAccountId(v string) *UpdateAccountResponseAccount {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetDisplayName(v string) *UpdateAccountResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetJoinTime(v string) *UpdateAccountResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetFolderId(v string) *UpdateAccountResponseAccount {
	s.FolderId = &v
	return s
}

func (s *UpdateAccountResponseAccount) SetAccountName(v string) *UpdateAccountResponseAccount {
	s.AccountName = &v
	return s
}

type GetFolderRequest struct {
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
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

type GetFolderResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Folder    *GetFolderResponseFolder `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true" type:"Struct"`
}

func (s GetFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponse) GoString() string {
	return s.String()
}

func (s *GetFolderResponse) SetRequestId(v string) *GetFolderResponse {
	s.RequestId = &v
	return s
}

func (s *GetFolderResponse) SetFolder(v *GetFolderResponseFolder) *GetFolderResponse {
	s.Folder = v
	return s
}

type GetFolderResponseFolder struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty" require:"true"`
}

func (s GetFolderResponseFolder) String() string {
	return tea.Prettify(s)
}

func (s GetFolderResponseFolder) GoString() string {
	return s.String()
}

func (s *GetFolderResponseFolder) SetCreateTime(v string) *GetFolderResponseFolder {
	s.CreateTime = &v
	return s
}

func (s *GetFolderResponseFolder) SetFolderId(v string) *GetFolderResponseFolder {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseFolder) SetFolderName(v string) *GetFolderResponseFolder {
	s.FolderName = &v
	return s
}

func (s *GetFolderResponseFolder) SetParentFolderId(v string) *GetFolderResponseFolder {
	s.ParentFolderId = &v
	return s
}

type ListAccountsForParentRequest struct {
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAccountsForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentRequest) SetParentFolderId(v string) *ListAccountsForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListAccountsForParentRequest) SetQueryKeyword(v string) *ListAccountsForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageNumber(v int) *ListAccountsForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentRequest) SetPageSize(v int) *ListAccountsForParentRequest {
	s.PageSize = &v
	return s
}

type ListAccountsForParentResponse struct {
	TotalCount *int                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Accounts   *ListAccountsForParentResponseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" require:"true" type:"Struct"`
}

func (s ListAccountsForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponse) SetTotalCount(v int) *ListAccountsForParentResponse {
	s.TotalCount = &v
	return s
}

func (s *ListAccountsForParentResponse) SetPageSize(v int) *ListAccountsForParentResponse {
	s.PageSize = &v
	return s
}

func (s *ListAccountsForParentResponse) SetRequestId(v string) *ListAccountsForParentResponse {
	s.RequestId = &v
	return s
}

func (s *ListAccountsForParentResponse) SetPageNumber(v int) *ListAccountsForParentResponse {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsForParentResponse) SetAccounts(v *ListAccountsForParentResponseAccounts) *ListAccountsForParentResponse {
	s.Accounts = v
	return s
}

type ListAccountsForParentResponseAccounts struct {
	Account []*ListAccountsForParentResponseAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Repeated"`
}

func (s ListAccountsForParentResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseAccounts) SetAccount(v []*ListAccountsForParentResponseAccountsAccount) *ListAccountsForParentResponseAccounts {
	s.Account = v
	return s
}

type ListAccountsForParentResponseAccountsAccount struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
}

func (s ListAccountsForParentResponseAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsForParentResponseAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsForParentResponseAccountsAccount) SetStatus(v string) *ListAccountsForParentResponseAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetModifyTime(v string) *ListAccountsForParentResponseAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetJoinMethod(v string) *ListAccountsForParentResponseAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsForParentResponseAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetType(v string) *ListAccountsForParentResponseAccountsAccount {
	s.Type = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetAccountId(v string) *ListAccountsForParentResponseAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetDisplayName(v string) *ListAccountsForParentResponseAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetJoinTime(v string) *ListAccountsForParentResponseAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsForParentResponseAccountsAccount) SetFolderId(v string) *ListAccountsForParentResponseAccountsAccount {
	s.FolderId = &v
	return s
}

type CreateResourceGroupRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetDisplayName(v string) *CreateResourceGroupRequest {
	s.DisplayName = &v
	return s
}

type CreateResourceGroupResponse struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceGroup *CreateResourceGroupResponseResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" require:"true" type:"Struct"`
}

func (s CreateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) SetRequestId(v string) *CreateResourceGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponse) SetResourceGroup(v *CreateResourceGroupResponseResourceGroup) *CreateResourceGroupResponse {
	s.ResourceGroup = v
	return s
}

type CreateResourceGroupResponseResourceGroup struct {
	Status         *string                                                 `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	AccountId      *string                                                 `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName    *string                                                 `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	Id             *string                                                 `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	CreateDate     *string                                                 `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	Name           *string                                                 `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionStatuses *CreateResourceGroupResponseResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" require:"true" type:"Struct"`
}

func (s CreateResourceGroupResponseResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseResourceGroup) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseResourceGroup) SetStatus(v string) *CreateResourceGroupResponseResourceGroup {
	s.Status = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetAccountId(v string) *CreateResourceGroupResponseResourceGroup {
	s.AccountId = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetDisplayName(v string) *CreateResourceGroupResponseResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetId(v string) *CreateResourceGroupResponseResourceGroup {
	s.Id = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetCreateDate(v string) *CreateResourceGroupResponseResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetName(v string) *CreateResourceGroupResponseResourceGroup {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroup) SetRegionStatuses(v *CreateResourceGroupResponseResourceGroupRegionStatuses) *CreateResourceGroupResponseResourceGroup {
	s.RegionStatuses = v
	return s
}

type CreateResourceGroupResponseResourceGroupRegionStatuses struct {
	RegionStatus []*CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" require:"true" type:"Repeated"`
}

func (s CreateResourceGroupResponseResourceGroupRegionStatuses) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseResourceGroupRegionStatuses) SetRegionStatus(v []*CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus) *CreateResourceGroupResponseResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

type CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *CreateResourceGroupResponseResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

type PromoteResourceAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
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

type PromoteResourceAccountResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *PromoteResourceAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s PromoteResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponse) SetRequestId(v string) *PromoteResourceAccountResponse {
	s.RequestId = &v
	return s
}

func (s *PromoteResourceAccountResponse) SetAccount(v *PromoteResourceAccountResponseAccount) *PromoteResourceAccountResponse {
	s.Account = v
	return s
}

type PromoteResourceAccountResponseAccount struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s PromoteResourceAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s PromoteResourceAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *PromoteResourceAccountResponseAccount) SetResourceDirectoryId(v string) *PromoteResourceAccountResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetAccountId(v string) *PromoteResourceAccountResponseAccount {
	s.AccountId = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetDisplayName(v string) *PromoteResourceAccountResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetFolderId(v string) *PromoteResourceAccountResponseAccount {
	s.FolderId = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetJoinMethod(v string) *PromoteResourceAccountResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetJoinTime(v string) *PromoteResourceAccountResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetType(v string) *PromoteResourceAccountResponseAccount {
	s.Type = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetStatus(v string) *PromoteResourceAccountResponseAccount {
	s.Status = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetRecordId(v string) *PromoteResourceAccountResponseAccount {
	s.RecordId = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetModifyTime(v string) *PromoteResourceAccountResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *PromoteResourceAccountResponseAccount) SetAccountName(v string) *PromoteResourceAccountResponseAccount {
	s.AccountName = &v
	return s
}

type ListAccountsRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetPageNumber(v int) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int) *ListAccountsRequest {
	s.PageSize = &v
	return s
}

type ListAccountsResponse struct {
	TotalCount *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Accounts   *ListAccountsResponseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" require:"true" type:"Struct"`
}

func (s ListAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) SetTotalCount(v int) *ListAccountsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListAccountsResponse) SetPageSize(v int) *ListAccountsResponse {
	s.PageSize = &v
	return s
}

func (s *ListAccountsResponse) SetRequestId(v string) *ListAccountsResponse {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponse) SetPageNumber(v int) *ListAccountsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsResponse) SetAccounts(v *ListAccountsResponseAccounts) *ListAccountsResponse {
	s.Accounts = v
	return s
}

type ListAccountsResponseAccounts struct {
	Account []*ListAccountsResponseAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Repeated"`
}

func (s ListAccountsResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseAccounts) SetAccount(v []*ListAccountsResponseAccountsAccount) *ListAccountsResponseAccounts {
	s.Account = v
	return s
}

type ListAccountsResponseAccountsAccount struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
}

func (s ListAccountsResponseAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseAccountsAccount) SetStatus(v string) *ListAccountsResponseAccountsAccount {
	s.Status = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetModifyTime(v string) *ListAccountsResponseAccountsAccount {
	s.ModifyTime = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetJoinMethod(v string) *ListAccountsResponseAccountsAccount {
	s.JoinMethod = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetResourceDirectoryId(v string) *ListAccountsResponseAccountsAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetType(v string) *ListAccountsResponseAccountsAccount {
	s.Type = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetAccountId(v string) *ListAccountsResponseAccountsAccount {
	s.AccountId = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetDisplayName(v string) *ListAccountsResponseAccountsAccount {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetJoinTime(v string) *ListAccountsResponseAccountsAccount {
	s.JoinTime = &v
	return s
}

func (s *ListAccountsResponseAccountsAccount) SetFolderId(v string) *ListAccountsResponseAccountsAccount {
	s.FolderId = &v
	return s
}

type CancelPromoteResourceAccountRequest struct {
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
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

type CancelPromoteResourceAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CancelPromoteResourceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPromoteResourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountResponse) SetRequestId(v string) *CancelPromoteResourceAccountResponse {
	s.RequestId = &v
	return s
}

type CreateFolderRequest struct {
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
}

func (s CreateFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) SetParentFolderId(v string) *CreateFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateFolderRequest) SetFolderName(v string) *CreateFolderRequest {
	s.FolderName = &v
	return s
}

type CreateFolderResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Folder    *CreateFolderResponseFolder `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true" type:"Struct"`
}

func (s CreateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateFolderResponse) SetRequestId(v string) *CreateFolderResponse {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponse) SetFolder(v *CreateFolderResponseFolder) *CreateFolderResponse {
	s.Folder = v
	return s
}

type CreateFolderResponseFolder struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty" require:"true"`
}

func (s CreateFolderResponseFolder) String() string {
	return tea.Prettify(s)
}

func (s CreateFolderResponseFolder) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseFolder) SetCreateTime(v string) *CreateFolderResponseFolder {
	s.CreateTime = &v
	return s
}

func (s *CreateFolderResponseFolder) SetFolderId(v string) *CreateFolderResponseFolder {
	s.FolderId = &v
	return s
}

func (s *CreateFolderResponseFolder) SetFolderName(v string) *CreateFolderResponseFolder {
	s.FolderName = &v
	return s
}

func (s *CreateFolderResponseFolder) SetParentFolderId(v string) *CreateFolderResponseFolder {
	s.ParentFolderId = &v
	return s
}

type DeleteFolderRequest struct {
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
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

type DeleteFolderResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponse) SetRequestId(v string) *DeleteFolderResponse {
	s.RequestId = &v
	return s
}

type GetAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
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

type GetAccountResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *GetAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s GetAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponse) GoString() string {
	return s.String()
}

func (s *GetAccountResponse) SetRequestId(v string) *GetAccountResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccountResponse) SetAccount(v *GetAccountResponseAccount) *GetAccountResponse {
	s.Account = v
	return s
}

type GetAccountResponseAccount struct {
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s GetAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s GetAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *GetAccountResponseAccount) SetIdentityInformation(v string) *GetAccountResponseAccount {
	s.IdentityInformation = &v
	return s
}

func (s *GetAccountResponseAccount) SetStatus(v string) *GetAccountResponseAccount {
	s.Status = &v
	return s
}

func (s *GetAccountResponseAccount) SetModifyTime(v string) *GetAccountResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *GetAccountResponseAccount) SetJoinMethod(v string) *GetAccountResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *GetAccountResponseAccount) SetResourceDirectoryId(v string) *GetAccountResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetAccountResponseAccount) SetType(v string) *GetAccountResponseAccount {
	s.Type = &v
	return s
}

func (s *GetAccountResponseAccount) SetAccountId(v string) *GetAccountResponseAccount {
	s.AccountId = &v
	return s
}

func (s *GetAccountResponseAccount) SetDisplayName(v string) *GetAccountResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *GetAccountResponseAccount) SetJoinTime(v string) *GetAccountResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *GetAccountResponseAccount) SetFolderId(v string) *GetAccountResponseAccount {
	s.FolderId = &v
	return s
}

func (s *GetAccountResponseAccount) SetAccountName(v string) *GetAccountResponseAccount {
	s.AccountName = &v
	return s
}

type GetResourceDirectoryRequest struct {
}

func (s GetResourceDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryRequest) GoString() string {
	return s.String()
}

type GetResourceDirectoryResponse struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceDirectory *GetResourceDirectoryResponseResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" require:"true" type:"Struct"`
}

func (s GetResourceDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponse) SetRequestId(v string) *GetResourceDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceDirectoryResponse) SetResourceDirectory(v *GetResourceDirectoryResponseResourceDirectory) *GetResourceDirectoryResponse {
	s.ResourceDirectory = v
	return s
}

type GetResourceDirectoryResponseResourceDirectory struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	MasterAccountName   *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	RootFolderId        *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty" require:"true"`
	ControlPolicyStatus *string `json:"ControlPolicyStatus,omitempty" xml:"ControlPolicyStatus,omitempty" require:"true"`
	MasterAccountId     *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty" require:"true"`
}

func (s GetResourceDirectoryResponseResourceDirectory) String() string {
	return tea.Prettify(s)
}

func (s GetResourceDirectoryResponseResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetResourceDirectoryId(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetMasterAccountName(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetCreateTime(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetRootFolderId(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.RootFolderId = &v
	return s
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetControlPolicyStatus(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.ControlPolicyStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseResourceDirectory) SetMasterAccountId(v string) *GetResourceDirectoryResponseResourceDirectory {
	s.MasterAccountId = &v
	return s
}

type UpdateFolderRequest struct {
	FolderId      *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	NewFolderName *string `json:"NewFolderName,omitempty" xml:"NewFolderName,omitempty" require:"true"`
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

type UpdateFolderResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Folder    *UpdateFolderResponseFolder `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true" type:"Struct"`
}

func (s UpdateFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponse) SetRequestId(v string) *UpdateFolderResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateFolderResponse) SetFolder(v *UpdateFolderResponseFolder) *UpdateFolderResponse {
	s.Folder = v
	return s
}

type UpdateFolderResponseFolder struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FolderId       *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	FolderName     *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty" require:"true"`
}

func (s UpdateFolderResponseFolder) String() string {
	return tea.Prettify(s)
}

func (s UpdateFolderResponseFolder) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseFolder) SetCreateTime(v string) *UpdateFolderResponseFolder {
	s.CreateTime = &v
	return s
}

func (s *UpdateFolderResponseFolder) SetFolderId(v string) *UpdateFolderResponseFolder {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderResponseFolder) SetFolderName(v string) *UpdateFolderResponseFolder {
	s.FolderName = &v
	return s
}

func (s *UpdateFolderResponseFolder) SetParentFolderId(v string) *UpdateFolderResponseFolder {
	s.ParentFolderId = &v
	return s
}

type MoveAccountRequest struct {
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DestinationFolderId *string `json:"DestinationFolderId,omitempty" xml:"DestinationFolderId,omitempty" require:"true"`
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

type MoveAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s MoveAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAccountResponse) GoString() string {
	return s.String()
}

func (s *MoveAccountResponse) SetRequestId(v string) *MoveAccountResponse {
	s.RequestId = &v
	return s
}

type ListAncestorsRequest struct {
	ChildId *string `json:"ChildId,omitempty" xml:"ChildId,omitempty" require:"true"`
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

type ListAncestorsResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Folders   *ListAncestorsResponseFolders `json:"Folders,omitempty" xml:"Folders,omitempty" require:"true" type:"Struct"`
}

func (s ListAncestorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponse) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponse) SetRequestId(v string) *ListAncestorsResponse {
	s.RequestId = &v
	return s
}

func (s *ListAncestorsResponse) SetFolders(v *ListAncestorsResponseFolders) *ListAncestorsResponse {
	s.Folders = v
	return s
}

type ListAncestorsResponseFolders struct {
	Folder []*ListAncestorsResponseFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true" type:"Repeated"`
}

func (s ListAncestorsResponseFolders) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseFolders) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseFolders) SetFolder(v []*ListAncestorsResponseFoldersFolder) *ListAncestorsResponseFolders {
	s.Folder = v
	return s
}

type ListAncestorsResponseFoldersFolder struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
}

func (s ListAncestorsResponseFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListAncestorsResponseFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseFoldersFolder) SetCreateTime(v string) *ListAncestorsResponseFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListAncestorsResponseFoldersFolder) SetFolderId(v string) *ListAncestorsResponseFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAncestorsResponseFoldersFolder) SetFolderName(v string) *ListAncestorsResponseFoldersFolder {
	s.FolderName = &v
	return s
}

type ResendCreateCloudAccountEmailRequest struct {
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
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

type ResendCreateCloudAccountEmailResponse struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *ResendCreateCloudAccountEmailResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s ResendCreateCloudAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponse) SetRequestId(v string) *ResendCreateCloudAccountEmailResponse {
	s.RequestId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) SetAccount(v *ResendCreateCloudAccountEmailResponseAccount) *ResendCreateCloudAccountEmailResponse {
	s.Account = v
	return s
}

type ResendCreateCloudAccountEmailResponseAccount struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s ResendCreateCloudAccountEmailResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponseAccount) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetResourceDirectoryId(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetAccountId(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.AccountId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetDisplayName(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetFolderId(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.FolderId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetJoinMethod(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetJoinTime(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetType(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.Type = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetStatus(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.Status = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetRecordId(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.RecordId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetModifyTime(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponseAccount) SetAccountName(v string) *ResendCreateCloudAccountEmailResponseAccount {
	s.AccountName = &v
	return s
}

type GetPayerForAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
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

type GetPayerForAccountResponse struct {
	PayerAccountName *string `json:"PayerAccountName,omitempty" xml:"PayerAccountName,omitempty" require:"true"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PayerAccountId   *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty" require:"true"`
}

func (s GetPayerForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPayerForAccountResponse) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountResponse) SetPayerAccountName(v string) *GetPayerForAccountResponse {
	s.PayerAccountName = &v
	return s
}

func (s *GetPayerForAccountResponse) SetRequestId(v string) *GetPayerForAccountResponse {
	s.RequestId = &v
	return s
}

func (s *GetPayerForAccountResponse) SetPayerAccountId(v string) *GetPayerForAccountResponse {
	s.PayerAccountId = &v
	return s
}

type ResendPromoteResourceAccountEmailRequest struct {
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
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

type ResendPromoteResourceAccountEmailResponse struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Account   *ResendPromoteResourceAccountEmailResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s ResendPromoteResourceAccountEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponse) SetRequestId(v string) *ResendPromoteResourceAccountEmailResponse {
	s.RequestId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) SetAccount(v *ResendPromoteResourceAccountEmailResponseAccount) *ResendPromoteResourceAccountEmailResponse {
	s.Account = v
	return s
}

type ResendPromoteResourceAccountEmailResponseAccount struct {
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty" require:"true"`
	AccountId           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	FolderId            *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	JoinMethod          *string `json:"JoinMethod,omitempty" xml:"JoinMethod,omitempty" require:"true"`
	JoinTime            *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty" require:"true"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RecordId            *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
	ModifyTime          *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
}

func (s ResendPromoteResourceAccountEmailResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponseAccount) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetResourceDirectoryId(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.ResourceDirectoryId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetAccountId(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.AccountId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetDisplayName(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.DisplayName = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetFolderId(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.FolderId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetJoinMethod(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.JoinMethod = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetJoinTime(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.JoinTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetType(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.Type = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetStatus(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.Status = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetRecordId(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.RecordId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetModifyTime(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.ModifyTime = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponseAccount) SetAccountName(v string) *ResendPromoteResourceAccountEmailResponseAccount {
	s.AccountName = &v
	return s
}

type ListFoldersForParentRequest struct {
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	QueryKeyword   *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFoldersForParentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentRequest) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentRequest) SetParentFolderId(v string) *ListFoldersForParentRequest {
	s.ParentFolderId = &v
	return s
}

func (s *ListFoldersForParentRequest) SetQueryKeyword(v string) *ListFoldersForParentRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageNumber(v int) *ListFoldersForParentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentRequest) SetPageSize(v int) *ListFoldersForParentRequest {
	s.PageSize = &v
	return s
}

type ListFoldersForParentResponse struct {
	TotalCount *int                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageSize   *int                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	Folders    *ListFoldersForParentResponseFolders `json:"Folders,omitempty" xml:"Folders,omitempty" require:"true" type:"Struct"`
}

func (s ListFoldersForParentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponse) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponse) SetTotalCount(v int) *ListFoldersForParentResponse {
	s.TotalCount = &v
	return s
}

func (s *ListFoldersForParentResponse) SetPageSize(v int) *ListFoldersForParentResponse {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentResponse) SetRequestId(v string) *ListFoldersForParentResponse {
	s.RequestId = &v
	return s
}

func (s *ListFoldersForParentResponse) SetPageNumber(v int) *ListFoldersForParentResponse {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentResponse) SetFolders(v *ListFoldersForParentResponseFolders) *ListFoldersForParentResponse {
	s.Folders = v
	return s
}

type ListFoldersForParentResponseFolders struct {
	Folder []*ListFoldersForParentResponseFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" require:"true" type:"Repeated"`
}

func (s ListFoldersForParentResponseFolders) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseFolders) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseFolders) SetFolder(v []*ListFoldersForParentResponseFoldersFolder) *ListFoldersForParentResponseFolders {
	s.Folder = v
	return s
}

type ListFoldersForParentResponseFoldersFolder struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty" require:"true"`
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty" require:"true"`
}

func (s ListFoldersForParentResponseFoldersFolder) String() string {
	return tea.Prettify(s)
}

func (s ListFoldersForParentResponseFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseFoldersFolder) SetCreateTime(v string) *ListFoldersForParentResponseFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListFoldersForParentResponseFoldersFolder) SetFolderId(v string) *ListFoldersForParentResponseFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListFoldersForParentResponseFoldersFolder) SetFolderName(v string) *ListFoldersForParentResponseFoldersFolder {
	s.FolderName = &v
	return s
}

type RemoveCloudAccountRequest struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
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

type RemoveCloudAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RemoveCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponse) SetRequestId(v string) *RemoveCloudAccountResponse {
	s.RequestId = &v
	return s
}

type CancelCreateCloudAccountRequest struct {
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty" require:"true"`
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

type CancelCreateCloudAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CancelCreateCloudAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCreateCloudAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountResponse) SetRequestId(v string) *CancelCreateCloudAccountResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) ListDelegatedServicesForAccountWithOptions(request *ListDelegatedServicesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.DoRequest(tea.String("ListDelegatedServicesForAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) RegisterDelegatedAdministratorWithOptions(request *RegisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.DoRequest(tea.String("RegisterDelegatedAdministrator"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListDelegatedAdministratorsWithOptions(request *ListDelegatedAdministratorsRequest, runtime *util.RuntimeOptions) (_result *ListDelegatedAdministratorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.DoRequest(tea.String("ListDelegatedAdministrators"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeregisterDelegatedAdministratorWithOptions(request *DeregisterDelegatedAdministratorRequest, runtime *util.RuntimeOptions) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.DoRequest(tea.String("DeregisterDelegatedAdministrator"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListControlPoliciesWithOptions(request *ListControlPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListControlPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.DoRequest(tea.String("ListControlPolicies"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.DoRequest(tea.String("ListControlPolicyAttachmentsForTarget"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateControlPolicyWithOptions(request *CreateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListTargetAttachmentsForControlPolicyWithOptions(request *ListTargetAttachmentsForControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("ListTargetAttachmentsForControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetControlPolicyWithOptions(request *GetControlPolicyRequest, runtime *util.RuntimeOptions) (_result *GetControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("GetControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) EnableControlPolicyWithOptions(request *EnableControlPolicyRequest, runtime *util.RuntimeOptions) (_result *EnableControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("EnableControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableControlPolicy(request *EnableControlPolicyRequest) (_result *EnableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.EnableControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetControlPolicyEnablementStatusWithOptions(request *GetControlPolicyEnablementStatusRequest, runtime *util.RuntimeOptions) (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetControlPolicyEnablementStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetControlPolicyEnablementStatus(request *GetControlPolicyEnablementStatusRequest) (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.GetControlPolicyEnablementStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableControlPolicyWithOptions(request *DisableControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DisableControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DisableControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableControlPolicy(request *DisableControlPolicyRequest) (_result *DisableControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.DisableControlPolicyWithOptions(request, runtime)
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
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DetachControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateControlPolicyWithOptions(request *UpdateControlPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AttachControlPolicyWithOptions(request *AttachControlPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("AttachControlPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetServiceLinkedRoleDeletionStatusWithOptions(request *GetServiceLinkedRoleDeletionStatusRequest, runtime *util.RuntimeOptions) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetServiceLinkedRoleDeletionStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetServiceLinkedRoleDeletionStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListTrustedServiceStatusWithOptions(request *ListTrustedServiceStatusRequest, runtime *util.RuntimeOptions) (_result *ListTrustedServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.DoRequest(tea.String("ListTrustedServiceStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteServiceLinkedRoleWithOptions(request *DeleteServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateCloudAccountWithOptions(request *CreateCloudAccountRequest, runtime *util.RuntimeOptions) (_result *CreateCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CreateCloudAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, runtime *util.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetRoleWithOptions(request *GetRoleRequest, runtime *util.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetRoleResponse{}
	_body, _err := client.DoRequest(tea.String("GetRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListRolesResponse{}
	_body, _err := client.DoRequest(tea.String("ListRoles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPolicyAttachmentsWithOptions(request *ListPolicyAttachmentsRequest, runtime *util.RuntimeOptions) (_result *ListPolicyAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPolicyAttachmentsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPolicyAttachments"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DetachPolicyWithOptions(request *DetachPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DetachPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DetachPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AttachPolicyWithOptions(request *AttachPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AttachPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("AttachPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetPolicyVersionWithOptions(request *GetPolicyVersionRequest, runtime *util.RuntimeOptions) (_result *GetPolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPolicyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("GetPolicyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SetDefaultPolicyVersionWithOptions(request *SetDefaultPolicyVersionRequest, runtime *util.RuntimeOptions) (_result *SetDefaultPolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetDefaultPolicyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("SetDefaultPolicyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteResourceGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("GetPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateResourceGroupWithOptions(request *UpdateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateResourceGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("ListResourceGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.DoRequest(tea.String("ListPolicies"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListPolicyVersionsWithOptions(request *ListPolicyVersionsRequest, runtime *util.RuntimeOptions) (_result *ListPolicyVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPolicyVersionsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPolicyVersions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateResourceAccountWithOptions(request *CreateResourceAccountRequest, runtime *util.RuntimeOptions) (_result *CreateResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CreateResourceAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListHandshakesForResourceDirectoryWithOptions(request *ListHandshakesForResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("ListHandshakesForResourceDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DestroyResourceDirectoryWithOptions(request *DestroyResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *DestroyResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("DestroyResourceDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DestroyResourceDirectory(request *DestroyResourceDirectoryRequest) (_result *DestroyResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.DestroyResourceDirectoryWithOptions(request, runtime)
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
	_result = &CreatePolicyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePolicyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeletePolicyVersionWithOptions(request *DeletePolicyVersionRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePolicyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePolicyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetResourceGroupWithOptions(request *GetResourceGroupRequest, runtime *util.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.DoRequest(tea.String("GetResourceGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) InitResourceDirectoryWithOptions(request *InitResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *InitResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("InitResourceDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitResourceDirectory(request *InitResourceDirectoryRequest) (_result *InitResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.InitResourceDirectoryWithOptions(request, runtime)
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
	_result = &GetHandshakeResponse{}
	_body, _err := client.DoRequest(tea.String("GetHandshake"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CancelHandshakeWithOptions(request *CancelHandshakeRequest, runtime *util.RuntimeOptions) (_result *CancelHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.DoRequest(tea.String("CancelHandshake"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeclineHandshakeWithOptions(request *DeclineHandshakeRequest, runtime *util.RuntimeOptions) (_result *DeclineHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.DoRequest(tea.String("DeclineHandshake"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListHandshakesForAccountWithOptions(request *ListHandshakesForAccountRequest, runtime *util.RuntimeOptions) (_result *ListHandshakesForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.DoRequest(tea.String("ListHandshakesForAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) InviteAccountToResourceDirectoryWithOptions(request *InviteAccountToResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("InviteAccountToResourceDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) AcceptHandshakeWithOptions(request *AcceptHandshakeRequest, runtime *util.RuntimeOptions) (_result *AcceptHandshakeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.DoRequest(tea.String("AcceptHandshake"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UpdateAccountWithOptions(request *UpdateAccountRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAccountResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *util.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFolderResponse{}
	_body, _err := client.DoRequest(tea.String("GetFolder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListAccountsForParentWithOptions(request *ListAccountsForParentRequest, runtime *util.RuntimeOptions) (_result *ListAccountsForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.DoRequest(tea.String("ListAccountsForParent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateResourceGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) PromoteResourceAccountWithOptions(request *PromoteResourceAccountRequest, runtime *util.RuntimeOptions) (_result *PromoteResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PromoteResourceAccountResponse{}
	_body, _err := client.DoRequest(tea.String("PromoteResourceAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *util.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.DoRequest(tea.String("ListAccounts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CancelPromoteResourceAccountWithOptions(request *CancelPromoteResourceAccountRequest, runtime *util.RuntimeOptions) (_result *CancelPromoteResourceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelPromoteResourceAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CancelPromoteResourceAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *util.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.DoRequest(tea.String("CreateFolder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *util.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteFolder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetAccountWithOptions(request *GetAccountRequest, runtime *util.RuntimeOptions) (_result *GetAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccountResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetResourceDirectoryWithOptions(request *GetResourceDirectoryRequest, runtime *util.RuntimeOptions) (_result *GetResourceDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("GetResourceDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceDirectory(request *GetResourceDirectoryRequest) (_result *GetResourceDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.GetResourceDirectoryWithOptions(request, runtime)
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
	_result = &UpdateFolderResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateFolder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) MoveAccountWithOptions(request *MoveAccountRequest, runtime *util.RuntimeOptions) (_result *MoveAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &MoveAccountResponse{}
	_body, _err := client.DoRequest(tea.String("MoveAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListAncestorsWithOptions(request *ListAncestorsRequest, runtime *util.RuntimeOptions) (_result *ListAncestorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAncestorsResponse{}
	_body, _err := client.DoRequest(tea.String("ListAncestors"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ResendCreateCloudAccountEmailWithOptions(request *ResendCreateCloudAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResendCreateCloudAccountEmailResponse{}
	_body, _err := client.DoRequest(tea.String("ResendCreateCloudAccountEmail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetPayerForAccountWithOptions(request *GetPayerForAccountRequest, runtime *util.RuntimeOptions) (_result *GetPayerForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.DoRequest(tea.String("GetPayerForAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ResendPromoteResourceAccountEmailWithOptions(request *ResendPromoteResourceAccountEmailRequest, runtime *util.RuntimeOptions) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResendPromoteResourceAccountEmailResponse{}
	_body, _err := client.DoRequest(tea.String("ResendPromoteResourceAccountEmail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListFoldersForParentWithOptions(request *ListFoldersForParentRequest, runtime *util.RuntimeOptions) (_result *ListFoldersForParentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.DoRequest(tea.String("ListFoldersForParent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) RemoveCloudAccountWithOptions(request *RemoveCloudAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.DoRequest(tea.String("RemoveCloudAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CancelCreateCloudAccountWithOptions(request *CancelCreateCloudAccountRequest, runtime *util.RuntimeOptions) (_result *CancelCreateCloudAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelCreateCloudAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CancelCreateCloudAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-31"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
