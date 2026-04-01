// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressWithRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *UnassociateEipAddressWithRCInstanceRequest
	GetAllocationId() *string
	SetInstanceId(v string) *UnassociateEipAddressWithRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *UnassociateEipAddressWithRCInstanceRequest
	GetRegionId() *string
}

type UnassociateEipAddressWithRCInstanceRequest struct {
	// The EIP ID.
	//
	// example:
	//
	// eip-bp166out2x4bpcf******
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-i322y2t562oh7o******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UnassociateEipAddressWithRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressWithRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressWithRCInstanceRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *UnassociateEipAddressWithRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateEipAddressWithRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateEipAddressWithRCInstanceRequest) SetAllocationId(v string) *UnassociateEipAddressWithRCInstanceRequest {
	s.AllocationId = &v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceRequest) SetInstanceId(v string) *UnassociateEipAddressWithRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceRequest) SetRegionId(v string) *UnassociateEipAddressWithRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
