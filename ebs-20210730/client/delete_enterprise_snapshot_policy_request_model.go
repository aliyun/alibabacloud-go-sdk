// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetPolicyId(v string) *DeleteEnterpriseSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *DeleteEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
}

type DeleteEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The id of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// esp-xxx
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID . You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which snapshot policy is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *DeleteEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
