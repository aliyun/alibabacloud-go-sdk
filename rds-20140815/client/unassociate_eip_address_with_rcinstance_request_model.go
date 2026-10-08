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
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
