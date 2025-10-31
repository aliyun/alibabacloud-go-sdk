// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLbReplica(v int32) *NetworkConfig
	GetLbReplica() *int32
	SetLoadBalanceConfig(v []*NetworkConfigLoadBalanceConfig) *NetworkConfig
	GetLoadBalanceConfig() []*NetworkConfigLoadBalanceConfig
	SetLoadBalanceType(v string) *NetworkConfig
	GetLoadBalanceType() *string
	SetType(v string) *NetworkConfig
	GetType() *string
	SetVpcId(v string) *NetworkConfig
	GetVpcId() *string
	SetVsArea(v string) *NetworkConfig
	GetVsArea() *string
	SetVswitchId(v string) *NetworkConfig
	GetVswitchId() *string
	SetWhiteIpGroupList(v []*WhiteIpGroup) *NetworkConfig
	GetWhiteIpGroupList() []*WhiteIpGroup
}

type NetworkConfig struct {
	LbReplica         *int32                            `json:"lbReplica,omitempty" xml:"lbReplica,omitempty"`
	LoadBalanceConfig []*NetworkConfigLoadBalanceConfig `json:"loadBalanceConfig,omitempty" xml:"loadBalanceConfig,omitempty" type:"Repeated"`
	// example:
	//
	// DEFAULT
	LoadBalanceType  *string         `json:"loadBalanceType,omitempty" xml:"loadBalanceType,omitempty"`
	Type             *string         `json:"type,omitempty" xml:"type,omitempty"`
	VpcId            *string         `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea           *string         `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId        *string         `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	WhiteIpGroupList []*WhiteIpGroup `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s NetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s NetworkConfig) GoString() string {
	return s.String()
}

func (s *NetworkConfig) GetLbReplica() *int32 {
	return s.LbReplica
}

func (s *NetworkConfig) GetLoadBalanceConfig() []*NetworkConfigLoadBalanceConfig {
	return s.LoadBalanceConfig
}

func (s *NetworkConfig) GetLoadBalanceType() *string {
	return s.LoadBalanceType
}

func (s *NetworkConfig) GetType() *string {
	return s.Type
}

func (s *NetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *NetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *NetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *NetworkConfig) GetWhiteIpGroupList() []*WhiteIpGroup {
	return s.WhiteIpGroupList
}

func (s *NetworkConfig) SetLbReplica(v int32) *NetworkConfig {
	s.LbReplica = &v
	return s
}

func (s *NetworkConfig) SetLoadBalanceConfig(v []*NetworkConfigLoadBalanceConfig) *NetworkConfig {
	s.LoadBalanceConfig = v
	return s
}

func (s *NetworkConfig) SetLoadBalanceType(v string) *NetworkConfig {
	s.LoadBalanceType = &v
	return s
}

func (s *NetworkConfig) SetType(v string) *NetworkConfig {
	s.Type = &v
	return s
}

func (s *NetworkConfig) SetVpcId(v string) *NetworkConfig {
	s.VpcId = &v
	return s
}

func (s *NetworkConfig) SetVsArea(v string) *NetworkConfig {
	s.VsArea = &v
	return s
}

func (s *NetworkConfig) SetVswitchId(v string) *NetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *NetworkConfig) SetWhiteIpGroupList(v []*WhiteIpGroup) *NetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

func (s *NetworkConfig) Validate() error {
	if s.LoadBalanceConfig != nil {
		for _, item := range s.LoadBalanceConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WhiteIpGroupList != nil {
		for _, item := range s.WhiteIpGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NetworkConfigLoadBalanceConfig struct {
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-xxxx
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s NetworkConfigLoadBalanceConfig) String() string {
	return dara.Prettify(s)
}

func (s NetworkConfigLoadBalanceConfig) GoString() string {
	return s.String()
}

func (s *NetworkConfigLoadBalanceConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *NetworkConfigLoadBalanceConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *NetworkConfigLoadBalanceConfig) SetVsArea(v string) *NetworkConfigLoadBalanceConfig {
	s.VsArea = &v
	return s
}

func (s *NetworkConfigLoadBalanceConfig) SetVswitchId(v string) *NetworkConfigLoadBalanceConfig {
	s.VswitchId = &v
	return s
}

func (s *NetworkConfigLoadBalanceConfig) Validate() error {
	return dara.Validate(s)
}
