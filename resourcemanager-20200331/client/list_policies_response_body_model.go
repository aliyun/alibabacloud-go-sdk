// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoliciesResponseBody
	GetPageSize() *int32
	SetPolicies(v *ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody
	GetPolicies() *ListPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPoliciesResponseBody
	GetTotalCount() *int32
}

type ListPoliciesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the permission policies.
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoliciesResponseBody) GetPolicies() *ListPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesResponseBodyPolicies struct {
	Policy []*ListPoliciesResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicy() []*ListPoliciesResponseBodyPoliciesPolicy {
	return s.Policy
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicy(v []*ListPoliciesResponseBodyPoliciesPolicy) *ListPoliciesResponseBodyPolicies {
	s.Policy = v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) Validate() error {
	if s.Policy != nil {
		for _, item := range s.Policy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPoliciesResponseBodyPoliciesPolicy struct {
	// The number of references to the permission policy.
	//
	// example:
	//
	// 1
	AttachmentCount *int32 `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the permission policy was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The default version of the permission policy.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the permission policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the permission policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the permission policy was updated.
	//
	// example:
	//
	// 2016-02-11T18:39:12Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetAttachmentCount() *int32 {
	return s.AttachmentCount
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetUpdateDate() *string {
	return s.UpdateDate
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

func (s *ListPoliciesResponseBodyPoliciesPolicy) Validate() error {
	return dara.Validate(s)
}
