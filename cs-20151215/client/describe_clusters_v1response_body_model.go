// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersV1ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeClustersV1ResponseBodyClusters) *DescribeClustersV1ResponseBody
	GetClusters() []*DescribeClustersV1ResponseBodyClusters
	SetPageInfo(v *DescribeClustersV1ResponseBodyPageInfo) *DescribeClustersV1ResponseBody
	GetPageInfo() *DescribeClustersV1ResponseBodyPageInfo
}

type DescribeClustersV1ResponseBody struct {
	// The queried cluster details.
	Clusters []*DescribeClustersV1ResponseBodyClusters `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeClustersV1ResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeClustersV1ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBody) GetClusters() []*DescribeClustersV1ResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClustersV1ResponseBody) GetPageInfo() *DescribeClustersV1ResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeClustersV1ResponseBody) SetClusters(v []*DescribeClustersV1ResponseBodyClusters) *DescribeClustersV1ResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersV1ResponseBody) SetPageInfo(v *DescribeClustersV1ResponseBodyPageInfo) *DescribeClustersV1ResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeClustersV1ResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClustersV1ResponseBodyClusters struct {
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
	// c3fb96524f9274b4495df0f12a6b5****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The specification of the cluster.
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The pod CIDR block and the configuration of the Flannel network plug-in.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The time at which the instance is created.
	//
	// example:
	//
	// 2020-08-20T10:51:29+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.16.9-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Specifies whether to enable cluster deletion protection. If you enable this option, the cluster cannot be deleted in the console or by calling API operations. Valid values:
	//
	// 	- `true`: enables deletion protection for the cluster. This way, the cluster cannot be deleted in the ACK console or by calling API operations.
	//
	// 	- `false`: disables deletion protection for the cluster. This way, the cluster can be deleted in the ACK console or by calling API operations.
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
	// The ID of the Server Load Balancer (SLB) instance that is used by the Ingresses of the cluster.
	//
	// The default SLB specification is slb.s1.small, which belongs to the high-performance instance type.
	//
	// example:
	//
	// lb-2vcrbmlevo6kjpgch****
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// The version of the cluster. For more information about the Kubernetes versions supported by ACK, see [Release notes for Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.16.9-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The IP stack of the cluster. Valid values:
	//
	// 	- ipv4: creates a cluster that supports only the IPv4 protocol stack.
	//
	// 	- dual: creates a cluster that supports IPv4/IPv6 dual-stack.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// The maintenance window of the cluster. This feature is available only for ACK managed clusters and ACK Serverless clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The address of the cluster API server. It includes an internal endpoint and a public endpoint.
	//
	// example:
	//
	// {\\"api_server_endpoint\\":\\"\\",\\"intranet_api_server_endpoint\\":\\"https://192.168.0.251:6443\\"}
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// The metadata of the cluster.
	//
	// example:
	//
	// {\\"Addons\\":[{\\"config\\":***}
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Deprecated
	//
	// The network mode of the cluster. Valid values:
	//
	// 	- `classic`: classic network.
	//
	// 	- `vpc`: virtual private cloud (VPC).
	//
	// 	- `overlay`: overlay network.
	//
	// 	- `calico`: network powered by Calico.
	//
	// example:
	//
	// vpc
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// The Kubernetes version to which the cluster can be updated.
	//
	// example:
	//
	// 1.18.8-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// The automatic O\\&M policy of the cluster.
	OperationPolicy *DescribeClustersV1ResponseBodyClustersOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// Indicates whether Alibaba Cloud DNS PrivateZone is enabled. Valid values:
	//
	// 	- `true`: Alibaba Cloud DNS PrivateZone is enabled.
	//
	// 	- `false`: Alibaba Cloud DNS PrivateZone is disabled.
	//
	// example:
	//
	// false
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// The subtype of the cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy mode.
	//
	// 	- `iptables`: a mature and stable mode that uses iptables rules to conduct service discovery and load balancing. The performance of this mode is limited by the size of the cluster. This mode is suitable for clusters that run a small number of Services.
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
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The ID of the security group of the cluster.
	//
	// example:
	//
	// sg-2vcgwsrwgt5mp0yi****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The Service CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The number of nodes in the cluster, including control planes and worker nodes.
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
	// 	- `upgrading`: The cluster is undergoing an upgrade.
	//
	// 	- `updating`: Cluster specification changes are being applied.
	//
	// 	- `removing`: Nodes are being removed from the node pool.
	//
	// 	- `draining`: Node draining is in progress.
	//
	// 	- `scaling`: Auto-scaling operation is in progress for the cluster.
	//
	// 	- `stopped`: The cluster has stopped running.
	//
	// 	- `deleting`: The cluster is being deleted.
	//
	// 	- `deleted`: The cluster has been deleted.
	//
	// 	- `delete_failed`: The cluster failed to be deleted.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the container_cidr parameter to obtain the pod CIDR block.
	//
	// example:
	//
	// 172.21.0.0/16
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// The label of the cluster.
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
	// 2020-09-16T11:09:55+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The ID of the virtual private cloud (VPC) that is used by the cluster.
	//
	// example:
	//
	// vpc-2vcg932hsxsxuqbgl****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch in the cluster.
	//
	// example:
	//
	// vsw-2vc41xuumx5z2rdma****,vsw-2vc41xuumx5z2rdma****
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// The vSwitches of the control planes.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The name of the worker Resource Access Management (RAM) role. The RAM role is assigned to the worker nodes of the cluster to allow the worker nodes to manage ECS instances.
	//
	// example:
	//
	// KubernetesWorkerRole-ec87d15b-edca-4302-933f-c8a16bf0****
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// Deprecated
	//
	// The ID of the zone where the cluster is deployed.
	//
	// example:
	//
	// cn-beijing-b
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClustersV1ResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClusters) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeClustersV1ResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeClustersV1ResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersV1ResponseBodyClusters) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *DescribeClustersV1ResponseBodyClusters) GetCreated() *string {
	return s.Created
}

func (s *DescribeClustersV1ResponseBodyClusters) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClustersV1ResponseBodyClusters) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeClustersV1ResponseBodyClusters) GetDockerVersion() *string {
	return s.DockerVersion
}

func (s *DescribeClustersV1ResponseBodyClusters) GetExternalLoadbalancerId() *string {
	return s.ExternalLoadbalancerId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetInitVersion() *string {
	return s.InitVersion
}

func (s *DescribeClustersV1ResponseBodyClusters) GetIpStack() *string {
	return s.IpStack
}

func (s *DescribeClustersV1ResponseBodyClusters) GetMaintenanceWindow() *MaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *DescribeClustersV1ResponseBodyClusters) GetMasterUrl() *string {
	return s.MasterUrl
}

func (s *DescribeClustersV1ResponseBodyClusters) GetMetaData() *string {
	return s.MetaData
}

func (s *DescribeClustersV1ResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeClustersV1ResponseBodyClusters) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *DescribeClustersV1ResponseBodyClusters) GetNextVersion() *string {
	return s.NextVersion
}

func (s *DescribeClustersV1ResponseBodyClusters) GetOperationPolicy() *DescribeClustersV1ResponseBodyClustersOperationPolicy {
	return s.OperationPolicy
}

func (s *DescribeClustersV1ResponseBodyClusters) GetPrivateZone() *bool {
	return s.PrivateZone
}

func (s *DescribeClustersV1ResponseBodyClusters) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersV1ResponseBodyClusters) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *DescribeClustersV1ResponseBodyClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *DescribeClustersV1ResponseBodyClusters) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClustersV1ResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeClustersV1ResponseBodyClusters) GetSubnetCidr() *string {
	return s.SubnetCidr
}

func (s *DescribeClustersV1ResponseBodyClusters) GetTags() []*Tag {
	return s.Tags
}

func (s *DescribeClustersV1ResponseBodyClusters) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeClustersV1ResponseBodyClusters) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClustersV1ResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClustersV1ResponseBodyClusters) GetWorkerRamRoleName() *string {
	return s.WorkerRamRoleName
}

func (s *DescribeClustersV1ResponseBodyClusters) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterDomain(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterSpec(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterType(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetContainerCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ContainerCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetCreated(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Created = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetCurrentVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDeletionProtection(v bool) *DescribeClustersV1ResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetDockerVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetExternalLoadbalancerId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetInitVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetIpStack(v string) *DescribeClustersV1ResponseBodyClusters {
	s.IpStack = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetMaintenanceWindow(v *MaintenanceWindow) *DescribeClustersV1ResponseBodyClusters {
	s.MaintenanceWindow = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetMasterUrl(v string) *DescribeClustersV1ResponseBodyClusters {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetMetaData(v string) *DescribeClustersV1ResponseBodyClusters {
	s.MetaData = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetName(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetNetworkMode(v string) *DescribeClustersV1ResponseBodyClusters {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetNextVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.NextVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetOperationPolicy(v *DescribeClustersV1ResponseBodyClustersOperationPolicy) *DescribeClustersV1ResponseBodyClusters {
	s.OperationPolicy = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetPrivateZone(v bool) *DescribeClustersV1ResponseBodyClusters {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetProfile(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetProxyMode(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ProxyMode = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetRegionId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetResourceGroupId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSecurityGroupId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetServiceCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ServiceCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSize(v int64) *DescribeClustersV1ResponseBodyClusters {
	s.Size = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetState(v string) *DescribeClustersV1ResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetSubnetCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetTags(v []*Tag) *DescribeClustersV1ResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetTimezone(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Timezone = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetUpdated(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Updated = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVpcId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVswitchId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VswitchId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVswitchIds(v []*string) *DescribeClustersV1ResponseBodyClusters {
	s.VswitchIds = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetWorkerRamRoleName(v string) *DescribeClustersV1ResponseBodyClusters {
	s.WorkerRamRoleName = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetZoneId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ZoneId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeClustersV1ResponseBodyClustersOperationPolicy struct {
	// The configurations of auto cluster update.
	ClusterAutoUpgrade *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade `json:"cluster_auto_upgrade,omitempty" xml:"cluster_auto_upgrade,omitempty" type:"Struct"`
}

func (s DescribeClustersV1ResponseBodyClustersOperationPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClustersOperationPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicy) GetClusterAutoUpgrade() *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade {
	return s.ClusterAutoUpgrade
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicy) SetClusterAutoUpgrade(v *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) *DescribeClustersV1ResponseBodyClustersOperationPolicy {
	s.ClusterAutoUpgrade = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade struct {
	// The frequency of auto cluster updates. For more information, see [Update frequency](https://help.aliyun.com/document_detail/2712866.html).
	//
	// Valid values:
	//
	// 	- patch: the latest patch version.
	//
	// 	- stables: the second-latest minor version.
	//
	// 	- rapid: the latest minor version.
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

func (s DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) GetChannel() *string {
	return s.Channel
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) SetChannel(v string) *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade {
	s.Channel = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) SetEnabled(v bool) *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade {
	s.Enabled = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade) Validate() error {
	return dara.Validate(s)
}

type DescribeClustersV1ResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClustersV1ResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClustersV1ResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClustersV1ResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetPageNumber(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetPageSize(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyPageInfo) SetTotalCount(v int32) *DescribeClustersV1ResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
