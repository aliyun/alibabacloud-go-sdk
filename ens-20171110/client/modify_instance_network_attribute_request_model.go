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
	// The ID of the ENS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5myukg7hnpbto7m024002****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new private IP address of the ECS instance. By default, if this parameter is empty, a private IP address is randomly assigned from the CIDR block of the specified vSwitch.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The vSwitch IDs.
	//
	// 	- If you set this parameter to the ID of the current vSwitch, the vSwitch of the ECS instance remains unchanged.
	//
	// 	- The input ID is a new vSwitch, and the new and old vSwitches must belong to the same node.
	//
	// 	- By default, if this parameter is not specified, a private IP address is randomly assigned from within the CIDR block of the specified vSwitch.
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
