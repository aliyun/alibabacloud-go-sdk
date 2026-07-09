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
	// The maximum number of entries per page for a paged query. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that marks the position from which to start reading. Leave this parameter empty to read from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of policy IDs.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The policy name.
	//
	// example:
	//
	// Default policy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The policy type.
	//
	// example:
	//
	// Instance
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
