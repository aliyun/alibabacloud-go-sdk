// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalPolicyGroupIds(v []*string) *DescribePolicyGroupsRequest
	GetExternalPolicyGroupIds() []*string
	SetMaxResults(v int32) *DescribePolicyGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePolicyGroupsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribePolicyGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolicyGroupsRequest
	GetPageSize() *int32
	SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest
	GetPolicyGroupId() []*string
	SetRegionId(v string) *DescribePolicyGroupsRequest
	GetRegionId() *string
	SetScope(v string) *DescribePolicyGroupsRequest
	GetScope() *string
}

type DescribePolicyGroupsRequest struct {
	// The array of cloud computer policy IDs to be excluded.
	ExternalPolicyGroupIds []*string `json:"ExternalPolicyGroupIds,omitempty" xml:"ExternalPolicyGroupIds,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 100
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the cloud computer policies.
	//
	// example:
	//
	// system-all-enabled-policy
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective scope of the cloud computer policy.
	//
	// Valid values:
	//
	// 	- ALL
	//
	// 	- IP
	//
	// 	- GLOBAL
	//
	// example:
	//
	// ALL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribePolicyGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsRequest) GetExternalPolicyGroupIds() []*string {
	return s.ExternalPolicyGroupIds
}

func (s *DescribePolicyGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePolicyGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePolicyGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolicyGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolicyGroupsRequest) GetPolicyGroupId() []*string {
	return s.PolicyGroupId
}

func (s *DescribePolicyGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolicyGroupsRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribePolicyGroupsRequest) SetExternalPolicyGroupIds(v []*string) *DescribePolicyGroupsRequest {
	s.ExternalPolicyGroupIds = v
	return s
}

func (s *DescribePolicyGroupsRequest) SetMaxResults(v int32) *DescribePolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetNextToken(v string) *DescribePolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPageNumber(v int32) *DescribePolicyGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPageSize(v int32) *DescribePolicyGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

func (s *DescribePolicyGroupsRequest) SetRegionId(v string) *DescribePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetScope(v string) *DescribePolicyGroupsRequest {
	s.Scope = &v
	return s
}

func (s *DescribePolicyGroupsRequest) Validate() error {
	return dara.Validate(s)
}
