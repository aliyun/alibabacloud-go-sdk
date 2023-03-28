// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	ClusterDescription    *string                           `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	ClusterName           *string                           `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string                           `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Components            []*CreateClusterRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	IgnoreFailedNodeTasks *bool                             `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	Networks              *CreateClusterRequestNetworks     `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	NodeGroups            []*CreateClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	ResourceGroupId       *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                   []*CreateClusterRequestTag        `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClusterDescription(v string) *CreateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetComponents(v []*CreateClusterRequestComponents) *CreateClusterRequest {
	s.Components = v
	return s
}

func (s *CreateClusterRequest) SetIgnoreFailedNodeTasks(v bool) *CreateClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *CreateClusterRequest) SetNetworks(v *CreateClusterRequestNetworks) *CreateClusterRequest {
	s.Networks = v
	return s
}

func (s *CreateClusterRequest) SetNodeGroups(v []*CreateClusterRequestNodeGroups) *CreateClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest {
	s.Tag = v
	return s
}

type CreateClusterRequestComponents struct {
	ComponentConfig *CreateClusterRequestComponentsComponentConfig `json:"ComponentConfig,omitempty" xml:"ComponentConfig,omitempty" type:"Struct"`
	ComponentType   *string                                        `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s CreateClusterRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponents) SetComponentConfig(v *CreateClusterRequestComponentsComponentConfig) *CreateClusterRequestComponents {
	s.ComponentConfig = v
	return s
}

func (s *CreateClusterRequestComponents) SetComponentType(v string) *CreateClusterRequestComponents {
	s.ComponentType = &v
	return s
}

type CreateClusterRequestComponentsComponentConfig struct {
	BasicArgs interface{}   `json:"BasicArgs,omitempty" xml:"BasicArgs,omitempty"`
	NodeUnits []interface{} `json:"NodeUnits,omitempty" xml:"NodeUnits,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestComponentsComponentConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestComponentsComponentConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponentsComponentConfig) SetBasicArgs(v interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.BasicArgs = v
	return s
}

func (s *CreateClusterRequestComponentsComponentConfig) SetNodeUnits(v []interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.NodeUnits = v
	return s
}

type CreateClusterRequestNetworks struct {
	IpAllocationPolicy []*CreateClusterRequestNetworksIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	NewVpdInfo         *CreateClusterRequestNetworksNewVpdInfo           `json:"NewVpdInfo,omitempty" xml:"NewVpdInfo,omitempty" type:"Struct"`
	VpdInfo            *CreateClusterRequestNetworksVpdInfo              `json:"VpdInfo,omitempty" xml:"VpdInfo,omitempty" type:"Struct"`
}

func (s CreateClusterRequestNetworks) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworks) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworks) SetIpAllocationPolicy(v []*CreateClusterRequestNetworksIpAllocationPolicy) *CreateClusterRequestNetworks {
	s.IpAllocationPolicy = v
	return s
}

func (s *CreateClusterRequestNetworks) SetNewVpdInfo(v *CreateClusterRequestNetworksNewVpdInfo) *CreateClusterRequestNetworks {
	s.NewVpdInfo = v
	return s
}

func (s *CreateClusterRequestNetworks) SetVpdInfo(v *CreateClusterRequestNetworksVpdInfo) *CreateClusterRequestNetworks {
	s.VpdInfo = v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicy struct {
	BondPolicy        *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy          `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	MachineTypePolicy []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	NodePolicy        []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy        `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetBondPolicy(v *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.BondPolicy = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetMachineTypePolicy(v []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.MachineTypePolicy = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetNodePolicy(v []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.NodePolicy = v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicy struct {
	BondDefaultSubnet *string                                                          `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	Bonds             []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.BondDefaultSubnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.Bonds = v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Subnet = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy struct {
	Bonds       []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	MachineType *string                                                                 `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.MachineType = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Subnet = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyNodePolicy struct {
	Bonds  []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	NodeId *string                                                          `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetNodeId(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.NodeId = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Subnet = &v
	return s
}

type CreateClusterRequestNetworksNewVpdInfo struct {
	CenId            *string                                             `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CloudLinkCidr    *string                                             `json:"CloudLinkCidr,omitempty" xml:"CloudLinkCidr,omitempty"`
	CloudLinkId      *string                                             `json:"CloudLinkId,omitempty" xml:"CloudLinkId,omitempty"`
	MonitorVpcId     *string                                             `json:"MonitorVpcId,omitempty" xml:"MonitorVpcId,omitempty"`
	MonitorVswitchId *string                                             `json:"MonitorVswitchId,omitempty" xml:"MonitorVswitchId,omitempty"`
	VpdCidr          *string                                             `json:"VpdCidr,omitempty" xml:"VpdCidr,omitempty"`
	VpdSubnets       []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksNewVpdInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCenId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CenId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCloudLinkCidr(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CloudLinkCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCloudLinkId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CloudLinkId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetMonitorVpcId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.MonitorVpcId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetMonitorVswitchId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.MonitorVswitchId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetVpdCidr(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.VpdCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetVpdSubnets(v []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets) *CreateClusterRequestNetworksNewVpdInfo {
	s.VpdSubnets = v
	return s
}

type CreateClusterRequestNetworksNewVpdInfoVpdSubnets struct {
	SubnetCidr *string `json:"SubnetCidr,omitempty" xml:"SubnetCidr,omitempty"`
	SubnetType *string `json:"SubnetType,omitempty" xml:"SubnetType,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNetworksNewVpdInfoVpdSubnets) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetSubnetCidr(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.SubnetCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetSubnetType(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.SubnetType = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetZoneId(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestNetworksVpdInfo struct {
	VpdId      *string   `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	VpdSubnets []*string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksVpdInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksVpdInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdId(v string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdId = &v
	return s
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdSubnets(v []*string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdSubnets = v
	return s
}

type CreateClusterRequestNodeGroups struct {
	ImageId              *string                                `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MachineType          *string                                `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	NodeGroupDescription *string                                `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	NodeGroupName        *string                                `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	Nodes                []*CreateClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	ZoneId               *string                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroups) SetImageId(v string) *CreateClusterRequestNodeGroups {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetMachineType(v string) *CreateClusterRequestNodeGroups {
	s.MachineType = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodeGroupDescription(v string) *CreateClusterRequestNodeGroups {
	s.NodeGroupDescription = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodeGroupName(v string) *CreateClusterRequestNodeGroups {
	s.NodeGroupName = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodes(v []*CreateClusterRequestNodeGroupsNodes) *CreateClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetZoneId(v string) *CreateClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestNodeGroupsNodes struct {
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsNodes) SetHostname(v string) *CreateClusterRequestNodeGroupsNodes {
	s.Hostname = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetLoginPassword(v string) *CreateClusterRequestNodeGroupsNodes {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetNodeId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

type CreateClusterRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

type CreateClusterShrinkRequest struct {
	ClusterDescription    *string                          `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	ClusterName           *string                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string                          `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ComponentsShrink      *string                          `json:"Components,omitempty" xml:"Components,omitempty"`
	IgnoreFailedNodeTasks *bool                            `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NetworksShrink        *string                          `json:"Networks,omitempty" xml:"Networks,omitempty"`
	NodeGroupsShrink      *string                          `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	ResourceGroupId       *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                   []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) SetClusterDescription(v string) *CreateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterType(v string) *CreateClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetComponentsShrink(v string) *CreateClusterShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *CreateClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNetworksShrink(v string) *CreateClusterShrinkRequest {
	s.NetworksShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetResourceGroupId(v string) *CreateClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest {
	s.Tag = v
	return s
}

type CreateClusterShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	ClusterDescription *string                                  `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	ClusterId          *string                                  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName        *string                                  `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType        *string                                  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Components         []*DescribeClusterResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	CreateTime         *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Networks           []*DescribeClusterResponseBodyNetworks   `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	NodeCount          *int64                                   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupCount     *int64                                   `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	OperatingState     *string                                  `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	RequestId          *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId    *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId             *string                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpdateTime         *string                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterDescription(v string) *DescribeClusterResponseBody {
	s.ClusterDescription = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterId(v string) *DescribeClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterName(v string) *DescribeClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterType(v string) *DescribeClusterResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterResponseBody) SetComponents(v []*DescribeClusterResponseBodyComponents) *DescribeClusterResponseBody {
	s.Components = v
	return s
}

func (s *DescribeClusterResponseBody) SetCreateTime(v string) *DescribeClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBody) SetNetworks(v []*DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeClusterResponseBody) SetNodeCount(v int64) *DescribeClusterResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterResponseBody) SetNodeGroupCount(v int64) *DescribeClusterResponseBody {
	s.NodeGroupCount = &v
	return s
}

func (s *DescribeClusterResponseBody) SetOperatingState(v string) *DescribeClusterResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetResourceGroupId(v string) *DescribeClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetTaskId(v string) *DescribeClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetUpdateTime(v string) *DescribeClusterResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeClusterResponseBodyComponents struct {
	ComponentId   *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s DescribeClusterResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyComponents) SetComponentId(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentId = &v
	return s
}

func (s *DescribeClusterResponseBodyComponents) SetComponentType(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentType = &v
	return s
}

type DescribeClusterResponseBodyNetworks struct {
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeClusterResponseBodyNetworks) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyNetworks) SetVpdId(v string) *DescribeClusterResponseBodyNetworks {
	s.VpdId = &v
	return s
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetStatusCode(v int32) *DescribeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeNodeRequest struct {
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeRequest) SetNodeId(v string) *DescribeNodeRequest {
	s.NodeId = &v
	return s
}

type DescribeNodeResponseBody struct {
	ClusterId      *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName    *string                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CreateTime     *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime    *string                             `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Hostname       *string                             `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	ImageId        *string                             `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MachineType    *string                             `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	Networks       []*DescribeNodeResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	NodeGroupId    *string                             `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName  *string                             `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeId         *string                             `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OperatingState *string                             `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sn             *string                             `json:"Sn,omitempty" xml:"Sn,omitempty"`
	ZoneId         *string                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBody) SetClusterId(v string) *DescribeNodeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetClusterName(v string) *DescribeNodeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetCreateTime(v string) *DescribeNodeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetExpiredTime(v string) *DescribeNodeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHostname(v string) *DescribeNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageId(v string) *DescribeNodeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetMachineType(v string) *DescribeNodeResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNetworks(v []*DescribeNodeResponseBodyNetworks) *DescribeNodeResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupId(v string) *DescribeNodeResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupName(v string) *DescribeNodeResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeId(v string) *DescribeNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetOperatingState(v string) *DescribeNodeResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeNodeResponseBody) SetRequestId(v string) *DescribeNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetSn(v string) *DescribeNodeResponseBody {
	s.Sn = &v
	return s
}

func (s *DescribeNodeResponseBody) SetZoneId(v string) *DescribeNodeResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeNodeResponseBodyNetworks struct {
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeNodeResponseBodyNetworks) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyNetworks) SetBondName(v string) *DescribeNodeResponseBodyNetworks {
	s.BondName = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetIp(v string) *DescribeNodeResponseBodyNetworks {
	s.Ip = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetSubnetId(v string) *DescribeNodeResponseBodyNetworks {
	s.SubnetId = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetVpdId(v string) *DescribeNodeResponseBodyNetworks {
	s.VpdId = &v
	return s
}

type DescribeNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponse) SetHeaders(v map[string]*string) *DescribeNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeResponse) SetStatusCode(v int32) *DescribeNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeResponse) SetBody(v *DescribeNodeResponseBody) *DescribeNodeResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskResponseBody struct {
	ClusterId   *string                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CreateTime  *string                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Message     *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Steps       []*DescribeTaskResponseBodySteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	TaskState   *string                          `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	TaskType    *string                          `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime  *string                          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) SetClusterId(v string) *DescribeTaskResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetClusterName(v string) *DescribeTaskResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCreateTime(v string) *DescribeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetMessage(v string) *DescribeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetSteps(v []*DescribeTaskResponseBodySteps) *DescribeTaskResponseBody {
	s.Steps = v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskState(v string) *DescribeTaskResponseBody {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskType(v string) *DescribeTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdateTime(v string) *DescribeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponseBodySteps struct {
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	StageTag   *string                                  `json:"StageTag,omitempty" xml:"StageTag,omitempty"`
	StartTime  *string                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepName   *string                                  `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepState  *string                                  `json:"StepState,omitempty" xml:"StepState,omitempty"`
	StepType   *string                                  `json:"StepType,omitempty" xml:"StepType,omitempty"`
	SubTasks   []*DescribeTaskResponseBodyStepsSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	UpdateTime *string                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodySteps) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBodySteps) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodySteps) SetMessage(v string) *DescribeTaskResponseBodySteps {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStageTag(v string) *DescribeTaskResponseBodySteps {
	s.StageTag = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStartTime(v string) *DescribeTaskResponseBodySteps {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepName(v string) *DescribeTaskResponseBodySteps {
	s.StepName = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepState(v string) *DescribeTaskResponseBodySteps {
	s.StepState = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepType(v string) *DescribeTaskResponseBodySteps {
	s.StepType = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetSubTasks(v []*DescribeTaskResponseBodyStepsSubTasks) *DescribeTaskResponseBodySteps {
	s.SubTasks = v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetUpdateTime(v string) *DescribeTaskResponseBodySteps {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponseBodyStepsSubTasks struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskState  *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	TaskType   *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodyStepsSubTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBodyStepsSubTasks) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetCreateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetMessage(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskId(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskState(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskType(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetUpdateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) SetHeaders(v map[string]*string) *DescribeTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskResponse) SetStatusCode(v int32) *DescribeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskResponse) SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ExtendClusterRequest struct {
	ClusterId             *string                                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool                                     `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	IpAllocationPolicy    []*ExtendClusterRequestIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	NodeGroups            []*ExtendClusterRequestNodeGroups         `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	VpdSubnets            []*string                                 `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequest) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequest) SetClusterId(v string) *ExtendClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ExtendClusterRequest) SetIgnoreFailedNodeTasks(v bool) *ExtendClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ExtendClusterRequest) SetIpAllocationPolicy(v []*ExtendClusterRequestIpAllocationPolicy) *ExtendClusterRequest {
	s.IpAllocationPolicy = v
	return s
}

func (s *ExtendClusterRequest) SetNodeGroups(v []*ExtendClusterRequestNodeGroups) *ExtendClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *ExtendClusterRequest) SetVpdSubnets(v []*string) *ExtendClusterRequest {
	s.VpdSubnets = v
	return s
}

type ExtendClusterRequestIpAllocationPolicy struct {
	BondPolicy        *ExtendClusterRequestIpAllocationPolicyBondPolicy          `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	MachineTypePolicy []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	NodePolicy        []*ExtendClusterRequestIpAllocationPolicyNodePolicy        `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestIpAllocationPolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetBondPolicy(v *ExtendClusterRequestIpAllocationPolicyBondPolicy) *ExtendClusterRequestIpAllocationPolicy {
	s.BondPolicy = v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetMachineTypePolicy(v []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) *ExtendClusterRequestIpAllocationPolicy {
	s.MachineTypePolicy = v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetNodePolicy(v []*ExtendClusterRequestIpAllocationPolicyNodePolicy) *ExtendClusterRequestIpAllocationPolicy {
	s.NodePolicy = v
	return s
}

type ExtendClusterRequestIpAllocationPolicyBondPolicy struct {
	BondDefaultSubnet *string                                                  `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	Bonds             []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
	s.BondDefaultSubnet = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
	s.Bonds = v
	return s
}

type ExtendClusterRequestIpAllocationPolicyBondPolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicy struct {
	Bonds       []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	MachineType *string                                                         `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
	s.Bonds = v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
	s.MachineType = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyNodePolicy struct {
	Bonds  []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	NodeId *string                                                  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
	s.Bonds = v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetNodeId(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
	s.NodeId = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyNodePolicyBonds struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestNodeGroups struct {
	NodeGroupId *string                                `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	Nodes       []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroups) SetNodeGroupId(v string) *ExtendClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodes(v []*ExtendClusterRequestNodeGroupsNodes) *ExtendClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

type ExtendClusterRequestNodeGroupsNodes struct {
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetHostname(v string) *ExtendClusterRequestNodeGroupsNodes {
	s.Hostname = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetLoginPassword(v string) *ExtendClusterRequestNodeGroupsNodes {
	s.LoginPassword = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetNodeId(v string) *ExtendClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

type ExtendClusterShrinkRequest struct {
	ClusterId                *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks    *bool   `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	IpAllocationPolicyShrink *string `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty"`
	NodeGroupsShrink         *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	VpdSubnetsShrink         *string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty"`
}

func (s ExtendClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExtendClusterShrinkRequest) SetClusterId(v string) *ExtendClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ExtendClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetIpAllocationPolicyShrink(v string) *ExtendClusterShrinkRequest {
	s.IpAllocationPolicyShrink = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetNodeGroupsShrink(v string) *ExtendClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest {
	s.VpdSubnetsShrink = &v
	return s
}

type ExtendClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExtendClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponseBody) SetRequestId(v string) *ExtendClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtendClusterResponseBody) SetTaskId(v string) *ExtendClusterResponseBody {
	s.TaskId = &v
	return s
}

type ExtendClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtendClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtendClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterResponse) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponse) SetHeaders(v map[string]*string) *ExtendClusterResponse {
	s.Headers = v
	return s
}

func (s *ExtendClusterResponse) SetStatusCode(v int32) *ExtendClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtendClusterResponse) SetBody(v *ExtendClusterResponseBody) *ExtendClusterResponse {
	s.Body = v
	return s
}

type ListClusterNodesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s ListClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequest) SetClusterId(v string) *ListClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterNodesRequest) SetMaxResults(v int64) *ListClusterNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClusterNodesRequest) SetNextToken(v string) *ListClusterNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListClusterNodesRequest) SetNodeGroupId(v string) *ListClusterNodesRequest {
	s.NodeGroupId = &v
	return s
}

type ListClusterNodesResponseBody struct {
	NextToken *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Nodes     []*ListClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBody) SetNextToken(v string) *ListClusterNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClusterNodesResponseBody) SetNodes(v []*ListClusterNodesResponseBodyNodes) *ListClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListClusterNodesResponseBody) SetRequestId(v string) *ListClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterNodesResponseBodyNodes struct {
	CreateTime     *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime    *string                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Hostname       *string                                      `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	ImageId        *string                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MachineType    *string                                      `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	Networks       []*ListClusterNodesResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	NodeGroupId    *string                                      `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName  *string                                      `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeId         *string                                      `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OperatingState *string                                      `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	Sn             *string                                      `json:"Sn,omitempty" xml:"Sn,omitempty"`
	ZoneId         *string                                      `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodes) SetCreateTime(v string) *ListClusterNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetExpiredTime(v string) *ListClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetHostname(v string) *ListClusterNodesResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetImageId(v string) *ListClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetMachineType(v string) *ListClusterNodesResponseBodyNodes {
	s.MachineType = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNetworks(v []*ListClusterNodesResponseBodyNodesNetworks) *ListClusterNodesResponseBodyNodes {
	s.Networks = v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeGroupId(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeGroupName(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeGroupName = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeId(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetOperatingState(v string) *ListClusterNodesResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetSn(v string) *ListClusterNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetZoneId(v string) *ListClusterNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListClusterNodesResponseBodyNodesNetworks struct {
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	VpdId    *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListClusterNodesResponseBodyNodesNetworks) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesNetworks) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetBondName(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.BondName = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetIp(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.Ip = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetSubnetId(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.SubnetId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetVpdId(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.VpdId = &v
	return s
}

type ListClusterNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponse) SetHeaders(v map[string]*string) *ListClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNodesResponse) SetStatusCode(v int32) *ListClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNodesResponse) SetBody(v *ListClusterNodesResponseBody) *ListClusterNodesResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	MaxResults      *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetMaxResults(v int64) *ListClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClustersRequest) SetNextToken(v string) *ListClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters  []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetNextToken(v string) *ListClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterDescription *string     `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	ClusterId          *string     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName        *string     `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType        *string     `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Components         interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	CreateTime         *string     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NodeCount          *int64      `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupCount     *int64      `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	OperatingState     *string     `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	ResourceGroupId    *string     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId             *string     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UpdateTime         *string     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterDescription(v string) *ListClustersResponseBodyClusters {
	s.ClusterDescription = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterName(v string) *ListClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterType(v string) *ListClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComponents(v interface{}) *ListClustersResponseBodyClusters {
	s.Components = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCreateTime(v string) *ListClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeGroupCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeGroupCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetOperatingState(v string) *ListClustersResponseBodyClusters {
	s.OperatingState = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetResourceGroupId(v string) *ListClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTaskId(v string) *ListClustersResponseBodyClusters {
	s.TaskId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUpdateTime(v string) *ListClustersResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListFreeNodesRequest struct {
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MaxResults  *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFreeNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequest) SetMachineType(v string) *ListFreeNodesRequest {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesRequest) SetMaxResults(v int64) *ListFreeNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreeNodesRequest) SetNextToken(v string) *ListFreeNodesRequest {
	s.NextToken = &v
	return s
}

type ListFreeNodesResponseBody struct {
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Nodes     []*ListFreeNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFreeNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBody) SetNextToken(v string) *ListFreeNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreeNodesResponseBody) SetNodes(v []*ListFreeNodesResponseBodyNodes) *ListFreeNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListFreeNodesResponseBody) SetRequestId(v string) *ListFreeNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListFreeNodesResponseBodyNodes struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	NodeId      *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Sn          *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodes) SetCreateTime(v string) *ListFreeNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetExpiredTime(v string) *ListFreeNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetMachineType(v string) *ListFreeNodesResponseBodyNodes {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetNodeId(v string) *ListFreeNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetSn(v string) *ListFreeNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetZoneId(v string) *ListFreeNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListFreeNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFreeNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponse) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponse) SetHeaders(v map[string]*string) *ListFreeNodesResponse {
	s.Headers = v
	return s
}

func (s *ListFreeNodesResponse) SetStatusCode(v int32) *ListFreeNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreeNodesResponse) SetBody(v *ListFreeNodesResponseBody) *ListFreeNodesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type RebootNodesRequest struct {
	ClusterId             *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool     `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	Nodes                 []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RebootNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesRequest) SetClusterId(v string) *RebootNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesRequest) SetNodes(v []*string) *RebootNodesRequest {
	s.Nodes = v
	return s
}

type RebootNodesShrinkRequest struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool   `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NodesShrink           *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s RebootNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesShrinkRequest) SetClusterId(v string) *RebootNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetNodesShrink(v string) *RebootNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

type RebootNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RebootNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootNodesResponseBody) SetRequestId(v string) *RebootNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootNodesResponseBody) SetTaskId(v string) *RebootNodesResponseBody {
	s.TaskId = &v
	return s
}

type RebootNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebootNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesResponse) GoString() string {
	return s.String()
}

func (s *RebootNodesResponse) SetHeaders(v map[string]*string) *RebootNodesResponse {
	s.Headers = v
	return s
}

func (s *RebootNodesResponse) SetStatusCode(v int32) *RebootNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootNodesResponse) SetBody(v *RebootNodesResponseBody) *RebootNodesResponse {
	s.Body = v
	return s
}

type ReimageNodesRequest struct {
	ClusterId             *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool                       `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	Nodes                 []*ReimageNodesRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ReimageNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequest) SetClusterId(v string) *ReimageNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesRequest) SetNodes(v []*ReimageNodesRequestNodes) *ReimageNodesRequest {
	s.Nodes = v
	return s
}

type ReimageNodesRequestNodes struct {
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ReimageNodesRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesRequestNodes) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequestNodes) SetHostname(v string) *ReimageNodesRequestNodes {
	s.Hostname = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetImageId(v string) *ReimageNodesRequestNodes {
	s.ImageId = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetLoginPassword(v string) *ReimageNodesRequestNodes {
	s.LoginPassword = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetNodeId(v string) *ReimageNodesRequestNodes {
	s.NodeId = &v
	return s
}

type ReimageNodesShrinkRequest struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool   `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NodesShrink           *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ReimageNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesShrinkRequest) SetClusterId(v string) *ReimageNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetNodesShrink(v string) *ReimageNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

type ReimageNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReimageNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponseBody) SetRequestId(v string) *ReimageNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReimageNodesResponseBody) SetTaskId(v string) *ReimageNodesResponseBody {
	s.TaskId = &v
	return s
}

type ReimageNodesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReimageNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReimageNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesResponse) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponse) SetHeaders(v map[string]*string) *ReimageNodesResponse {
	s.Headers = v
	return s
}

func (s *ReimageNodesResponse) SetStatusCode(v int32) *ReimageNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReimageNodesResponse) SetBody(v *ReimageNodesResponseBody) *ReimageNodesResponse {
	s.Body = v
	return s
}

type ShrinkClusterRequest struct {
	ClusterId             *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool                             `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NodeGroups            []*ShrinkClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequest) SetClusterId(v string) *ShrinkClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterRequest) SetNodeGroups(v []*ShrinkClusterRequestNodeGroups) *ShrinkClusterRequest {
	s.NodeGroups = v
	return s
}

type ShrinkClusterRequestNodeGroups struct {
	NodeGroupId *string                                `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	Nodes       []*ShrinkClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroups) SetNodeGroupId(v string) *ShrinkClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) SetNodes(v []*ShrinkClusterRequestNodeGroupsNodes) *ShrinkClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

type ShrinkClusterRequestNodeGroupsNodes struct {
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ShrinkClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroupsNodes) SetNodeId(v string) *ShrinkClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

type ShrinkClusterShrinkRequest struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IgnoreFailedNodeTasks *bool   `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	NodeGroupsShrink      *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
}

func (s ShrinkClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterShrinkRequest) SetClusterId(v string) *ShrinkClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetNodeGroupsShrink(v string) *ShrinkClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

type ShrinkClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ShrinkClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponseBody) SetRequestId(v string) *ShrinkClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkClusterResponseBody) SetTaskId(v string) *ShrinkClusterResponseBody {
	s.TaskId = &v
	return s
}

type ShrinkClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShrinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShrinkClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterResponse) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponse) SetHeaders(v map[string]*string) *ShrinkClusterResponse {
	s.Headers = v
	return s
}

func (s *ShrinkClusterResponse) SetStatusCode(v int32) *ShrinkClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ShrinkClusterResponse) SetBody(v *ShrinkClusterResponseBody) *ShrinkClusterResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eflo-controller"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Networks)) {
		request.NetworksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Networks, tea.String("Networks"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterDescription)) {
		body["ClusterDescription"] = request.ClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		body["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		body["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NetworksShrink)) {
		body["Networks"] = request.NetworksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeWithOptions(request *DescribeNodeRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNode"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNode(request *DescribeNodeRequest) (_result *DescribeNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeResponse{}
	_body, _err := client.DescribeNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskWithOptions(request *DescribeTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTask(request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtendClusterWithOptions(tmpReq *ExtendClusterRequest, runtime *util.RuntimeOptions) (_result *ExtendClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExtendClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IpAllocationPolicy)) {
		request.IpAllocationPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpAllocationPolicy, tea.String("IpAllocationPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VpdSubnets)) {
		request.VpdSubnetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpdSubnets, tea.String("VpdSubnets"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.IpAllocationPolicyShrink)) {
		body["IpAllocationPolicy"] = request.IpAllocationPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpdSubnetsShrink)) {
		body["VpdSubnets"] = request.VpdSubnetsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtendCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtendClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtendCluster(request *ExtendClusterRequest) (_result *ExtendClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtendClusterResponse{}
	_body, _err := client.ExtendClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterNodesWithOptions(request *ListClusterNodesRequest, runtime *util.RuntimeOptions) (_result *ListClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterNodes(request *ListClusterNodesRequest) (_result *ListClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterNodesResponse{}
	_body, _err := client.ListClusterNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFreeNodesWithOptions(request *ListFreeNodesRequest, runtime *util.RuntimeOptions) (_result *ListFreeNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MachineType)) {
		body["MachineType"] = request.MachineType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFreeNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFreeNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFreeNodes(request *ListFreeNodesRequest) (_result *ListFreeNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFreeNodesResponse{}
	_body, _err := client.ListFreeNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootNodesWithOptions(tmpReq *RebootNodesRequest, runtime *util.RuntimeOptions) (_result *RebootNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RebootNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodesShrink)) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootNodes(request *RebootNodesRequest) (_result *RebootNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootNodesResponse{}
	_body, _err := client.RebootNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReimageNodesWithOptions(tmpReq *ReimageNodesRequest, runtime *util.RuntimeOptions) (_result *ReimageNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReimageNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodesShrink)) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReimageNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReimageNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReimageNodes(request *ReimageNodesRequest) (_result *ReimageNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReimageNodesResponse{}
	_body, _err := client.ReimageNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShrinkClusterWithOptions(tmpReq *ShrinkClusterRequest, runtime *util.RuntimeOptions) (_result *ShrinkClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ShrinkClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ShrinkCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ShrinkClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShrinkCluster(request *ShrinkClusterRequest) (_result *ShrinkClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ShrinkClusterResponse{}
	_body, _err := client.ShrinkClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
