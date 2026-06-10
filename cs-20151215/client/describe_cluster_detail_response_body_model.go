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
	// Smart managed mode configuration.
	AutoMode *DescribeClusterDetailResponseBodyAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// Local domain name of the cluster.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// Cluster ID.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Cluster specification when `cluster_type` is set to `ManagedKubernetes` and `profile` is configured. Valid values:
	//
	// - `ack.standard`: Basic Edition (default if left empty)
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (requires whitelist approval from customer service)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers offered by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). These tiers pre-allocate and dedicate control plane resources to ensure consistent high performance for API concurrency and pod scheduling, making them suitable for AI training and inference, large-scale clusters, and mission-critical workloads.
	//
	// For cluster management fees of Pro Edition and provisioned control plane clusters, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Cluster type.
	//
	// - `Kubernetes`: ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: ACK managed clusters, including ACK Pro Edition and Basic Edition clusters, ACK Serverless clusters (Pro and Basic), ACK Edge clusters (Pro and Basic), and ACK LINGJUN clusters (Pro).
	//
	// - `ExternalKubernetes`: registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// CIDR block for pod networks, used with Flannel.
	//
	// example:
	//
	// 172.20.xx.xx/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// Control plane configuration for dedicated clusters.
	ControlPlaneConfig *DescribeClusterDetailResponseBodyControlPlaneConfig `json:"control_plane_config,omitempty" xml:"control_plane_config,omitempty" type:"Struct"`
	// Cluster connection configuration.
	ControlPlaneEndpointsConfig *DescribeClusterDetailResponseBodyControlPlaneEndpointsConfig `json:"control_plane_endpoints_config,omitempty" xml:"control_plane_endpoints_config,omitempty" type:"Struct"`
	// Time when the cluster was created.
	//
	// example:
	//
	// 2025-04-07T09:57:26+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Current Kubernetes version of the cluster. For supported Kubernetes versions in ACK, see [Overview of Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.32.1-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Deletion protection for the cluster prevents accidental deletion through the console or API. Valid values:
	//
	// - `true`: Deletion protection is enabled. You cannot delete the cluster through the console or API.
	//
	// - `false`: Deletion protection is disabled. You can delete the cluster through the console or API.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// Docker version used in the cluster.
	//
	// example:
	//
	// 19.03.5
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// Deprecated
	//
	// ID of the Server Load Balancer instance used for the cluster Ingress.
	//
	// example:
	//
	// lb-2zehc05z3b8dwiifh****
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// Custom Subject Alternative Names (SANs) for the API server certificate.
	ExtraSans []*string `json:"extra_sans,omitempty" xml:"extra_sans,omitempty" type:"Repeated"`
	// Initial Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// IP protocol stack of the cluster. Valid values:
	//
	// - ipv4: Creates a cluster that supports IPv4 only.
	//
	// - dual: Creates a cluster that supports both IPv4 and IPv6.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// Maintenance window configuration for the cluster. This setting applies only to managed clusters (ACK Pro clusters).
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// Cluster endpoint, including internal and public endpoints.
	//
	// example:
	//
	// {\\"intranet_api_server_endpoint\\":\\"https://192.168.xx.xx:6443\\"***}
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// Metadata of the cluster.
	//
	// example:
	//
	// \\"Addons\\":***
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// Cluster name.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Deprecated
	//
	// Network type used by the cluster, such as VPC.
	//
	// example:
	//
	// vpc
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// Next available Kubernetes version for upgrade.
	//
	// example:
	//
	// 1.xx.x-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// Applies only to Flannel network plugin.
	//
	// Subnet mask size allocated to each node, which controls the number of IP addresses assignable to the node.
	//
	// example:
	//
	// 26
	NodeCidrMask *string `json:"node_cidr_mask,omitempty" xml:"node_cidr_mask,omitempty"`
	// Automatic O\\&M policy for the cluster.
	OperationPolicy *DescribeClusterDetailResponseBodyOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// Collection of ROS parameters for the cluster.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// Deprecated
	//
	// Indicates whether PrivateZone is enabled for the cluster.
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// Cluster subtype.
	//
	// - `Default`: ACK managed cluster, including ACK Pro Edition and Basic Edition.
	//
	// - `Edge`: ACK Edge cluster, including ACK Edge Pro Edition and Basic Edition.
	//
	// - `Serverless`: ACK Serverless cluster, including ACK Serverless Pro Edition and Basic Edition.
	//
	// - `Lingjun`: ACK LINGJUN cluster, available in Pro Edition.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// kube-proxy proxy mode.
	//
	// - `iptables`: A mature and stable kube-proxy mode that uses iptables rules for Kubernetes service discovery and load balancing. Performance is moderate and degrades at scale. Suitable for clusters with a small number of services.
	//
	// - `ipvs`: A high-performance kube-proxy mode that uses the Linux IPVS module for Kubernetes service discovery and load balancing. Suitable for clusters with many services and high load balancing demands.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// Region ID where the cluster is deployed.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// Resource group ID of the cluster.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// RRSA configuration.
	RrsaConfig *DescribeClusterDetailResponseBodyRrsaConfig `json:"rrsa_config,omitempty" xml:"rrsa_config,omitempty" type:"Struct"`
	// Security group ID of the cluster.
	//
	// example:
	//
	// sg-25yq****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// CIDR block for service networks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.xx.xx/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// Total number of nodes in the cluster, including master and worker nodes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Cluster status. Valid values:
	//
	// - `initial`: The cluster is being created.
	//
	// - `failed`: Cluster creation failed.
	//
	// - `running`: The cluster is running.
	//
	// - `updating`: The cluster is being updated.
	//
	// - `upgrading`: The cluster is being upgraded.
	//
	// - `removing`: Nodes are being removed.
	//
	// - `draining`: Nodes are being drained.
	//
	// - `scaling`: The cluster is scaling.
	//
	// - `inactive`: The cluster is inactive.
	//
	// - `unavailable`: The cluster is unavailable.
	//
	// - `deleting`: The cluster is being deleted.
	//
	// - `deleted`: The cluster has been deleted.
	//
	// - `delete_failed`: Cluster deletion failed.
	//
	// - `waiting`: Waiting for access.
	//
	// - `disconnected`: Disconnected.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Deprecated
	//
	// CIDR block for pod networks.
	//
	// example:
	//
	// 172.20.xx.xx/16
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// Tags associated with the cluster.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Last time the cluster was updated.
	//
	// example:
	//
	// 2025-04-10T13:28:09+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// VPC ID of the cluster. This parameter is required when creating a cluster.
	//
	// example:
	//
	// vpc-2zecuu62b9zw7a7qn****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// Deprecated
	//
	// vSwitch ID. This field is deprecated. Use vswitch_ids to query control plane vSwitches and node pool vswitch_ids to query data plane vSwitches.
	//
	// example:
	//
	// vsw-2zete8s4qocqg0mf6****,vsw-2zete8s4qocqg0mf6****
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// vSwitches for the cluster control plane.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Name of the RAM role assigned to ECS instances acting as worker nodes in the cluster.
	//
	// example:
	//
	// KubernetesWorkerRole-ec87d15b-edca-4302-933f-c8a16bf0****
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// Deprecated
	//
	// Zone ID within the region where the cluster is deployed.
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
	// Indicates whether smart managed mode is enabled.
	//
	// example:
	//
	// false
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
	// Indicates whether auto-renewal is enabled for nodes.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Auto-renewal duration for nodes.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Billing method for control plane nodes.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// Indicates whether Cloud Monitor is installed on nodes.
	//
	// example:
	//
	// true
	CloudMonitorFlags *bool `json:"cloud_monitor_flags,omitempty" xml:"cloud_monitor_flags,omitempty"`
	// CPU management policy for nodes.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Deployment set ID.
	//
	// example:
	//
	// ds-bp10b35imuam5amw****
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// Image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20240819
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Operating system image type.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// Metadata access configuration for ECS instances.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// Instance types for control plane nodes.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// Key pair name. Specify either this parameter or login_password.
	//
	// example:
	//
	// ack
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Port range for node services.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"node_port_range,omitempty" xml:"node_port_range,omitempty"`
	// Subscription duration for nodes.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Time unit for node subscription.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Runtime name.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Indicates whether Alibaba Cloud OS security hardening is enabled.
	//
	// example:
	//
	// true
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Number of control plane nodes.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Indicates whether security hardening for compliance is enabled.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Indicates whether burst performance is enabled for node system disks.
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// System disk category for nodes.
	//
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// Disk performance level for node system disks. Applies only to ESSD disks.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// Provisioned IOPS for node system disks.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// System disk size for nodes, in GB. Minimum value: 40.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// Automatic snapshot backup policy for node system disks.
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
	// Internal domain name configuration for the cluster, applicable to ACK managed clusters. The internal domain name allows node-side system components such as kubelet and kube-proxy to access the API server. If internal domain name access is disabled, these components access the API server through the CLB IP address.
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
	// VPCs where the internal domain name resolution takes effect. By default, this includes the VPC where the cluster resides.
	BindVpcs []*string `json:"bind_vpcs,omitempty" xml:"bind_vpcs,omitempty" type:"Repeated"`
	// Indicates whether internal domain name access is enabled.
	//
	// - true: Internal domain name access is enabled. Node-side components (kubelet, kube-proxy) access the API server through the internal domain name.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
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
	// Automatic cluster upgrade settings.
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
	// Frequency of automatic cluster upgrades. For more information, see [Upgrade frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// - patch: Latest patch version.
	//
	// - stable: Second latest minor version.
	//
	// - rapid: Latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Indicates whether automatic cluster upgrade is enabled.
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
	// Default audience for the OIDC token. Multiple values are separated by commas (,). These values appear as an array in the aud field of the OIDC token.
	//
	// example:
	//
	// https://kubernetes.default.svc,https://example.***.com
	Audience *string `json:"audience,omitempty" xml:"audience,omitempty"`
	// Indicates whether RRSA is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Issuer of the OIDC token. Multiple values are separated by commas (,). The first value appears in the iss field of the OIDC token and serves as the issuer URL for the OIDC identity provider.
	//
	// example:
	//
	// https://oidc-ack-***,https://kubernetes.default.svc
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// URL of the OIDC public key information.
	JwksUrl *string `json:"jwks_url,omitempty" xml:"jwks_url,omitempty"`
	// Maximum validity period configurable for the OIDC token.
	//
	// example:
	//
	// 12h
	MaxOidcTokenExpiration *string `json:"max_oidc_token_expiration,omitempty" xml:"max_oidc_token_expiration,omitempty"`
	// ARN of the OIDC identity provider.
	//
	// example:
	//
	// acs:ram::1138***:oidc-provider/ack-rrsa-***
	OidcArn *string `json:"oidc_arn,omitempty" xml:"oidc_arn,omitempty"`
	// Name of the OIDC identity provider.
	//
	// example:
	//
	// ack-rrsa-***
	OidcName *string `json:"oidc_name,omitempty" xml:"oidc_name,omitempty"`
	// URL of the OIDC configuration document.
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
