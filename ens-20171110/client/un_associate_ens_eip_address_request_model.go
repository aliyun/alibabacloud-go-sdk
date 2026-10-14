// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateEnsEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *UnAssociateEnsEipAddressRequest
	GetAllocationId() *string
	SetForce(v bool) *UnAssociateEnsEipAddressRequest
	GetForce() *bool
}

type UnAssociateEnsEipAddressRequest struct {
	// The instance ID of the edge elastic IP address (EIP).
	//
	// If the EIP is referenced by an SNAT entry, invoke DeleteSnatIpForSnatEntry to remove it first. If the EIP is referenced by a DNAT entry, invoke DeleteForwardEntry to delete it first, or set Force to true under Settings as needed.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-5sqa431nx3vee8heqxfxp****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// Specifies whether to forcibly disassociate the EIP when the EIP is bound to a NAT gateway that has DNAT or SNAT entries. Valid values:
	//
	// - **false*	- (default): does not forcibly disassociate the EIP.
	//
	// - **true**: forcibly disassociates the EIP.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s UnAssociateEnsEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateEnsEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *UnAssociateEnsEipAddressRequest) GetForce() *bool {
	return s.Force
}

func (s *UnAssociateEnsEipAddressRequest) SetAllocationId(v string) *UnAssociateEnsEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *UnAssociateEnsEipAddressRequest) SetForce(v bool) *UnAssociateEnsEipAddressRequest {
	s.Force = &v
	return s
}

func (s *UnAssociateEnsEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
