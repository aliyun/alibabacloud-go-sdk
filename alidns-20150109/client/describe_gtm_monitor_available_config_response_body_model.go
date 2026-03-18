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
	CityCode        *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CityName        *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	DefaultSelected *bool   `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType       *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IspCode         *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	IspName         *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	Mainland        *bool   `json:"Mainland,omitempty" xml:"Mainland,omitempty"`
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
