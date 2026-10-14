// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEnsEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AssociateEnsEipAddressRequest
	GetAllocationId() *string
	SetInstanceId(v string) *AssociateEnsEipAddressRequest
	GetInstanceId() *string
	SetInstanceType(v string) *AssociateEnsEipAddressRequest
	GetInstanceType() *string
	SetStandby(v bool) *AssociateEnsEipAddressRequest
	GetStandby() *bool
}

type AssociateEnsEipAddressRequest struct {
	// The ID of the EIP to be associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-5sc1sgcrsrwgwdvx44hru****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the cloud product instance to attach the EIP to.
	//
	// >Load balancing (ELB) instance IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5saivuir6b1mupxjfbhmk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the EIP instance. Valid values:
	//
	// - **Nat**: NAT gateway.
	//
	// - **SlbInstance**: load balancing (ELB).
	//
	// - **NetworkInterface**: secondary network interface controller (NIC).
	//
	// - **NatSlbInstance**: select this value when attaching multiple EIPs to an ELB.
	//
	// - **EnsInstance*	- (default): ENS instance.
	//
	// example:
	//
	// SlbInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the EIP is a standby EIP. Valid values:
	//
	// - true: standby.
	//
	// - false: not standby.
	//
	// example:
	//
	// true
	Standby *bool `json:"Standby,omitempty" xml:"Standby,omitempty"`
}

func (s AssociateEnsEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateEnsEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AssociateEnsEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateEnsEipAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AssociateEnsEipAddressRequest) GetStandby() *bool {
	return s.Standby
}

func (s *AssociateEnsEipAddressRequest) SetAllocationId(v string) *AssociateEnsEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) SetInstanceId(v string) *AssociateEnsEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) SetInstanceType(v string) *AssociateEnsEipAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) SetStandby(v bool) *AssociateEnsEipAddressRequest {
	s.Standby = &v
	return s
}

func (s *AssociateEnsEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
