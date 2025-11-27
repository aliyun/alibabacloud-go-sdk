// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressWithRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AssociateEipAddressWithRCInstanceRequest
	GetAllocationId() *string
	SetInstanceId(v string) *AssociateEipAddressWithRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *AssociateEipAddressWithRCInstanceRequest
	GetRegionId() *string
}

type AssociateEipAddressWithRCInstanceRequest struct {
	// The EIP ID.
	//
	// >  If no EIP is available, create an EIP. For more information, see [Create an EIP](https://help.aliyun.com/document_detail/292841.html).
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

func (s AssociateEipAddressWithRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressWithRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressWithRCInstanceRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AssociateEipAddressWithRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateEipAddressWithRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateEipAddressWithRCInstanceRequest) SetAllocationId(v string) *AssociateEipAddressWithRCInstanceRequest {
	s.AllocationId = &v
	return s
}

func (s *AssociateEipAddressWithRCInstanceRequest) SetInstanceId(v string) *AssociateEipAddressWithRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEipAddressWithRCInstanceRequest) SetRegionId(v string) *AssociateEipAddressWithRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateEipAddressWithRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
