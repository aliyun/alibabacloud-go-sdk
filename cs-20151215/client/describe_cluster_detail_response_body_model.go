// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoMode(v *DescribeClusterDetailResponseBodyAutoMode) *DescribeClusterDetailResponseBody
	GetAutoMode() *DescribeClusterDetailResponseBodyAutoMode
	SetClusterDomain(v string) *DescribeClusterDetailResponseBody
	GetClusterDomain() *string
	SetClusterId(v string) *DescribeClusterDetailResponseBody
	GetClusterId() *string
	SetClusterSpec(v string) *DescribeClusterDetailResponseBody
	GetClusterSpec() *string
	SetClusterType(v string) *DescribeClusterDetailResponseBody
	GetClusterType() *string
	SetContainerCidr(v string) *DescribeClusterDetailResponseBody
	GetContainerCidr() *string
	SetControlPlaneConfig(v *DescribeClusterDetailResponseBodyControlPlaneConfig) *DescribeClusterDetailResponseBody
	GetControlPlaneConfig() *DescribeClusterDetailResponseBodyControlPlaneConfig
	SetControlPlaneEndpointsConfig(v *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) *DescribeClusterDetailResponseBody
	GetControlPlaneEndpointsConfig() *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig
	SetCreated(v string) *DescribeClusterDetailResponseBody
	GetCreated() *string
	SetCurrentVersion(v string) *DescribeClusterDetailResponseBody
	GetCurrentVersion() *string
	SetDeletionProtection(v bool) *DescribeClusterDetailResponseBody
	GetDeletionProtection() *bool
	SetDockerVersion(v string) *DescribeClusterDetailResponseBody
	GetDockerVersion() *string
	SetExternalLoadbalancerId(v string) *DescribeClusterDetailResponseBody
	GetExternalLoadbalancerId() *string
	SetExtraSans(v []*string) *DescribeClusterDetailResponseBody
	GetExtraSans() []*string
	SetInitVersion(v string) *DescribeClusterDetailResponseBody
	GetInitVersion() *string
	SetIpStack(v string) *DescribeClusterDetailResponseBody
	GetIpStack() *string
	SetMaintenanceWindow(v *MaintenanceWindow) *DescribeClusterDetailResponseBody
	GetMaintenanceWindow() *MaintenanceWindow
	SetMasterUrl(v string) *DescribeClusterDetailResponseBody
	GetMasterUrl() *string
	SetMetaData(v string) *DescribeClusterDetailResponseBody
	GetMetaData() *string
	SetName(v string) *DescribeClusterDetailResponseBody
	GetName() *string
	SetNetworkMode(v string) *DescribeClusterDetailResponseBody
	GetNetworkMode() *string
	SetNextVersion(v string) *DescribeClusterDetailResponseBody
	GetNextVersion() *string
	SetNodeCidrMask(v string) *DescribeClusterDetailResponseBody
	GetNodeCidrMask() *string
	SetOperationPolicy(v *DescribeClusterDetailResponseBodyOperationPolicy) *DescribeClusterDetailResponseBody
	GetOperationPolicy() *DescribeClusterDetailResponseBodyOperationPolicy
	SetParameters(v map[string]*string) *DescribeClusterDetailResponseBody
	GetParameters() map[string]*string
	SetPrivateZone(v bool) *DescribeClusterDetailResponseBody
	GetPrivateZone() *bool
	SetProfile(v string) *DescribeClusterDetailResponseBody
	GetProfile() *string
	SetProxyMode(v string) *DescribeClusterDetailResponseBody
	GetProxyMode() *string
	SetRegionId(v string) *DescribeClusterDetailResponseBody
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeClusterDetailResponseBody
	GetResourceGroupId() *string
	SetRrsaConfig(v *DescribeClusterDetailResponseBodyRrsaConfig) *DescribeClusterDetailResponseBody
	GetRrsaConfig() *DescribeClusterDetailResponseBodyRrsaConfig
	SetSecurityGroupId(v string) *DescribeClusterDetailResponseBody
	GetSecurityGroupId() *string
	SetServiceCidr(v string) *DescribeClusterDetailResponseBody
	GetServiceCidr() *string
	SetSize(v int64) *DescribeClusterDetailResponseBody
	GetSize() *int64
	SetState(v string) *DescribeClusterDetailResponseBody
	GetState() *string
	SetSubnetCidr(v string) *DescribeClusterDetailResponseBody
	GetSubnetCidr() *string
	SetTags(v []*Tag) *DescribeClusterDetailResponseBody
	GetTags() []*Tag
	SetTimezone(v string) *DescribeClusterDetailResponseBody
	GetTimezone() *string
	SetUpdated(v string) *DescribeClusterDetailResponseBody
	GetUpdated() *string
	SetVpcId(v string) *DescribeClusterDetailResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *DescribeClusterDetailResponseBody
	GetVswitchId() *string
	SetVswitchIds(v []*string) *DescribeClusterDetailResponseBody
	GetVswitchIds() []*string
	SetWorkerRamRoleName(v string) *DescribeClusterDetailResponseBody
	GetWorkerRamRoleName() *string
	SetZoneId(v string) *DescribeClusterDetailResponseBody
	GetZoneId() *string
}

type DescribeClusterDetailResponseBody struct {
	AutoMode *DescribeClusterDetailResponseBodyAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The domain name of the cluster.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The edition of the cluster
	//
	// 	- `ack.pro.small`: the Pro edition.
	//
	// 	- `ack.standard`: the Basic edition.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the instance.
	//
	// 	- `Kubernetes`: ACK dedicated cluster.
	//
	// 	- `ManagedKubernetes`: ACK managed cluster. ACK managed clusters include ACK managed Basic clusters, ACK managed Pro clusters, ACK Serverless Pro clusters, ACK Serverless Basic clusters, ACK Edge Pro clusters, ACK Edge Basic clusters, and ACK Lingjun Pro clusters.
	//
	// 	- `ExternalKubernetes`: registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The pod CIDR block. The configuration of the Flannel network plug-in.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The control plane configurations in an ACK dedicated cluster.
	ControlPlaneConfig          *DescribeClusterDetailResponseBodyControlPlaneConfig          `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	ControlPlaneEndpointsConfig *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2019-11-25T15:50:20+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The Kubernetes version of the cluster. For more information about the Kubernetes versions supported by ACK, see [Release notes for Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.16.6-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster. If deletion protection is enabled, the cluster cannot be deleted in the Container Service console or by calling API operations. Valid values:
	//
	// 	- `true`: deletion protection is enabled for the cluster. This way, the cluster cannot be deleted in the Container Service console or by calling API operations.
	//
	// 	- `false`: deletion protection is disabled for the cluster. This way, the cluster can be deleted in the Container Service console or by calling API operations.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// The Docker version that is used by the cluster.
	//
	// example:
	//
	// 19.03.5
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// Deprecated
	//
	// The ID of the Server Load Balancer (SLB) instance that is created for the Ingress of the cluster.
	//
	// example:
	//
	// lb-2zehc05z3b8dwiifh****
	ExternalLoadbalancerId *string   `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	ExtraSans              []*string `json:"extra_sans,omitempty" xml:"extra_sans,omitempty" type:"Repeated"`
	// The initial Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.16.6-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The IP stack of the cluster. Valid values:
	//
	// 	- ipv4: The cluster is an IPv4 cluster.
	//
	// 	- dual: The cluster is an IPv4/IPv6 dual-stack cluster.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// The maintenance window of the cluster. This feature is available only in ACK Pro clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The endpoints of the cluster, including an internal endpoint and a public endpoint.
	//
	// example:
	//
	// {\\"intranet_api_server_endpoint\\":\\"https://192.168.0.251:6443\\"***}
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// The metadata of the cluster.
	//
	// example:
	//
	// \\"Addons\\":***
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Deprecated
	//
	// The network type of the cluster. Example: Virtual Private Cloud (VPC).
	//
	// example:
	//
	// vpc
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// The Kubernetes version to which the cluster can be upgraded.
	//
	// example:
	//
	// 1.18.8-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// This parameter is available only for Flannel.
	//
	// The subnet mask length of the node CIDR block. This parameter indicates the maximum number of IP addresses that can be assigned to nodes.
	//
	// example:
	//
	// 26
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// The automatic O\\&M policy of the cluster.
	OperationPolicy *DescribeClusterDetailResponseBodyOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// The Resource Orchestration Service (ROS) parameters of the cluster.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// Deprecated
	//
	// Indicates whether Alibaba Cloud DNS PrivateZone (PrivateZone) is enabled for the cluster. Valid values:
	//
	// 	- `true`: PrivateZone is enabled.
	//
	// 	- `false`: PrivateZone is dislabled.
	//
	// Default value: false
	//
	// example:
	//
	// false
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// The subtype of the cluster.
	//
	// 	- `Default`. ACK managed cluster. ACK managed clusters include ACK Basic clusters and ACK Pro clusters.
	//
	// 	- `Edge`: ACK Edge cluster. ACK Edge clusters include ACK Edge Basic clusters and ACK Edge Pro clusters.
	//
	// 	- `Serverless`: ACK Serverless cluster. ACK Serverless clusters include ACK Serverless Basic clusters and ACK Serverless Pro clusters.
	//
	// 	- `Lingjun`: ACK Lingjun Pro cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy mode. Valid values:
	//
	// 	- `iptables`: a mature and stable kube-proxy mode that uses iptables rules to conduct Service discovery and load balancing. The performance of this mode is limited by the size of the cluster. This mode is suitable for clusters that run a small number of Services.
	//
	// 	- `ipvs`: provides high performance and uses IP Virtual Server (IPVS). This allows you to configure service discovery and load balancing. This mode is suitable for clusters that are required to run a large number of services. We recommend that you use this mode in scenarios that require high load balancing performance.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string                                      `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	RrsaConfig      *DescribeClusterDetailResponseBodyRrsaConfig `json:"rrsa_config,omitempty" xml:"rrsa_config,omitempty" type:"Struct"`
	// The ID of the security group to which the cluster belongs.
	//
	// example:
	//
	// sg-25yq****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The Service CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The number of nodes in the cluster. Master nodes and worker nodes are included.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- `initial`: The cluster is being created.
	//
	// 	- `failed`: The cluster failed to be created.
	//
	// 	- `running`: The cluster is running.
	//
	// 	- `updating`: The cluster is being updated.
	//
	// 	- `updating_failed`: The cluster failed to be updated.
	//
	// 	- `scaling`: The cluster is being scaled.
	//
	// 	- `waiting`: The cluster is waiting for connection requests.
	//
	// 	- `disconnected`: The cluster is disconnected.
	//
	// 	- `stopped`: The cluster is stopped.
	//
	// 	- `deleting`: The cluster is being deleted.
	//
	// 	- `deleted`: The cluster is deleted.
	//
	// 	- `delete_failed`: The cluster failed to be deleted.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Deprecated
	//
	// The pod CIDR block.
	//
	// example:
	//
	// 172.20.0.0/16
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// The resource labels of the cluster.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time zone
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2020-01-13T23:01:03+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The ID of the VPC where the cluster is deployed. This parameter is required when you create a cluster.
	//
	// example:
	//
	// vpc-2zecuu62b9zw7a7qn****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitche. This field is deprecated. Use vswitch_ids to query the vSwitches on the control plane and vswitch_ids to query the vSwitches on the data plane.
	//
	// example:
	//
	// vsw-2zete8s4qocqg0mf6****,vsw-2zete8s4qocqg0mf6****
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// The vSwitch for the control plane of the cluster.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// The name of the worker Resource Access Management (RAM) role. The RAM role is assigned to the worker nodes of the cluster to allow the worker nodes to manage Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// KubernetesWorkerRole-ec87d15b-edca-4302-933f-c8a16bf0****
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// Deprecated
	//
	// The ID of the zone within the region where the cluster is located.
	//
	// example:
	//
	// cn-beijing-a
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClusterDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBody) GetAutoMode() *DescribeClusterDetailResponseBodyAutoMode {
	return s.AutoMode
}

func (s *DescribeClusterDetailResponseBody) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeClusterDetailResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterDetailResponseBody) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeClusterDetailResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterDetailResponseBody) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *DescribeClusterDetailResponseBody) GetControlPlaneConfig() *DescribeClusterDetailResponseBodyControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *DescribeClusterDetailResponseBody) GetControlPlaneEndpointsConfig() *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig {
	return s.ControlPlaneEndpointsConfig
}

func (s *DescribeClusterDetailResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterDetailResponseBody) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClusterDetailResponseBody) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeClusterDetailResponseBody) GetDockerVersion() *string {
	return s.DockerVersion
}

func (s *DescribeClusterDetailResponseBody) GetExternalLoadbalancerId() *string {
	return s.ExternalLoadbalancerId
}

func (s *DescribeClusterDetailResponseBody) GetExtraSans() []*string {
	return s.ExtraSans
}

func (s *DescribeClusterDetailResponseBody) GetInitVersion() *string {
	return s.InitVersion
}

func (s *DescribeClusterDetailResponseBody) GetIpStack() *string {
	return s.IpStack
}

func (s *DescribeClusterDetailResponseBody) GetMaintenanceWindow() *MaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *DescribeClusterDetailResponseBody) GetMasterUrl() *string {
	return s.MasterUrl
}

func (s *DescribeClusterDetailResponseBody) GetMetaData() *string {
	return s.MetaData
}

func (s *DescribeClusterDetailResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeClusterDetailResponseBody) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *DescribeClusterDetailResponseBody) GetNextVersion() *string {
	return s.NextVersion
}

func (s *DescribeClusterDetailResponseBody) GetNodeCidrMask() *string {
	return s.NodeCidrMask
}

func (s *DescribeClusterDetailResponseBody) GetOperationPolicy() *DescribeClusterDetailResponseBodyOperationPolicy {
	return s.OperationPolicy
}

func (s *DescribeClusterDetailResponseBody) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *DescribeClusterDetailResponseBody) GetPrivateZone() *bool {
	return s.PrivateZone
}

func (s *DescribeClusterDetailResponseBody) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClusterDetailResponseBody) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *DescribeClusterDetailResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterDetailResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterDetailResponseBody) GetRrsaConfig() *DescribeClusterDetailResponseBodyRrsaConfig {
	return s.RrsaConfig
}

func (s *DescribeClusterDetailResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClusterDetailResponseBody) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *DescribeClusterDetailResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClusterDetailResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeClusterDetailResponseBody) GetSubnetCidr() *string {
	return s.SubnetCidr
}

func (s *DescribeClusterDetailResponseBody) GetTags() []*Tag {
	return s.Tags
}

func (s *DescribeClusterDetailResponseBody) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeClusterDetailResponseBody) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClusterDetailResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterDetailResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeClusterDetailResponseBody) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClusterDetailResponseBody) GetWorkerRamRoleName() *string {
	return s.WorkerRamRoleName
}

func (s *DescribeClusterDetailResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterDetailResponseBody) SetAutoMode(v *DescribeClusterDetailResponseBodyAutoMode) *DescribeClusterDetailResponseBody {
	s.AutoMode = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterDomain(v string) *DescribeClusterDetailResponseBody {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterId(v string) *DescribeClusterDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterSpec(v string) *DescribeClusterDetailResponseBody {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterType(v string) *DescribeClusterDetailResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetContainerCidr(v string) *DescribeClusterDetailResponseBody {
	s.ContainerCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetControlPlaneConfig(v *DescribeClusterDetailResponseBodyControlPlaneConfig) *DescribeClusterDetailResponseBody {
	s.ControlPlaneConfig = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetControlPlaneEndpointsConfig(v *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) *DescribeClusterDetailResponseBody {
	s.ControlPlaneEndpointsConfig = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCreated(v string) *DescribeClusterDetailResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCurrentVersion(v string) *DescribeClusterDetailResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDeletionProtection(v bool) *DescribeClusterDetailResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDockerVersion(v string) *DescribeClusterDetailResponseBody {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetExternalLoadbalancerId(v string) *DescribeClusterDetailResponseBody {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetExtraSans(v []*string) *DescribeClusterDetailResponseBody {
	s.ExtraSans = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetInitVersion(v string) *DescribeClusterDetailResponseBody {
	s.InitVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetIpStack(v string) *DescribeClusterDetailResponseBody {
	s.IpStack = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMaintenanceWindow(v *MaintenanceWindow) *DescribeClusterDetailResponseBody {
	s.MaintenanceWindow = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMasterUrl(v string) *DescribeClusterDetailResponseBody {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMetaData(v string) *DescribeClusterDetailResponseBody {
	s.MetaData = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetName(v string) *DescribeClusterDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNetworkMode(v string) *DescribeClusterDetailResponseBody {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNextVersion(v string) *DescribeClusterDetailResponseBody {
	s.NextVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNodeCidrMask(v string) *DescribeClusterDetailResponseBody {
	s.NodeCidrMask = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetOperationPolicy(v *DescribeClusterDetailResponseBodyOperationPolicy) *DescribeClusterDetailResponseBody {
	s.OperationPolicy = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetParameters(v map[string]*string) *DescribeClusterDetailResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPrivateZone(v bool) *DescribeClusterDetailResponseBody {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetProfile(v string) *DescribeClusterDetailResponseBody {
	s.Profile = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetProxyMode(v string) *DescribeClusterDetailResponseBody {
	s.ProxyMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRegionId(v string) *DescribeClusterDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetResourceGroupId(v string) *DescribeClusterDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRrsaConfig(v *DescribeClusterDetailResponseBodyRrsaConfig) *DescribeClusterDetailResponseBody {
	s.RrsaConfig = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSecurityGroupId(v string) *DescribeClusterDetailResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetServiceCidr(v string) *DescribeClusterDetailResponseBody {
	s.ServiceCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSize(v int64) *DescribeClusterDetailResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetState(v string) *DescribeClusterDetailResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSubnetCidr(v string) *DescribeClusterDetailResponseBody {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetTags(v []*Tag) *DescribeClusterDetailResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetTimezone(v string) *DescribeClusterDetailResponseBody {
	s.Timezone = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetUpdated(v string) *DescribeClusterDetailResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVpcId(v string) *DescribeClusterDetailResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchId(v string) *DescribeClusterDetailResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchIds(v []*string) *DescribeClusterDetailResponseBody {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetWorkerRamRoleName(v string) *DescribeClusterDetailResponseBody {
	s.WorkerRamRoleName = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetZoneId(v string) *DescribeClusterDetailResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) Validate() error {
	if s.AutoMode != nil {
		if err := s.AutoMode.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneConfig != nil {
		if err := s.ControlPlaneConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneEndpointsConfig != nil {
		if err := s.ControlPlaneEndpointsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MaintenanceWindow != nil {
		if err := s.MaintenanceWindow.Validate(); err != nil {
			return err
		}
	}
	if s.OperationPolicy != nil {
		if err := s.OperationPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.RrsaConfig != nil {
		if err := s.RrsaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyAutoMode struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s DescribeClusterDetailResponseBodyAutoMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyAutoMode) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyAutoMode) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterDetailResponseBodyAutoMode) SetEnable(v bool) *DescribeClusterDetailResponseBodyAutoMode {
	s.Enable = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyAutoMode) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyControlPlaneConfig struct {
	// Indicates whether auto-renewal is enabled for the nodes.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal duration for the nodes.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// The billing method of the control plane node.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Indicates whether to install CloudMonitor for the node.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// The CPU management policy of nodes.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp10b35imuam5amw****
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the OS image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType               *string                  `json:"image_type,omitempty" xml:"image_type,omitempty"`
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance types of the nodes.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The name of the key pair. You must set key_pair or login_password.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The node port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// The subscription duration of nodes in the node pool.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The runtime.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Indicates whether to enable Alibaba Cloud Linux Security Hardening.
	//
	// example:
	//
	// true
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// The number of control plane nodes.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Indicates whether to enable Multi-Level Protection Scheme (MLPS) security hardening.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Indicates whether to enable the burst feature for the system disk.
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The category of the system disk for nodes.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The performance level (PL) of the system disk that you want to use for the node. This parameter takes effect only for ESSDs.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The preset read/write IOPS of the system disk.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The system disk size of the node. The value must be at least 40 GB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The automatic snapshot policy of the node.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetCloudMonitorFlags() *bool {
	return s.CloudMonitorFlags
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetInstanceMetadataOptions() *InstanceMetadataOptions {
	return s.InstanceMetadataOptions
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetKeyPair() *string {
	return s.KeyPair
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetAutoRenew(v bool) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetAutoRenewPeriod(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetChargeType(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetCloudMonitorFlags(v bool) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.CloudMonitorFlags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetCpuPolicy(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.CpuPolicy = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetDeploymentsetId(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.DeploymentsetId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetImageId(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetImageType(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.ImageType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetInstanceMetadataOptions(v *InstanceMetadataOptions) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.InstanceMetadataOptions = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetInstanceTypes(v []*string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.InstanceTypes = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetKeyPair(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.KeyPair = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetNodePortRange(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetPeriod(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.Period = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetPeriodUnit(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetRuntime(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.Runtime = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSecurityHardeningOs(v bool) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SecurityHardeningOs = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSize(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSocEnabled(v bool) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SocEnabled = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskBurstingEnabled(v bool) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskCategory(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskPerformanceLevel(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskProvisionedIops(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskSize(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskSnapshotPolicyId(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) Validate() error {
	if s.InstanceMetadataOptions != nil {
		if err := s.InstanceMetadataOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig struct {
	InternalDnsConfig *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig `json:"internal_dns_config,omitempty" xml:"internal_dns_config,omitempty" type:"Struct"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) GetInternalDnsConfig() *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig {
	return s.InternalDnsConfig
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) SetInternalDnsConfig(v *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig {
	s.InternalDnsConfig = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig) Validate() error {
	if s.InternalDnsConfig != nil {
		if err := s.InternalDnsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig struct {
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	Enabled  *bool     `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) GetBindVpcs() []*string {
	return s.BindVpcs
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) SetBindVpcs(v []*string) *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig {
	s.BindVpcs = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) SetEnabled(v bool) *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfigInternalDnsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyOperationPolicy struct {
	// The configurations of auto cluster update.
	ClusterAutoUpgrade *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade `json:"cluster_auto_upgrade,omitempty" xml:"cluster_auto_upgrade,omitempty" type:"Struct"`
}

func (s DescribeClusterDetailResponseBodyOperationPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyOperationPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyOperationPolicy) GetClusterAutoUpgrade() *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade {
	return s.ClusterAutoUpgrade
}

func (s *DescribeClusterDetailResponseBodyOperationPolicy) SetClusterAutoUpgrade(v *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) *DescribeClusterDetailResponseBodyOperationPolicy {
	s.ClusterAutoUpgrade = v
	return s
}

func (s *DescribeClusterDetailResponseBodyOperationPolicy) Validate() error {
	if s.ClusterAutoUpgrade != nil {
		if err := s.ClusterAutoUpgrade.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade struct {
	// The frequency of auto cluster updates. For more information, see [Update frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// 	- patch: specifies the latest patch version.
	//
	// 	- stable: specifies the second-latest minor version.
	//
	// 	- rapid: specifies the latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether to enable auto cluster update.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) GetChannel() *string {
	return s.Channel
}

func (s *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) SetChannel(v string) *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade {
	s.Channel = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) SetEnabled(v bool) *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade {
	s.Enabled = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyOperationPolicyClusterAutoUpgrade) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyRrsaConfig struct {
	Audience                *string `json:"audience,omitempty" xml:"audience,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Issuer                  *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	JwksUrl                 *string `json:"jwks_url,omitempty" xml:"jwks_url,omitempty"`
	MaxOidcTokenExpiration  *string `json:"max_oidc_token_expiration,omitempty" xml:"max_oidc_token_expiration,omitempty"`
	OidcArn                 *string `json:"oidc_arn,omitempty" xml:"oidc_arn,omitempty"`
	OidcName                *string `json:"oidc_name,omitempty" xml:"oidc_name,omitempty"`
	OpenApiConfigurationUrl *string `json:"open_api_configuration_url,omitempty" xml:"open_api_configuration_url,omitempty"`
}

func (s DescribeClusterDetailResponseBodyRrsaConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyRrsaConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetAudience() *string {
	return s.Audience
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetJwksUrl() *string {
	return s.JwksUrl
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetMaxOidcTokenExpiration() *string {
	return s.MaxOidcTokenExpiration
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetOidcArn() *string {
	return s.OidcArn
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetOidcName() *string {
	return s.OidcName
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) GetOpenApiConfigurationUrl() *string {
	return s.OpenApiConfigurationUrl
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetAudience(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.Audience = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetEnabled(v bool) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetIssuer(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.Issuer = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetJwksUrl(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.JwksUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetMaxOidcTokenExpiration(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.MaxOidcTokenExpiration = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetOidcArn(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.OidcArn = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetOidcName(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.OidcName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) SetOpenApiConfigurationUrl(v string) *DescribeClusterDetailResponseBodyRrsaConfig {
	s.OpenApiConfigurationUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyRrsaConfig) Validate() error {
	return dara.Validate(s)
}
