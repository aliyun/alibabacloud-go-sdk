// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNetworkInterfaceToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStart(v bool) *AddNetworkInterfaceToInstanceRequest
	GetAutoStart() *bool
	SetInstanceId(v string) *AddNetworkInterfaceToInstanceRequest
	GetInstanceId() *string
	SetNetworks(v string) *AddNetworkInterfaceToInstanceRequest
	GetNetworks() *string
}

type AddNetworkInterfaceToInstanceRequest struct {
	// Specifies whether to automatically restart the instance.
	//
	// example:
	//
	// true
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourInstance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network. The value is a JSON string. Only IPv6 is supported. Sample code of an IPv6 network: [{ "ipType": "public", "ipAddressType": "ipv6" }]
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ipType": "public", "ipAddressType": "ipv6" }]
	Networks *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
}

func (s AddNetworkInterfaceToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *AddNetworkInterfaceToInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddNetworkInterfaceToInstanceRequest) GetNetworks() *string {
	return s.Networks
}

func (s *AddNetworkInterfaceToInstanceRequest) SetAutoStart(v bool) *AddNetworkInterfaceToInstanceRequest {
	s.AutoStart = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceRequest) SetInstanceId(v string) *AddNetworkInterfaceToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceRequest) SetNetworks(v string) *AddNetworkInterfaceToInstanceRequest {
	s.Networks = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
