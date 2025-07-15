// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPolicyGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPolicyGroupsRequest
	GetNextToken() *string
	SetPolicyGroupIds(v []*string) *ListPolicyGroupsRequest
	GetPolicyGroupIds() []*string
	SetPolicyGroupName(v string) *ListPolicyGroupsRequest
	GetPolicyGroupName() *string
	SetPolicyType(v string) *ListPolicyGroupsRequest
	GetPolicyType() *string
}

type ListPolicyGroupsRequest struct {
	// The maximum number of entries per page. Value range: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the policies.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicyGroup
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicyGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPolicyGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicyGroupsRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *ListPolicyGroupsRequest) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *ListPolicyGroupsRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyGroupsRequest) SetMaxResults(v int32) *ListPolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPolicyGroupsRequest) SetNextToken(v string) *ListPolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPolicyGroupsRequest) SetPolicyGroupIds(v []*string) *ListPolicyGroupsRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *ListPolicyGroupsRequest) SetPolicyGroupName(v string) *ListPolicyGroupsRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *ListPolicyGroupsRequest) SetPolicyType(v string) *ListPolicyGroupsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyGroupsRequest) Validate() error {
	return dara.Validate(s)
}
