// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecurityPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSecurityPoliciesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListSecurityPoliciesRequest
	GetResourceGroupId() *string
	SetSecurityPolicyIds(v []*string) *ListSecurityPoliciesRequest
	GetSecurityPolicyIds() []*string
	SetSecurityPolicyNames(v []*string) *ListSecurityPoliciesRequest
	GetSecurityPolicyNames() []*string
	SetTag(v []*ListSecurityPoliciesRequestTag) *ListSecurityPoliciesRequest
	GetTag() []*ListSecurityPoliciesRequestTag
}

type ListSecurityPoliciesRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security policy IDs. You can specify at most 20 security policies.
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
	// The names of the security policies. You can specify up to 10 names.
	SecurityPolicyNames []*string `json:"SecurityPolicyNames,omitempty" xml:"SecurityPolicyNames,omitempty" type:"Repeated"`
	// The tags.
	Tag []*ListSecurityPoliciesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecurityPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecurityPoliciesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSecurityPoliciesRequest) GetSecurityPolicyIds() []*string {
	return s.SecurityPolicyIds
}

func (s *ListSecurityPoliciesRequest) GetSecurityPolicyNames() []*string {
	return s.SecurityPolicyNames
}

func (s *ListSecurityPoliciesRequest) GetTag() []*ListSecurityPoliciesRequestTag {
	return s.Tag
}

func (s *ListSecurityPoliciesRequest) SetMaxResults(v int32) *ListSecurityPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetNextToken(v string) *ListSecurityPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetResourceGroupId(v string) *ListSecurityPoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyIds = v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyNames(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyNames = v
	return s
}

func (s *ListSecurityPoliciesRequest) SetTag(v []*ListSecurityPoliciesRequestTag) *ListSecurityPoliciesRequest {
	s.Tag = v
	return s
}

func (s *ListSecurityPoliciesRequest) Validate() error {
	return dara.Validate(s)
}

type ListSecurityPoliciesRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecurityPoliciesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesRequestTag) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListSecurityPoliciesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListSecurityPoliciesRequestTag) SetKey(v string) *ListSecurityPoliciesRequestTag {
	s.Key = &v
	return s
}

func (s *ListSecurityPoliciesRequestTag) SetValue(v string) *ListSecurityPoliciesRequestTag {
	s.Value = &v
	return s
}

func (s *ListSecurityPoliciesRequestTag) Validate() error {
	return dara.Validate(s)
}
