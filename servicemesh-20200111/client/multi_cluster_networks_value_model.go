// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiClusterNetworksValue interface {
	dara.Model
	String() string
	GoString() string
	SetNetwork(v string) *MultiClusterNetworksValue
	GetNetwork() *string
	SetEnableGateway(v bool) *MultiClusterNetworksValue
	GetEnableGateway() *bool
	SetCustomGatewayAddress(v string) *MultiClusterNetworksValue
	GetCustomGatewayAddress() *string
	SetGatewayName(v string) *MultiClusterNetworksValue
	GetGatewayName() *string
}

type MultiClusterNetworksValue struct {
	// example:
	//
	// network1
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// true
	EnableGateway *bool `json:"EnableGateway,omitempty" xml:"EnableGateway,omitempty"`
	// example:
	//
	// 8.16x.1x.1x:15443
	CustomGatewayAddress *string `json:"CustomGatewayAddress,omitempty" xml:"CustomGatewayAddress,omitempty"`
	// example:
	//
	// asm-cross-network-ccb37ff104***
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
}

func (s MultiClusterNetworksValue) String() string {
	return dara.Prettify(s)
}

func (s MultiClusterNetworksValue) GoString() string {
	return s.String()
}

func (s *MultiClusterNetworksValue) GetNetwork() *string {
	return s.Network
}

func (s *MultiClusterNetworksValue) GetEnableGateway() *bool {
	return s.EnableGateway
}

func (s *MultiClusterNetworksValue) GetCustomGatewayAddress() *string {
	return s.CustomGatewayAddress
}

func (s *MultiClusterNetworksValue) GetGatewayName() *string {
	return s.GatewayName
}

func (s *MultiClusterNetworksValue) SetNetwork(v string) *MultiClusterNetworksValue {
	s.Network = &v
	return s
}

func (s *MultiClusterNetworksValue) SetEnableGateway(v bool) *MultiClusterNetworksValue {
	s.EnableGateway = &v
	return s
}

func (s *MultiClusterNetworksValue) SetCustomGatewayAddress(v string) *MultiClusterNetworksValue {
	s.CustomGatewayAddress = &v
	return s
}

func (s *MultiClusterNetworksValue) SetGatewayName(v string) *MultiClusterNetworksValue {
	s.GatewayName = &v
	return s
}

func (s *MultiClusterNetworksValue) Validate() error {
	return dara.Validate(s)
}
