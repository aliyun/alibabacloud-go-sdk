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
	// A list of permission policies.
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// b66ad557-9c00-4064-9c8d-b621c3263308
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type ListPoliciesResponseBodyPoliciesPolicy struct {
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetName() *string {
	return s.Name
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetName(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Name = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) Validate() error {
	return dara.Validate(s)
}
