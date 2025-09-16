// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworks(v *DescribeNetworksResponseBodyNetworks) *DescribeNetworksResponseBody
	GetNetworks() *DescribeNetworksResponseBodyNetworks
	SetPageNumber(v int32) *DescribeNetworksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNetworksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworksResponseBody
	GetTotalCount() *int32
}

type DescribeNetworksResponseBody struct {
	// The VPCs.
	Networks *DescribeNetworksResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBody) GetNetworks() *DescribeNetworksResponseBodyNetworks {
	return s.Networks
}

func (s *DescribeNetworksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworksResponseBody) SetNetworks(v *DescribeNetworksResponseBodyNetworks) *DescribeNetworksResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeNetworksResponseBody) SetPageNumber(v int32) *DescribeNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetPageSize(v int32) *DescribeNetworksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetRequestId(v string) *DescribeNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworksResponseBody) SetTotalCount(v int32) *DescribeNetworksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworks struct {
	Network []*DescribeNetworksResponseBodyNetworksNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworks) GetNetwork() []*DescribeNetworksResponseBodyNetworksNetwork {
	return s.Network
}

func (s *DescribeNetworksResponseBodyNetworks) SetNetwork(v []*DescribeNetworksResponseBodyNetworksNetwork) *DescribeNetworksResponseBodyNetworks {
	s.Network = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworks) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetwork struct {
	// The IPv4 CIDR block of the network.
	//
	// example:
	//
	// 10.0.xx.xx/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The timestamp when the instance was created. Unit: milliseconds.
	//
	// example:
	//
	// 2020-06-16T06:33:15Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the network.
	//
	// example:
	//
	// exampleDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the gateway route table.
	//
	// example:
	//
	// rt-5*****tbs
	GatewayRouteTableId *string `json:"GatewayRouteTableId,omitempty" xml:"GatewayRouteTableId,omitempty"`
	// The ID of the network access control list (ACL).
	//
	// example:
	//
	// nacl-a2do9e413e0spxscd****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5***
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The name of the network.
	//
	// example:
	//
	// example
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// rt-5*****pks
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The IDs of the route tables.
	RouteTableIds *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds `json:"RouteTableIds,omitempty" xml:"RouteTableIds,omitempty" type:"Struct"`
	// The route table ID.
	//
	// example:
	//
	// rtb-5**
	RouterTableId       *string                                                         `json:"RouterTableId,omitempty" xml:"RouterTableId,omitempty"`
	SecondaryCidrBlocks *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Struct"`
	// The status of the network. Valid values:
	//
	// 	- Pending
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeNetworksResponseBodyNetworksNetworkTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The list of vSwitches in the network.
	VSwitchIds *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
}

func (s DescribeNetworksResponseBodyNetworksNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetwork) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetGatewayRouteTableId() *string {
	return s.GatewayRouteTableId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetNetworkName() *string {
	return s.NetworkName
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetRouteTableIds() *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds {
	return s.RouteTableIds
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetRouterTableId() *string {
	return s.RouterTableId
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetSecondaryCidrBlocks() *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks {
	return s.SecondaryCidrBlocks
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetTags() *DescribeNetworksResponseBodyNetworksNetworkTags {
	return s.Tags
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) GetVSwitchIds() *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetCidrBlock(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetCreatedTime(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.CreatedTime = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetDescription(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.Description = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetEnsRegionId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetGatewayRouteTableId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.GatewayRouteTableId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetNetworkAclId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetNetworkId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetNetworkName(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetRouteTableId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.RouteTableId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetRouteTableIds(v *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) *DescribeNetworksResponseBodyNetworksNetwork {
	s.RouteTableIds = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetRouterTableId(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.RouterTableId = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetSecondaryCidrBlocks(v *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) *DescribeNetworksResponseBodyNetworksNetwork {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetStatus(v string) *DescribeNetworksResponseBodyNetworksNetwork {
	s.Status = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetTags(v *DescribeNetworksResponseBodyNetworksNetworkTags) *DescribeNetworksResponseBodyNetworksNetwork {
	s.Tags = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) SetVSwitchIds(v *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) *DescribeNetworksResponseBodyNetworksNetwork {
	s.VSwitchIds = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetworkRouteTableIds struct {
	RouteTableId []*string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) GetRouteTableId() []*string {
	return s.RouteTableId
}

func (s *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) SetRouteTableId(v []*string) *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds {
	s.RouteTableId = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkRouteTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks struct {
	SecondaryCidrBlock []*string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) GetSecondaryCidrBlock() []*string {
	return s.SecondaryCidrBlock
}

func (s *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) SetSecondaryCidrBlock(v []*string) *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks {
	s.SecondaryCidrBlock = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkSecondaryCidrBlocks) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetworkTags struct {
	Tag []*DescribeNetworksResponseBodyNetworksNetworkTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTags) GetTag() []*DescribeNetworksResponseBodyNetworksNetworkTagsTag {
	return s.Tag
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTags) SetTag(v []*DescribeNetworksResponseBodyNetworksNetworkTagsTag) *DescribeNetworksResponseBodyNetworksNetworkTags {
	s.Tag = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTags) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetworkTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The bandwidth.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) SetKey(v string) *DescribeNetworksResponseBodyNetworksNetworkTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) SetTagKey(v string) *DescribeNetworksResponseBodyNetworksNetworkTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) SetTagValue(v string) *DescribeNetworksResponseBodyNetworksNetworkTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) SetValue(v string) *DescribeNetworksResponseBodyNetworksNetworkTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworksResponseBodyNetworksNetworkVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) SetVSwitchId(v []*string) *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeNetworksResponseBodyNetworksNetworkVSwitchIds) Validate() error {
	return dara.Validate(s)
}
