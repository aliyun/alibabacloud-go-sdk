// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApproveOperationRequest struct {
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The O\\&M operation type
	//
	// Valid value:
	//
	// 	- RepairMachine
	//
	// example:
	//
	// RepairMachine
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s ApproveOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationRequest) GoString() string {
	return s.String()
}

func (s *ApproveOperationRequest) SetNodeId(v string) *ApproveOperationRequest {
	s.NodeId = &v
	return s
}

func (s *ApproveOperationRequest) SetOperationType(v string) *ApproveOperationRequest {
	s.OperationType = &v
	return s
}

type ApproveOperationResponseBody struct {
	// The error message.
	//
	// example:
	//
	// Resource not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponseBody) SetErrorMessage(v string) *ApproveOperationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOperationResponseBody) SetRequestId(v string) *ApproveOperationResponseBody {
	s.RequestId = &v
	return s
}

type ApproveOperationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationResponse) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponse) SetHeaders(v map[string]*string) *ApproveOperationResponse {
	s.Headers = v
	return s
}

func (s *ApproveOperationResponse) SetStatusCode(v int32) *ApproveOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOperationResponse) SetBody(v *ApproveOperationResponseBody) *ApproveOperationResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the resource group into which you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzyqdwnfabx6q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i118099391667548921125
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Node
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CloseSessionRequest struct {
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
}

func (s CloseSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseSessionRequest) SetSessionId(v string) *CloseSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseSessionRequest) SetSessionToken(v string) *CloseSessionRequest {
	s.SessionToken = &v
	return s
}

type CloseSessionResponseBody struct {
	// example:
	//
	// 07AA3A1F-321E-50D8-B834-88C411331C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// i206495551737511455528
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// ClosingActive
	//
	// example:
	//
	// Inactive
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CloseSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseSessionResponseBody) SetRequestId(v string) *CloseSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseSessionResponseBody) SetSessionId(v string) *CloseSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CloseSessionResponseBody) SetState(v string) *CloseSessionResponseBody {
	s.State = &v
	return s
}

type CloseSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseSessionResponse) SetHeaders(v map[string]*string) *CloseSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseSessionResponse) SetStatusCode(v int32) *CloseSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseSessionResponse) SetBody(v *CloseSessionResponseBody) *CloseSessionResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// The cluster description.
	//
	// example:
	//
	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// Lite
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The components (software instance).
	Components []*CreateClusterRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The network information.
	Networks *CreateClusterRequestNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// The node vSwitch.
	NimizVSwitches []*string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty" type:"Repeated"`
	// The node groups.
	NodeGroups []*CreateClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// Specifies whether the elastic network interface (ENI) supports the Jumbo Frames feature.
	//
	// example:
	//
	// false
	OpenEniJumboFrame *bool `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags.
	Tag []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateClusterRequest) SetHpnZone(v string) *CreateClusterRequest {
	s.HpnZone = &v
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

func (s *CreateClusterRequest) SetNimizVSwitches(v []*string) *CreateClusterRequest {
	s.NimizVSwitches = v
	return s
}

func (s *CreateClusterRequest) SetNodeGroups(v []*CreateClusterRequestNodeGroups) *CreateClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *CreateClusterRequest) SetOpenEniJumboFrame(v bool) *CreateClusterRequest {
	s.OpenEniJumboFrame = &v
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
	// The component configurations.
	ComponentConfig *CreateClusterRequestComponentsComponentConfig `json:"ComponentConfig,omitempty" xml:"ComponentConfig,omitempty" type:"Struct"`
	// The component type.
	//
	// Valid values:
	//
	// 	- ARMS
	//
	// 	- ACKEdge
	//
	// example:
	//
	// ACKEdge
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
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
	// The basic parameters of the component.
	//
	// example:
	//
	// {
	//
	//       "EndpointPublicAccess": false,
	//
	//       "ContainerCidr": "10.4.0.0/24",
	//
	//       "KeyPair": "test",
	//
	//       "NodeCidrMask": "25",
	//
	//       "ResourceGroupId": "rg-axsadm3sdzsdvdsndstdisd",
	//
	//       "WorkerSystemDiskCategory": "da",
	//
	//       "WorkerSystemDiskSize": 40,
	//
	//       "DeletionProtection": false,
	//
	//       "KubeProxy": "iptables",
	//
	//       "Name": "da",
	//
	//       "LoadBalancerSpec": "slb.s1.small",
	//
	//       "Runtime": {
	//
	//             "Version": "19.03.15",
	//
	//             "Name": "docker"
	//
	//       },
	//
	//       "IsEnterpriseSecurityGroup": true,
	//
	//       "Vpcid": "192.168.23.0/24",
	//
	//       "NumOfNodes": 1,
	//
	//       "VswitchIds": [
	//
	//             "dad"
	//
	//       ],
	//
	//       "ServiceCidr": "10.0.0.0/16",
	//
	//       "SnatEntry": false,
	//
	//       "kubernetesVersion": "1.20.11-aliyunedge.1",
	//
	//       "WorkerInstanceTypes": [
	//
	//             "da"
	//
	//       ]
	//
	// }
	BasicArgs interface{} `json:"BasicArgs,omitempty" xml:"BasicArgs,omitempty"`
	// The node pool configurations, which are used to establish the mappings between node groups and node pools. This parameter is required when ComponentType is set to ACKEdge. Otherwise, this parameter is left empty.
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
	// The IP allocation policy.
	IpAllocationPolicy []*CreateClusterRequestNetworksIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	// The virtual private domain (VPD) configuration information.
	NewVpdInfo *CreateClusterRequestNetworksNewVpdInfo `json:"NewVpdInfo,omitempty" xml:"NewVpdInfo,omitempty" type:"Struct"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1d3dvbh9by7j5rujax
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IP version.
	//
	// example:
	//
	// IPv4
	TailIpVersion *string `json:"TailIpVersion,omitempty" xml:"TailIpVersion,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-asjdfklj
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-0jl36lqzmc06qogy0t5ll
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPD information.
	VpdInfo *CreateClusterRequestNetworksVpdInfo `json:"VpdInfo,omitempty" xml:"VpdInfo,omitempty" type:"Struct"`
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

func (s *CreateClusterRequestNetworks) SetSecurityGroupId(v string) *CreateClusterRequestNetworks {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetTailIpVersion(v string) *CreateClusterRequestNetworks {
	s.TailIpVersion = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVSwitchId(v string) *CreateClusterRequestNetworks {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVSwitchZoneId(v string) *CreateClusterRequestNetworks {
	s.VSwitchZoneId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVpcId(v string) *CreateClusterRequestNetworks {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVpdInfo(v *CreateClusterRequestNetworksVpdInfo) *CreateClusterRequestNetworks {
	s.VpdInfo = v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicy struct {
	// The bond policy.
	BondPolicy *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// The allocation policy for the instance type.
	MachineTypePolicy []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// The node allocation policy.
	NodePolicy []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
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
	// The default bond subnet for the cluster.
	//
	// example:
	//
	// 172.168.0.0/24
	BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	// The bond information.
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// 172.16.0.0/24
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
	// The bond information.
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga8n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// 192.168.1.0/24
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
	// The bond information.
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-2r42vq62001
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// 10.0.0.0/24
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
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-1gb1eftc5qp2ao75fo
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The CIDR block for Cloud Link.
	//
	// example:
	//
	// 172.16.0.0/24
	CloudLinkCidr *string `json:"CloudLinkCidr,omitempty" xml:"CloudLinkCidr,omitempty"`
	// The Cloud Link ID.
	//
	// example:
	//
	// vcc-cn-c4dtycm5i08
	CloudLinkId *string `json:"CloudLinkId,omitempty" xml:"CloudLinkId,omitempty"`
	// The VPC.
	//
	// example:
	//
	// vpc-0jl2x45apm6odc2c10h25
	MonitorVpcId *string `json:"MonitorVpcId,omitempty" xml:"MonitorVpcId,omitempty"`
	// The vSwitch.
	//
	// example:
	//
	// vsw-0jl2w3ffbghkss0x2foff
	MonitorVswitchId *string `json:"MonitorVswitchId,omitempty" xml:"MonitorVswitchId,omitempty"`
	// The CIDR block for the cluster.
	//
	// example:
	//
	// 192.168.0.0/16
	VpdCidr *string `json:"VpdCidr,omitempty" xml:"VpdCidr,omitempty"`
	// The cluster subnet.
	VpdSubnets []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
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
	// The subnet CIDR block.
	//
	// example:
	//
	// 10.0.1.8/24
	SubnetCidr *string `json:"SubnetCidr,omitempty" xml:"SubnetCidr,omitempty"`
	// The subnet type.
	//
	// example:
	//
	// 10.0.2.8/24
	SubnetType *string `json:"SubnetType,omitempty" xml:"SubnetType,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The VPC ID.
	//
	// example:
	//
	// vpd-vfuz6ejv
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The IDs of the subnets for a cluster.
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
	// Specifies whether to support file system mounting.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The system image ID.
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the AccessKey pair.
	//
	// example:
	//
	// sc-key
	KeyPairName   *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The node group description.
	//
	// example:
	//
	// Node group description
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// The node group name.
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The nodes.
	Nodes []*CreateClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The system disk information.
	SystemDisk *CreateClusterRequestNodeGroupsSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The user data of the instance. The user data must be Base64-encoded. The raw data can be up to 16 KB in size.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroups) SetFileSystemMountEnabled(v bool) *CreateClusterRequestNodeGroups {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetImageId(v string) *CreateClusterRequestNodeGroups {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetKeyPairName(v string) *CreateClusterRequestNodeGroups {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetLoginPassword(v string) *CreateClusterRequestNodeGroups {
	s.LoginPassword = &v
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

func (s *CreateClusterRequestNodeGroups) SetSystemDisk(v *CreateClusterRequestNodeGroupsSystemDisk) *CreateClusterRequestNodeGroups {
	s.SystemDisk = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetUserData(v string) *CreateClusterRequestNodeGroups {
	s.UserData = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetZoneId(v string) *CreateClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestNodeGroupsNodes struct {
	DataDisk []*CreateClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// 8d13b784-17a9-11ed-bc7b-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The logon password.
	//
	// example:
	//
	// ***
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01poc-cn-i7m2wnivf0d
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp169pi5fj151rrms4sia
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-0jlasms92fdxqd3wlf8ny
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsNodes) SetDataDisk(v []*CreateClusterRequestNodeGroupsNodesDataDisk) *CreateClusterRequestNodeGroupsNodes {
	s.DataDisk = v
	return s
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

func (s *CreateClusterRequestNodeGroupsNodes) SetVSwitchId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetVpcId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.VpcId = &v
	return s
}

type CreateClusterRequestNodeGroupsNodesDataDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithNode   *bool   `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsNodesDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodesDataDisk) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetCategory(v string) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetDeleteWithNode(v bool) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.DeleteWithNode = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetPerformanceLevel(v string) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetSize(v int32) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.Size = &v
	return s
}

type CreateClusterRequestNodeGroupsSystemDisk struct {
	// The disk category. Valid values:
	//
	// 	- cloud_essd
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Unit: GB
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetCategory(v string) *CreateClusterRequestNodeGroupsSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetPerformanceLevel(v string) *CreateClusterRequestNodeGroupsSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetSize(v int32) *CreateClusterRequestNodeGroupsSystemDisk {
	s.Size = &v
	return s
}

type CreateClusterRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env-name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
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
	// The cluster description.
	//
	// example:
	//
	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// Lite
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The components (software instance).
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The network information.
	NetworksShrink *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
	// The node vSwitch.
	NimizVSwitchesShrink *string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty"`
	// The node groups.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// Specifies whether the elastic network interface (ENI) supports the Jumbo Frames feature.
	//
	// example:
	//
	// false
	OpenEniJumboFrame *bool `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags.
	Tag []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateClusterShrinkRequest) SetHpnZone(v string) *CreateClusterShrinkRequest {
	s.HpnZone = &v
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

func (s *CreateClusterShrinkRequest) SetNimizVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.NimizVSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetOpenEniJumboFrame(v bool) *CreateClusterShrinkRequest {
	s.OpenEniJumboFrame = &v
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
	// The tag key.
	//
	// example:
	//
	// env-name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
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
	// The cluster ID.
	//
	// example:
	//
	// i116913051663373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i159809891662373011015
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateDiagnosticTaskRequest struct {
	// The log information.
	AiJobLogInfo *CreateDiagnosticTaskRequestAiJobLogInfo `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The diagnostics type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The IDs of the nodes.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s CreateDiagnosticTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequest) SetAiJobLogInfo(v *CreateDiagnosticTaskRequestAiJobLogInfo) *CreateDiagnosticTaskRequest {
	s.AiJobLogInfo = v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetClusterId(v string) *CreateDiagnosticTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetNodeIds(v []*string) *CreateDiagnosticTaskRequest {
	s.NodeIds = v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfo struct {
	// The task logs.
	AiJobLogs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs `json:"AiJobLogs,omitempty" xml:"AiJobLogs,omitempty" type:"Repeated"`
	// The end time. The value is in the timestamp format. Unit: seconds.
	//
	// >  This timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time. The value is in the timestamp format. Unit: seconds.
	//
	// >  This timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 2024-10-11T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetAiJobLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.AiJobLogs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetEndTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetStartTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.StartTime = &v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs struct {
	// The instance ID.
	//
	// example:
	//
	// null
	AiInstance *string `json:"AiInstance,omitempty" xml:"AiInstance,omitempty"`
	// The logs.
	Logs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-p2p2al5u1hn
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetAiInstance(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.AiInstance = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.Logs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetNodeId(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.NodeId = &v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs struct {
	// The sending date in the yyyymmdd format.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	Datetime *string `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	// The log content.
	//
	// example:
	//
	// success
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetDatetime(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.Datetime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetLogContent(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.LogContent = &v
	return s
}

type CreateDiagnosticTaskShrinkRequest struct {
	// The log information.
	AiJobLogInfoShrink *string `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The diagnostics type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The IDs of the nodes.
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s CreateDiagnosticTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskShrinkRequest) SetAiJobLogInfoShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.AiJobLogInfoShrink = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetClusterId(v string) *CreateDiagnosticTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskShrinkRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetNodeIdsShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

type CreateDiagnosticTaskResponseBody struct {
	// The ID of the diagnostics task.
	//
	// example:
	//
	// diag-i150553931717380274931
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A511C02A-0127-51AA-A9F9-966382C9A1B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosticTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponseBody) SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) SetRequestId(v string) *CreateDiagnosticTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiagnosticTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponse) SetHeaders(v map[string]*string) *CreateDiagnosticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetStatusCode(v int32) *CreateDiagnosticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetBody(v *CreateDiagnosticTaskResponseBody) *CreateDiagnosticTaskResponse {
	s.Body = v
	return s
}

type CreateNetTestTaskRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specify when NetTestType is CommTest.
	CommTest *CreateNetTestTaskRequestCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// Specify when NetTestType is DelayTest.
	DelayTest *CreateNetTestTaskRequestDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// The type of the network test. Valid values: DelayTest, TrafficTest, and CommTest.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The network mode.
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// If the TrafficModel is Fullmesh, leave this parameter empty.
	TrafficTest *CreateNetTestTaskRequestTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s CreateNetTestTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequest) SetClusterId(v string) *CreateNetTestTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetClusterName(v string) *CreateNetTestTaskRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetCommTest(v *CreateNetTestTaskRequestCommTest) *CreateNetTestTaskRequest {
	s.CommTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetDelayTest(v *CreateNetTestTaskRequestDelayTest) *CreateNetTestTaskRequest {
	s.DelayTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetTestType(v string) *CreateNetTestTaskRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetworkMode(v string) *CreateNetTestTaskRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetPort(v string) *CreateNetTestTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetTrafficTest(v *CreateNetTestTaskRequestTrafficTest) *CreateNetTestTaskRequest {
	s.TrafficTest = v
	return s
}

type CreateNetTestTaskRequestCommTest struct {
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GPUNum *int64 `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// The host IDs.
	Hosts []*CreateNetTestTaskRequestCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The communication library model.
	//
	// example:
	//
	// intention_v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The CommTest type, which can be ACCL or NCCL.
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetTestTaskRequestCommTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTest) SetGPUNum(v int64) *CreateNetTestTaskRequestCommTest {
	s.GPUNum = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetHosts(v []*CreateNetTestTaskRequestCommTestHosts) *CreateNetTestTaskRequestCommTest {
	s.Hosts = v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetModel(v string) *CreateNetTestTaskRequestCommTest {
	s.Model = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetType(v string) *CreateNetTestTaskRequestCommTest {
	s.Type = &v
	return s
}

type CreateNetTestTaskRequestCommTestHosts struct {
	// The IP address.
	//
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetIP(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetServerName(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestDelayTest struct {
	// The hosts of the test node.
	Hosts []*CreateNetTestTaskRequestDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s CreateNetTestTaskRequestDelayTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTest) SetHosts(v []*CreateNetTestTaskRequestDelayTestHosts) *CreateNetTestTaskRequestDelayTest {
	s.Hosts = v
	return s
}

type CreateNetTestTaskRequestDelayTestHosts struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-fou43an0a05
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetBond(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetIP(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetServerName(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestTrafficTest struct {
	// The client IDs.
	Clients []*CreateNetTestTaskRequestTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The running duration of the pipeline job. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// If the protocol is RDMA, enter True or False. If the protocol is TCP, leave this field empty.
	//
	// example:
	//
	// False
	GDR *bool `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// The network protocol, which can be RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// If the protocol is TCP, enter the number of concurrent connections. If the protocol is RDMA, enter the configured QP value.
	//
	// example:
	//
	// 1
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// The services.
	Servers []*CreateNetTestTaskRequestTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The traffic model, which can be MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTest) SetClients(v []*CreateNetTestTaskRequestTrafficTestClients) *CreateNetTestTaskRequestTrafficTest {
	s.Clients = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetDuration(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.Duration = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetGDR(v bool) *CreateNetTestTaskRequestTrafficTest {
	s.GDR = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetProtocol(v string) *CreateNetTestTaskRequestTrafficTest {
	s.Protocol = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetQP(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.QP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetServers(v []*CreateNetTestTaskRequestTrafficTestServers) *CreateNetTestTaskRequestTrafficTest {
	s.Servers = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetTrafficModel(v string) *CreateNetTestTaskRequestTrafficTest {
	s.TrafficModel = &v
	return s
}

type CreateNetTestTaskRequestTrafficTestClients struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-w5elqg7pw18
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-20s41p6cx01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestClients) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetBond(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetIP(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestTrafficTestServers struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestServers) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetBond(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetIP(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specify when NetTestType is CommTest.
	CommTestShrink *string `json:"CommTest,omitempty" xml:"CommTest,omitempty"`
	// Specify when NetTestType is DelayTest.
	DelayTestShrink *string `json:"DelayTest,omitempty" xml:"DelayTest,omitempty"`
	// The type of the network test. Valid values: DelayTest, TrafficTest, and CommTest.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The network mode.
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// If the TrafficModel is Fullmesh, leave this parameter empty.
	TrafficTestShrink *string `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty"`
}

func (s CreateNetTestTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterId(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterName(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetCommTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.CommTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetDelayTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.DelayTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetTestType(v string) *CreateNetTestTaskShrinkRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetworkMode(v string) *CreateNetTestTaskShrinkRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetPort(v string) *CreateNetTestTaskShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetTrafficTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.TrafficTestShrink = &v
	return s
}

type CreateNetTestTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the test task. The unique identifier of a network test task.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s CreateNetTestTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponseBody) SetRequestId(v string) *CreateNetTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) SetTestId(v string) *CreateNetTestTaskResponseBody {
	s.TestId = &v
	return s
}

type CreateNetTestTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetTestTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponse) SetHeaders(v map[string]*string) *CreateNetTestTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateNetTestTaskResponse) SetStatusCode(v int32) *CreateNetTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetTestTaskResponse) SetBody(v *CreateNetTestTaskResponseBody) *CreateNetTestTaskResponse {
	s.Body = v
	return s
}

type CreateNodeGroupRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NodeGroup *CreateNodeGroupRequestNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Struct"`
	NodeUnit  map[string]interface{}           `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) SetClusterId(v string) *CreateNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeGroup(v *CreateNodeGroupRequestNodeGroup) *CreateNodeGroupRequest {
	s.NodeGroup = v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeUnit(v map[string]interface{}) *CreateNodeGroupRequest {
	s.NodeUnit = v
	return s
}

type CreateNodeGroupRequestNodeGroup struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	Az *string `json:"Az,omitempty" xml:"Az,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i191887641687336652616
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// test-keypair
	KeyPairName   *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-machine-type3
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// describe for node group
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PAI-LINGJUN
	NodeGroupName *string                                    `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	SystemDisk    *CreateNodeGroupRequestNodeGroupSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroup) SetAz(v string) *CreateNodeGroupRequestNodeGroup {
	s.Az = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetFileSystemMountEnabled(v bool) *CreateNodeGroupRequestNodeGroup {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetImageId(v string) *CreateNodeGroupRequestNodeGroup {
	s.ImageId = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetKeyPairName(v string) *CreateNodeGroupRequestNodeGroup {
	s.KeyPairName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetLoginPassword(v string) *CreateNodeGroupRequestNodeGroup {
	s.LoginPassword = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetMachineType(v string) *CreateNodeGroupRequestNodeGroup {
	s.MachineType = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupDescription(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupDescription = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupName(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetSystemDisk(v *CreateNodeGroupRequestNodeGroupSystemDisk) *CreateNodeGroupRequestNodeGroup {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetUserData(v string) *CreateNodeGroupRequestNodeGroup {
	s.UserData = &v
	return s
}

type CreateNodeGroupRequestNodeGroupSystemDisk struct {
	// *
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// *
	//
	// *
	//
	// example:
	//
	// PL!
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 250
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetCategory(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetPerformanceLevel(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetSize(v int32) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Size = &v
	return s
}

type CreateNodeGroupShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NodeGroupShrink *string `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	NodeUnitShrink  *string `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupShrinkRequest) SetClusterId(v string) *CreateNodeGroupShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeGroupShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeGroupShrink = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeUnitShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeUnitShrink = &v
	return s
}

type CreateNodeGroupResponseBody struct {
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupId(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupName(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetRequestId(v string) *CreateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponse) SetHeaders(v map[string]*string) *CreateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeGroupResponse) SetStatusCode(v int32) *CreateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeGroupResponse) SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse {
	s.Body = v
	return s
}

type CreateSessionRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The type of the session corresponding to the session package.
	//
	// example:
	//
	// Valid values: Sol (default): based on serial port Assistant: based on cloud assistant
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The start time. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1669340937156
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) SetNodeId(v string) *CreateSessionRequest {
	s.NodeId = &v
	return s
}

func (s *CreateSessionRequest) SetSessionType(v string) *CreateSessionRequest {
	s.SessionType = &v
	return s
}

func (s *CreateSessionRequest) SetStartTime(v string) *CreateSessionRequest {
	s.StartTime = &v
	return s
}

type CreateSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 2A59143F1
	ServerSn *string `json:"ServerSn,omitempty" xml:"ServerSn,omitempty"`
	// The session ID.
	//
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session credential.
	//
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	// The WebSocket address.
	//
	// example:
	//
	// ws://x.x.x.x:xx/calypso_web_console
	WssEndpoint *string `json:"WssEndpoint,omitempty" xml:"WssEndpoint,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponseBody) SetServerSn(v string) *CreateSessionResponseBody {
	s.ServerSn = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionId(v string) *CreateSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionToken(v string) *CreateSessionResponseBody {
	s.SessionToken = &v
	return s
}

func (s *CreateSessionResponseBody) SetWssEndpoint(v string) *CreateSessionResponseBody {
	s.WssEndpoint = &v
	return s
}

type CreateSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) SetHeaders(v map[string]*string) *CreateSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionResponse) SetStatusCode(v int32) *CreateSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSessionResponse) SetBody(v *CreateSessionResponseBody) *CreateSessionResponse {
	s.Body = v
	return s
}

type CreateVscRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags.
	Tag []*CreateVscRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The custom name of the VSC, which is unique on a compute node.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type. Valid values: primary and standard. Default value: primary.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s CreateVscRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVscRequest) GoString() string {
	return s.String()
}

func (s *CreateVscRequest) SetClientToken(v string) *CreateVscRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVscRequest) SetNodeId(v string) *CreateVscRequest {
	s.NodeId = &v
	return s
}

func (s *CreateVscRequest) SetResourceGroupId(v string) *CreateVscRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVscRequest) SetTag(v []*CreateVscRequestTag) *CreateVscRequest {
	s.Tag = v
	return s
}

func (s *CreateVscRequest) SetVscName(v string) *CreateVscRequest {
	s.VscName = &v
	return s
}

func (s *CreateVscRequest) SetVscType(v string) *CreateVscRequest {
	s.VscType = &v
	return s
}

type CreateVscRequestTag struct {
	// The resource tag key.
	//
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The resource tag value.
	//
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVscRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVscRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVscRequestTag) SetKey(v string) *CreateVscRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVscRequestTag) SetValue(v string) *CreateVscRequestTag {
	s.Value = &v
	return s
}

type CreateVscResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s CreateVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVscResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscResponseBody) SetRequestId(v string) *CreateVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVscResponseBody) SetVscId(v string) *CreateVscResponseBody {
	s.VscId = &v
	return s
}

type CreateVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVscResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVscResponse) GoString() string {
	return s.String()
}

func (s *CreateVscResponse) SetHeaders(v map[string]*string) *CreateVscResponse {
	s.Headers = v
	return s
}

func (s *CreateVscResponse) SetStatusCode(v int32) *CreateVscResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVscResponse) SetBody(v *CreateVscResponseBody) *CreateVscResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i116913051662373010974
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
	// The request ID.
	//
	// example:
	//
	// 0FC4A1C7-421C-5EAB-9361-4C0338EFA287
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteNodeGroupRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i114444141733395242745
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// i121824791737080429819
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) SetClusterId(v string) *DeleteNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type DeleteNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponse) SetHeaders(v map[string]*string) *DeleteNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeGroupResponse) SetStatusCode(v int32) *DeleteNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeGroupResponse) SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse {
	s.Body = v
	return s
}

type DeleteVscRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the VSC that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

type DeleteVscResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscResponseBody) SetRequestId(v string) *DeleteVscResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVscResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscResponse) SetHeaders(v map[string]*string) *DeleteVscResponse {
	s.Headers = v
	return s
}

func (s *DeleteVscResponse) SetStatusCode(v int32) *DeleteVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVscResponse) SetBody(v *DeleteVscResponseBody) *DeleteVscResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	// 集群id。
	//
	// This parameter is required.
	//
	// example:
	//
	// i119982311660892626523
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
	// 集群描述
	//
	// example:
	//
	// 测试集群
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// 集群id
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群名称
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 集群类型
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// 组件信息
	Components []*DescribeClusterResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// 计算网络的IP类型
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2022-06-08T07:05:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集群编号
	//
	// example:
	//
	// A2
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// 网络信息
	Networks *DescribeClusterResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// 节点数
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// 节点组数量
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// 网络接口巨帧
	//
	// example:
	//
	// unsupported
	OpenEniJumboFrame *string `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// 集群状态
	//
	// example:
	//
	// running
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// 请求id。
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源组id
	//
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 任务id
	//
	// example:
	//
	// i152609221670466904596
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2022-08-23T06:36:17.000Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 专有网络ID
	//
	// example:
	//
	// vpc-0jlkqysom5dmcviymep3f
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeClusterResponseBody) SetComputingIpVersion(v string) *DescribeClusterResponseBody {
	s.ComputingIpVersion = &v
	return s
}

func (s *DescribeClusterResponseBody) SetCreateTime(v string) *DescribeClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBody) SetHpnZone(v string) *DescribeClusterResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeClusterResponseBody) SetNetworks(v *DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody {
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

func (s *DescribeClusterResponseBody) SetOpenEniJumboFrame(v string) *DescribeClusterResponseBody {
	s.OpenEniJumboFrame = &v
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

func (s *DescribeClusterResponseBody) SetVpcId(v string) *DescribeClusterResponseBody {
	s.VpcId = &v
	return s
}

type DescribeClusterResponseBodyComponents struct {
	// 组件id
	//
	// example:
	//
	// i149549021660892626529
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// 组件类型
	//
	// example:
	//
	// ACKEdge
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
	// 集群网段id
	//
	// example:
	//
	// vpd-iqd7xunc
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDiagnosticResultRequest struct {
	// The diagnostic task ID.
	//
	// example:
	//
	// diag-i151942361720577788844
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
}

func (s DescribeDiagnosticResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultRequest) SetDiagnosticId(v string) *DescribeDiagnosticResultRequest {
	s.DiagnosticId = &v
	return s
}

type DescribeDiagnosticResultResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-06-15T10:17:56
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The diagnostic task ID.
	//
	// example:
	//
	// diag-i155363241720059671316
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// The diagnostic information.
	DiagnosticResults []interface{} `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// The diagnostic status.
	//
	// example:
	//
	// Fault
	DiagnosticState *string `json:"DiagnosticState,omitempty" xml:"DiagnosticState,omitempty"`
	// The type of the diagnostic task.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The end time of the instance exception. The time format with time zone based on the ISO8601 standard. The format is yyyy-MM-ddTHH:mm:ss +0800.
	//
	// example:
	//
	// 2024-06-11T10:00:30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponseBody) SetClusterId(v string) *DescribeDiagnosticResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetCreatedTime(v string) *DescribeDiagnosticResultResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticId(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticResults(v []interface{}) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticState(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticState = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticType(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticType = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetEndTime(v string) *DescribeDiagnosticResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetNodeIds(v []*string) *DescribeDiagnosticResultResponseBody {
	s.NodeIds = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetRequestId(v string) *DescribeDiagnosticResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiagnosticResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetStatusCode(v int32) *DescribeDiagnosticResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetBody(v *DescribeDiagnosticResultResponseBody) *DescribeDiagnosticResultResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	// The encoding mode of the `CommandContent` and `Output` response parameters. Valid values:
	//
	// 	- PlainText: returns the original command content and command outputs.
	//
	// 	- Base64 (default): returns the Base64-encoded command content and command output.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to return the command outputs in the response.
	//
	// 	- true: returns the command outputs. When this parameter is set to true, you must specify `InvokeId`, `InstanceId`, or both.
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-cd03crwys0lrls0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNodeId(v string) *DescribeInvocationsRequest {
	s.NodeId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	// The command execution record.
	Invocations *DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The file sending records.
	Invocation []*DescribeInvocationsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocation(v []*DescribeInvocationsResponseBodyInvocationsInvocation) *DescribeInvocationsResponseBodyInvocations {
	s.Invocation = v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocation struct {
	// The executed command.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command content is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command content is returned.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The command name.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The time when the command task was created.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The schedule on which the command was run.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The overall execution state of the command task. The value of this parameter depends on the execution states of the command task on all the involved instances. Valid values:
	//
	// 	- Pending: The command was being verified or sent. If the execution state on at least one instance is Pending, the overall execution state is Pending.
	//
	// 	- Scheduled: The command that is set to run on a schedule is sent and waiting to be run. If the execution state on at least one instance is Scheduled, the overall execution state is Scheduled.
	//
	// 	- Running: The command is being run on the instance. When the execution state on at least one instance is Running, the overall execution state is Running.
	//
	// 	- Success: When the execution state on at least one instance is Success and the execution state on the other instances is Stopped or Success, the overall execution state is Success.
	//
	//     	- One-time task: The execution is complete, and the exit code is 0.
	//
	//     	- Scheduled task: The last execution was complete, the exit code was 0, and the specified period ended.
	//
	// 	- Failed: When the execution state on all instances is Stopped or Failed, the overall execution state is Failed. When the execution state on an instance is one of the following values, Failed is returned as the overall execution state:
	//
	//     	- Invalid: The command is invalid.
	//
	//     	- Aborted: The command failed to be sent.
	//
	//     	- Failed: The execution was complete, but the exit code was not 0.
	//
	//     	- Timeout: The execution timed out.
	//
	//     	- Error: An error occurred while the command was being run.
	//
	// 	- Stopping: The command task is being stopped. When the execution state on at least one instance is Stopping, the overall execution state is Stopping.
	//
	// 	- Stopped: The task was stopped. When the execution state on all instances is Stopped, the overall execution state is Stopped. When the execution state on an instance is one of the following values, Stopped is returned as the overall execution state:
	//
	//     	- Cancelled: The task was canceled.
	//
	//     	- Terminated: The task was terminated.
	//
	// 	- PartialFailed: The execution was complete on some instances and failed on other instances. When the execution state is Success on some instances and is Failed or Stopped on the other instances, the overall execution state is PartialFailed.
	//
	// >  The value of the `InvokeStatus` response parameter is similar to the value of InvocationStatus. We recommend that you ignore InvokeStatus and check the value of InvocationStatus.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// t-ind3k9ytvvduoe8
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The command execution records.
	InvokeNodes *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The overall execution status of the command task. The value of this parameter depends on the execution states of the command task on all involved instances. Valid values:
	//
	// 	- Running:
	//
	//     	- Scheduled task: Before you stop the scheduled execution of the command, the overall execution state is always Running.
	//
	//     	- One-time task: If the command is being run on instances, the overall execution state is Running.
	//
	// 	- Finished:
	//
	//     	- Scheduled task: The overall execution state can never be Finished.
	//
	//     	- One-time task: The execution is complete on all instances, or the execution is stopped on some instances and is complete on the other instances.
	//
	// 	- Failed:
	//
	//     	- Scheduled task: The overall execution state can never be Failed.
	//
	//     	- One-time task: The execution failed on all instances.
	//
	// 	- Stopped: The task is stopped.
	//
	// 	- Stopping: The task is being stopped.
	//
	// 	- PartialFailed: The task fails on some instances. If you specify both this parameter and `InstanceId`, this parameter does not take effect.
	//
	// example:
	//
	// Running
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The custom parameters in the command.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The execution mode of the command. Valid values:
	//
	// 	- Once: The command is run immediately.
	//
	// 	- Period: The command is run on a schedule.
	//
	// 	- NextRebootOnly: The command is run the next time the instances start.
	//
	// 	- EveryReboot: runs the command every time the instances start.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The timeout period for the command execution. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username that is used to run the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working directory of the command on the instance.
	//
	// example:
	//
	// /home
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandDescription(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandDescription = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetFrequency(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Frequency = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetParameters(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Parameters = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetRepeatMode(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.RepeatMode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTimeout(v int32) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Timeout = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetUsername(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Username = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetWorkingDir(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.WorkingDir = &v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The command execution records of the node.
	InvokeNode []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The start time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The size of the Output text that was truncated and discarded because the Output value exceeded 24 KB in size.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The error code returned when the file failed to be sent to the instance. Valid values:
	//
	// 	- Null: The file is sent to the instance.
	//
	// 	- NodeNotExists: The specified instance does not exist or has been released.
	//
	// 	- NodeReleased: The instance was released while the file was being sent.
	//
	// 	- NodeNotRunning: The instance was not running while the file sending task was being created.
	//
	// 	- AccountNotExists: The specified account does not exist.
	//
	// 	- ClientNotRunning: Cloud Assistant Agent is not running.
	//
	// 	- ClientNotResponse: Cloud Assistant Agent does not respond.
	//
	// 	- ClientIsUpgrading: Cloud Assistant Agent is being upgraded.
	//
	// 	- ClientNeedUpgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// 	- DeliveryTimeout: The file sending task timed out.
	//
	// 	- FileCreateFail: The file failed to be created.
	//
	// 	- FileAlreadyExists: A file with the same name exists in the specified directory.
	//
	// 	- FileContentInvalid: The file content is invalid.
	//
	// 	- FileNameInvalid: The file name is invalid.
	//
	// 	- FilePathInvalid: The specified directory is invalid.
	//
	// 	- FileAuthorityInvalid: The specified permissions on the file are invalid.
	//
	// 	- UserGroupNotExists: The specified user group does not exist.
	//
	// example:
	//
	// NodeNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the command cannot be sent or run.
	//
	// 	- If this parameter is empty, the command was run as expected.
	//
	// 	- the specified node does not exists: The specified instance does not exist or is released.
	//
	// 	- the node has node when create task: The instance is released when the command is being run.
	//
	// 	- the node is not running when create task: The instance is not in the Running state while the command is being run.
	//
	// 	- the command is not applicable: The command is not applicable to the specified instance.
	//
	// 	- the specified account does not exists: The specified account does not exist.
	//
	// 	- the specified directory does not exists: The specified directory does not exist.
	//
	// 	- the cron job expression is invalid: The cron expression that specifies the execution time is invalid.
	//
	// 	- the aliyun service is not running on the instance: Cloud Assistant Agent is not running.
	//
	// 	- the aliyun service in the instance does not response: Cloud Assistant Agent does not respond.
	//
	// 	- the aliyun service in the node is upgrading now: Cloud Assistant Agent is being upgraded.
	//
	// 	- the aliyun service in the node need upgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// 	- the command delivery has been timeout: The request to send the command timed out.
	//
	// 	- the command execution has been timeout: The command execution timed out.
	//
	// 	- the command execution got an exception: An exception occurred when the command is being run.
	//
	// 	- the command execution has been interrupted: The command execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The command execution completes, but the exit code is not 0.
	//
	// 	- the specified node has been released: The instance is released while the file is being sent.
	//
	// example:
	//
	// the specified node does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the execution. Valid values:
	//
	// 	- For Linux instances, the value is the exit code of the shell process.
	//
	// 	- For Windows instances, the value is the exit code of the batch or PowerShell process.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The execution status of the command on a single instance. Valid values:
	//
	// 	- Pending: The command was being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command failed to be sent to the instance. To send a command to an instance, make sure that the instance is in the Running state and the command can be sent to the instance within 1 minute.
	//
	// 	- Running: The command is being run on the instance.
	//
	// 	- Success:
	//
	//     	- One-time task: The execution was complete, and the exit code was 0.
	//
	//     	- Recurring execution: The previous occurrence of the execution is complete, and the exit code is 0. The specified recurring period during which the command is run ends.
	//
	// 	- Failed:
	//
	//     	- One-time task: The execution was complete, but the exit code was not 0.
	//
	//     	- Recurring execution: The previous occurrence of the execution is complete, but the exit code is not 0. The specified recurring period during which the command is run is about to end.
	//
	// 	- Error: The execution cannot proceed due to an exception.
	//
	// 	- Timeout: The execution timed out.
	//
	// 	- Cancelled: The execution was canceled before it started.
	//
	// 	- Stopping: The command task is being stopped.
	//
	// 	- Terminated: The execution was terminated before completion.
	//
	// 	- Scheduled:
	//
	//     	- One-time task: The execution state can never be Scheduled.
	//
	//     	- Recurring execution: The command is waiting to be run.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-lbj36wkp70b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The execution status of the command on a single instance.
	//
	// example:
	//
	// Finished
	NodeInvokeStatus *string `json:"NodeInvokeStatus,omitempty" xml:"NodeInvokeStatus,omitempty"`
	// The command output.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command output is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command output is returned.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times that the command was run on the instance.
	//
	// 	- If the command is set to run only once, the value is 0 or 1.
	//
	// 	- If the command is set to run on a schedule, the value is the number of times that the command has been run on the instance.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the command task was stopped. If you call the StopInvocation operation to stop the command task, the value of this parameter is the time when the operation is called.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// Indicates whether the command is to be automatically run. Valid values:
	//
	// 	- true: The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Period`, `NextRebootOnly`, or `EveryReboot`.
	//
	// 	- false (default): The command meets the following requirements.
	//
	//     	- The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Once`.
	//
	//     	- The command task is canceled, stopped, or completed.
	//
	// example:
	//
	// false
	Timed *string `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The update time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetExitCode(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeInvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetTimed(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeNetTestResultRequest struct {
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s DescribeNetTestResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultRequest) SetTestId(v string) *DescribeNetTestResultRequest {
	s.TestId = &v
	return s
}

type DescribeNetTestResultResponseBody struct {
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// Standard_Cluster
	ClusterName *string                                    `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CommTest    *DescribeNetTestResultResponseBodyCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// example:
	//
	// 2024-10-15T10:25:42+08:00
	CreationTime *string                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DelayTest    *DescribeNetTestResultResponseBodyDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {}
	ResultDetial *string `json:"ResultDetial,omitempty" xml:"ResultDetial,omitempty"`
	// *
	//
	// *
	//
	// *
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// af35ce53-7c35-4277-834a-fbf49c316a96
	TestId      *string                                       `json:"TestId,omitempty" xml:"TestId,omitempty"`
	TrafficTest *DescribeNetTestResultResponseBodyTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s DescribeNetTestResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBody) SetClusterId(v string) *DescribeNetTestResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetClusterName(v string) *DescribeNetTestResultResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCommTest(v *DescribeNetTestResultResponseBodyCommTest) *DescribeNetTestResultResponseBody {
	s.CommTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCreationTime(v string) *DescribeNetTestResultResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetDelayTest(v *DescribeNetTestResultResponseBodyDelayTest) *DescribeNetTestResultResponseBody {
	s.DelayTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetFinishedTime(v string) *DescribeNetTestResultResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetNetTestType(v string) *DescribeNetTestResultResponseBody {
	s.NetTestType = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetPort(v string) *DescribeNetTestResultResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetRequestId(v string) *DescribeNetTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetResultDetial(v string) *DescribeNetTestResultResponseBody {
	s.ResultDetial = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetStatus(v string) *DescribeNetTestResultResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTestId(v string) *DescribeNetTestResultResponseBody {
	s.TestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTrafficTest(v *DescribeNetTestResultResponseBodyTrafficTest) *DescribeNetTestResultResponseBody {
	s.TrafficTest = v
	return s
}

type DescribeNetTestResultResponseBodyCommTest struct {
	// example:
	//
	// 1
	GPUNum *string                                           `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	Hosts  []*DescribeNetTestResultResponseBodyCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// example:
	//
	// intention_v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetTestResultResponseBodyCommTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetGPUNum(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.GPUNum = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetHosts(v []*DescribeNetTestResultResponseBodyCommTestHosts) *DescribeNetTestResultResponseBodyCommTest {
	s.Hosts = v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetModel(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Model = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetType(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Type = &v
	return s
}

type DescribeNetTestResultResponseBodyCommTestHosts struct {
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyDelayTest struct {
	Hosts []*DescribeNetTestResultResponseBodyDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeNetTestResultResponseBodyDelayTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTest) SetHosts(v []*DescribeNetTestResultResponseBodyDelayTestHosts) *DescribeNetTestResultResponseBodyDelayTest {
	s.Hosts = v
	return s
}

type DescribeNetTestResultResponseBodyDelayTestHosts struct {
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// example:
	//
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetBond(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTest struct {
	Clients []*DescribeNetTestResultResponseBodyTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// False
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 1
	QP      *int64                                                 `json:"QP,omitempty" xml:"QP,omitempty"`
	Servers []*DescribeNetTestResultResponseBodyTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetClients(v []*DescribeNetTestResultResponseBodyTrafficTestClients) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Clients = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetDuration(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Duration = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetGDR(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.GDR = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetProtocol(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Protocol = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetQP(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.QP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetServers(v []*DescribeNetTestResultResponseBodyTrafficTestServers) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Servers = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetTrafficModel(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.TrafficModel = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTestClients struct {
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// e01-cn-20s41p6cx01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTestServers struct {
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetTestResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponse) SetHeaders(v map[string]*string) *DescribeNetTestResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetTestResultResponse) SetStatusCode(v int32) *DescribeNetTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetTestResultResponse) SetBody(v *DescribeNetTestResultResponseBody) *DescribeNetTestResultResponse {
	s.Body = v
	return s
}

type DescribeNodeRequest struct {
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-sn-2060
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
	// The cluster ID.
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-09-30T03:35:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The disks.
	Disks []*DescribeNodeResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The expiration time.
	//
	// example:
	//
	// 2022-06-23T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The hostname.
	//
	// example:
	//
	// 31d38530-241e-11ed-bc63-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Centos7.9_all_0811
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The network information.
	Networks []*DescribeNodeResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node status.
	//
	// Valid values:
	//
	// 	- Extending
	//
	// 	- UnusedNodeStopped
	//
	// 	- UnusedNodeStopping
	//
	// 	- Unused
	//
	// 	- Using
	//
	// 	- ReleaseLocking
	//
	// 	- Operating
	//
	// 	- Cutting
	//
	// 	- ClusterNodeStopped
	//
	// 	- UnusedNodeRecovering
	//
	// 	- ClusterNodeStopping
	//
	// 	- ClusterNodeRecovering
	//
	// 	- Replacing
	//
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sag42ckf4jx
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The custom script.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeNodeResponseBody) SetDisks(v []*DescribeNodeResponseBodyDisks) *DescribeNodeResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeNodeResponseBody) SetExpiredTime(v string) *DescribeNodeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetFileSystemMountEnabled(v bool) *DescribeNodeResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHostname(v string) *DescribeNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHpnZone(v string) *DescribeNodeResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageId(v string) *DescribeNodeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageName(v string) *DescribeNodeResponseBody {
	s.ImageName = &v
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

func (s *DescribeNodeResponseBody) SetResourceGroupId(v string) *DescribeNodeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetSn(v string) *DescribeNodeResponseBody {
	s.Sn = &v
	return s
}

func (s *DescribeNodeResponseBody) SetUserData(v string) *DescribeNodeResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeNodeResponseBody) SetZoneId(v string) *DescribeNodeResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeNodeResponseBodyDisks struct {
	// The disk type. Valid values:
	//
	// 	- cloud_essd
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp1fi88ryk4yah8a6yos
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The disk type. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNodeResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyDisks) SetCategory(v string) *DescribeNodeResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetDiskId(v string) *DescribeNodeResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetPerformanceLevel(v string) *DescribeNodeResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetSize(v int32) *DescribeNodeResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetType(v string) *DescribeNodeResponseBodyDisks {
	s.Type = &v
	return s
}

type DescribeNodeResponseBodyNetworks struct {
	// The port information of the elastic network interface (ENI).
	//
	// example:
	//
	// Bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 47.254.235.44
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The subnet ID.
	//
	// example:
	//
	// vsw-uf68v51fldm5egmui5a6k
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The ID of the cluster network.
	//
	// example:
	//
	// vpd-xcuhjyrj
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The natural language that is used to filter responses. For more information, see RFC 7231.
	//
	// zh-CN en-US Default value: zh-CN.
	//
	// Valid values:
	//
	// 	- en-US
	//
	// 	- zh-CN
	//
	// example:
	//
	// zh-CN
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
	// The regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1D2FBB36-C39B-5EBB-9928-FCC1A236D65D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The region name.
	//
	// example:
	//
	// Hang Zhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeSendFileResultsRequest struct {
	// The command execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bj038i0d6r8zoqo
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSendFileResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequest) SetInvokeId(v string) *DescribeSendFileResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetNodeId(v string) *DescribeSendFileResultsRequest {
	s.NodeId = &v
	return s
}

type DescribeSendFileResultsResponseBody struct {
	// The file sending records.
	Invocations *DescribeSendFileResultsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the commands.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSendFileResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBody) SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetRequestId(v string) *DescribeSendFileResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetTotalCount(v string) *DescribeSendFileResultsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSendFileResultsResponseBodyInvocations struct {
	// The command execution ID.
	Invocation []*DescribeSendFileResultsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocations) SetInvocation(v []*DescribeSendFileResultsResponseBodyInvocationsInvocation) *DescribeSendFileResultsResponseBodyInvocations {
	s.Invocation = v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocation struct {
	// The command output.
	//
	// If ContentEncoding is set to PlainText in the request, the original command output is returned. If ContentEncoding is set to Base64 in the request, the Base64-encoded command output is returned.
	//
	// example:
	//
	// Base64
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file.
	//
	// PlainText: The file content is not encoded. Base64: The file content is encoded in Base64. Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The time when the file sending task was created.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The command description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file.
	//
	// example:
	//
	// root
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The overall sending status of the file. The overall sending status of the file varies based on the sending status of the file on all destination instances. Valid values:
	//
	// 	- Pending: The file is being verified or sent. If the sending state of the file on at least one instance is Pending, the overall sending state of the file is Pending.
	//
	// 	- Running: The file is being sent to the instance. If the sending state of the file on at least one instance is Running, the overall sending state of the file is Running.
	//
	// 	- Success: If the sending state of the file on all instances is Success, the overall sending state of the file is Success.
	//
	// 	- Failed: If the sending state of the file on all instances is Failed, the overall sending state of the file is Failed. If the sending state of the file on one or more instances is one of the following values, the overall sending state of the file is Failed:
	//
	//     	- Invalid: The file is invalid.
	//
	//     	- Aborted: The file failed to be sent to the instances.
	//
	//     	- Failed: The file failed to be created on the instances.
	//
	//     	- Timeout: The file sending task timed out.
	//
	//     	- Error: An error occurred and interrupted the file sending task.
	//
	// 	- PartialFailed: The file sending task was completed on some instances but failed on other instances. If the sending state of the file is Success on some instances and is Failed on other instances, the overall sending state of the file is PartialFailed.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The file sending records.
	InvokeNodes *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The name of the file sending task.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Indicates whether a file was overwritten in the destination directory if the file has the same name as the sent file.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The destination directory.
	//
	// example:
	//
	// /home/user
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContent(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Content = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContentType(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.ContentType = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetDescription(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Description = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileGroup(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileGroup = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileMode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileMode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileOwner(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileOwner = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetName(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Name = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetNodeCount(v int32) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.NodeCount = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetOverwrite(v bool) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Overwrite = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetTargetDir(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.TargetDir = &v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The file sending records on a node.
	InvokeNode []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The time when the file sending task was created.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code returned when the file failed to be sent to the instance. Valid values:
	//
	// Null: The file is sent to the instance. NodeNotExists: The specified instance does not exist or has been released. NodeReleased: The instance was released while the file was being sent. NodeNotRunning: The instance was not running while the file sending task was being created. AccountNotExists: The specified account does not exist. ClientNotRunning: Cloud Assistant Agent is not running. ClientNotResponse: Cloud Assistant Agent did not respond. ClientIsUpgrading: Cloud Assistant Agent was being upgraded. ClientNeedUpgrade: Cloud Assistant Agent needs to be upgraded. DeliveryTimeout: The file sending task timed out. FileCreateFail: The file failed to be created. FileAlreadyExists: A file with the same name exists in the specified directory. FileContentInvalid: The file content is invalid. FileNameInvalid: The file name is invalid. FilePathInvalid: The specified directory is invalid. FileAuthorityInvalid: The specified permissions on the file are invalid. UserGroupNotExists: The specified user group does not exist.
	//
	// example:
	//
	// AccountNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the command failed to be sent or run. Valid values:
	//
	// 	- null: The command is run as expected.
	//
	// 	- the specified instance does not exists: The specified instance does not exist or is released.
	//
	// 	- the node has released when create task: The instance is released when the command is being run.
	//
	// 	- the node is not running when create task: The instance is not in the Running state while the command is being run.
	//
	// 	- the command is not applicable: The command is not applicable to the specified instance.
	//
	// 	- the specified account does not exists: The specified account does not exist.
	//
	// 	- the specified directory does not exists: The specified directory does not exist.
	//
	// 	- the cron job expression is invalid: The cron expression that specifies the execution time is invalid.
	//
	// 	- the aliyun service is not running on the instance: Cloud Assistant Agent is not running.
	//
	// 	- the aliyun service in the instance does not response: Cloud Assistant Agent does not respond.
	//
	// 	- the aliyun service in the node is upgrading now: Cloud Assistant Agent is being upgraded.
	//
	// 	- the aliyun service in the node need upgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// 	- the command delivery has been timeout: The request to send the command timed out.
	//
	// 	- the command execution has been timeout: The command execution timed out.
	//
	// 	- the command execution got an exception: An exception occurred when the command is being run.
	//
	// 	- the command execution has been interrupted: The command execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The command execution completes, but the exit code is not 0.
	//
	// 	- the specified instance has been released: The instance is released while the file is being sent.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The time when the file sending task ends. The time is in the 2020-05-22T17:04:18 format.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The status of the file sending task on an instance. Valid values:
	//
	// 	- Pending: The file is being verified or sent.
	//
	// 	- Invalid: The file is invalid.
	//
	// 	- Running: The file is being sent to the instance.
	//
	// 	- Aborted: The file failed to be sent to the instance.
	//
	// 	- Success: The file is sent.
	//
	// 	- Failed: The file failed to be created on the instance.
	//
	// 	- Error: An error occurred and interrupted the file sending task.
	//
	// 	- Timeout: The file sending task timed out.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-9lb3c15m81j
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

type DescribeSendFileResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSendFileResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSendFileResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponse) SetHeaders(v map[string]*string) *DescribeSendFileResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSendFileResultsResponse) SetStatusCode(v int32) *DescribeSendFileResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSendFileResultsResponse) SetBody(v *DescribeSendFileResultsResponseBody) *DescribeSendFileResultsResponse {
	s.Body = v
	return s
}

type DescribeTaskRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i156331731670384438138
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
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The create time.
	//
	// example:
	//
	// 2022-11-30T02:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned for failed tasks.
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The IDs of the nodes.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A7FD7411-9395-52E8-AF42-EB3A4A55446D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The steps.
	Steps []*DescribeTaskResponseBodySteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// The task status.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- execution_success
	//
	// 	- execution_fail
	//
	// 	- waiting_to_run
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The task type.
	//
	// Valid values:
	//
	// 	- reclone_node_sub_task
	//
	// 	- initialize_bare_cluster
	//
	// 	- extend_bare_cluster
	//
	// 	- reclone_node
	//
	// 	- reboot_node
	//
	// 	- extend_ack_edge_cluster
	//
	// 	- extend_cluster
	//
	// 	- initialize_ack_edge_cluster
	//
	// 	- cut_node_sub_task
	//
	// 	- reboot_node_sub_task
	//
	// 	- reclone_ack_edge_node
	//
	// 	- initialize_cluster
	//
	// 	- cut_cluster
	//
	// 	- reclone_bare_node
	//
	// 	- cut_bare_cluster
	//
	// example:
	//
	// cut_cluster
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-11-30T03:40:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *DescribeTaskResponseBody) SetNodeIds(v []*string) *DescribeTaskResponseBody {
	s.NodeIds = v
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
	// The error message of the step.
	//
	// example:
	//
	// get taskinfo failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The stage marker.
	//
	// Valid values:
	//
	// 	- 机器释放: Machine release.
	//
	// 	- 节点并发初始化: Node concurrent initialization.
	//
	// 	- 节点释放: Node release.
	//
	// 	- 机器替换: Machine replacement.
	//
	// 	- 节点缩容: Node scale-in.
	//
	// 	- 提前续费: Early renewal.
	//
	// 	- 物理机清理: Physical machine cleanup.
	//
	// 	- 节点清理: Node cleanup.
	//
	// 	- 创建K8s集群: Create Kubernetes cluster.
	//
	// 	- 网络初始化: Network initialization.
	//
	// 	- 节点重启: Node restart.
	//
	// 	- 节点退订: Node unsubscribe.
	//
	// 	- 集群扩容: Cluster scale-out.
	//
	// 	- 异常机器释放: Abnormal machine release.
	//
	// example:
	//
	// 节点缩容
	StageTag *string `json:"StageTag,omitempty" xml:"StageTag,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the step.
	//
	// example:
	//
	// create_vpd
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The step status.
	//
	// Valid values:
	//
	// 	- execution_success
	//
	// 	- execution_failed
	//
	// example:
	//
	// execution_success
	StepState *string `json:"StepState,omitempty" xml:"StepState,omitempty"`
	// The type of the step.
	//
	// Valid values:
	//
	// 	- normal: A normal step has only one successor step.
	//
	// 	- dispersive: A dispersive step has multiple successor steps.
	//
	// example:
	//
	// normal
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
	// The sub tasks.
	SubTasks []*DescribeTaskResponseBodyStepsSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	// The update time.
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// The creation time.
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned for failed sub tasks.
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task ID.
	//
	// example:
	//
	// i158805051661047928377
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The task type.
	//
	// Valid values:
	//
	// 	- reclone_node_sub_task
	//
	// 	- initialize_bare_cluster
	//
	// 	- extend_bare_cluster
	//
	// 	- reclone_node
	//
	// 	- reboot_node
	//
	// 	- extend_ack_edge_cluster
	//
	// 	- extend_cluster
	//
	// 	- initialize_ack_edge_cluster
	//
	// 	- cut_node_sub_task
	//
	// 	- reboot_node_sub_task
	//
	// 	- reclone_ack_edge_node
	//
	// 	- initialize_cluster
	//
	// 	- cut_cluster
	//
	// 	- reclone_bare_node
	//
	// 	- cut_bare_cluster
	//
	// example:
	//
	// cut_node_sub_task
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeVscRequest struct {
	// The VSC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeVscRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscRequest) SetVscId(v string) *DescribeVscRequest {
	s.VscId = &v
	return s
}

type DescribeVscResponseBody struct {
	// The ID of the compute node in which the VSC resides.
	//
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The VSC status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Normal
	//
	// 	- Deleting
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The custom name of the VSC.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscResponseBody) SetNodeId(v string) *DescribeVscResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeVscResponseBody) SetRequestId(v string) *DescribeVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscResponseBody) SetResourceGroupId(v string) *DescribeVscResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscResponseBody) SetStatus(v string) *DescribeVscResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscId(v string) *DescribeVscResponseBody {
	s.VscId = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscName(v string) *DescribeVscResponseBody {
	s.VscName = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscType(v string) *DescribeVscResponseBody {
	s.VscType = &v
	return s
}

type DescribeVscResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscResponse) SetHeaders(v map[string]*string) *DescribeVscResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscResponse) SetStatusCode(v int32) *DescribeVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscResponse) SetBody(v *DescribeVscResponseBody) *DescribeVscResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The natural language that is used to filter responses. For more information, see RFC 7231. Valid values:
	//
	// zh-CN en-US Default value: zh-CN.
	//
	// Valid values:
	//
	// 	- en-US
	//
	// 	- zh-CN
	//
	// example:
	//
	// zh-CN
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
	// The request ID.
	//
	// example:
	//
	// E9116F2D-82F8-501E-9ADB-2BE0C02B6A84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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
	// The zone name.
	//
	// example:
	//
	// Hang Zhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The combined policies for assigning IP addresses: Only one policy type can be selected for each policy, and multiple policies can be combined.
	IpAllocationPolicy []*ExtendClusterRequestIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	// The node groups.
	NodeGroups []*ExtendClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// The subnets of the cluster.
	VpdSubnets []*string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
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

func (s *ExtendClusterRequest) SetVSwitchZoneId(v string) *ExtendClusterRequest {
	s.VSwitchZoneId = &v
	return s
}

func (s *ExtendClusterRequest) SetVpdSubnets(v []*string) *ExtendClusterRequest {
	s.VpdSubnets = v
	return s
}

type ExtendClusterRequestIpAllocationPolicy struct {
	// The bond policy that you specify the cluster subnet ID based on the bond name.
	BondPolicy *ExtendClusterRequestIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// The allocation policies for the instance type.
	MachineTypePolicy []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// The node allocation policies.
	NodePolicy []*ExtendClusterRequestIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
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
	// The default bond subnet for the cluster.
	//
	// example:
	//
	// subnet-3od2fe
	BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	// The bonds.
	Bonds []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
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
	// The bond name.
	//
	// example:
	//
	// Bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// subnet-3od2fe
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
	// The bonds.
	Bonds []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
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
	// The bond name.
	//
	// example:
	//
	// Bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// subnet-fdo3dv
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
	// The bonds.
	Bonds []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// i22c11282.eu95sqa
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i-3fdodw2
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetHostname(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
	s.Hostname = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetNodeId(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
	s.NodeId = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyNodePolicyBonds struct {
	// The bond name.
	//
	// example:
	//
	// Bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster subnet from which the IP address originates.
	//
	// example:
	//
	// subnet-fdo3dv
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
	// The number of nodes to be purchased. Valid values: 0 to 500. If you set the value of the Amount parameter to 0, you do not want to purchase nodes and scale out the cluster by using existing nodes. If you set the value of the Amount parameter to a value ranging from 1 to 500, you want to purchase a certain number of nodes for cluster scale-out. Default value: 0.
	//
	// example:
	//
	// 4
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal for the purchased nodes. This parameter takes effect only when the Amount parameter is not set to 0 and the ChargeType parameter is set to PrePaid. Valid values: true and false. Default value: False.
	//
	// example:
	//
	// True
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The billing method of the node. This parameter does not take effect if you set the Amount parameter to 0. Valid values: PrePaid (subscription) and PostPaid (pay-as-you-go). Default value: PrePaid.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The hostname for the purchased node. This parameter does not take effect if you set the Amount parameter to 0.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The logon password of the purchased node. This parameter does not take effect if you set the Amount parameter to 0.
	//
	// example:
	//
	// Addk(*78
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// i16d4883a46cbadeb4bc9
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The tags.
	NodeTag []*ExtendClusterRequestNodeGroupsNodeTag `json:"NodeTag,omitempty" xml:"NodeTag,omitempty" type:"Repeated"`
	// The nodes.
	Nodes []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The duration of the purchased node. Unit: months. Valid values: 1, 6, 12, 24, 36, and 48. This parameter takes effect only when the Amount parameter is not set to 0 and the ChargeType parameter is set to PrePaid.
	//
	// example:
	//
	// 6
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The user data.
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0jly2d537ejphyq6h13ke
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-zq1econyv63tvyci5hefw
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ExtendClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroups) SetAmount(v int64) *ExtendClusterRequestNodeGroups {
	s.Amount = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetAutoRenew(v bool) *ExtendClusterRequestNodeGroups {
	s.AutoRenew = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetChargeType(v string) *ExtendClusterRequestNodeGroups {
	s.ChargeType = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetHostnames(v []*string) *ExtendClusterRequestNodeGroups {
	s.Hostnames = v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetLoginPassword(v string) *ExtendClusterRequestNodeGroups {
	s.LoginPassword = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodeGroupId(v string) *ExtendClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodeTag(v []*ExtendClusterRequestNodeGroupsNodeTag) *ExtendClusterRequestNodeGroups {
	s.NodeTag = v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodes(v []*ExtendClusterRequestNodeGroupsNodes) *ExtendClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetPeriod(v int64) *ExtendClusterRequestNodeGroups {
	s.Period = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetUserData(v string) *ExtendClusterRequestNodeGroups {
	s.UserData = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetVSwitchId(v string) *ExtendClusterRequestNodeGroups {
	s.VSwitchId = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetVpcId(v string) *ExtendClusterRequestNodeGroups {
	s.VpcId = &v
	return s
}

func (s *ExtendClusterRequestNodeGroups) SetZoneId(v string) *ExtendClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

type ExtendClusterRequestNodeGroupsNodeTag struct {
	// The tag key.
	//
	// example:
	//
	// my_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// my_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodeTag) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodeTag) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetKey(v string) *ExtendClusterRequestNodeGroupsNodeTag {
	s.Key = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetValue(v string) *ExtendClusterRequestNodeGroupsNodeTag {
	s.Value = &v
	return s
}

type ExtendClusterRequestNodeGroupsNodes struct {
	DataDisk []*ExtendClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// d044d220-33fd-11ed-86a6
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The logon password.
	//
	// example:
	//
	// ***
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2zdpy601
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp169pi5fj151rrms4sia
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-0jlasms92fdxqd3wlf8ny
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetDataDisk(v []*ExtendClusterRequestNodeGroupsNodesDataDisk) *ExtendClusterRequestNodeGroupsNodes {
	s.DataDisk = v
	return s
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

func (s *ExtendClusterRequestNodeGroupsNodes) SetVSwitchId(v string) *ExtendClusterRequestNodeGroupsNodes {
	s.VSwitchId = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetVpcId(v string) *ExtendClusterRequestNodeGroupsNodes {
	s.VpcId = &v
	return s
}

type ExtendClusterRequestNodeGroupsNodesDataDisk struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithNode   *bool   `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodesDataDisk) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodesDataDisk) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetCategory(v string) *ExtendClusterRequestNodeGroupsNodesDataDisk {
	s.Category = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetDeleteWithNode(v bool) *ExtendClusterRequestNodeGroupsNodesDataDisk {
	s.DeleteWithNode = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetPerformanceLevel(v string) *ExtendClusterRequestNodeGroupsNodesDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetSize(v int32) *ExtendClusterRequestNodeGroupsNodesDataDisk {
	s.Size = &v
	return s
}

type ExtendClusterShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The combined policies for assigning IP addresses: Only one policy type can be selected for each policy, and multiple policies can be combined.
	IpAllocationPolicyShrink *string `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty"`
	// The node groups.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// The subnets of the cluster.
	VpdSubnetsShrink *string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty"`
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

func (s *ExtendClusterShrinkRequest) SetVSwitchZoneId(v string) *ExtendClusterShrinkRequest {
	s.VSwitchZoneId = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest {
	s.VpdSubnetsShrink = &v
	return s
}

type ExtendClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExtendClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxkxkllss
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClusterNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListClusterNodesRequest) SetResourceGroupId(v string) *ListClusterNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterNodesRequest) SetTags(v []*ListClusterNodesRequestTags) *ListClusterNodesRequest {
	s.Tags = v
	return s
}

type ListClusterNodesRequestTags struct {
	// The tag key for the node.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the node.
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequestTags) SetKey(v string) *ListClusterNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesRequestTags) SetValue(v string) *ListClusterNodesRequestTags {
	s.Value = &v
	return s
}

type ListClusterNodesResponseBody struct {
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAXW/ZB9TBvH+0ZK0phtCibQgQmu1RbqplAI6Velo2OKR
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The nodes.
	Nodes []*ListClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2BA76272-6608-5AEC-BBA8-B6F0D3D14CDB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The commodity code.
	//
	// example:
	//
	// bcccluster
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1642472468000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 1762185600000
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The hostname.
	//
	// example:
	//
	// 72432f80-273e-11ed-b57a-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The system image ID.
	//
	// example:
	//
	// i190297201669099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Alinux3_x86_AMD_R_Host_D3_E3_24.13.00_UEFI_N_250121
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The node type.
	//
	// example:
	//
	// cn-wulanchabu-b11
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The network information.
	Networks []*ListClusterNodesResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-e9b74f4d450cf18d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// emr_master
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-2r42tmj4z02
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node status.
	//
	// Valid values:
	//
	// 	- Extending
	//
	// 	- UnusedNodeStopped
	//
	// 	- UnusedNodeStopping
	//
	// 	- Unused
	//
	// 	- Using
	//
	// 	- ReleaseLocking
	//
	// 	- Operating
	//
	// 	- Cutting
	//
	// 	- ClusterNodeStopped
	//
	// 	- UnusedNodeRecovering
	//
	// 	- ClusterNodeStopping
	//
	// 	- ClusterNodeRecovering
	//
	// 	- Replacing
	//
	// example:
	//
	// Extending
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sn_tOuUk
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The tags.
	Tags []*ListClusterNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// i28ddkdkkdkdd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1mxqhw8o20tgv3xk47h
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-0jltf9vinjz3if3lltdy7
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodes) SetCommodityCode(v string) *ListClusterNodesResponseBodyNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetCreateTime(v string) *ListClusterNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetExpiredTime(v string) *ListClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetFileSystemMountEnabled(v bool) *ListClusterNodesResponseBodyNodes {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetHostname(v string) *ListClusterNodesResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetHpnZone(v string) *ListClusterNodesResponseBodyNodes {
	s.HpnZone = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetImageId(v string) *ListClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetImageName(v string) *ListClusterNodesResponseBodyNodes {
	s.ImageName = &v
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

func (s *ListClusterNodesResponseBodyNodes) SetTags(v []*ListClusterNodesResponseBodyNodesTags) *ListClusterNodesResponseBodyNodes {
	s.Tags = v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetTaskId(v string) *ListClusterNodesResponseBodyNodes {
	s.TaskId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetVSwitchId(v string) *ListClusterNodesResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetVpcId(v string) *ListClusterNodesResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetZoneId(v string) *ListClusterNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListClusterNodesResponseBodyNodesNetworks struct {
	// The name of the network port for the node.
	//
	// example:
	//
	// bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// The IP address of the node in the virtual private cloud (VPC).
	//
	// example:
	//
	// 192.168.22.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The subnet ID.
	//
	// example:
	//
	// subnet-fwekrvg9
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
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

type ListClusterNodesResponseBodyNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesResponseBodyNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesTags) SetKey(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesTags) SetValue(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

type ListClusterNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bg6wyoox6jq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClustersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListClustersRequest) SetTags(v []*ListClustersRequestTags) *ListClustersRequest {
	s.Tags = v
	return s
}

type ListClustersRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// key_aa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value_aa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequestTags) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTags) SetKey(v string) *ListClustersRequestTags {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTags) SetValue(v string) *ListClustersRequestTags {
	s.Value = &v
	return s
}

type ListClustersResponseBody struct {
	// The clusters.
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// f4f9a292c17072a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The cluster description.
	//
	// example:
	//
	// PPU-cluster2 bz
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// i137590131672134915401
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cnp_test_cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// Valid values:
	//
	// 	- AckEdgePro
	//
	// 	- ExclusiveBareCluster
	//
	// 	- Lite
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The component information.
	//
	// example:
	//
	// {}
	Components interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// The IP type of the computing network.
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1672134938
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 12
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of node groups.
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// The cluster status.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- expanding
	//
	// 	- shrinking
	//
	// 	- initializing
	//
	// example:
	//
	// initializing
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2ajbjoloa23q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClustersResponseBodyClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// i156365121663149566024
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1672134968
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-0jlx4hol2bjboafzmffvd
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *ListClustersResponseBodyClusters) SetComputingIpVersion(v string) *ListClustersResponseBodyClusters {
	s.ComputingIpVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCreateTime(v string) *ListClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetHpnZone(v string) *ListClustersResponseBodyClusters {
	s.HpnZone = &v
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

func (s *ListClustersResponseBodyClusters) SetTags(v []*ListClustersResponseBodyClustersTags) *ListClustersResponseBodyClusters {
	s.Tags = v
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

func (s *ListClustersResponseBodyClusters) SetVpcId(v string) *ListClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

type ListClustersResponseBodyClustersTags struct {
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersResponseBodyClustersTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersTags) SetKey(v string) *ListClustersResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) SetValue(v string) *ListClustersResponseBodyClustersTags {
	s.Value = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListDiagnosticResultsRequest struct {
	// example:
	//
	// NetDiag
	DiagType *string `json:"DiagType,omitempty" xml:"DiagType,omitempty"`
	// *
	//
	// *
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListDiagnosticResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsRequest) SetDiagType(v string) *ListDiagnosticResultsRequest {
	s.DiagType = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetMaxResults(v int64) *ListDiagnosticResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetNextToken(v string) *ListDiagnosticResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetResourceGroupId(v string) *ListDiagnosticResultsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListDiagnosticResultsResponseBody struct {
	DiagnosticResults []*ListDiagnosticResultsResponseBodyDiagnosticResults `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// *
	//
	// *
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDiagnosticResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBody) SetDiagnosticResults(v []*ListDiagnosticResultsResponseBodyDiagnosticResults) *ListDiagnosticResultsResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetMaxResults(v int64) *ListDiagnosticResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetNextToken(v string) *ListDiagnosticResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetRequestId(v string) *ListDiagnosticResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListDiagnosticResultsResponseBodyDiagnosticResults struct {
	// example:
	//
	// i118578141694745246055
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// pjlab-lingjun
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 2024-01-15T02:01:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// diagcontent
	DiagContent *string `json:"DiagContent,omitempty" xml:"DiagContent,omitempty"`
	// example:
	//
	// 123
	DiagId *string `json:"DiagId,omitempty" xml:"DiagId,omitempty"`
	// example:
	//
	// Success
	DiagResult *string `json:"DiagResult,omitempty" xml:"DiagResult,omitempty"`
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// example:
	//
	// e01-cn-bl03ofg6206
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// proxy-rps.mos.csvw.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetCreationTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.CreationTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagContent(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagContent = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagResult(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagResult = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetFinishedTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.FinishedTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetResourceId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ResourceId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetServerName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ServerName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetStatus(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.Status = &v
	return s
}

type ListDiagnosticResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosticResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosticResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponse) SetHeaders(v map[string]*string) *ListDiagnosticResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosticResultsResponse) SetStatusCode(v int32) *ListDiagnosticResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosticResultsResponse) SetBody(v *ListDiagnosticResultsResponseBody) *ListDiagnosticResultsResponse {
	s.Body = v
	return s
}

type ListFreeNodesRequest struct {
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The instance type.
	//
	// example:
	//
	// mock-machine-type2
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The types of the returned nodes that are not used.
	OperatingStates []*string `json:"OperatingStates,omitempty" xml:"OperatingStates,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListFreeNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFreeNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequest) SetHpnZone(v string) *ListFreeNodesRequest {
	s.HpnZone = &v
	return s
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

func (s *ListFreeNodesRequest) SetOperatingStates(v []*string) *ListFreeNodesRequest {
	s.OperatingStates = v
	return s
}

func (s *ListFreeNodesRequest) SetResourceGroupId(v string) *ListFreeNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesRequest) SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest {
	s.Tags = v
	return s
}

type ListFreeNodesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// key_aa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value_aa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequestTags) SetKey(v string) *ListFreeNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesRequestTags) SetValue(v string) *ListFreeNodesRequestTags {
	s.Value = &v
	return s
}

type ListFreeNodesResponseBody struct {
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The nodes.
	Nodes []*ListFreeNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AA14CB86-70C4-5CB7-9E7B-6CCA77F3512B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The commodity code.
	//
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1652321554
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 1673107200
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-7pp2x193801
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node status.
	//
	// example:
	//
	// Unused
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzkkbrpl4owgy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sn_pozkHBgicd
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The tags.
	Tags []*ListFreeNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodes) SetCommodityCode(v string) *ListFreeNodesResponseBodyNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetCreateTime(v string) *ListFreeNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetExpiredTime(v string) *ListFreeNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetHpnZone(v string) *ListFreeNodesResponseBodyNodes {
	s.HpnZone = &v
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

func (s *ListFreeNodesResponseBodyNodes) SetOperatingState(v string) *ListFreeNodesResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetResourceGroupId(v string) *ListFreeNodesResponseBodyNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetSn(v string) *ListFreeNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetTags(v []*ListFreeNodesResponseBodyNodesTags) *ListFreeNodesResponseBodyNodes {
	s.Tags = v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetZoneId(v string) *ListFreeNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListFreeNodesResponseBodyNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_vakye
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesResponseBodyNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodesTags) SetKey(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) SetValue(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

type ListFreeNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListImagesRequest struct {
	// The architecture.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image version.
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The platform.
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetArchitecture(v string) *ListImagesRequest {
	s.Architecture = &v
	return s
}

func (s *ListImagesRequest) SetImageVersion(v string) *ListImagesRequest {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesRequest) SetPlatform(v string) *ListImagesRequest {
	s.Platform = &v
	return s
}

type ListImagesResponseBody struct {
	// The image details.
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0FC4A1C7-421C-5EAB-9361-4C0338EFA287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetNextToken(v string) *ListImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	// The architecture.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The description.
	//
	// example:
	//
	// alibaba cloud linux 3 full for H800
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i190951671671438639388
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image version.
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 20
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The platform.
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The MD5 hash value of the file.
	//
	// example:
	//
	// 40741292480fc6d778138adcf8c
	ReleaseFileMd5 *string `json:"ReleaseFileMd5,omitempty" xml:"ReleaseFileMd5,omitempty"`
	// The image size.
	//
	// example:
	//
	// 5.8G
	ReleaseFileSize *string `json:"ReleaseFileSize,omitempty" xml:"ReleaseFileSize,omitempty"`
	// The image type.
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetArchitecture(v string) *ListImagesResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageVersion(v string) *ListImagesResponseBodyImages {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetNodeCount(v int64) *ListImagesResponseBodyImages {
	s.NodeCount = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileMd5(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileMd5 = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileSize(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetType(v string) *ListImagesResponseBodyImages {
	s.Type = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListMachineNetworkInfoRequest struct {
	MachineHpnInfo []*ListMachineNetworkInfoRequestMachineHpnInfo `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty" type:"Repeated"`
}

func (s ListMachineNetworkInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequest) SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest {
	s.MachineHpnInfo = v
	return s
}

type ListMachineNetworkInfoRequestMachineHpnInfo struct {
	// example:
	//
	// C1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// efg2.C48cNHmcn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetHpnZone(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetMachineType(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetRegionId(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.RegionId = &v
	return s
}

type ListMachineNetworkInfoShrinkRequest struct {
	MachineHpnInfoShrink *string `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty"`
}

func (s ListMachineNetworkInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoShrinkRequest) SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest {
	s.MachineHpnInfoShrink = &v
	return s
}

type ListMachineNetworkInfoResponseBody struct {
	MachineNetworkInfo []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo `json:"MachineNetworkInfo,omitempty" xml:"MachineNetworkInfo,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBody) SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody {
	s.MachineNetworkInfo = v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) SetRequestId(v string) *ListMachineNetworkInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListMachineNetworkInfoResponseBodyMachineNetworkInfo struct {
	// example:
	//
	// vpc/acl
	ClusterNet *string `json:"ClusterNet,omitempty" xml:"ClusterNet,omitempty"`
	// example:
	//
	// true
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// true
	IsDpuMode *bool `json:"IsDpuMode,omitempty" xml:"IsDpuMode,omitempty"`
	// example:
	//
	// efg1.nvga8n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// example:
	//
	// XX-7.0
	NetArch *string `json:"NetArch,omitempty" xml:"NetArch,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetClusterNet(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.ClusterNet = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetEnableJumboFrame(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.EnableJumboFrame = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetHpnZone(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetIsDpuMode(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.IsDpuMode = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetMachineType(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetNetArch(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.NetArch = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetRegionId(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.RegionId = &v
	return s
}

type ListMachineNetworkInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineNetworkInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineNetworkInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponse) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponse) SetHeaders(v map[string]*string) *ListMachineNetworkInfoResponse {
	s.Headers = v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetStatusCode(v int32) *ListMachineNetworkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetBody(v *ListMachineNetworkInfoResponseBody) *ListMachineNetworkInfoResponse {
	s.Body = v
	return s
}

type ListMachineTypesRequest struct {
	// The name of the instance type.
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMachineTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesRequest) GoString() string {
	return s.String()
}

func (s *ListMachineTypesRequest) SetName(v string) *ListMachineTypesRequest {
	s.Name = &v
	return s
}

type ListMachineTypesResponseBody struct {
	// The instance types.
	MachineTypes []*ListMachineTypesResponseBodyMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F16BA4D8-FF50-53B6-A026-F443FE31006C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBody) SetMachineTypes(v []*ListMachineTypesResponseBodyMachineTypes) *ListMachineTypesResponseBody {
	s.MachineTypes = v
	return s
}

func (s *ListMachineTypesResponseBody) SetNextToken(v string) *ListMachineTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMachineTypesResponseBody) SetRequestId(v string) *ListMachineTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListMachineTypesResponseBodyMachineTypes struct {
	// The number of bonds.
	//
	// example:
	//
	// 2
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// The CPU information.
	//
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// The disk information.
	//
	// example:
	//
	// 2x 480GB SATA SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// The GPU information.
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// The storage information.
	//
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// The name of the instance type.
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network information.
	//
	// example:
	//
	// 2x 100Gbps DP NIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 10
	NodeCount *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 48
	TotalCpuCore *int32 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// The access type.
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypes) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetBondNum(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.BondNum = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetCpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.CpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetDiskInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.DiskInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetGpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.GpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetMemoryInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.MemoryInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetName(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Name = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNetworkInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NetworkInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNodeCount(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NodeCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetTotalCpuCore(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.TotalCpuCore = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetType(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Type = &v
	return s
}

type ListMachineTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponse) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponse) SetHeaders(v map[string]*string) *ListMachineTypesResponse {
	s.Headers = v
	return s
}

func (s *ListMachineTypesResponse) SetStatusCode(v int32) *ListMachineTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineTypesResponse) SetBody(v *ListMachineTypesResponseBody) *ListMachineTypesResponse {
	s.Body = v
	return s
}

type ListNetTestResultsRequest struct {
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value:
	//
	// 	- If you do not configure this parameter or if you set this parameter to a value less than 20, the default value is 20.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the network test.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNetTestResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsRequest) SetMaxResults(v int64) *ListNetTestResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNetTestType(v string) *ListNetTestResultsRequest {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNextToken(v string) *ListNetTestResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsRequest) SetResourceGroupId(v string) *ListNetTestResultsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListNetTestResultsResponseBody struct {
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value:
	//
	// 	- If you do not configure this parameter or if you set this parameter to a value less than 20, the default value is 20.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The results.
	NetTestResults []*ListNetTestResultsResponseBodyNetTestResults `json:"NetTestResults,omitempty" xml:"NetTestResults,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetTestResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBody) SetMaxResults(v int64) *ListNetTestResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNetTestResults(v []*ListNetTestResultsResponseBodyNetTestResults) *ListNetTestResultsResponseBody {
	s.NetTestResults = v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNextToken(v string) *ListNetTestResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetRequestId(v string) *ListNetTestResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResults struct {
	// The cluster ID.
	//
	// example:
	//
	// i110667211718265012218
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Returned when NetTestType is CommTest.
	CommTest *ListNetTestResultsResponseBodyNetTestResultsCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2024-01-19T02:18:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Returned when NetTestType is DelayTest.
	DelayTest *ListNetTestResultsResponseBodyNetTestResultsDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// The finish time.
	//
	// example:
	//
	// 2024-10-30T02:07Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The type of the network test.
	//
	// example:
	//
	// NetDiag
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The network mode.
	//
	// example:
	//
	// 01
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the network test task. Valid values:\\
	//
	// ● InProgress\\
	//
	// ● Finished\\
	//
	// ● Failed
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The test ID. The unique identifier of the resource test task.
	//
	// example:
	//
	// String	i-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	// Returned when NetTestType is TrafficTest.
	TrafficTest *ListNetTestResultsResponseBodyNetTestResultsTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s ListNetTestResultsResponseBodyNetTestResults) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResults) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterName(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCommTest(v *ListNetTestResultsResponseBodyNetTestResultsCommTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.CommTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCreationTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.CreationTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetDelayTest(v *ListNetTestResultsResponseBodyNetTestResultsDelayTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.DelayTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetFinishedTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.FinishedTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetTestType(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetworkMode(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetworkMode = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetPort(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Port = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetStatus(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Status = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTestId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.TestId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTrafficTest(v *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.TrafficTest = v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsCommTest struct {
	// The number of GPUs.
	//
	// example:
	//
	// 4
	GPUNum *string `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// The hosts of the test node.
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The communication library model.
	//
	// example:
	//
	// AllToAll
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The CommTest type, which can be ACCL or NCCL.
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetGPUNum(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.GPUNum = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Hosts = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetModel(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Model = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetType(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Type = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsCommTestHosts struct {
	// The IP address of the node.
	//
	// example:
	//
	// 10.51.16.21
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// www.xinjiaoyu.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTest struct {
	// The hosts.
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	s.Hosts = v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// pgm-bp174z988a27wre71o.pg.rds.aliyuncs.com
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// WrF
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTest struct {
	// The clients.
	Clients []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The running duration of the pipeline job. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// If the protocol is RDMA, can be True or False. If the protocol is TCP, this field is empty.
	//
	// example:
	//
	// True
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// The network protocol, which can be RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// If the protocol is TCP, the number of concurrent connections. If the protocol is RDMA, the configured QP value.
	//
	// example:
	//
	// RDMA
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// If the TrafficModel is Fullmesh, this parameter is empty.
	Servers []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The traffic model, which can be MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetClients(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Clients = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetDuration(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Duration = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetGDR(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.GDR = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetProtocol(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Protocol = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetQP(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.QP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetServers(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Servers = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetTrafficModel(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.TrafficModel = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 74.73.100.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-20p36bqet39
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.1.168.183
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetTestResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponse) SetHeaders(v map[string]*string) *ListNetTestResultsResponse {
	s.Headers = v
	return s
}

func (s *ListNetTestResultsResponse) SetStatusCode(v int32) *ListNetTestResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetTestResultsResponse) SetBody(v *ListNetTestResultsResponseBody) *ListNetTestResultsResponse {
	s.Body = v
	return s
}

type ListNodeGroupsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value:
	//
	// • If you do not configure this parameter or if you set this parameter to a value less than 20, the default value is 20.
	//
	// • If you set this parameter to a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) SetClusterId(v string) *ListNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetMaxResults(v int32) *ListNodeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNextToken(v string) *ListNodeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupId(v string) *ListNodeGroupsRequest {
	s.NodeGroupId = &v
	return s
}

type ListNodeGroupsResponseBody struct {
	// The node groups.
	Groups []*ListNodeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) SetGroups(v []*ListNodeGroupsResponseBodyGroups) *ListNodeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNextToken(v string) *ListNodeGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetRequestId(v string) *ListNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeGroupsResponseBodyGroups struct {
	// The cluster ID.
	//
	// example:
	//
	// i113952461729854708648
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// wzq-exclusivelite-71
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-02-27T13:16:31.599
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// created by ga2_prepare
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The group ID.
	//
	// example:
	//
	// 238276221
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// backend-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i194015071707321240258
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-09-22T00:03:05.114
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shenzhen-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyGroups) SetClusterId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetClusterName(v string) *ListNodeGroupsResponseBodyGroups {
	s.ClusterName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetCreateTime(v string) *ListNodeGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetDescription(v string) *ListNodeGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetFileSystemMountEnabled(v bool) *ListNodeGroupsResponseBodyGroups {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetGroupId(v string) *ListNodeGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetGroupName(v string) *ListNodeGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetImageId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ImageId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetImageName(v string) *ListNodeGroupsResponseBodyGroups {
	s.ImageName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetMachineType(v string) *ListNodeGroupsResponseBodyGroups {
	s.MachineType = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetNodeCount(v int64) *ListNodeGroupsResponseBodyGroups {
	s.NodeCount = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetUpdateTime(v string) *ListNodeGroupsResponseBodyGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetZoneId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ZoneId = &v
	return s
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) SetHeaders(v map[string]*string) *ListNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupsResponse) SetStatusCode(v int32) *ListNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupsResponse) SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Node
	//
	// 	- Vcc
	//
	// 	- Cluster
	//
	// 	- Subnet
	//
	// 	- Vpd
	//
	// This parameter is required.
	//
	// example:
	//
	// Node
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key.
	//
	// example:
	//
	// PodName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// WFT-OTC
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
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F208B6D-4C42-5FD3-B6BE-E826E92A44DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
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
	// The resource ID.
	//
	// example:
	//
	// i15azeddnvf7uhw2oij57o0
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Node
	//
	// 	- Cluster
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListUserClusterTypesResponseBody struct {
	ClusterTypes []*ListUserClusterTypesResponseBodyClusterTypes `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserClusterTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBody) SetClusterTypes(v []*ListUserClusterTypesResponseBodyClusterTypes) *ListUserClusterTypesResponseBody {
	s.ClusterTypes = v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetNextToken(v string) *ListUserClusterTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetRequestId(v string) *ListUserClusterTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserClusterTypesResponseBodyClusterTypes struct {
	// example:
	//
	// Public
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// AckEdgePro
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListUserClusterTypesResponseBodyClusterTypes) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponseBodyClusterTypes) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetAccessType(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.AccessType = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetTypeName(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.TypeName = &v
	return s
}

type ListUserClusterTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserClusterTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponse) SetHeaders(v map[string]*string) *ListUserClusterTypesResponse {
	s.Headers = v
	return s
}

func (s *ListUserClusterTypesResponse) SetStatusCode(v int32) *ListUserClusterTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserClusterTypesResponse) SetBody(v *ListUserClusterTypesResponseBody) *ListUserClusterTypesResponse {
	s.Body = v
	return s
}

type ListVscsRequest struct {
	// The maximum number of data entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the nodes.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VSC name.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVscsRequest) GoString() string {
	return s.String()
}

func (s *ListVscsRequest) SetMaxResults(v int32) *ListVscsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsRequest) SetNextToken(v string) *ListVscsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsRequest) SetNodeIds(v []*string) *ListVscsRequest {
	s.NodeIds = v
	return s
}

func (s *ListVscsRequest) SetResourceGroupId(v string) *ListVscsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsRequest) SetTag(v []*ListVscsRequestTag) *ListVscsRequest {
	s.Tag = v
	return s
}

func (s *ListVscsRequest) SetVscName(v string) *ListVscsRequest {
	s.VscName = &v
	return s
}

type ListVscsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVscsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsRequestTag) SetKey(v string) *ListVscsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsRequestTag) SetValue(v string) *ListVscsRequestTag {
	s.Value = &v
	return s
}

type ListVscsShrinkRequest struct {
	// The maximum number of data entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the nodes.
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListVscsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VSC name.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVscsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequest) SetMaxResults(v int32) *ListVscsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNextToken(v string) *ListVscsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNodeIdsShrink(v string) *ListVscsShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *ListVscsShrinkRequest) SetResourceGroupId(v string) *ListVscsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsShrinkRequest) SetTag(v []*ListVscsShrinkRequestTag) *ListVscsShrinkRequest {
	s.Tag = v
	return s
}

func (s *ListVscsShrinkRequest) SetVscName(v string) *ListVscsShrinkRequest {
	s.VscName = &v
	return s
}

type ListVscsShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVscsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequestTag) SetKey(v string) *ListVscsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsShrinkRequestTag) SetValue(v string) *ListVscsShrinkRequestTag {
	s.Value = &v
	return s
}

type ListVscsResponseBody struct {
	// No response is returned. The TotalCount parameter is used.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token. It can be used in the next request to retrieve a new page of results. If this parameter is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VSCs.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VSCs.
	Vscs []*ListVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s ListVscsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBody) SetMaxResults(v int32) *ListVscsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVscsResponseBody) SetNextToken(v string) *ListVscsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVscsResponseBody) SetRequestId(v string) *ListVscsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVscsResponseBody) SetTotalCount(v int32) *ListVscsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVscsResponseBody) SetVscs(v []*ListVscsResponseBodyVscs) *ListVscsResponseBody {
	s.Vscs = v
	return s
}

type ListVscsResponseBodyVscs struct {
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-fzh47xd7u08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm2zkwhkns57i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The VSC status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Normal
	//
	// 	- Deleting
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListVscsResponseBodyVscsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The custom name of the VSC.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type. Valid values: primary and standard.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s ListVscsResponseBodyVscs) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscs) SetNodeId(v string) *ListVscsResponseBodyVscs {
	s.NodeId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetResourceGroupId(v string) *ListVscsResponseBodyVscs {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetStatus(v string) *ListVscsResponseBodyVscs {
	s.Status = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetTags(v []*ListVscsResponseBodyVscsTags) *ListVscsResponseBodyVscs {
	s.Tags = v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscId(v string) *ListVscsResponseBodyVscs {
	s.VscId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscName(v string) *ListVscsResponseBodyVscs {
	s.VscName = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscType(v string) *ListVscsResponseBodyVscs {
	s.VscType = &v
	return s
}

type ListVscsResponseBodyVscsTags struct {
	// The tag key.
	//
	// example:
	//
	// key001
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value001
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVscsResponseBodyVscsTags) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBodyVscsTags) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscsTags) SetTagKey(v string) *ListVscsResponseBodyVscsTags {
	s.TagKey = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) SetTagValue(v string) *ListVscsResponseBodyVscsTags {
	s.TagValue = &v
	return s
}

type ListVscsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVscsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponse) GoString() string {
	return s.String()
}

func (s *ListVscsResponse) SetHeaders(v map[string]*string) *ListVscsResponse {
	s.Headers = v
	return s
}

func (s *ListVscsResponse) SetStatusCode(v int32) *ListVscsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVscsResponse) SetBody(v *ListVscsResponseBody) *ListVscsResponse {
	s.Body = v
	return s
}

type RebootNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
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
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*ReimageNodesRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The user data.
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *ReimageNodesRequest) SetUserData(v string) *ReimageNodesRequest {
	s.UserData = &v
	return s
}

type ReimageNodesRequestNodes struct {
	// The hostname.
	//
	// example:
	//
	// 457db5ca-241d-11ed-9fd7-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The system image ID.
	//
	// example:
	//
	// m-8vbf8rpv2nn14y7oybjy
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The logon password.
	//
	// example:
	//
	// ***
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-zvp2tgykr0b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The user data.
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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

func (s *ReimageNodesShrinkRequest) SetUserData(v string) *ReimageNodesShrinkRequest {
	s.UserData = &v
	return s
}

type ReimageNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15FBCD9B-C93F-54E8-A168-AADE7E66DAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i158782151663841517926
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReimageNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type RunCommandRequest struct {
	// The client token to ensure the idempotency of the request. Use your client to generate the token that is unique among requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. Take note of the following:
	//
	// 	- When `EnableParameter` is set to true, you can use custom parameters in the command.
	//
	// 	- Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	// 	- You can specify up to 20 custom parameters.
	//
	// 	- A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is not case-sensitive.
	//
	// 	- Each custom parameter name is up to 64 bytes in length.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The ID of the command.
	//
	// example:
	//
	// c-e996287206324975b5fbe1d***
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding mode of the command content. Valid values:
	//
	// 	- PlainText
	//
	// 	- Base64
	//
	// Default value: PlainText. If the specified value of this parameter is invalid, PlainText is used by default.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The schedule to run the command. Supported schedule types: at a fixed interval based on a rate expression, run only once at a specific time, or run at specific times based on a cron expression.
	//
	// 	- To run the command at a fixed interval, use a rate expression to specify the interval. The interval can be in seconds, minutes, hours, or days. This option is suitable for scenarios in which tasks need to be executed at a fixed interval. Format: rate(\\<Execution interval value> \\<Execution interval unit>). For example, rate(5m) means to run the command every 5 minutes. When you specify an interval, take note of the following limits:
	//
	//     	- The interval can be anywhere from 60 seconds to 7 days, but must be longer than the timeout period of the scheduled task.
	//
	//     	- The interval is the time between two consecutive executions, irrelevant of the time required to run the command. For example, assume that you set the interval to 5 minutes and that it takes 2 minutes to run the command each time. The system waits 3 minutes before running the command again.
	//
	//     	- The command is not immediately executed after the task is created. For example, assume that you set the interval to 5 minutes. The task begins to be executed 5 minutes after it is created.
	//
	// 	- To run a command only once at a specific point in time, specify a point in time and a time zone. Format: at(yyyy-MM-dd HH:mm:ss \\<Time zone>). If you do not specify a time zone, the Coordinated Universal Time (UTC) time zone is used by default. The time zone name supports the following formats: Full name, such as Asia/Shanghai and America/Los_Angeles. The time offset from GMT. Examples: GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value. The time zone abbreviation. Only UTC is supported. For example, to configure a command to run only once at 13:15:30 on June 6, 2022 (Shanghai time), set the time to at(2022-06-06 13:15:30 Asia/Shanghai). To configure a command to run only once at 13:15:30 on June 6, 2022 (UTC-7), set the time to at(2022-06-06 13:15:30 GMT-7:00).
	//
	// 	- To run a command at designated points in time, use a cron expression to define the schedule. Format: \\<Cron expression> \\<Time zone>. Example: \\<Seconds> \\<Minutes> \\<Hours> \\<Day of the month> \\<Month> \\<Day of the week> \\<Year (optional)> \\<Time zone>. The system calculates the execution times of the command based on the specified cron expression and time zone and runs the command as scheduled. If you do not specify a time zone, the system time zone of the instance is used by default. For more information, see Cron expressions. The time zone name supports the following formats:
	//
	//     	- Full name, such as Asia/Shanghai and America/Los_Angeles.
	//
	//     	- The time offset from GMT. Examples: GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported.
	//
	//     For example, if you specify a command to run at 10:15:00 every day in 2022 in China/Shanghai time, set the time to 0 15 10 ? \\	- \\	- 2022 Asia/Shanghai. To configure a command to run every half an hour from 10:00:00 to 11:30:00 every day in 2022 (UTC+8), set the schedule to 0 0/30 10-11 \\	- \\	- ? 2022 GMT+8:00. To configure a command to run every 5 minutes from 14:00:00 to 14:55:00 every October every two years from 2022 in UTC, set the schedule to 0 0/5 14 \\	- 10 ? 2022/2 UTC.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The launcher for script execution. Cannot exceed 1 KB.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node list.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// The key-value pairs of custom parameters to pass in when the command includes custom parameters. For example, the command content is `echo {{name}}`. You can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced by the paired Jack value to generate a new command. As a result, the `echo Jack` command is run.
	//
	// You can specify 0 to 10 custom parameters. Take note of the following:
	//
	// 	- The key of a custom parameter can be up to 64 characters in length, and cannot be an empty string.
	//
	// 	- The value of a custom parameter can be an empty string.
	//
	// 	- If you want to retain a command, make sure that the command after Base64 encoding, including custom parameters and original command content, does not exceed 18 KB in size. If you do not want to retain the command, make sure that the command after Base64 encoding does not exceed 24 KB in size. You can set `KeepCommand` to specify whether to retain the command.
	//
	// 	- The specified custom parameter names must be included in the custom parameter names that you specify when you create the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is left empty by default, which indicates that the custom parameter feature is disabled.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The mode to run the command. Valid values:
	//
	// 	- Once: runs the command immediately.
	//
	// 	- Period: runs the command based on a schedule. When set to `Period`, `Frequency` is required.
	//
	// 	- NextRebootOnly: runs the command the next time the instances is started.
	//
	// 	- EveryReboot: runs the command every time the instance is started.
	//
	// Default value:
	//
	// 	- If you do not specify `Frequency`, the default value is `Once`.
	//
	// 	- If you specify `Frequency`, RepeatMode is set to `Period` regardless of whether a value is already specified.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// Indicates how the command task is stopped when a command execution is manually stopped or times out. Valid values:
	//
	// Process: The process of the command is stopped. ProcessTree: The process tree of the command is stopped. In this case, the process of the command and all subprocesses are stopped.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// The timeout period for the command. Unit: seconds. A timeout error occurs if the command cannot be run because the process slows down or because a specific module or Cloud Assistant Agent does not exist. When the specified timeout period ends, the command process is forcefully terminated. Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username that you use to run the command. The name can be up to 255 characters in length. For Linux instances, the root user is used by default.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working path of the command. You can specify a custom path. Default path:
	//
	// Linux instances: By default, the path is in the /home directory of the root user.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetClientToken(v string) *RunCommandRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetCommandId(v string) *RunCommandRequest {
	s.CommandId = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDescription(v string) *RunCommandRequest {
	s.Description = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetFrequency(v string) *RunCommandRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandRequest) SetLauncher(v string) *RunCommandRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetNodeIdList(v []*string) *RunCommandRequest {
	s.NodeIdList = v
	return s
}

func (s *RunCommandRequest) SetParameters(v map[string]interface{}) *RunCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunCommandRequest) SetRepeatMode(v string) *RunCommandRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandRequest) SetTerminationMode(v string) *RunCommandRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int32) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetUsername(v string) *RunCommandRequest {
	s.Username = &v
	return s
}

func (s *RunCommandRequest) SetWorkingDir(v string) *RunCommandRequest {
	s.WorkingDir = &v
	return s
}

type RunCommandShrinkRequest struct {
	// The client token to ensure the idempotency of the request. Use your client to generate the token that is unique among requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command content. Take note of the following:
	//
	// 	- When `EnableParameter` is set to true, you can use custom parameters in the command.
	//
	// 	- Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	// 	- You can specify up to 20 custom parameters.
	//
	// 	- A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is not case-sensitive.
	//
	// 	- Each custom parameter name is up to 64 bytes in length.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The ID of the command.
	//
	// example:
	//
	// c-e996287206324975b5fbe1d***
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The encoding mode of the command content. Valid values:
	//
	// 	- PlainText
	//
	// 	- Base64
	//
	// Default value: PlainText. If the specified value of this parameter is invalid, PlainText is used by default.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The schedule to run the command. Supported schedule types: at a fixed interval based on a rate expression, run only once at a specific time, or run at specific times based on a cron expression.
	//
	// 	- To run the command at a fixed interval, use a rate expression to specify the interval. The interval can be in seconds, minutes, hours, or days. This option is suitable for scenarios in which tasks need to be executed at a fixed interval. Format: rate(\\<Execution interval value> \\<Execution interval unit>). For example, rate(5m) means to run the command every 5 minutes. When you specify an interval, take note of the following limits:
	//
	//     	- The interval can be anywhere from 60 seconds to 7 days, but must be longer than the timeout period of the scheduled task.
	//
	//     	- The interval is the time between two consecutive executions, irrelevant of the time required to run the command. For example, assume that you set the interval to 5 minutes and that it takes 2 minutes to run the command each time. The system waits 3 minutes before running the command again.
	//
	//     	- The command is not immediately executed after the task is created. For example, assume that you set the interval to 5 minutes. The task begins to be executed 5 minutes after it is created.
	//
	// 	- To run a command only once at a specific point in time, specify a point in time and a time zone. Format: at(yyyy-MM-dd HH:mm:ss \\<Time zone>). If you do not specify a time zone, the Coordinated Universal Time (UTC) time zone is used by default. The time zone name supports the following formats: Full name, such as Asia/Shanghai and America/Los_Angeles. The time offset from GMT. Examples: GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value. The time zone abbreviation. Only UTC is supported. For example, to configure a command to run only once at 13:15:30 on June 6, 2022 (Shanghai time), set the time to at(2022-06-06 13:15:30 Asia/Shanghai). To configure a command to run only once at 13:15:30 on June 6, 2022 (UTC-7), set the time to at(2022-06-06 13:15:30 GMT-7:00).
	//
	// 	- To run a command at designated points in time, use a cron expression to define the schedule. Format: \\<Cron expression> \\<Time zone>. Example: \\<Seconds> \\<Minutes> \\<Hours> \\<Day of the month> \\<Month> \\<Day of the week> \\<Year (optional)> \\<Time zone>. The system calculates the execution times of the command based on the specified cron expression and time zone and runs the command as scheduled. If you do not specify a time zone, the system time zone of the instance is used by default. For more information, see Cron expressions. The time zone name supports the following formats:
	//
	//     	- Full name, such as Asia/Shanghai and America/Los_Angeles.
	//
	//     	- The time offset from GMT. Examples: GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported.
	//
	//     For example, if you specify a command to run at 10:15:00 every day in 2022 in China/Shanghai time, set the time to 0 15 10 ? \\	- \\	- 2022 Asia/Shanghai. To configure a command to run every half an hour from 10:00:00 to 11:30:00 every day in 2022 (UTC+8), set the schedule to 0 0/30 10-11 \\	- \\	- ? 2022 GMT+8:00. To configure a command to run every 5 minutes from 14:00:00 to 14:55:00 every October every two years from 2022 in UTC, set the schedule to 0 0/5 14 \\	- 10 ? 2022/2 UTC.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The launcher for script execution. Cannot exceed 1 KB.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node list.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	// The key-value pairs of custom parameters to pass in when the command includes custom parameters. For example, the command content is `echo {{name}}`. You can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced by the paired Jack value to generate a new command. As a result, the `echo Jack` command is run.
	//
	// You can specify 0 to 10 custom parameters. Take note of the following:
	//
	// 	- The key of a custom parameter can be up to 64 characters in length, and cannot be an empty string.
	//
	// 	- The value of a custom parameter can be an empty string.
	//
	// 	- If you want to retain a command, make sure that the command after Base64 encoding, including custom parameters and original command content, does not exceed 18 KB in size. If you do not want to retain the command, make sure that the command after Base64 encoding does not exceed 24 KB in size. You can set `KeepCommand` to specify whether to retain the command.
	//
	// 	- The specified custom parameter names must be included in the custom parameter names that you specify when you create the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is left empty by default, which indicates that the custom parameter feature is disabled.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The mode to run the command. Valid values:
	//
	// 	- Once: runs the command immediately.
	//
	// 	- Period: runs the command based on a schedule. When set to `Period`, `Frequency` is required.
	//
	// 	- NextRebootOnly: runs the command the next time the instances is started.
	//
	// 	- EveryReboot: runs the command every time the instance is started.
	//
	// Default value:
	//
	// 	- If you do not specify `Frequency`, the default value is `Once`.
	//
	// 	- If you specify `Frequency`, RepeatMode is set to `Period` regardless of whether a value is already specified.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// Indicates how the command task is stopped when a command execution is manually stopped or times out. Valid values:
	//
	// Process: The process of the command is stopped. ProcessTree: The process tree of the command is stopped. In this case, the process of the command and all subprocesses are stopped.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// The timeout period for the command. Unit: seconds. A timeout error occurs if the command cannot be run because the process slows down or because a specific module or Cloud Assistant Agent does not exist. When the specified timeout period ends, the command process is forcefully terminated. Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username that you use to run the command. The name can be up to 255 characters in length. For Linux instances, the root user is used by default.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working path of the command. You can specify a custom path. Default path:
	//
	// Linux instances: By default, the path is in the /home directory of the root user.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) SetClientToken(v string) *RunCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandId(v string) *RunCommandShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContentEncoding(v string) *RunCommandShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandShrinkRequest) SetDescription(v string) *RunCommandShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetFrequency(v string) *RunCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandShrinkRequest) SetLauncher(v string) *RunCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetNodeIdListShrink(v string) *RunCommandShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRepeatMode(v string) *RunCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTerminationMode(v string) *RunCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int32) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetUsername(v string) *RunCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

type RunCommandResponseBody struct {
	// The ID of the execution.
	//
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d*
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type SendFileRequest struct {
	// The content of the file. The file must not exceed 32 KB in size after it is encoded in Base64.
	//
	// 	- If `ContentType` is set to `PlainText`, the value of Content is in plaintext.
	//
	// 	- If `ContentType` is set to `Base64`, the value of Content is Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash echo "Current User is :" echo $(ps | grep "$$" | awk \\"{print $2}\\") -------- oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file. Valid values:
	//
	// PlainText Base64 Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the file. The description can be up to 512 characters in length and can contain any characters.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file. This parameter takes effect only on Linux instances. Default value: root. The value can be up to 64 characters in length.
	//
	// Note If you want to use a non-root user group, make sure that the user group exists in the instances.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file. This parameter takes effect only on Linux instances. You can configure this parameter in the same way as you configure the chmod command.
	//
	// Default value: 0644: the owner of the file has the read and write permission. The user group of the file and other users have read-only permission.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The file name. The name can be up to 255 characters in length and can contain any characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node list.
	//
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// Specifies whether to overwrite file with the same name in the destination directory.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The directory in the Lingjun node to which the file is sent. If the specified directory does not exist, the system creates the directory automatically.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout period for the file sending task. Unit: seconds.
	//
	// 	- A timeout error occurs when a file cannot be sent because the process slows down or because a specific module or Cloud Assistant Agent does not exist.
	//
	// 	- If the specified timeout period is less than 10 seconds, the system sets the timeout period to 10 seconds to ensure that the file can be sent.
	//
	// Default value: 60.
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) SetContent(v string) *SendFileRequest {
	s.Content = &v
	return s
}

func (s *SendFileRequest) SetContentType(v string) *SendFileRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileRequest) SetDescription(v string) *SendFileRequest {
	s.Description = &v
	return s
}

func (s *SendFileRequest) SetFileGroup(v string) *SendFileRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileRequest) SetFileMode(v string) *SendFileRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileRequest) SetFileOwner(v string) *SendFileRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileRequest) SetName(v string) *SendFileRequest {
	s.Name = &v
	return s
}

func (s *SendFileRequest) SetNodeIdList(v []*string) *SendFileRequest {
	s.NodeIdList = v
	return s
}

func (s *SendFileRequest) SetOverwrite(v bool) *SendFileRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileRequest) SetTargetDir(v string) *SendFileRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileRequest) SetTimeout(v int32) *SendFileRequest {
	s.Timeout = &v
	return s
}

type SendFileShrinkRequest struct {
	// The content of the file. The file must not exceed 32 KB in size after it is encoded in Base64.
	//
	// 	- If `ContentType` is set to `PlainText`, the value of Content is in plaintext.
	//
	// 	- If `ContentType` is set to `Base64`, the value of Content is Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash echo "Current User is :" echo $(ps | grep "$$" | awk \\"{print $2}\\") -------- oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file. Valid values:
	//
	// PlainText Base64 Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the file. The description can be up to 512 characters in length and can contain any characters.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file. This parameter takes effect only on Linux instances. Default value: root. The value can be up to 64 characters in length.
	//
	// Note If you want to use a non-root user group, make sure that the user group exists in the instances.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file. This parameter takes effect only on Linux instances. You can configure this parameter in the same way as you configure the chmod command.
	//
	// Default value: 0644: the owner of the file has the read and write permission. The user group of the file and other users have read-only permission.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. This parameter takes effect only on Linux instances. Default value: root.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The file name. The name can be up to 255 characters in length and can contain any characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node list.
	//
	// This parameter is required.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	// Specifies whether to overwrite file with the same name in the destination directory.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The directory in the Lingjun node to which the file is sent. If the specified directory does not exist, the system creates the directory automatically.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout period for the file sending task. Unit: seconds.
	//
	// 	- A timeout error occurs when a file cannot be sent because the process slows down or because a specific module or Cloud Assistant Agent does not exist.
	//
	// 	- If the specified timeout period is less than 10 seconds, the system sets the timeout period to 10 seconds to ensure that the file can be sent.
	//
	// Default value: 60.
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendFileShrinkRequest) SetContent(v string) *SendFileShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendFileShrinkRequest) SetContentType(v string) *SendFileShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileShrinkRequest) SetDescription(v string) *SendFileShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileGroup(v string) *SendFileShrinkRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileMode(v string) *SendFileShrinkRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileOwner(v string) *SendFileShrinkRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileShrinkRequest) SetName(v string) *SendFileShrinkRequest {
	s.Name = &v
	return s
}

func (s *SendFileShrinkRequest) SetNodeIdListShrink(v string) *SendFileShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *SendFileShrinkRequest) SetOverwrite(v bool) *SendFileShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileShrinkRequest) SetTargetDir(v string) *SendFileShrinkRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileShrinkRequest) SetTimeout(v int32) *SendFileShrinkRequest {
	s.Timeout = &v
	return s
}

type SendFileResponseBody struct {
	// The ID of the execution.
	//
	// example:
	//
	// t-hz03la52z1zkvls
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) SetInvokeId(v string) *SendFileResponseBody {
	s.InvokeId = &v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) SetHeaders(v map[string]*string) *SendFileResponse {
	s.Headers = v
	return s
}

func (s *SendFileResponse) SetStatusCode(v int32) *SendFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFileResponse) SetBody(v *SendFileResponseBody) *SendFileResponse {
	s.Body = v
	return s
}

type ShrinkClusterRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The node groups.
	NodeGroups []*ShrinkClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
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
	// The node group ID.
	//
	// example:
	//
	// ng-3b6fbd24b1b845a0
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The nodes.
	Nodes []*ShrinkClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
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
	// The node ID.
	//
	// example:
	//
	// e01poc-cn-zmb2ypjdc01
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
	// The cluster ID.
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The node groups.
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// CC9FEF89-9BE5-5E03-845E-238B48D7599B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i159136551662516768776
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShrinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type StopInvocationRequest struct {
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The nodes.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetNodeIdList(v []*string) *StopInvocationRequest {
	s.NodeIdList = v
	return s
}

type StopInvocationShrinkRequest struct {
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The nodes.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
}

func (s StopInvocationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationShrinkRequest) SetInvokeId(v string) *StopInvocationShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationShrinkRequest) SetNodeIdListShrink(v string) *StopInvocationShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

type StopInvocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A7FD7411-9395-52E8-AF42-EB3A4A55446D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

type StopInvocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetHeaders(v map[string]*string) *StopInvocationResponse {
	s.Headers = v
	return s
}

func (s *StopInvocationResponse) SetStatusCode(v int32) *StopInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInvocationResponse) SetBody(v *StopInvocationResponseBody) *StopInvocationResponse {
	s.Body = v
	return s
}

type StopNodesRequest struct {
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s StopNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesRequest) SetNodes(v []*string) *StopNodesRequest {
	s.Nodes = v
	return s
}

type StopNodesShrinkRequest struct {
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s StopNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesShrinkRequest) SetNodesShrink(v string) *StopNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

type StopNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// i155847351716171893489
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

type StopNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetStatusCode(v int32) *StopNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Node
	//
	// 	- Vcc
	//
	// 	- Cluster
	//
	// 	- Vpd
	//
	// 	- Subnet
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key.
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v3
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
	// The request ID.
	//
	// example:
	//
	// E7BB53E1-0B08-5C4E-BA66-9225548C3151
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to remove all tags. This parameter takes effect only when TagKey.N is not specified. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// Default value: false.
	//
	// example:
	//
	// False
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Node
	//
	// 	- Cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	// The request ID.
	//
	// example:
	//
	// 81F648D0-5570-5351-AE98-6F501C7E957F
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateNodeGroupRequest struct {
	// Whether file storage mounting is supported.
	//
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The default image ID of the node group. if you do not set this parameter, the image ID will not change.
	//
	// example:
	//
	// i123847249284734
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// sc-key
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// 节点组内机器的登录密码
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node group name
	//
	// example:
	//
	// test-update
	NewNodeGroupName *string `json:"NewNodeGroupName,omitempty" xml:"NewNodeGroupName,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// i120021051733814190732
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The custom script.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetImageId(v string) *UpdateNodeGroupRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetKeyPairName(v string) *UpdateNodeGroupRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetLoginPassword(v string) *UpdateNodeGroupRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNewNodeGroupName(v string) *UpdateNodeGroupRequest {
	s.NewNodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNodeGroupId(v string) *UpdateNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetUserData(v string) *UpdateNodeGroupRequest {
	s.UserData = &v
	return s
}

type UpdateNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// i15374011238111706
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponseBody) SetRequestId(v string) *UpdateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) SetTaskId(v string) *UpdateNodeGroupResponseBody {
	s.TaskId = &v
	return s
}

type UpdateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponse) SetHeaders(v map[string]*string) *UpdateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeGroupResponse) SetStatusCode(v int32) *UpdateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeGroupResponse) SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse {
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

// Summary:
//
// Approves an O\\&M operation.
//
// @param request - ApproveOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperationWithOptions(request *ApproveOperationRequest, runtime *util.RuntimeOptions) (_result *ApproveOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		body["OperationType"] = request.OperationType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveOperation"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Approves an O\\&M operation.
//
// @param request - ApproveOperationRequest
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperation(request *ApproveOperationRequest) (_result *ApproveOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOperationResponse{}
	_body, _err := client.ApproveOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a resource from one resource group to another.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
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

// Summary:
//
// Moves a resource from one resource group to another.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
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

// Summary:
//
// 断开连接
//
// @param request - CloseSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseSessionResponse
func (client *Client) CloseSessionWithOptions(request *CloseSessionRequest, runtime *util.RuntimeOptions) (_result *CloseSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		body["SessionToken"] = request.SessionToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseSession"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 断开连接
//
// @param request - CloseSessionRequest
//
// @return CloseSessionResponse
func (client *Client) CloseSession(request *CloseSessionRequest) (_result *CloseSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseSessionResponse{}
	_body, _err := client.CloseSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Lingjun cluster.
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
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

	if !tea.BoolValue(util.IsUnset(tmpReq.NimizVSwitches)) {
		request.NimizVSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NimizVSwitches, tea.String("NimizVSwitches"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.HpnZone)) {
		body["HpnZone"] = request.HpnZone
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NetworksShrink)) {
		body["Networks"] = request.NetworksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NimizVSwitchesShrink)) {
		body["NimizVSwitches"] = request.NimizVSwitchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpenEniJumboFrame)) {
		body["OpenEniJumboFrame"] = request.OpenEniJumboFrame
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

// Summary:
//
// Creates a Lingjun cluster.
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
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

// Summary:
//
// Creates a diagnostics task.
//
// @param tmpReq - CreateDiagnosticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTaskWithOptions(tmpReq *CreateDiagnosticTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosticTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDiagnosticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AiJobLogInfo)) {
		request.AiJobLogInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AiJobLogInfo, tea.String("AiJobLogInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIds)) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, tea.String("NodeIds"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiJobLogInfoShrink)) {
		body["AiJobLogInfo"] = request.AiJobLogInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DiagnosticType)) {
		body["DiagnosticType"] = request.DiagnosticType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdsShrink)) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnosticTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosticTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostics task.
//
// @param request - CreateDiagnosticTaskRequest
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTask(request *CreateDiagnosticTaskRequest) (_result *CreateDiagnosticTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosticTaskResponse{}
	_body, _err := client.CreateDiagnosticTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a network test task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param tmpReq - CreateNetTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetTestTaskResponse
func (client *Client) CreateNetTestTaskWithOptions(tmpReq *CreateNetTestTaskRequest, runtime *util.RuntimeOptions) (_result *CreateNetTestTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNetTestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CommTest)) {
		request.CommTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommTest, tea.String("CommTest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DelayTest)) {
		request.DelayTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DelayTest, tea.String("DelayTest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TrafficTest)) {
		request.TrafficTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrafficTest, tea.String("TrafficTest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.CommTestShrink)) {
		body["CommTest"] = request.CommTestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DelayTestShrink)) {
		body["DelayTest"] = request.DelayTestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NetTestType)) {
		body["NetTestType"] = request.NetTestType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkMode)) {
		body["NetworkMode"] = request.NetworkMode
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficTestShrink)) {
		body["TrafficTest"] = request.TrafficTestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetTestTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetTestTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network test task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - CreateNetTestTaskRequest
//
// @return CreateNetTestTaskResponse
func (client *Client) CreateNetTestTask(request *CreateNetTestTaskRequest) (_result *CreateNetTestTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetTestTaskResponse{}
	_body, _err := client.CreateNetTestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建集群下的节点分组
//
// @param tmpReq - CreateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroupWithOptions(tmpReq *CreateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodeGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroup)) {
		request.NodeGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroup, tea.String("NodeGroup"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeUnit)) {
		request.NodeUnitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeUnit, tea.String("NodeUnit"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupShrink)) {
		body["NodeGroup"] = request.NodeGroupShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeUnitShrink)) {
		body["NodeUnit"] = request.NodeUnitShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建集群下的节点分组
//
// @param request - CreateNodeGroupRequest
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroup(request *CreateNodeGroupRequest) (_result *CreateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CreateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Web Terminal session.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - CreateSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, runtime *util.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		body["SessionType"] = request.SessionType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSession"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Web Terminal session.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - CreateSessionRequest
//
// @return CreateSessionResponse
func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual storage channel (VSC).
//
// @param request - CreateVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVscResponse
func (client *Client) CreateVscWithOptions(request *CreateVscRequest, runtime *util.RuntimeOptions) (_result *CreateVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VscName)) {
		body["VscName"] = request.VscName
	}

	if !tea.BoolValue(util.IsUnset(request.VscType)) {
		body["VscType"] = request.VscType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual storage channel (VSC).
//
// @param request - CreateVscRequest
//
// @return CreateVscResponse
func (client *Client) CreateVsc(request *CreateVscRequest) (_result *CreateVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVscResponse{}
	_body, _err := client.CreateVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Lingjun cluster.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
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

// Summary:
//
// Deletes a Lingjun cluster.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
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

// Summary:
//
// 删除节点分组
//
// @param request - DeleteNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroupWithOptions(request *DeleteNodeGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除节点分组
//
// @param request - DeleteNodeGroupRequest
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroup(request *DeleteNodeGroupRequest) (_result *DeleteNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.DeleteNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual storage channel (VSC).
//
// @param request - DeleteVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVscResponse
func (client *Client) DeleteVscWithOptions(request *DeleteVscRequest, runtime *util.RuntimeOptions) (_result *DeleteVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VscId)) {
		body["VscId"] = request.VscId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual storage channel (VSC).
//
// @param request - DeleteVscRequest
//
// @return DeleteVscResponse
func (client *Client) DeleteVsc(request *DeleteVscRequest) (_result *DeleteVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVscResponse{}
	_body, _err := client.DeleteVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询灵骏集群详情。
//
// @param request - DescribeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResponse
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

// Summary:
//
// 查询灵骏集群详情。
//
// @param request - DescribeClusterRequest
//
// @return DescribeClusterResponse
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

// Summary:
//
// Queries the results of a diagnostic task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - DescribeDiagnosticResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResultWithOptions(request *DescribeDiagnosticResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosticResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnosticId)) {
		body["DiagnosticId"] = request.DiagnosticId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosticResult"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosticResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a diagnostic task.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - DescribeDiagnosticResultRequest
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResult(request *DescribeDiagnosticResultRequest) (_result *DescribeDiagnosticResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosticResultResponse{}
	_body, _err := client.DescribeDiagnosticResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution list and status of O\\&M Assistant commands.
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOutput)) {
		body["IncludeOutput"] = request.IncludeOutput
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution list and status of O\\&M Assistant commands.
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网络测试结果
//
// @param request - DescribeNetTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetTestResultResponse
func (client *Client) DescribeNetTestResultWithOptions(request *DescribeNetTestResultRequest, runtime *util.RuntimeOptions) (_result *DescribeNetTestResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TestId)) {
		body["TestId"] = request.TestId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNetTestResult"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNetTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网络测试结果
//
// @param request - DescribeNetTestResultRequest
//
// @return DescribeNetTestResultResponse
func (client *Client) DescribeNetTestResult(request *DescribeNetTestResultRequest) (_result *DescribeNetTestResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetTestResultResponse{}
	_body, _err := client.DescribeNetTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of nodes.
//
// @param request - DescribeNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeResponse
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

// Summary:
//
// Queries a list of nodes.
//
// @param request - DescribeNodeRequest
//
// @return DescribeNodeResponse
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

// Summary:
//
// Queries a list of regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries a list of regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries the files that are sent by an O\\&M assistant and the status of the files.
//
// @param request - DescribeSendFileResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResultsWithOptions(request *DescribeSendFileResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeSendFileResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSendFileResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSendFileResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the files that are sent by an O\\&M assistant and the status of the files.
//
// @param request - DescribeSendFileResultsRequest
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResults(request *DescribeSendFileResultsRequest) (_result *DescribeSendFileResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSendFileResultsResponse{}
	_body, _err := client.DescribeSendFileResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a task.
//
// @param request - DescribeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
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

// Summary:
//
// Queries the details of a task.
//
// @param request - DescribeTaskRequest
//
// @return DescribeTaskResponse
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

// Summary:
//
// Queries information about a virtual storage channel (VSC).
//
// @param request - DescribeVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVscResponse
func (client *Client) DescribeVscWithOptions(request *DescribeVscRequest, runtime *util.RuntimeOptions) (_result *DescribeVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VscId)) {
		body["VscId"] = request.VscId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a virtual storage channel (VSC).
//
// @param request - DescribeVscRequest
//
// @return DescribeVscResponse
func (client *Client) DescribeVsc(request *DescribeVscRequest) (_result *DescribeVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVscResponse{}
	_body, _err := client.DescribeVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of zones.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
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

// Summary:
//
// Queries a list of zones.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
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

// Summary:
//
// Scales out a cluster.
//
// @param tmpReq - ExtendClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtendClusterResponse
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

	if !tea.BoolValue(util.IsUnset(request.VSwitchZoneId)) {
		body["VSwitchZoneId"] = request.VSwitchZoneId
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

// Summary:
//
// Scales out a cluster.
//
// @param request - ExtendClusterRequest
//
// @return ExtendClusterResponse
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

// Summary:
//
// Queries a list of nodes in a cluster.
//
// @param request - ListClusterNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterNodesResponse
func (client *Client) ListClusterNodesWithOptions(request *ListClusterNodesRequest, runtime *util.RuntimeOptions) (_result *ListClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

// Summary:
//
// Queries a list of nodes in a cluster.
//
// @param request - ListClusterNodesRequest
//
// @return ListClusterNodesResponse
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

// Summary:
//
// Queries a list of clusters.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

// Summary:
//
// Queries a list of clusters.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
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

// Summary:
//
// 诊断任务列表
//
// @param request - ListDiagnosticResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosticResultsResponse
func (client *Client) ListDiagnosticResultsWithOptions(request *ListDiagnosticResultsRequest, runtime *util.RuntimeOptions) (_result *ListDiagnosticResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagType)) {
		body["DiagType"] = request.DiagType
	}

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
		Action:      tea.String("ListDiagnosticResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnosticResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 诊断任务列表
//
// @param request - ListDiagnosticResultsRequest
//
// @return ListDiagnosticResultsResponse
func (client *Client) ListDiagnosticResults(request *ListDiagnosticResultsRequest) (_result *ListDiagnosticResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiagnosticResultsResponse{}
	_body, _err := client.ListDiagnosticResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of nodes that are not used.
//
// @param request - ListFreeNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreeNodesResponse
func (client *Client) ListFreeNodesWithOptions(request *ListFreeNodesRequest, runtime *util.RuntimeOptions) (_result *ListFreeNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HpnZone)) {
		body["HpnZone"] = request.HpnZone
	}

	if !tea.BoolValue(util.IsUnset(request.MachineType)) {
		body["MachineType"] = request.MachineType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingStates)) {
		body["OperatingStates"] = request.OperatingStates
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

// Summary:
//
// Queries a list of nodes that are not used.
//
// @param request - ListFreeNodesRequest
//
// @return ListFreeNodesResponse
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

// Summary:
//
// Lists available images.
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Architecture)) {
		body["Architecture"] = request.Architecture
	}

	if !tea.BoolValue(util.IsUnset(request.ImageVersion)) {
		body["ImageVersion"] = request.ImageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists available images.
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 使用HPNZone和机型查询机型网络配置
//
// @param tmpReq - ListMachineNetworkInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineNetworkInfoResponse
func (client *Client) ListMachineNetworkInfoWithOptions(tmpReq *ListMachineNetworkInfoRequest, runtime *util.RuntimeOptions) (_result *ListMachineNetworkInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMachineNetworkInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MachineHpnInfo)) {
		request.MachineHpnInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MachineHpnInfo, tea.String("MachineHpnInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MachineHpnInfoShrink)) {
		body["MachineHpnInfo"] = request.MachineHpnInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineNetworkInfo"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineNetworkInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 使用HPNZone和机型查询机型网络配置
//
// @param request - ListMachineNetworkInfoRequest
//
// @return ListMachineNetworkInfoResponse
func (client *Client) ListMachineNetworkInfo(request *ListMachineNetworkInfoRequest) (_result *ListMachineNetworkInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMachineNetworkInfoResponse{}
	_body, _err := client.ListMachineNetworkInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instance types that are available to users.
//
// @param request - ListMachineTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypesWithOptions(request *ListMachineTypesRequest, runtime *util.RuntimeOptions) (_result *ListMachineTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineTypes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instance types that are available to users.
//
// @param request - ListMachineTypesRequest
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypes(request *ListMachineTypesRequest) (_result *ListMachineTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMachineTypesResponse{}
	_body, _err := client.ListMachineTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the results of network test results.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - ListNetTestResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetTestResultsResponse
func (client *Client) ListNetTestResultsWithOptions(request *ListNetTestResultsRequest, runtime *util.RuntimeOptions) (_result *ListNetTestResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NetTestType)) {
		body["NetTestType"] = request.NetTestType
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
		Action:      tea.String("ListNetTestResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetTestResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the results of network test results.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - ListNetTestResultsRequest
//
// @return ListNetTestResultsResponse
func (client *Client) ListNetTestResults(request *ListNetTestResultsRequest) (_result *ListNetTestResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetTestResultsResponse{}
	_body, _err := client.ListNetTestResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries node groups in a cluster.
//
// @param request - ListNodeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
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
		Action:      tea.String("ListNodeGroups"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries node groups in a cluster.
//
// @param request - ListNodeGroupsRequest
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries the tags of resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// 查询用户可以使用的集群类型
//
// @param request - ListUserClusterTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserClusterTypesResponse
func (client *Client) ListUserClusterTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListUserClusterTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListUserClusterTypes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserClusterTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户可以使用的集群类型
//
// @return ListUserClusterTypesResponse
func (client *Client) ListUserClusterTypes() (_result *ListUserClusterTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserClusterTypesResponse{}
	_body, _err := client.ListUserClusterTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of virtual storage channels (VSC).
//
// @param tmpReq - ListVscsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVscsResponse
func (client *Client) ListVscsWithOptions(tmpReq *ListVscsRequest, runtime *util.RuntimeOptions) (_result *ListVscsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListVscsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIds)) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, tea.String("NodeIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdsShrink)) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VscName)) {
		body["VscName"] = request.VscName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVscs"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVscsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of virtual storage channels (VSC).
//
// @param request - ListVscsRequest
//
// @return ListVscsResponse
func (client *Client) ListVscs(request *ListVscsRequest) (_result *ListVscsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVscsResponse{}
	_body, _err := client.ListVscsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts nodes.
//
// @param tmpReq - RebootNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootNodesResponse
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

// Summary:
//
// Restarts nodes.
//
// @param request - RebootNodesRequest
//
// @return RebootNodesResponse
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

// Summary:
//
// Reinstall a node.
//
// @param tmpReq - ReimageNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReimageNodesResponse
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

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
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

// Summary:
//
// Reinstall a node.
//
// @param request - ReimageNodesRequest
//
// @return ReimageNodesResponse
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

// Summary:
//
// Runs a Shell script on one or more Lingjun nodes.
//
// @param tmpReq - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithOptions(tmpReq *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		body["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		body["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableParameter)) {
		body["EnableParameter"] = request.EnableParameter
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		body["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.Launcher)) {
		body["Launcher"] = request.Launcher
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatMode)) {
		body["RepeatMode"] = request.RepeatMode
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationMode)) {
		body["TerminationMode"] = request.TerminationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		body["WorkingDir"] = request.WorkingDir
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a Shell script on one or more Lingjun nodes.
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a file to one or more Lingjun nodes.
//
// @param tmpReq - SendFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFileResponse
func (client *Client) SendFileWithOptions(tmpReq *SendFileRequest, runtime *util.RuntimeOptions) (_result *SendFileResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileGroup)) {
		body["FileGroup"] = request.FileGroup
	}

	if !tea.BoolValue(util.IsUnset(request.FileMode)) {
		body["FileMode"] = request.FileMode
	}

	if !tea.BoolValue(util.IsUnset(request.FileOwner)) {
		body["FileOwner"] = request.FileOwner
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Overwrite)) {
		body["Overwrite"] = request.Overwrite
	}

	if !tea.BoolValue(util.IsUnset(request.TargetDir)) {
		body["TargetDir"] = request.TargetDir
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendFile"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a file to one or more Lingjun nodes.
//
// @param request - SendFileRequest
//
// @return SendFileResponse
func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendFileResponse{}
	_body, _err := client.SendFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales in a cluster.
//
// @param tmpReq - ShrinkClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkClusterResponse
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

// Summary:
//
// Scales in a cluster.
//
// @param request - ShrinkClusterRequest
//
// @return ShrinkClusterResponse
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

// Summary:
//
// Stops the O\\&M assistant command execution.
//
// @param tmpReq - StopInvocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInvocationResponse
func (client *Client) StopInvocationWithOptions(tmpReq *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopInvocationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInvocation"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the O\\&M assistant command execution.
//
// @param request - StopInvocationRequest
//
// @return StopInvocationResponse
func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops nodes.
//
// @param tmpReq - StopNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopNodesResponse
func (client *Client) StopNodesWithOptions(tmpReq *StopNodesRequest, runtime *util.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
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
		Action:      tea.String("StopNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops nodes.
//
// @param request - StopNodesRequest
//
// @return StopNodesResponse
func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopNodesResponse{}
	_body, _err := client.StopNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tags resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
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

// Summary:
//
// Tags resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Deletes a custom tag from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
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

// Summary:
//
// Deletes a custom tag from a resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// Updates a node group.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - UpdateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroupWithOptions(request *UpdateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemMountEnabled)) {
		body["FileSystemMountEnabled"] = request.FileSystemMountEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		body["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.LoginPassword)) {
		body["LoginPassword"] = request.LoginPassword
	}

	if !tea.BoolValue(util.IsUnset(request.NewNodeGroupName)) {
		body["NewNodeGroupName"] = request.NewNodeGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a node group.
//
// Description:
//
// The API creates a session, returns the frontend endpoint, and starts a periodic task to track the session status.
//
// @param request - UpdateNodeGroupRequest
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroup(request *UpdateNodeGroupRequest) (_result *UpdateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.UpdateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
