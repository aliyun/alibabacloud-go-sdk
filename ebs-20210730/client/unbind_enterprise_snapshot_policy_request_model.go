// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnbindEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetDiskTargets(v []*string) *UnbindEnterpriseSnapshotPolicyRequest
	GetDiskTargets() []*string
	SetPolicyId(v string) *UnbindEnterpriseSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *UnbindEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
}

type UnbindEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of disks.
	DiskTargets []*string `json:"DiskTargets,omitempty" xml:"DiskTargets,omitempty" type:"Repeated"`
	// The id of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// esp-xxs
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

func (s UnbindEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) GetDiskTargets() []*string {
	return s.DiskTargets
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetDiskTargets(v []*string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.DiskTargets = v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *UnbindEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
