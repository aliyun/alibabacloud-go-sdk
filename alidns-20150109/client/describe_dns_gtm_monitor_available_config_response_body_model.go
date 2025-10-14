// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainIpv4IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody
	GetDomainIpv4IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes
	SetDomainIpv6IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody
	GetDomainIpv6IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes
	SetIpv4IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody
	GetIpv4IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes
	SetIpv6IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody
	GetIpv6IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes
	SetRequestId(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBody
	GetRequestId() *string
}

type DescribeDnsGtmMonitorAvailableConfigResponseBody struct {
	// The nodes that perform health checks on domain names that use public IPv4 addresses.
	DomainIpv4IspCityNodes *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes `json:"DomainIpv4IspCityNodes,omitempty" xml:"DomainIpv4IspCityNodes,omitempty" type:"Struct"`
	// The nodes that perform health checks on domain names that use public IPv6 addresses.
	DomainIpv6IspCityNodes *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes `json:"DomainIpv6IspCityNodes,omitempty" xml:"DomainIpv6IspCityNodes,omitempty" type:"Struct"`
	// The nodes that perform health checks on public IPv4 addresses.
	Ipv4IspCityNodes *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes `json:"Ipv4IspCityNodes,omitempty" xml:"Ipv4IspCityNodes,omitempty" type:"Struct"`
	// The nodes that perform health checks on public IPv6 addresses.
	Ipv6IspCityNodes *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes `json:"Ipv6IspCityNodes,omitempty" xml:"Ipv6IspCityNodes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) GetDomainIpv4IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	return s.DomainIpv4IspCityNodes
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) GetDomainIpv6IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	return s.DomainIpv6IspCityNodes
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) GetIpv4IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	return s.Ipv4IspCityNodes
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) GetIpv6IspCityNodes() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	return s.Ipv6IspCityNodes
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetDomainIpv4IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.DomainIpv4IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetDomainIpv6IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.DomainIpv6IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetIpv4IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.Ipv4IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetIpv6IspCityNodes(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.Ipv6IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBody) Validate() error {
	if s.DomainIpv4IspCityNodes != nil {
		if err := s.DomainIpv4IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	if s.DomainIpv6IspCityNodes != nil {
		if err := s.DomainIpv6IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv4IspCityNodes != nil {
		if err := s.Ipv4IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6IspCityNodes != nil {
		if err := s.Ipv6IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes struct {
	DomainIpv4IspCityNode []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode `json:"DomainIpv4IspCityNode,omitempty" xml:"DomainIpv4IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) GetDomainIpv4IspCityNode() []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	return s.DomainIpv4IspCityNode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) SetDomainIpv4IspCityNode(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes {
	s.DomainIpv4IspCityNode = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodes) Validate() error {
	if s.DomainIpv4IspCityNode != nil {
		for _, item := range s.DomainIpv4IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode struct {
	// The city code.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Indicates whether the health check node is selected by default.
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// 	- BGP: BGP node
	//
	// 	- OVERSEAS: node outside the Chinese mainland
	//
	// 	- ISP: ISP node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The IP addresses of the health check nodes.
	Ips *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// The ISP code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetIps() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps {
	return s.Ips
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetIps(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.Ips = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNode) Validate() error {
	if s.Ips != nil {
		if err := s.Ips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) SetIp(v []*string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv4IspCityNodesDomainIpv4IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes struct {
	DomainIpv6IspCityNode []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode `json:"DomainIpv6IspCityNode,omitempty" xml:"DomainIpv6IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) GetDomainIpv6IspCityNode() []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	return s.DomainIpv6IspCityNode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) SetDomainIpv6IspCityNode(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes {
	s.DomainIpv6IspCityNode = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodes) Validate() error {
	if s.DomainIpv6IspCityNode != nil {
		for _, item := range s.DomainIpv6IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode struct {
	// The city code.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Indicates whether the health check node is selected by default.
	//
	// example:
	//
	// false
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// 	- BGP: BGP node
	//
	// 	- OVERSEAS: node outside the Chinese mainland
	//
	// 	- ISP: ISP node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// This parameter is not returned.
	Ips *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// The ISP code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetIps() *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps {
	return s.Ips
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetIps(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.Ips = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNode) Validate() error {
	if s.Ips != nil {
		if err := s.Ips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps struct {
	Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) SetIp(v []*string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyDomainIpv6IspCityNodesDomainIpv6IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes struct {
	Ipv4IspCityNode []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode `json:"Ipv4IspCityNode,omitempty" xml:"Ipv4IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) GetIpv4IspCityNode() []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	return s.Ipv4IspCityNode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) SetIpv4IspCityNode(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes {
	s.Ipv4IspCityNode = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodes) Validate() error {
	if s.Ipv4IspCityNode != nil {
		for _, item := range s.Ipv4IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode struct {
	// The city code.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Indicates whether the health check node is selected by default.
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// 	- BGP: Border Gateway Protocol (BGP) node
	//
	// 	- OVERSEAS: node outside the Chinese mainland
	//
	// 	- ISP: ISP node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The IP addresses of the health check nodes.
	Ips *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// The Internet service provider (ISP) code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIps() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps {
	return s.Ips
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIps(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.Ips = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNode) Validate() error {
	if s.Ips != nil {
		if err := s.Ips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) SetIp(v []*string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes struct {
	Ipv6IspCityNode []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode `json:"Ipv6IspCityNode,omitempty" xml:"Ipv6IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) GetIpv6IspCityNode() []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	return s.Ipv6IspCityNode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) SetIpv6IspCityNode(v []*DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes {
	s.Ipv6IspCityNode = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodes) Validate() error {
	if s.Ipv6IspCityNode != nil {
		for _, item := range s.Ipv6IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode struct {
	// The city code.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Indicates whether the health check node is selected by default.
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// 	- BGP: BGP node
	//
	// 	- OVERSEAS: node outside the Chinese mainland
	//
	// 	- ISP: ISP node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// This parameter is not returned.
	Ips *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// The ISP code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIps() *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps {
	return s.Ips
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCityCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCityName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetDefaultSelected(v bool) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetGroupName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetGroupType(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIps(v *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.Ips = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIspCode(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIspName(v string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNode) Validate() error {
	if s.Ips != nil {
		if err := s.Ips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) SetIp(v []*string) *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}
