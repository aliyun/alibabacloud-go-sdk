// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v []*string) *DeleteAutoSnapshotPolicyRequest
	GetPolicyId() []*string
	SetRegionId(v string) *DeleteAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type DeleteAutoSnapshotPolicyRequest struct {
	// The IDs of the automatic snapshot policies that you want to delete.
	//
	// This parameter is required.
	PolicyId []*string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) GetPolicyId() []*string {
	return s.PolicyId
}

func (s *DeleteAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAutoSnapshotPolicyRequest) SetPolicyId(v []*string) *DeleteAutoSnapshotPolicyRequest {
	s.PolicyId = v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetRegionId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
