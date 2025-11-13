// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckTypes(v []*string) *ListCheckPoliciesRequest
	GetCheckTypes() []*string
	SetCurrentPage(v int32) *ListCheckPoliciesRequest
	GetCurrentPage() *int32
	SetDependentPolicyId(v int64) *ListCheckPoliciesRequest
	GetDependentPolicyId() *int64
	SetLang(v string) *ListCheckPoliciesRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckPoliciesRequest
	GetPageSize() *int32
	SetPolicyId(v int64) *ListCheckPoliciesRequest
	GetPolicyId() *int64
	SetPolicyShowName(v string) *ListCheckPoliciesRequest
	GetPolicyShowName() *string
	SetPolicyType(v string) *ListCheckPoliciesRequest
	GetPolicyType() *string
	SetType(v string) *ListCheckPoliciesRequest
	GetType() *string
}

type ListCheckPoliciesRequest struct {
	// The types of policies to be queried (default queries both custom and system predefined policies).
	CheckTypes []*string `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty" type:"Repeated"`
	// Specifies the page number from which to start displaying the query results. The starting value is **1**. The default value is **1**, indicating that the display starts from the **1st*	- page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// ID of the associated parent policy.
	//
	// (The specific dependency relationship from low to high is: Section -> Requirement -> Standard)
	//
	// example:
	//
	// 1000000000002
	DependentPolicyId *int64 `json:"DependentPolicyId,omitempty" xml:"DependentPolicyId,omitempty"`
	// Language type for request and response messages, with a default value of **zh**. Possible values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Number of check item information entries displayed per page during pagination. The default value is **50**, indicating 50 entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of the newly added classification setting.
	//
	// example:
	//
	// 1000000000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Name of the newly added classification setting.
	//
	// example:
	//
	// testPolicyName
	PolicyShowName *string `json:"PolicyShowName,omitempty" xml:"PolicyShowName,omitempty"`
	// Policy type of the custom check item rule:
	//
	// - **STANDARD**: New standard
	//
	// - **REQUIREMENT**: New requirement
	//
	// - **SECTION**: New section
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// Name of the associated major policy category (required when PolicyType is STANDARD):
	//
	// - **AISPM**: AI Configuration Management (AI-SPM)
	//
	// - **IDENTITY_PERMISSION**: Identity and Permission Management (CIEM)
	//
	// - **RISK**: Security Risk
	//
	// - **COMPLIANCE**: Compliance Risk
	//
	// example:
	//
	// AISPM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCheckPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCheckPoliciesRequest) GetCheckTypes() []*string {
	return s.CheckTypes
}

func (s *ListCheckPoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckPoliciesRequest) GetDependentPolicyId() *int64 {
	return s.DependentPolicyId
}

func (s *ListCheckPoliciesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckPoliciesRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ListCheckPoliciesRequest) GetPolicyShowName() *string {
	return s.PolicyShowName
}

func (s *ListCheckPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListCheckPoliciesRequest) GetType() *string {
	return s.Type
}

func (s *ListCheckPoliciesRequest) SetCheckTypes(v []*string) *ListCheckPoliciesRequest {
	s.CheckTypes = v
	return s
}

func (s *ListCheckPoliciesRequest) SetCurrentPage(v int32) *ListCheckPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetDependentPolicyId(v int64) *ListCheckPoliciesRequest {
	s.DependentPolicyId = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetLang(v string) *ListCheckPoliciesRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetPageSize(v int32) *ListCheckPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetPolicyId(v int64) *ListCheckPoliciesRequest {
	s.PolicyId = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetPolicyShowName(v string) *ListCheckPoliciesRequest {
	s.PolicyShowName = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetPolicyType(v string) *ListCheckPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListCheckPoliciesRequest) SetType(v string) *ListCheckPoliciesRequest {
	s.Type = &v
	return s
}

func (s *ListCheckPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
