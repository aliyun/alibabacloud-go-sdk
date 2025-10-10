// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecurityPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSecurityPoliciesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSecurityPoliciesResponseBody
	GetRequestId() *string
	SetSecurityPolicies(v []*ListSecurityPoliciesResponseBodySecurityPolicies) *ListSecurityPoliciesResponseBody
	GetSecurityPolicies() []*ListSecurityPoliciesResponseBodySecurityPolicies
	SetTotalCount(v int32) *ListSecurityPoliciesResponseBody
	GetTotalCount() *int32
}

type ListSecurityPoliciesResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported security policies.
	SecurityPolicies []*ListSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecurityPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecurityPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityPoliciesResponseBody) GetSecurityPolicies() []*ListSecurityPoliciesResponseBodySecurityPolicies {
	return s.SecurityPolicies
}

func (s *ListSecurityPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecurityPoliciesResponseBody) SetMaxResults(v int32) *ListSecurityPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetNextToken(v string) *ListSecurityPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetRequestId(v string) *ListSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSecurityPoliciesResponseBodySecurityPolicies) *ListSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetTotalCount(v int32) *ListSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecurityPoliciesResponseBodySecurityPolicies struct {
	// The supported cipher suites.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The time when the ACL was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2023-02-15T07:37:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security policy.
	//
	// example:
	//
	// rg-atstuj3rtop****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the security policy.
	//
	// example:
	//
	// test-secrity
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The status of the security policy. Valid values:
	//
	// 	- **Configuring**
	//
	// 	- **Available**
	//
	// example:
	//
	// Available
	SecurityPolicyStatus *string `json:"SecurityPolicyStatus,omitempty" xml:"SecurityPolicyStatus,omitempty"`
	// The supported TLS protocol versions.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
	// The tags.
	Tags []*ListSecurityPoliciesResponseBodySecurityPoliciesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetCiphers() []*string {
	return s.Ciphers
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetSecurityPolicyName() *string {
	return s.SecurityPolicyName
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetSecurityPolicyStatus() *string {
	return s.SecurityPolicyStatus
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) GetTags() []*ListSecurityPoliciesResponseBodySecurityPoliciesTags {
	return s.Tags
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetCreateTime(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetResourceGroupId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyName(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyStatus(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyStatus = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetTags(v []*ListSecurityPoliciesResponseBodySecurityPoliciesTags) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.Tags = v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) Validate() error {
	return dara.Validate(s)
}

type ListSecurityPoliciesResponseBodySecurityPoliciesTags struct {
	// The tag key. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecurityPoliciesResponseBodySecurityPoliciesTags) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesResponseBodySecurityPoliciesTags) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBodySecurityPoliciesTags) GetKey() *string {
	return s.Key
}

func (s *ListSecurityPoliciesResponseBodySecurityPoliciesTags) GetValue() *string {
	return s.Value
}

func (s *ListSecurityPoliciesResponseBodySecurityPoliciesTags) SetKey(v string) *ListSecurityPoliciesResponseBodySecurityPoliciesTags {
	s.Key = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPoliciesTags) SetValue(v string) *ListSecurityPoliciesResponseBodySecurityPoliciesTags {
	s.Value = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPoliciesTags) Validate() error {
	return dara.Validate(s)
}
