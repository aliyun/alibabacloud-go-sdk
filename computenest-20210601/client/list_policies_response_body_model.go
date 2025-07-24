// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPoliciesResponseBody
	GetNextToken() *string
	SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody
	GetPolicies() []*ListPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
}

type ListPoliciesResponseBody struct {
	// 分页大小。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	//
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Permission policy list
	Policies []*ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesResponseBody) GetPolicies() []*ListPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) SetMaxResults(v int32) *ListPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPoliciesResponseBodyPolicies struct {
	// Permission policy description.
	//
	// example:
	//
	// Only read permission policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Policy content.
	//
	// example:
	//
	// {"Action":["*:Describe*","*:List*","*:Get*","*:BatchGet*","*:Query*","*:BatchQuery*","actiontrail:LookupEvents"],"Resource":["*"],"Effect":"Allow"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Permission policy name.
	//
	// example:
	//
	// AliyunComputeNestPolicyForReadOnly
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Permission policy type.
	//
	// - Custom: Custom policy.
	//
	// - System: System policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesResponseBodyPolicies) SetDescription(v string) *ListPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyDocument(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyDocument = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyType(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
