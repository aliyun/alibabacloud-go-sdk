// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEnterpriseSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *BindEnterpriseSnapshotPolicyRequest
	GetClientToken() *string
	SetDiskTargets(v []*string) *BindEnterpriseSnapshotPolicyRequest
	GetDiskTargets() []*string
	SetPolicyId(v string) *BindEnterpriseSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *BindEnterpriseSnapshotPolicyRequest
	GetRegionId() *string
}

type BindEnterpriseSnapshotPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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

func (s BindEnterpriseSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BindEnterpriseSnapshotPolicyRequest) GetDiskTargets() []*string {
	return s.DiskTargets
}

func (s *BindEnterpriseSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *BindEnterpriseSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetClientToken(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetDiskTargets(v []*string) *BindEnterpriseSnapshotPolicyRequest {
	s.DiskTargets = v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetPolicyId(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) SetRegionId(v string) *BindEnterpriseSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
