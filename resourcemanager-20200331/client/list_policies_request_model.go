// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListPoliciesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *ListPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoliciesRequest
	GetPageSize() *int32
	SetPolicyType(v string) *ListPoliciesRequest
	GetPolicyType() *string
}

type ListPoliciesRequest struct {
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
	// The type of the access policy. If this parameter is left empty, all types of access policies are listed.
	//
	// - Custom: custom policy.
	//
	// - System: system policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
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

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
