// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListPolicyAttachmentsRequest
	GetLanguage() *string
	SetPageNumber(v int32) *ListPolicyAttachmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPolicyAttachmentsRequest
	GetPageSize() *int32
	SetPolicyName(v string) *ListPolicyAttachmentsRequest
	GetPolicyName() *string
	SetPolicyType(v string) *ListPolicyAttachmentsRequest
	GetPolicyType() *string
	SetPrincipalName(v string) *ListPolicyAttachmentsRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *ListPolicyAttachmentsRequest
	GetPrincipalType() *string
	SetResourceGroupId(v string) *ListPolicyAttachmentsRequest
	GetResourceGroupId() *string
}

type ListPolicyAttachmentsRequest struct {
	// The language type. The description of the system access policy is returned in the specified language.
	//
	// - en: English.
	//
	// - zh-CN: Chinese.
	//
	// - ja: Japanese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Policy Name of the access policy.
	//
	// Policy Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// AdministratorAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access policy. If this parameter is left empty, all types of access policies are listed.
	//
	// - Custom: custom policy.
	//
	// - System: system policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the principal.
	//
	// example:
	//
	// alice@demo.onaliyun.com
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The authorization object type. If this parameter is left empty, all types of authorization objects are listed.
	//
	// - IMSUser: Resource Access Management (RAM) user.
	//
	// - IMSGroup: RAM user group.
	//
	// - ServiceRole: RAM role.
	//
	// example:
	//
	// IMSUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the resource group or the Alibaba Cloud account ID to which the resource group belongs. If this parameter is left empty, all authorization records under the current account are listed.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListPolicyAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListPolicyAttachmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyAttachmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyAttachmentsRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyAttachmentsRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyAttachmentsRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListPolicyAttachmentsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPolicyAttachmentsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *ListPolicyAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
