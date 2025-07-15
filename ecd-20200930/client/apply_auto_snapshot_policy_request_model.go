// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *ApplyAutoSnapshotPolicyRequest
	GetDesktopId() []*string
	SetPolicyId(v string) *ApplyAutoSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *ApplyAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type ApplyAutoSnapshotPolicyRequest struct {
	// The IDs of the cloud computers. You can specify 1 to 20 IDs.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the automatic snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-ejtum8j5tfcw7****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyAutoSnapshotPolicyRequest) SetDesktopId(v []*string) *ApplyAutoSnapshotPolicyRequest {
	s.DesktopId = v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetRegionId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
