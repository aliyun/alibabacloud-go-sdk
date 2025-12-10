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
	// The language in which you want to return the descriptions of the system permission policies. Valid values:
	//
	// 	- en: English.
	//
	// 	- zh-CN: Chinese
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
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
	// The type of the permission policy. If you do not configure this parameter, all types of permission policies are returned. Valid values:
	//
	// 	- Custom
	//
	// 	- System
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
