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
	SetTag(v []*ListControlPoliciesRequestTag) *ListControlPoliciesRequest
	GetTag() []*ListControlPoliciesRequestTag
}

type ListControlPoliciesRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// > This parameter is available only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number.
	//
	// Page starts from page 1. Default value: 1.
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
	// The type of the access control policies. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	Tag []*ListControlPoliciesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListControlPoliciesRequest) GetTag() []*ListControlPoliciesRequestTag {
	return s.Tag
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

func (s *ListControlPoliciesRequest) SetTag(v []*ListControlPoliciesRequestTag) *ListControlPoliciesRequest {
	s.Tag = v
	return s
}

func (s *ListControlPoliciesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListControlPoliciesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListControlPoliciesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesRequestTag) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListControlPoliciesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListControlPoliciesRequestTag) SetKey(v string) *ListControlPoliciesRequestTag {
	s.Key = &v
	return s
}

func (s *ListControlPoliciesRequestTag) SetValue(v string) *ListControlPoliciesRequestTag {
	s.Value = &v
	return s
}

func (s *ListControlPoliciesRequestTag) Validate() error {
	return dara.Validate(s)
}
