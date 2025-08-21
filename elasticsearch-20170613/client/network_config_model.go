// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
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
	return dara.Validate(s)
}
