// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *CancelAutoSnapshotPolicyRequest
	GetDesktopId() []*string
	SetPolicyId(v string) *CancelAutoSnapshotPolicyRequest
	GetPolicyId() *string
	SetRegionId(v string) *CancelAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type CancelAutoSnapshotPolicyRequest struct {
	// The IDs of the cloud computers. You can specify 1 to 50 IDs. The IDs cannot be an empty string. The IDs can be up to 64 characters in length and cannot contain `http://` or `https://`. The IDs cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-78lhzpe7kjfnd****
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

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *CancelAutoSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CancelAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelAutoSnapshotPolicyRequest) SetDesktopId(v []*string) *CancelAutoSnapshotPolicyRequest {
	s.DesktopId = v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetPolicyId(v string) *CancelAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetRegionId(v string) *CancelAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
