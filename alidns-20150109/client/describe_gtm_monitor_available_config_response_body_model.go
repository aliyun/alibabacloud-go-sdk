// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorAvailableConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIspCityNodes(v *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) *DescribeGtmMonitorAvailableConfigResponseBody
	GetIspCityNodes() *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes
	SetRequestId(v string) *DescribeGtmMonitorAvailableConfigResponseBody
	GetRequestId() *string
}

type DescribeGtmMonitorAvailableConfigResponseBody struct {
	// The monitored nodes.
	IspCityNodes *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) GetIspCityNodes() *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	return s.IspCityNodes
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) SetIspCityNodes(v *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) *DescribeGtmMonitorAvailableConfigResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) SetRequestId(v string) *DescribeGtmMonitorAvailableConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBody) Validate() error {
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes struct {
	IspCityNode []*DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) GetIspCityNode() []*DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) SetIspCityNode(v []*DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodes) Validate() error {
	if s.IspCityNode != nil {
		for _, item := range s.IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode struct {
	// The code of the city where the monitored node is deployed.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city where the monitored node is deployed.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Indicates whether the monitored node is selected for the health check by default.
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// The name of the group to which the monitored node belongs.
	//
	// Valid values: Overseas Nodes, BGP Nodes, and ISP Nodes.
	//
	// example:
	//
	// Overseas Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group to which the monitored node belongs.
	//
	// Valid values: BGP, OVERSEAS, and ISP.
	//
	// example:
	//
	// OVERSEAS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The code of the Internet service provider (ISP) to which the monitored node belongs.
	//
	// 	- If the value of the GroupType parameter is BGP or OVERSEAS, the value of IspCode is 465 by default.
	//
	// 	- If the value of the GroupType parameter is not BGP or OVERSEAS, valid values of IspCode are 232, 132, and 5. and is used together with CityCode.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP to which the monitored node belongs.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Indicates whether the monitored node is deployed in the Chinese mainland.
	//
	// example:
	//
	// true
	Mainland *bool `json:"Mainland,omitempty" xml:"Mainland,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) GetMainland() *bool {
	return s.Mainland
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetCityCode(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetCityName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetDefaultSelected(v bool) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetGroupName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetGroupType(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetIspCode(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetIspName(v string) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) SetMainland(v bool) *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode {
	s.Mainland = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigResponseBodyIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
