// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest
	GetDescription() *string
	SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceName() *string
}

type ModifyNetworkInterfaceAttributeRequest struct {
	// The description. The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-5f6533jbifugr5fo***
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI. The name must be 1 to 128 characters in length, The name cannot start with http:// or https://.
	//
	// example:
	//
	// test-01
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
