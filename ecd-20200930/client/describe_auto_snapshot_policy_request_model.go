// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeAutoSnapshotPolicyRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAutoSnapshotPolicyRequest
	GetNextToken() *string
	SetPolicyId(v string) *DescribeAutoSnapshotPolicyRequest
	GetPolicyId() *string
	SetPolicyName(v string) *DescribeAutoSnapshotPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *DescribeAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type DescribeAutoSnapshotPolicyRequest struct {
	// The maximum number of entries per page for a paged query.
	//
	// - Maximum value: 50
	//
	// - Default value: 50
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mqIGdPhID94rjhZFGsvpJo
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-c5tv9d64ebjnj****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the automatic snapshot policy.
	//
	// example:
	//
	// Test1201
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAutoSnapshotPolicyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeAutoSnapshotPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyRequest) SetMaxResults(v int32) *DescribeAutoSnapshotPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyRequest) SetNextToken(v string) *DescribeAutoSnapshotPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyRequest) SetPolicyId(v string) *DescribeAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyRequest) SetPolicyName(v string) *DescribeAutoSnapshotPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyRequest) SetRegionId(v string) *DescribeAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
