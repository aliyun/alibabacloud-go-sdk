// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayNetworkConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkMode(v string) *GatewayNetworkConfiguration
	GetNetworkMode() *string
	SetVpcId(v string) *GatewayNetworkConfiguration
	GetVpcId() *string
	SetVswitchIds(v []*string) *GatewayNetworkConfiguration
	GetVswitchIds() []*string
}

type GatewayNetworkConfiguration struct {
	NetworkMode *string   `json:"networkMode,omitempty" xml:"networkMode,omitempty"`
	VpcId       *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VswitchIds  []*string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty" type:"Repeated"`
}

func (s GatewayNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GatewayNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *GatewayNetworkConfiguration) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *GatewayNetworkConfiguration) GetVpcId() *string {
	return s.VpcId
}

func (s *GatewayNetworkConfiguration) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *GatewayNetworkConfiguration) SetNetworkMode(v string) *GatewayNetworkConfiguration {
	s.NetworkMode = &v
	return s
}

func (s *GatewayNetworkConfiguration) SetVpcId(v string) *GatewayNetworkConfiguration {
	s.VpcId = &v
	return s
}

func (s *GatewayNetworkConfiguration) SetVswitchIds(v []*string) *GatewayNetworkConfiguration {
	s.VswitchIds = v
	return s
}

func (s *GatewayNetworkConfiguration) Validate() error {
	return dara.Validate(s)
}
