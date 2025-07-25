// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4IspCityNodes(v *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) *ListCloudGtmMonitorNodesResponseBody
	GetIpv4IspCityNodes() *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes
	SetIpv6IspCityNodes(v *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) *ListCloudGtmMonitorNodesResponseBody
	GetIpv6IspCityNodes() *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes
	SetRequestId(v string) *ListCloudGtmMonitorNodesResponseBody
	GetRequestId() *string
}

type ListCloudGtmMonitorNodesResponseBody struct {
	// Public IPv4 monitoring node list.
	Ipv4IspCityNodes *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes `json:"Ipv4IspCityNodes,omitempty" xml:"Ipv4IspCityNodes,omitempty" type:"Struct"`
	// List of public IPv6 monitoring nodes.
	Ipv6IspCityNodes *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes `json:"Ipv6IspCityNodes,omitempty" xml:"Ipv6IspCityNodes,omitempty" type:"Struct"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCloudGtmMonitorNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBody) GetIpv4IspCityNodes() *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes {
	return s.Ipv4IspCityNodes
}

func (s *ListCloudGtmMonitorNodesResponseBody) GetIpv6IspCityNodes() *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes {
	return s.Ipv6IspCityNodes
}

func (s *ListCloudGtmMonitorNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmMonitorNodesResponseBody) SetIpv4IspCityNodes(v *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) *ListCloudGtmMonitorNodesResponseBody {
	s.Ipv4IspCityNodes = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBody) SetIpv6IspCityNodes(v *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) *ListCloudGtmMonitorNodesResponseBody {
	s.Ipv6IspCityNodes = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBody) SetRequestId(v string) *ListCloudGtmMonitorNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes struct {
	Ipv4IspCityNode []*ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode `json:"Ipv4IspCityNode,omitempty" xml:"Ipv4IspCityNode,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) GetIpv4IspCityNode() []*ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	return s.Ipv4IspCityNode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) SetIpv4IspCityNode(v []*ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes {
	s.Ipv4IspCityNode = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodes) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode struct {
	// City code.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// City name.
	//
	// example:
	//
	// Beijing
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Country code.
	//
	// example:
	//
	// 629
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Country name.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// Monitor node default selection:
	//
	// - true: Selected by default
	//
	// - false: Not selected by default
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// Monitor probe group name.
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Monitoring node group type, currently supported:
	//
	// - BGP: BGP node
	//
	// - OVERSEAS: International node
	//
	// - ISP: Carrier node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// List of node IP addresses.
	Ips *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// Operator code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// Operator name.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Unique identifier ID of the probe node.
	//
	// example:
	//
	// node-ewze1bysndy4gf**j8
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIps() *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps {
	return s.Ips
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) GetNodeId() *string {
	return s.NodeId
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCityCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CityCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCityName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CityName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCountryCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CountryCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetCountryName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.CountryName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetDefaultSelected(v bool) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetGroupName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.GroupName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetGroupType(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.GroupType = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIps(v *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.Ips = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIspCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.IspCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetIspName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.IspName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) SetNodeId(v string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode {
	s.NodeId = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNode) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) SetIp(v []*string) *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv4IspCityNodesIpv4IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes struct {
	Ipv6IspCityNode []*ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode `json:"Ipv6IspCityNode,omitempty" xml:"Ipv6IspCityNode,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) GetIpv6IspCityNode() []*ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	return s.Ipv6IspCityNode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) SetIpv6IspCityNode(v []*ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes {
	s.Ipv6IspCityNode = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodes) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode struct {
	// City code.
	//
	// example:
	//
	// 357
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// City name.
	//
	// example:
	//
	// Shanghai
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Country code.
	//
	// example:
	//
	// 629
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Country name.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// Monitor node default selection:
	//
	// - true: Selected by default
	//
	// - false: Not selected by default
	//
	// example:
	//
	// true
	DefaultSelected *bool `json:"DefaultSelected,omitempty" xml:"DefaultSelected,omitempty"`
	// Monitoring probe group name.
	//
	// example:
	//
	// BGP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Monitoring node group type, currently supported:
	//
	// - BGP: BGP node
	//
	// - OVERSEAS: International node
	//
	// - ISP: Carrier node
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// List of node IP addresses.
	Ips *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Struct"`
	// Operator code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// Operator name.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Unique identifier ID of the probe node.
	//
	// example:
	//
	// node-ewze1bysndy4gf**j8
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetDefaultSelected() *bool {
	return s.DefaultSelected
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIps() *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps {
	return s.Ips
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) GetNodeId() *string {
	return s.NodeId
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCityCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CityCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCityName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CityName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCountryCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CountryCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetCountryName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.CountryName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetDefaultSelected(v bool) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.DefaultSelected = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetGroupName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.GroupName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetGroupType(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.GroupType = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIps(v *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.Ips = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIspCode(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.IspCode = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetIspName(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.IspName = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) SetNodeId(v string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode {
	s.NodeId = &v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNode) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) GetIp() []*string {
	return s.Ip
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) SetIp(v []*string) *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps {
	s.Ip = v
	return s
}

func (s *ListCloudGtmMonitorNodesResponseBodyIpv6IspCityNodesIpv6IspCityNodeIps) Validate() error {
	return dara.Validate(s)
}
