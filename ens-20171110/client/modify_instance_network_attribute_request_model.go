// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceNetworkAttributeRequest
	GetInstanceId() *string
	SetPrivateIpAddress(v string) *ModifyInstanceNetworkAttributeRequest
	GetPrivateIpAddress() *string
	SetVSwitchId(v string) *ModifyInstanceNetworkAttributeRequest
	GetVSwitchId() *string
}

type ModifyInstanceNetworkAttributeRequest struct {
	// The ENS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5myukg7hnpbto7m024002****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new private IP address.
	//
	// Default value: If this parameter is not specified, an IP address is randomly assigned from the CIDR block of the vSwitch.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The vSwitch ID.
	//
	// - If the specified ID is the current vSwitch of the instance, the vSwitch of the instance remains unchanged.
	//
	// - If the specified ID is a new vSwitch, the new and old vSwitches must belong to the same edge zone.
	//
	// - If this parameter is not specified, a value is randomly assigned from the CIDR block of the original vSwitch.
	//
	// example:
	//
	// vsw-5rllcjb3ol6duzjdnbm1o****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyInstanceNetworkAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetworkAttributeRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ModifyInstanceNetworkAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyInstanceNetworkAttributeRequest) SetInstanceId(v string) *ModifyInstanceNetworkAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetworkAttributeRequest) SetPrivateIpAddress(v string) *ModifyInstanceNetworkAttributeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ModifyInstanceNetworkAttributeRequest) SetVSwitchId(v string) *ModifyInstanceNetworkAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyInstanceNetworkAttributeRequest) Validate() error {
	return dara.Validate(s)
}
