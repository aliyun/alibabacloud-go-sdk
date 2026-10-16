// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListControlPoliciesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *ListControlPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListControlPoliciesRequest
	GetPageSize() *int32
	SetPolicyType(v string) *ListControlPoliciesRequest
	GetPolicyType() *string
}

type ListControlPoliciesRequest struct {
	// The language in which you want the description of the control policy to be returned. Valid values:
	//
	// - zh-CN (default): Chinese.
	//
	// - en: English.
	//
	// - ja: Japanese.
	//
	// > This parameter is valid only for system control policies.
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
	// The type of the control policy. Valid values:
	//
	// - System: system control policy.
	//
	// - Custom: custom control policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListControlPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListControlPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListControlPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListControlPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
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

func (s *ListControlPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
