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
	// The list of cluster information.
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
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersV1ResponseBodyClusters struct {
	// The local domain name of the cluster.
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
	// The cluster specifications.
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The Pod network CIDR block for Flannel network configuration.
	//
	// example:
	//
	// 172.20.xx.xx/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2025-04-07T09:57:26+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The current version of the cluster.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Specifies whether deletion protection is enabled for the cluster to prevent accidental deletion through the console or API. Valid values:
	//
	// - true: Deletion protection is enabled. The cluster cannot be deleted through the console or API.
	//
	// - false: Deletion protection is disabled. The cluster can be deleted through the console or API.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// Deprecated
	//
	// The Docker version used by the cluster.
	//
	// example:
	//
	// 19.03.5
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// Deprecated
	//
	// The ID of the Server Load Balancer (SLB) instance for the Ingress in the cluster.
	//
	// Default instance type: guaranteed-performance instance (slb.s1.small).
	//
	// example:
	//
	// lb-2vcrbmlevo6kjpgch****
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// The initial version of the cluster. For the Kubernetes versions supported by ACK, refer to [Overview of Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.32.1-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The IP protocol stack of the cluster. Valid values:
	//
	// - ipv4: Creates a cluster that supports only the IPv4 protocol stack.
	//
	// - dual: Creates a cluster that supports both IPv4 and IPv6 protocol stacks.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// The maintenance window of the cluster. This feature takes effect only in ACK managed clusters and ACK Serverless clusters.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty" xml:"maintenance_window,omitempty"`
	// The access endpoints of the cluster API server, including the internal network access address and the public network access address.
	//
	// example:
	//
	// {\\"api_server_endpoint\\":\\"\\",\\"intranet_api_server_endpoint\\":\\"https://192.168.xx.xx:6443\\"}
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// The metadata of the cluster.
	//
	// example:
	//
	// {\\"Addons\\":[{\\"config\\":***}}
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// The cluster name. Fuzzy match is supported.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Deprecated
	//
	// The network mode of the cluster. Valid values:
	//
	// - classic: classic network
	//
	// - vpc: VPC
	//
	// - overlay: overlay network
	//
	// - calico: Calico network
	//
	// example:
	//
	// vpc
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// The version to which the cluster can be upgraded.
	//
	// example:
	//
	// 1.xx.x-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// The automated O&M policy of the cluster.
	OperationPolicy *DescribeClustersV1ResponseBodyClustersOperationPolicy `json:"operation_policy,omitempty" xml:"operation_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// The Private Zone configuration of the cluster. Valid values:
	//
	// - true: Private Zone is enabled.
	//
	// - false: Private Zone is disabled.
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
	// The kube-proxy mode. Valid values:
	//
	// - iptables: The mature and stable kube-proxy mode. Service discovery and load balancing for Kubernetes services are configured by using iptables rules. This mode provides average performance and is affected by the cluster scale. It is suitable for clusters with a small number of services.
	//
	// - ipvs: The high-performance kube-proxy mode. Service discovery and load balancing for Kubernetes services are configured by using the Linux IPVS module. This mode is suitable for clusters with a large number of services and scenarios that require high-performance load balancing.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// The ID of the region where the cluster is deployed.
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
	// The ID of the security group to which the cluster belongs.
	//
	// example:
	//
	// sg-2vcgwsrwgt5mp0yi****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The CIDR block of the service network.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.xx.xx/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The current number of nodes in the cluster, including the master node and worker nodes.
	//
	// example:
	//
	// 5
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The running status of the cluster. Valid values:
	//
	// - initial: The cluster is being created.
	//
	// - failed: The cluster failed to be created.
	//
	// - running: The cluster is running.
	//
	// - updating: The cluster is being updated.
	//
	// - upgrading: The cluster is being upgraded.
	//
	// - removing: Nodes are being removed.
	//
	// - draining: Nodes are being drained.
	//
	// - scaling: The cluster is being scaled.
	//
	// - inactive: The cluster is inactive.
	//
	// - unavailable: The cluster is unavailable.
	//
	// - deleting: The cluster is being deleted.
	//
	// - deleted: The cluster is deleted.
	//
	// - delete_failed: The cluster failed to be deleted.
	//
	// - waiting: The cluster is waiting for the accessed state.
	//
	// - disconnected: The cluster is disconnected.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use container_cidr to obtain the Pod network CIDR block.
	//
	// example:
	//
	// null
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// The resource tags of the cluster.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The time when the cluster was last updated.
	//
	// example:
	//
	// 2025-04-07T09:57:26+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The ID of the virtual private cloud (VPC) where the cluster is deployed.
	//
	// example:
	//
	// vpc-2vcg932hsxsxuqbgl****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch used by the cluster.
	//
	// example:
	//
	// vsw-2vc41xuumx5z2rdma****,vsw-2vc41xuumx5z2rdma****
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// The vSwitches of the cluster control plane.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The name of the worker RAM role, which authorizes ECS instances to serve as worker nodes in the cluster.
	//
	// example:
	//
	// KubernetesWorkerRole-ec87d15b-edca-4302-933f-c8a16bf0****
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// Deprecated
	//
	// The ID of the zone to which the cluster belongs.
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

type DescribeClustersV1ResponseBodyClustersOperationPolicy struct {
	// The auto-upgrade settings for the cluster.
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
	if s.ClusterAutoUpgrade != nil {
		if err := s.ClusterAutoUpgrade.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersV1ResponseBodyClustersOperationPolicyClusterAutoUpgrade struct {
	// The auto-upgrade channel for the cluster. For more information, refer to [Upgrade channels](https://help.aliyun.com/document_detail/2712866.html). Valid values:
	//
	// - patch: The latest patch version.
	//
	// - stable: The second latest minor version.
	//
	// - rapid: The latest minor version.
	//
	// example:
	//
	// patch
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Specifies whether auto-upgrade is enabled for the cluster.
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
	// The total number of results.
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
