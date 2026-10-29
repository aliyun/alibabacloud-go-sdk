// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoMode(v *CreateClusterNodePoolRequestAutoMode) *CreateClusterNodePoolRequest
	GetAutoMode() *CreateClusterNodePoolRequestAutoMode
	SetAutoScaling(v *CreateClusterNodePoolRequestAutoScaling) *CreateClusterNodePoolRequest
	GetAutoScaling() *CreateClusterNodePoolRequestAutoScaling
	SetCount(v int64) *CreateClusterNodePoolRequest
	GetCount() *int64
	SetEfloNodeGroup(v *CreateClusterNodePoolRequestEfloNodeGroup) *CreateClusterNodePoolRequest
	GetEfloNodeGroup() *CreateClusterNodePoolRequestEfloNodeGroup
	SetHostNetwork(v bool) *CreateClusterNodePoolRequest
	GetHostNetwork() *bool
	SetInterconnectConfig(v *CreateClusterNodePoolRequestInterconnectConfig) *CreateClusterNodePoolRequest
	GetInterconnectConfig() *CreateClusterNodePoolRequestInterconnectConfig
	SetInterconnectMode(v string) *CreateClusterNodePoolRequest
	GetInterconnectMode() *string
	SetIntranet(v bool) *CreateClusterNodePoolRequest
	GetIntranet() *bool
	SetKubernetesConfig(v *CreateClusterNodePoolRequestKubernetesConfig) *CreateClusterNodePoolRequest
	GetKubernetesConfig() *CreateClusterNodePoolRequestKubernetesConfig
	SetManagement(v *CreateClusterNodePoolRequestManagement) *CreateClusterNodePoolRequest
	GetManagement() *CreateClusterNodePoolRequestManagement
	SetMaxNodes(v int64) *CreateClusterNodePoolRequest
	GetMaxNodes() *int64
	SetNodeComponents(v []*CreateClusterNodePoolRequestNodeComponents) *CreateClusterNodePoolRequest
	GetNodeComponents() []*CreateClusterNodePoolRequestNodeComponents
	SetNodeConfig(v *CreateClusterNodePoolRequestNodeConfig) *CreateClusterNodePoolRequest
	GetNodeConfig() *CreateClusterNodePoolRequestNodeConfig
	SetNodepoolInfo(v *CreateClusterNodePoolRequestNodepoolInfo) *CreateClusterNodePoolRequest
	GetNodepoolInfo() *CreateClusterNodePoolRequestNodepoolInfo
	SetScalingGroup(v *CreateClusterNodePoolRequestScalingGroup) *CreateClusterNodePoolRequest
	GetScalingGroup() *CreateClusterNodePoolRequestScalingGroup
	SetTeeConfig(v *CreateClusterNodePoolRequestTeeConfig) *CreateClusterNodePoolRequest
	GetTeeConfig() *CreateClusterNodePoolRequestTeeConfig
}

type CreateClusterNodePoolRequest struct {
	// The auto mode configuration of the node pool.
	AutoMode *CreateClusterNodePoolRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The elastic scaling configuration.
	AutoScaling *CreateClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- Use desired_size instead.
	//
	// The number of nodes in the node pool.
	//
	// example:
	//
	// null
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The Lingjun node pool configuration.
	EfloNodeGroup *CreateClusterNodePoolRequestEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// Specifies whether the pod network mode uses the host network mode. Valid values:
	//
	// - `true`: host network. Pods directly use the network stack of the host and share the IP address and ports with the host.
	//
	// - `false`: container network. Pods have independent network stacks and do not occupy host network ports.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// The configuration of the edge node pool.
	InterconnectConfig *CreateClusterNodePoolRequestInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter takes effect only when `type` is set to `edge`. Valid values:
	//
	// - `basic`: Internet. Nodes in cloud node pool interact with cloud nodes over the Internet. Applications in cloud node pool cannot directly access the internal network of the cloud VPC.
	//
	// - `private`: private network. Nodes in cloud node pool connect to the cloud network through Express Connect circuits, VPNs, or CEN, which provides higher cloud-edge communication quality and more effective security protection.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Specifies whether Layer 3 network connectivity is enabled between nodes in the edge node pool. Valid values:
	//
	// - `true`: enabled. Layer 3 network connectivity is enabled between all nodes in the node pool.
	//
	// - `false`: disabled. Layer 3 network connectivity is disabled between all hosts in the node pool.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The cluster configurations.
	KubernetesConfig *CreateClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations for the managed node pool feature.
	Management *CreateClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- The maximum number of nodes that the edge node pool can accommodate.
	//
	// example:
	//
	// null
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The list of node components.
	NodeComponents []*CreateClusterNodePoolRequestNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// The node configuration.
	NodeConfig *CreateClusterNodePoolRequestNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The node pool configuration.
	NodepoolInfo *CreateClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The scaling group configuration of the node pool.
	ScalingGroup *CreateClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The confidential computing cluster configuration.
	TeeConfig *CreateClusterNodePoolRequestTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequest) GetAutoMode() *CreateClusterNodePoolRequestAutoMode {
	return s.AutoMode
}

func (s *CreateClusterNodePoolRequest) GetAutoScaling() *CreateClusterNodePoolRequestAutoScaling {
	return s.AutoScaling
}

func (s *CreateClusterNodePoolRequest) GetCount() *int64 {
	return s.Count
}

func (s *CreateClusterNodePoolRequest) GetEfloNodeGroup() *CreateClusterNodePoolRequestEfloNodeGroup {
	return s.EfloNodeGroup
}

func (s *CreateClusterNodePoolRequest) GetHostNetwork() *bool {
	return s.HostNetwork
}

func (s *CreateClusterNodePoolRequest) GetInterconnectConfig() *CreateClusterNodePoolRequestInterconnectConfig {
	return s.InterconnectConfig
}

func (s *CreateClusterNodePoolRequest) GetInterconnectMode() *string {
	return s.InterconnectMode
}

func (s *CreateClusterNodePoolRequest) GetIntranet() *bool {
	return s.Intranet
}

func (s *CreateClusterNodePoolRequest) GetKubernetesConfig() *CreateClusterNodePoolRequestKubernetesConfig {
	return s.KubernetesConfig
}

func (s *CreateClusterNodePoolRequest) GetManagement() *CreateClusterNodePoolRequestManagement {
	return s.Management
}

func (s *CreateClusterNodePoolRequest) GetMaxNodes() *int64 {
	return s.MaxNodes
}

func (s *CreateClusterNodePoolRequest) GetNodeComponents() []*CreateClusterNodePoolRequestNodeComponents {
	return s.NodeComponents
}

func (s *CreateClusterNodePoolRequest) GetNodeConfig() *CreateClusterNodePoolRequestNodeConfig {
	return s.NodeConfig
}

func (s *CreateClusterNodePoolRequest) GetNodepoolInfo() *CreateClusterNodePoolRequestNodepoolInfo {
	return s.NodepoolInfo
}

func (s *CreateClusterNodePoolRequest) GetScalingGroup() *CreateClusterNodePoolRequestScalingGroup {
	return s.ScalingGroup
}

func (s *CreateClusterNodePoolRequest) GetTeeConfig() *CreateClusterNodePoolRequestTeeConfig {
	return s.TeeConfig
}

func (s *CreateClusterNodePoolRequest) SetAutoMode(v *CreateClusterNodePoolRequestAutoMode) *CreateClusterNodePoolRequest {
	s.AutoMode = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetAutoScaling(v *CreateClusterNodePoolRequestAutoScaling) *CreateClusterNodePoolRequest {
	s.AutoScaling = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetCount(v int64) *CreateClusterNodePoolRequest {
	s.Count = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetEfloNodeGroup(v *CreateClusterNodePoolRequestEfloNodeGroup) *CreateClusterNodePoolRequest {
	s.EfloNodeGroup = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetHostNetwork(v bool) *CreateClusterNodePoolRequest {
	s.HostNetwork = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetInterconnectConfig(v *CreateClusterNodePoolRequestInterconnectConfig) *CreateClusterNodePoolRequest {
	s.InterconnectConfig = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetInterconnectMode(v string) *CreateClusterNodePoolRequest {
	s.InterconnectMode = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetIntranet(v bool) *CreateClusterNodePoolRequest {
	s.Intranet = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetKubernetesConfig(v *CreateClusterNodePoolRequestKubernetesConfig) *CreateClusterNodePoolRequest {
	s.KubernetesConfig = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetManagement(v *CreateClusterNodePoolRequestManagement) *CreateClusterNodePoolRequest {
	s.Management = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetMaxNodes(v int64) *CreateClusterNodePoolRequest {
	s.MaxNodes = &v
	return s
}

func (s *CreateClusterNodePoolRequest) SetNodeComponents(v []*CreateClusterNodePoolRequestNodeComponents) *CreateClusterNodePoolRequest {
	s.NodeComponents = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetNodeConfig(v *CreateClusterNodePoolRequestNodeConfig) *CreateClusterNodePoolRequest {
	s.NodeConfig = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetNodepoolInfo(v *CreateClusterNodePoolRequestNodepoolInfo) *CreateClusterNodePoolRequest {
	s.NodepoolInfo = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetScalingGroup(v *CreateClusterNodePoolRequestScalingGroup) *CreateClusterNodePoolRequest {
	s.ScalingGroup = v
	return s
}

func (s *CreateClusterNodePoolRequest) SetTeeConfig(v *CreateClusterNodePoolRequestTeeConfig) *CreateClusterNodePoolRequest {
	s.TeeConfig = v
	return s
}

func (s *CreateClusterNodePoolRequest) Validate() error {
	if s.AutoMode != nil {
		if err := s.AutoMode.Validate(); err != nil {
			return err
		}
	}
	if s.AutoScaling != nil {
		if err := s.AutoScaling.Validate(); err != nil {
			return err
		}
	}
	if s.EfloNodeGroup != nil {
		if err := s.EfloNodeGroup.Validate(); err != nil {
			return err
		}
	}
	if s.InterconnectConfig != nil {
		if err := s.InterconnectConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KubernetesConfig != nil {
		if err := s.KubernetesConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Management != nil {
		if err := s.Management.Validate(); err != nil {
			return err
		}
	}
	if s.NodeComponents != nil {
		for _, item := range s.NodeComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeConfig != nil {
		if err := s.NodeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodepoolInfo != nil {
		if err := s.NodepoolInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingGroup != nil {
		if err := s.ScalingGroup.Validate(); err != nil {
			return err
		}
	}
	if s.TeeConfig != nil {
		if err := s.TeeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolRequestAutoMode struct {
	// Specifies whether to enable auto mode. Valid values:
	//
	// - true: enables auto mode. This value is valid only when auto mode is enabled for the cluster.
	//
	// - false: disables auto mode.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateClusterNodePoolRequestAutoMode) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestAutoMode) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestAutoMode) GetEnable() *bool {
	return s.Enable
}

func (s *CreateClusterNodePoolRequestAutoMode) SetEnable(v bool) *CreateClusterNodePoolRequestAutoMode {
	s.Enable = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoMode) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestAutoScaling struct {
	// Deprecated
	//
	// **[Deprecated]*	- Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The peak bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// null
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The billing method of the EIP. Valid values:
	//
	// - `PayByBandwidth`: pay-by-bandwidth.
	//
	// - `PayByTraffic`: pay-by-data-transfer.
	//
	// Default value: `PayByBandwidth`.
	//
	// example:
	//
	// null
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable auto scaling. Valid values:
	//
	// - `true`: enables auto scaling for the node pool. When the capacity planning of the cluster cannot meet the scheduling requirements of application pods, ACK automatically scales node resources based on the configured minimum and maximum numbers of instances. For clusters of version 1.24 and later, node instant elasticity is enabled by default. For clusters earlier than version 1.24, node auto scaling is enabled by default. For more information, see [Node auto scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: disables auto scaling. ACK adjusts the number of nodes in the node pool based on the configured expected number of nodes to keep the number of nodes at the expected value.
	//
	// If this parameter is set to false, other configuration parameters in `auto_scaling` do not take effect.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// Specifies whether to associate an EIP. Valid values:
	//
	// - `true`: associates an EIP.
	//
	// - `false`: does not associate an EIP.
	//
	// Default value: `false`.
	//
	// example:
	//
	// null
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of scalable instances in the node pool, excluding existing instances. This parameter takes effect only when `enable` is set to true.
	//
	// Valid values: [min_instances, 2000]. Default value: 0.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of scalable instances in the node pool, excluding existing instances. This parameter takes effect only when `enable` is set to true.
	//
	// Valid values: [0, max_instances]. Default value: 0.
	//
	// > - If the minimum number of instances is not 0, the automatic creation of the corresponding number of ECS instances occurs after the scaling group takes effect.
	//
	// > - The recommended Settings for the maximum number of instances should not be less than the current number of nodes in the node pool. Otherwise, the elastic scaling feature scales in the nodes in the node pool immediately after it takes effect.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The instance type for elastic scaling. This parameter takes effect only when `enable` is set to true. Valid values:
	//
	// - `cpu`: regular instance.
	//
	// - `gpu`: GPU-accelerated instance.
	//
	// - `gpushare`: GPU-shared instance.
	//
	// - `spot`: spot instance.
	//
	// Default value: `cpu`.
	//
	// 	Notice: This parameter cannot be modified after the node pool is created.
	//
	// example:
	//
	// cpu
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateClusterNodePoolRequestAutoScaling) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestAutoScaling) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetEipBandwidth() *int64 {
	return s.EipBandwidth
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetEipInternetChargeType() *string {
	return s.EipInternetChargeType
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetEnable() *bool {
	return s.Enable
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetIsBondEip() *bool {
	return s.IsBondEip
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *CreateClusterNodePoolRequestAutoScaling) GetType() *string {
	return s.Type
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetEipBandwidth(v int64) *CreateClusterNodePoolRequestAutoScaling {
	s.EipBandwidth = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetEipInternetChargeType(v string) *CreateClusterNodePoolRequestAutoScaling {
	s.EipInternetChargeType = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetEnable(v bool) *CreateClusterNodePoolRequestAutoScaling {
	s.Enable = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetIsBondEip(v bool) *CreateClusterNodePoolRequestAutoScaling {
	s.IsBondEip = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetMaxInstances(v int64) *CreateClusterNodePoolRequestAutoScaling {
	s.MaxInstances = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetMinInstances(v int64) *CreateClusterNodePoolRequestAutoScaling {
	s.MinInstances = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) SetType(v string) *CreateClusterNodePoolRequestAutoScaling {
	s.Type = &v
	return s
}

func (s *CreateClusterNodePoolRequestAutoScaling) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestEfloNodeGroup struct {
	// Specifies whether to enable auto-attach for the Lingjun node pool.
	AutoAttachEnabled *bool `json:"auto_attach_enabled,omitempty" xml:"auto_attach_enabled,omitempty"`
	// The ID of the Lingjun cluster to associate when creating the Lingjun node pool.
	//
	// example:
	//
	// i1169130516633730****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The ID of the Lingjun group in the Lingjun cluster to associate when creating the Lingjun node pool.
	//
	// example:
	//
	// ng-ec3c96ff0aa****
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The worker RAM role used by the Lingjun node pool.
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
}

func (s CreateClusterNodePoolRequestEfloNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestEfloNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetAutoAttachEnabled() *bool {
	return s.AutoAttachEnabled
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetWorkerRamRoleName() *string {
	return s.WorkerRamRoleName
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetAutoAttachEnabled(v bool) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.AutoAttachEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetClusterId(v string) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetGroupId(v string) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.GroupId = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetWorkerRamRoleName(v string) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.WorkerRamRoleName = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestInterconnectConfig struct {
	// **[Deprecated]**
	//
	// The network bandwidth of the edge-enhanced node pool. Unit: Mbit/s.
	//
	// example:
	//
	// null
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// **[Deprecated]**
	//
	// The instance ID of the Cloud Connect Network (CCN) attached to the edge-enhanced node pool.
	//
	// example:
	//
	// null
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// **[Deprecated]**
	//
	// The region of the CCN instance associated with the edge-enhanced node pool.
	//
	// example:
	//
	// null
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// **[Deprecated]**
	//
	// The instance ID of the Cloud Enterprise Network (CEN) attached to the edge-enhanced node pool.
	//
	// example:
	//
	// null
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// **[Deprecated]**
	//
	// The subscription duration of the edge-enhanced node pool. Unit: months.
	//
	// example:
	//
	// null
	ImprovedPeriod *string `json:"improved_period,omitempty" xml:"improved_period,omitempty"`
}

func (s CreateClusterNodePoolRequestInterconnectConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestInterconnectConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) GetCcnId() *string {
	return s.CcnId
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) GetCcnRegionId() *string {
	return s.CcnRegionId
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) GetCenId() *string {
	return s.CenId
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) GetImprovedPeriod() *string {
	return s.ImprovedPeriod
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) SetBandwidth(v int64) *CreateClusterNodePoolRequestInterconnectConfig {
	s.Bandwidth = &v
	return s
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) SetCcnId(v string) *CreateClusterNodePoolRequestInterconnectConfig {
	s.CcnId = &v
	return s
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) SetCcnRegionId(v string) *CreateClusterNodePoolRequestInterconnectConfig {
	s.CcnRegionId = &v
	return s
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) SetCenId(v string) *CreateClusterNodePoolRequestInterconnectConfig {
	s.CenId = &v
	return s
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) SetImprovedPeriod(v string) *CreateClusterNodePoolRequestInterconnectConfig {
	s.ImprovedPeriod = &v
	return s
}

func (s *CreateClusterNodePoolRequestInterconnectConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestKubernetesConfig struct {
	// Specifies whether to install the CloudMonitor agent on ECS nodes. After the agent is installed, you can view the monitoring information of the created ECS instances in the CloudMonitor console. Valid values:
	//
	// - `true`: The CloudMonitor agent is installed on ECS nodes.
	//
	// - `false`: The CloudMonitor agent is not installed on ECS nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy for nodes. The following two policies are supported when the cluster version is 1.12.6 or later:
	//
	// - `static`: Allows you to enhance CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - `none`: Uses the default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels to be added to the nodes in the Kubernetes cluster.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name. After you specify a custom node name, the node name, ECS instance name, and ECS instance hostname are changed accordingly.
	//
	// > For Windows instances with custom node names enabled, the hostname is fixed to the IP address, with hyphens (-) replacing periods (.) in the IP address, and without prefixes or suffixes.
	//
	// The node name consists of a prefix, the node IP address, and a suffix:
	//
	// - The total length is 2 to 64 characters. The node name must start and end with a lowercase letter or a digit.
	//
	// - The prefix and suffix can contain uppercase and lowercase letters, digits, hyphens (-), and periods (.). They must start with a letter and cannot start or end with a hyphen (-) or a period (.). Consecutive hyphens (-) or periods (.) are not allowed.
	//
	// - The prefix is required due to ECS limits, and the suffix is optional.
	//
	// - The node IP address is the complete private IP address of the node.
	//
	// For example, if the node IP address is 192.XX.YY.55, the prefix is aliyun.com, and the suffix is test:
	//
	// - If the node is a Linux node, the node name, ECS instance name, and ECS instance hostname are all aliyun.com192.XX.YY.55test.
	//
	// - If the node is a Windows node, the ECS instance hostname is 192-XX-YY-55, and both the node name and ECS instance name are aliyun.com192.XX.YY.55test.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-custom data of the instance. Before the node is added to the cluster, the specified pre-custom data script is run. For more information, see [User data scripts](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. ACK supports the following three container runtimes:
	//
	// - `containerd`: Recommended. Supported by all cluster versions.
	//
	// - `Sandboxed-Container.runv`: Secure sandboxed containers that provide higher isolation. Supported by clusters of version 1.31 and earlier.
	//
	// - `docker`: Deprecated. Supported by clusters of version 1.22 and earlier.
	//
	// Default value: `containerd`.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The container runtime version.
	//
	// example:
	//
	// 1.6.38
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// The taint configuration.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Specifies whether the scaled-out nodes are unschedulable. Valid values:
	//
	// - `true`: The nodes are unschedulable.
	//
	// - `false`: The nodes are schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After the node is added to the cluster, the specified instance user data script is run. For more information, see [User data scripts](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU=
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s CreateClusterNodePoolRequestKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestKubernetesConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetCmsEnabled() *bool {
	return s.CmsEnabled
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetLabels() []*Tag {
	return s.Labels
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetTaints() []*Taint {
	return s.Taints
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetUnschedulable() *bool {
	return s.Unschedulable
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetCmsEnabled(v bool) *CreateClusterNodePoolRequestKubernetesConfig {
	s.CmsEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetCpuPolicy(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.CpuPolicy = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetLabels(v []*Tag) *CreateClusterNodePoolRequestKubernetesConfig {
	s.Labels = v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetNodeNameMode(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.NodeNameMode = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetPreUserData(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetRuntime(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.Runtime = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetRuntimeVersion(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetTaints(v []*Taint) *CreateClusterNodePoolRequestKubernetesConfig {
	s.Taints = v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetUnschedulable(v bool) *CreateClusterNodePoolRequestKubernetesConfig {
	s.Unschedulable = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) SetUserData(v string) *CreateClusterNodePoolRequestKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *CreateClusterNodePoolRequestKubernetesConfig) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Taints != nil {
		for _, item := range s.Taints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterNodePoolRequestManagement struct {
	// Specifies whether to enable ECS fault detection for node auto-healing.
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to automatically repair nodes. This parameter takes effect only when enable is set to true. Valid values:
	//
	// - `true`: Automatically repairs nodes.
	//
	// - `false`: Does not automatically repair nodes.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto-repair policy for nodes.
	AutoRepairPolicy *CreateClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically upgrade nodes. This parameter takes effect only when enable is set to true. Valid values:
	//
	// - `true`: Enables auto-upgrade.
	//
	// - `false`: Disables auto-upgrade.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto-upgrade policy for nodes.
	AutoUpgradePolicy *CreateClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically repair CVE vulnerabilities. This parameter takes effect only when enable is set to true. Valid values:
	//
	// - `true`: Allows automatic CVE repair.
	//
	// - `false`: Does not allow automatic CVE repair.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The auto-repair policy for CVE vulnerabilities.
	AutoVulFixPolicy *CreateClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable the node rotation feature. This feature is supported only by intelligent managed node pools and is enabled by default. It is not supported by regular node pools.
	DriftEnabled *bool `json:"drift_enabled,omitempty" xml:"drift_enabled,omitempty"`
	// Specifies whether to enable the managed node pool feature. Valid values:
	//
	// - `true`: Enables the managed node pool feature.
	//
	// - `false`: Disables the managed node pool feature. Other related configurations take effect only when this parameter is set to true.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the auto_upgrade parameter at the upper level instead. The auto-upgrade configuration. This parameter takes effect only when enable is set to true.
	UpgradeConfig *CreateClusterNodePoolRequestManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolRequestManagement) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagement) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoFaultDiagnosis() *bool {
	return s.AutoFaultDiagnosis
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoRepair() *bool {
	return s.AutoRepair
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoRepairPolicy() *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	return s.AutoRepairPolicy
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoUpgradePolicy() *CreateClusterNodePoolRequestManagementAutoUpgradePolicy {
	return s.AutoUpgradePolicy
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoVulFix() *bool {
	return s.AutoVulFix
}

func (s *CreateClusterNodePoolRequestManagement) GetAutoVulFixPolicy() *CreateClusterNodePoolRequestManagementAutoVulFixPolicy {
	return s.AutoVulFixPolicy
}

func (s *CreateClusterNodePoolRequestManagement) GetDriftEnabled() *bool {
	return s.DriftEnabled
}

func (s *CreateClusterNodePoolRequestManagement) GetEnable() *bool {
	return s.Enable
}

func (s *CreateClusterNodePoolRequestManagement) GetUpgradeConfig() *CreateClusterNodePoolRequestManagementUpgradeConfig {
	return s.UpgradeConfig
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoFaultDiagnosis(v bool) *CreateClusterNodePoolRequestManagement {
	s.AutoFaultDiagnosis = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoRepair(v bool) *CreateClusterNodePoolRequestManagement {
	s.AutoRepair = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoRepairPolicy(v *CreateClusterNodePoolRequestManagementAutoRepairPolicy) *CreateClusterNodePoolRequestManagement {
	s.AutoRepairPolicy = v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoUpgrade(v bool) *CreateClusterNodePoolRequestManagement {
	s.AutoUpgrade = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoUpgradePolicy(v *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) *CreateClusterNodePoolRequestManagement {
	s.AutoUpgradePolicy = v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoVulFix(v bool) *CreateClusterNodePoolRequestManagement {
	s.AutoVulFix = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetAutoVulFixPolicy(v *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) *CreateClusterNodePoolRequestManagement {
	s.AutoVulFixPolicy = v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetDriftEnabled(v bool) *CreateClusterNodePoolRequestManagement {
	s.DriftEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetEnable(v bool) *CreateClusterNodePoolRequestManagement {
	s.Enable = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) SetUpgradeConfig(v *CreateClusterNodePoolRequestManagementUpgradeConfig) *CreateClusterNodePoolRequestManagement {
	s.UpgradeConfig = v
	return s
}

func (s *CreateClusterNodePoolRequestManagement) Validate() error {
	if s.AutoRepairPolicy != nil {
		if err := s.AutoRepairPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.AutoUpgradePolicy != nil {
		if err := s.AutoUpgradePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.AutoVulFixPolicy != nil {
		if err := s.AutoVulFixPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.UpgradeConfig != nil {
		if err := s.UpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolRequestManagementAutoRepairPolicy struct {
	// Specifies whether manual approval is required for node repair.
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// The ID of the auto-repair policy. If you do not specify this parameter or specify an empty string, no auto-healing rule is bound.
	//
	// example:
	//
	// r-xxxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// The maximum number of nodes or the percentage of nodes that can be repaired simultaneously when a large number of abnormal nodes exist in the node pool. You can specify a number (for example, 5, valid values: 1 to 100000) or a percentage (for example, 10%, valid values: 1% to 100%). Default value: 1.
	//
	// example:
	//
	// 5
	MaxParallelRepairingNodes *string `json:"max_parallel_repairing_nodes,omitempty" xml:"max_parallel_repairing_nodes,omitempty"`
	// The circuit breaker threshold for auto-healing. When the number or percentage of faulty nodes exceeds this threshold, auto-healing enters the circuit breaker state and stops initiating new auto-healing repair actions. You can specify a number (for example, 10, valid values: 1 to 100000) or a percentage (for example, 20%, valid values: 1% to 100%). Default value: 100%.
	//
	// example:
	//
	// 20%
	MaxUnhealthyNodesThreshold *string `json:"max_unhealthy_nodes_threshold,omitempty" xml:"max_unhealthy_nodes_threshold,omitempty"`
	// Specifies whether to allow node restarts. This parameter takes effect only when auto_repair is set to true. Valid values:
	//
	// - `true`: Allows node restarts.
	//
	// - `false`: Does not allow node restarts.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
}

func (s CreateClusterNodePoolRequestManagementAutoRepairPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagementAutoRepairPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetApprovalRequired() *bool {
	return s.ApprovalRequired
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetAutoRepairPolicyId() *string {
	return s.AutoRepairPolicyId
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetMaxParallelRepairingNodes() *string {
	return s.MaxParallelRepairingNodes
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetMaxUnhealthyNodesThreshold() *string {
	return s.MaxUnhealthyNodesThreshold
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetApprovalRequired(v bool) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetAutoRepairPolicyId(v string) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.AutoRepairPolicyId = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetMaxParallelRepairingNodes(v string) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.MaxParallelRepairingNodes = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetMaxUnhealthyNodesThreshold(v string) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.MaxUnhealthyNodesThreshold = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetRestartNode(v bool) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestManagementAutoUpgradePolicy struct {
	// Specifies whether to allow automatic kubelet upgrades. This parameter takes effect only when auto_upgrade is set to true. Valid values:
	//
	// - `true`: Allows automatic kubelet upgrades.
	//
	// - `false`: Does not allow automatic kubelet upgrades.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow automatic operating system upgrades. This parameter takes effect only when auto_upgrade is set to true. Valid values:
	//
	// - `true`: Allows automatic operating system upgrades.
	//
	// - `false`: Does not allow automatic operating system upgrades.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow automatic runtime upgrades. This parameter takes effect only when auto_upgrade is set to true. Valid values:
	//
	// - `true`: Allows automatic runtime upgrades.
	//
	// - `false`: Does not allow automatic runtime upgrades.
	//
	// Default value: `true`.
	//
	// example:
	//
	// false
	AutoUpgradeRuntime *bool `json:"auto_upgrade_runtime,omitempty" xml:"auto_upgrade_runtime,omitempty"`
}

func (s CreateClusterNodePoolRequestManagementAutoUpgradePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagementAutoUpgradePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeKubelet() *bool {
	return s.AutoUpgradeKubelet
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeOs() *bool {
	return s.AutoUpgradeOs
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeRuntime() *bool {
	return s.AutoUpgradeRuntime
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeKubelet(v bool) *CreateClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeKubelet = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeOs(v bool) *CreateClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeOs = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeRuntime(v bool) *CreateClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeRuntime = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoUpgradePolicy) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestManagementAutoVulFixPolicy struct {
	// The packages to be excluded during vulnerability repair. Default value: `kernel`.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Specifies whether to allow node restarts. This parameter takes effect only when auto_vul_fix is set to true. Valid values:
	//
	// - `true`: Allows node restarts.
	//
	// - `false`: Does not allow node restarts.
	//
	// Default value: `true`.
	//
	// example:
	//
	// false
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels that are allowed for automatic repair, separated by commas. Example: `asap,later`. Valid values:
	//
	// - `asap`: High.
	//
	// - `later`: Medium.
	//
	// - `nntf`: Low.
	//
	// Default value: `asap`.
	//
	// example:
	//
	// asap,nntf
	VulLevel *string `json:"vul_level,omitempty" xml:"vul_level,omitempty"`
}

func (s CreateClusterNodePoolRequestManagementAutoVulFixPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagementAutoVulFixPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) GetExcludePackages() *string {
	return s.ExcludePackages
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) GetVulLevel() *string {
	return s.VulLevel
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) SetExcludePackages(v string) *CreateClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.ExcludePackages = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) SetRestartNode(v bool) *CreateClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.RestartNode = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) SetVulLevel(v string) *CreateClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.VulLevel = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementAutoVulFixPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestManagementUpgradeConfig struct {
	// Deprecated
	//
	// **[Deprecated]*	- Use the auto_upgrade parameter at the upper level instead. Specifies whether to enable auto-upgrade. Valid values:
	//
	// - `true`: Enables auto-upgrade.
	//
	// - `false`: Disables auto-upgrade.
	//
	// example:
	//
	// null
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes. Valid values: 1 to 1000. Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes. This parameter and surge_percentage are mutually exclusive. Nodes become unavailable during upgrades. You can create extra nodes to compensate for the cluster load.
	//
	// > The number of extra nodes should not exceed the current number of nodes.
	//
	// example:
	//
	// 0
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. This parameter and surge are mutually exclusive. Number of extra nodes = Percentage of extra nodes × Number of nodes. For example, if the percentage of extra nodes is set to 50% and the current number of nodes is 6, the number of extra nodes generated is 50% × 6 = 3.
	//
	// example:
	//
	// 0
	SurgePercentage *int64 `json:"surge_percentage,omitempty" xml:"surge_percentage,omitempty"`
}

func (s CreateClusterNodePoolRequestManagementUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagementUpgradeConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) GetMaxUnavailable() *int64 {
	return s.MaxUnavailable
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) GetSurge() *int64 {
	return s.Surge
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) GetSurgePercentage() *int64 {
	return s.SurgePercentage
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) SetAutoUpgrade(v bool) *CreateClusterNodePoolRequestManagementUpgradeConfig {
	s.AutoUpgrade = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) SetMaxUnavailable(v int64) *CreateClusterNodePoolRequestManagementUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) SetSurge(v int64) *CreateClusterNodePoolRequestManagementUpgradeConfig {
	s.Surge = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) SetSurgePercentage(v int64) *CreateClusterNodePoolRequestManagementUpgradeConfig {
	s.SurgePercentage = &v
	return s
}

func (s *CreateClusterNodePoolRequestManagementUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestNodeComponents struct {
	// The configuration of the node component.
	Config *CreateClusterNodePoolRequestNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the node component.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the node component.
	//
	// example:
	//
	// 1.33.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateClusterNodePoolRequestNodeComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeComponents) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeComponents) GetConfig() *CreateClusterNodePoolRequestNodeComponentsConfig {
	return s.Config
}

func (s *CreateClusterNodePoolRequestNodeComponents) GetName() *string {
	return s.Name
}

func (s *CreateClusterNodePoolRequestNodeComponents) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterNodePoolRequestNodeComponents) SetConfig(v *CreateClusterNodePoolRequestNodeComponentsConfig) *CreateClusterNodePoolRequestNodeComponents {
	s.Config = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponents) SetName(v string) *CreateClusterNodePoolRequestNodeComponents {
	s.Name = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponents) SetVersion(v string) *CreateClusterNodePoolRequestNodeComponents {
	s.Version = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolRequestNodeComponentsConfig struct {
	// The custom configuration of the node component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
	// The environment variables of the node component.
	Envs []*CreateClusterNodePoolRequestNodeComponentsConfigEnvs `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
}

func (s CreateClusterNodePoolRequestNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) GetEnvs() []*CreateClusterNodePoolRequestNodeComponentsConfigEnvs {
	return s.Envs
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) SetCustomConfig(v map[string]interface{}) *CreateClusterNodePoolRequestNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) SetEnvs(v []*CreateClusterNodePoolRequestNodeComponentsConfigEnvs) *CreateClusterNodePoolRequestNodeComponentsConfig {
	s.Envs = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) Validate() error {
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterNodePoolRequestNodeComponentsConfigEnvs struct {
	// The name of the environment variable.
	//
	// example:
	//
	// LOG_LEVEL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// info
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateClusterNodePoolRequestNodeComponentsConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeComponentsConfigEnvs) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfigEnvs) GetName() *string {
	return s.Name
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfigEnvs) SetName(v string) *CreateClusterNodePoolRequestNodeComponentsConfigEnvs {
	s.Name = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfigEnvs) SetValue(v string) *CreateClusterNodePoolRequestNodeComponentsConfigEnvs {
	s.Value = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestNodeConfig struct {
	// The Kubelet parameter settings.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
	// The operating system configuration of the node.
	NodeOsConfig *CreateClusterNodePoolRequestNodeConfigNodeOsConfig `json:"node_os_config,omitempty" xml:"node_os_config,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolRequestNodeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeConfig) GetKubeletConfiguration() *KubeletConfig {
	return s.KubeletConfiguration
}

func (s *CreateClusterNodePoolRequestNodeConfig) GetNodeOsConfig() *CreateClusterNodePoolRequestNodeConfigNodeOsConfig {
	return s.NodeOsConfig
}

func (s *CreateClusterNodePoolRequestNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *CreateClusterNodePoolRequestNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeConfig) SetNodeOsConfig(v *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) *CreateClusterNodePoolRequestNodeConfig {
	s.NodeOsConfig = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeConfig) Validate() error {
	if s.KubeletConfiguration != nil {
		if err := s.KubeletConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NodeOsConfig != nil {
		if err := s.NodeOsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolRequestNodeConfigNodeOsConfig struct {
	// The Transparent HugePages configuration.
	Hugepage *Hugepage `json:"hugepage,omitempty" xml:"hugepage,omitempty"`
	// The sysctl kernel parameter settings. Only specific kernel parameters are supported, and the cluster version must be 1.28 or later.
	//
	// example:
	//
	// {"fs.file-max":"1048576"}
	Sysctl map[string]interface{} `json:"sysctl,omitempty" xml:"sysctl,omitempty"`
}

func (s CreateClusterNodePoolRequestNodeConfigNodeOsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeConfigNodeOsConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) GetHugepage() *Hugepage {
	return s.Hugepage
}

func (s *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) GetSysctl() map[string]interface{} {
	return s.Sysctl
}

func (s *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) SetHugepage(v *Hugepage) *CreateClusterNodePoolRequestNodeConfigNodeOsConfig {
	s.Hugepage = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) SetSysctl(v map[string]interface{}) *CreateClusterNodePoolRequestNodeConfigNodeOsConfig {
	s.Sysctl = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeConfigNodeOsConfig) Validate() error {
	if s.Hugepage != nil {
		if err := s.Hugepage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterNodePoolRequestNodepoolInfo struct {
	// The name of the node pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodepool-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group to which the node pool belongs. Instances added to the node pool belong to this resource group.
	//
	// A resource can belong to only one resource group. You can map resource groups to concepts such as projects, applications, or organizations based on your business scenarios.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of the node pool. Valid values:
	//
	// - `ess`: standard node pool, which supports managed features and automatic elastic scaling.
	//
	// - `edge`: edge node pool.
	//
	// - `lingjun`: Lingjun node pool.
	//
	// - `hybrid`: hybrid cloud node pool.
	//
	// example:
	//
	// ess
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateClusterNodePoolRequestNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodepoolInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) GetType() *string {
	return s.Type
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) SetName(v string) *CreateClusterNodePoolRequestNodepoolInfo {
	s.Name = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) SetResourceGroupId(v string) *CreateClusterNodePoolRequestNodepoolInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) SetType(v string) *CreateClusterNodePoolRequestNodepoolInfo {
	s.Type = &v
	return s
}

func (s *CreateClusterNodePoolRequestNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroup struct {
	// Specifies whether to enable auto-renewal for nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: enables auto-renewal.
	//
	// - `false`: disables auto-renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The duration of each auto-renewal cycle. Valid values:
	//
	// - When PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - When PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `security_hardening_os` parameter instead.
	//
	// example:
	//
	// null
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the required number of ECS instances when spot instances cannot be created due to price or inventory issues, provided that `multi_az_policy` is set to `COST_OPTIMIZED`. Valid values:
	//
	// - `true`: automatically creates pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: does not automatically create pay-as-you-go instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The CPU-related configuration options.
	CpuOptions *CreateClusterNodePoolRequestScalingGroupCpuOptions `json:"cpu_options,omitempty" xml:"cpu_options,omitempty" type:"Struct"`
	// The data cloud disk configuration for nodes in the node pool.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The ID of the deployment set. You can use a deployment set to distribute ECS instances created by the node pool across different physical servers to ensure high availability and underlying disaster recovery capabilities. When you create ECS instances in a deployment set, the instances are distributed and started in the specified region based on the predefined deployment policy.
	//
	//
	// 	Notice: After you select a deployment set, the maximum number of nodes in the node pool is limited. By default, a deployment set supports a maximum of 20 × the number of zones nodes. The number of zones is determined by the vSwitches. Proceed with caution and ensure that the quota in the deployment set is sufficient to prevent node scale-out failures.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The expected number of nodes in the node pool.
	//
	// This is the total number of nodes that the node pool should maintain. We recommend that you configure at least two nodes to ensure the normal operation of cluster components. You can scale out or scale in the node pool by adjusting the expected number of nodes.
	//
	// If you do not need to create nodes, set this parameter to 0 and manually increase the number of nodes later.
	//
	// example:
	//
	// 0
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// Specifies whether to enable the high-density cloud disk mode. This mode is supported only when the node pool uses instance types. After this mode is enabled, the total number of system cloud disks and data cloud disks does not exceed the maximum number of high-density cloud disks supported by the instance type.
	//
	// example:
	//
	// false
	EnableHighDensityMode *bool `json:"enable_high_density_mode,omitempty" xml:"enable_high_density_mode,omitempty"`
	// The ID of the custom image. By default, the system-provided image is used.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the operating system image. Valid values:
	//
	// - `AliyunLinux`: Alinux2 image.
	//
	// - `AliyunLinuxSecurity`: Alinux2 UEFI image.
	//
	// - `AliyunLinux3`: Alinux3 image.
	//
	// - `AliyunLinux3Arm64`: Alinux3 ARM image.
	//
	// - `AliyunLinux3Security`: Alinux3 UEFI image.
	//
	// - `CentOS`: CentOS image.
	//
	// - `Windows`: Windows image.
	//
	// - `WindowsCore`: Windows Core image.
	//
	// - `ContainerOS`: container-optimized image.
	//
	// - `AliyunLinux3ContainerOptimized`: Alinux3 container-optimized image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of nodes in the node pool. Valid values:
	//
	//
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The instance metadata access configuration for ECS instances.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance property configuration.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of instance types for nodes in the node pool. When the node pool scales out nodes, it selects instance types that meet the requirements from the specified list to create instances.
	//
	// Valid values for the number of instance types: 1 to 10.
	//
	//
	// > To ensure high availability, we recommend that you select multiple instance types.
	//
	// This parameter is required.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method for public IP addresses. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-data-transfer.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth for public IP addresses of nodes. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair used for password-free logon. You must specify either this parameter or `login_password`.
	//
	// > If the node pool uses the ContainerOS operating system, only `key_pair` is supported.
	//
	// example:
	//
	// np-key-name
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether to log on to the created ECS instances as a non-root user.
	//
	//
	//
	// - true: logs on as a non-root user (ecs-user).
	//
	// - false: logs on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. You must specify either this parameter or `key_pair`. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// ****
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: scales ECS instances based on the vSwitches (VSwitchIds.N) that you specify. If ECS instances cannot be created in the zone of the vSwitch with the highest priority, the system automatically uses the vSwitch with the next highest priority to create ECS instances.
	//
	// - `COST_OPTIMIZED`: creates ECS instances based on the vCPU unit price in ascending order. When the scaling configuration specifies multiple instance types with the preemptible billing method, the system preferentially creates spot instances. You can use the `CompensateWithOnDemand` parameter to specify if it is to automatically create pay-as-you-go instances when spot instances cannot be created due to inventory shortages or other reasons.
	//
	//   > `COST_OPTIMIZED` takes effect only when the scaling configuration specifies multiple instance types or uses spot instances.
	//
	// - `BALANCE`: evenly allocates ECS instances across the zones specified in the scaling group. If the zones become unbalanced due to inventory shortages or other reasons, you can call the [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) API operation to balance the resources.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances required by the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than this value, the system preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the excess instances after the scaling group meets the minimum number of pay-as-you-go instances (`on_demand_base_capacity`). Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`.
	//
	// - When `period_unit` is set to `Week`, valid values for `period` are 1, 2, 3, and 4.
	//
	// - When `period_unit` is set to `Month`, valid values for `period` are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`.
	//
	// - `Month`: month.
	//
	// - `Week`: week.
	//
	// Default value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `image_type` parameter instead.
	//
	// The operating system distribution. Valid values:
	//
	// - `CentOS`
	//
	// - `AliyunLinux`
	//
	// - `Windows`
	//
	// - `WindowsCore`
	//
	// Default value: `AliyunLinux`.
	//
	// example:
	//
	// null
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The private pool configuration.
	PrivatePoolOptions *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker RAM role.
	//
	// 	- If this parameter is left empty, the default worker RAM role created for the cluster is used.
	//
	// 	- If this parameter is specified, the RAM role must be a **normal service role*	- and its **trusted service*	- must be set to **ECS**. For more information, see [Create a normal service role](https://help.aliyun.com/document_detail/116800.html). If the specified RAM role is not the default worker RAM role created for the cluster, the role name cannot start with `KubernetesMasterRole-` or `KubernetesWorkerRole-`.
	//
	// 	Notice: This parameter is supported only by ACK managed clusters of version 1.22 and later.
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// The list of ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy used for instance creation. Take note of the following items when you configure this parameter:
	//
	// This parameter takes effect only when you create pay-as-you-go instances.
	//
	// This parameter cannot be specified together with private_pool_options.match_criteria or private_pool_options.id.
	ResourcePoolOptions *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling group mode. Valid values:
	//
	// - `release`: Standard mode. Instances are scaled by creating and releasing ECS instances based on the requested resource usage.
	//
	// - `recycle`: Extreme mode. Instances are scaled by creating, stopping, and starting instances to improve the speed of subsequent scaling operations. When an instance is stopped, compute resources are not billed, and only storage fees are charged, except for instances with local disks.
	//
	// Default value: `release`.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// Deprecated
	//
	// The security group ID of the node pool. Specify either this parameter or security_group_ids. We recommend that you use security_group_ids.
	//
	// example:
	//
	// sg-wz9a8g2mt6x5llu0****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The list of security group IDs. Specify either this parameter or security_group_id. We recommend that you use security_group_ids. If both security_group_id and security_group_ids are specified, security_group_ids takes precedence.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// Specifies whether to enable Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: Enables Alibaba Cloud OS security hardening.
	//
	// - `false`: Does not enable Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. You can enable MLPS 2.0 security hardening for nodes only when the system image is Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3. Alibaba Cloud provides baseline check standards and scanning programs for classified protection compliance for Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 MLPS 2.0 Level 3 images.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of available instance types. The scaling group evenly creates spot instances across multiple instance types with the lowest costs. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable spot instance replenishment. After this feature is enabled, when the scaling group receives a system message that a spot instance will be reclaimed, it attempts to create a new instance to replace the spot instance to be reclaimed. Valid values:
	//
	// - `true`: Enables spot instance replenishment.
	//
	// - `false`: Does not enable spot instance replenishment.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The market price range configuration for a single spot instance type.
	SpotPriceLimit []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The spot instance strategy. Valid values:
	//
	// - `NoSpot`: Non-spot instance.
	//
	// - `SpotWithPriceLimit`: Sets a maximum price for spot instances.
	//
	// - `SpotAsPriceGo`: The system automatically bids based on the current market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/165053.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable performance burst for the node system disk. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// This parameter is supported only when system_disk_category is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The categories of the system disk. If a higher-priority disk category is unavailable, the system automatically attempts to use the next priority disk category to create the system disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The category of the node system disk. Valid values:
	//
	// - `cloud_efficiency`: Ultra cloud disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud_auto`: ESSD AutoPL cloud disk.
	//
	// - `cloud_essd_entry`: ESSD Entry cloud disk.
	//
	// Default value: `cloud_efficiency`.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm used for the system disk. Valid values: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: Encrypts the system disk.
	//
	// - false: Does not encrypt the system disk.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The ID of the KMS key used for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the node system disk. This parameter takes effect only for ESSD cloud disks. The performance level is related to the disk size. For more information, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html). Valid values:
	//
	// - PL0: Medium concurrent I/O performance and relatively stable read/write latency.
	//
	// - PL1: Medium concurrent I/O performance and relatively stable read/write latency.
	//
	// - PL2: High concurrent I/O performance and stable read/write latency.
	//
	// - PL3: Extremely high concurrent I/O performance and highly stable read/write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS for the node system disk. Valid values: 0 to min{50,000, 1000 × capacity - baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50000}.
	//
	// This parameter is supported only when system_disk_category is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the node system disk. Unit: GiB. Valid values: 20 to 2048.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The snapshot policy for the system disk.
	//
	// example:
	//
	// sp-0jl6xnmme8v7o935****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// Adds tags only to ECS instances. Tag keys must be unique and can be up to 128 characters in length. Neither tag keys nor tag values can start with "aliyun" or "acs:", or contain "https://" or "http://".
	Tags []*CreateClusterNodePoolRequestScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of vSwitch IDs. Valid values: 1 to 8.
	//
	// > To ensure high availability, we recommend that you select vSwitches in different zones.
	//
	// This parameter is required.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s CreateClusterNodePoolRequestScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetCisEnabled() *bool {
	return s.CisEnabled
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetCpuOptions() *CreateClusterNodePoolRequestScalingGroupCpuOptions {
	return s.CpuOptions
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDesiredSize() *int64 {
	return s.DesiredSize
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetDiskInit() []*DiskInit {
	return s.DiskInit
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetEnableHighDensityMode() *bool {
	return s.EnableHighDensityMode
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetImageType() *string {
	return s.ImageType
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstanceMetadataOptions() *InstanceMetadataOptions {
	return s.InstanceMetadataOptions
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstancePatterns() []*InstancePatterns {
	return s.InstancePatterns
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetKeyPair() *string {
	return s.KeyPair
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetOnDemandBaseCapacity() *int64 {
	return s.OnDemandBaseCapacity
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int64 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetPlatform() *string {
	return s.Platform
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetPrivatePoolOptions() *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetResourcePoolOptions() *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSpotInstancePools() *int64 {
	return s.SpotInstancePools
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSpotPriceLimit() []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit {
	return s.SpotPriceLimit
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskKmsKeyId() *string {
	return s.SystemDiskKmsKeyId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetTags() []*CreateClusterNodePoolRequestScalingGroupTags {
	return s.Tags
}

func (s *CreateClusterNodePoolRequestScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetAutoRenew(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetAutoRenewPeriod(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetCisEnabled(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.CisEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetCompensateWithOnDemand(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetCpuOptions(v *CreateClusterNodePoolRequestScalingGroupCpuOptions) *CreateClusterNodePoolRequestScalingGroup {
	s.CpuOptions = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDataDisks(v []*DataDisk) *CreateClusterNodePoolRequestScalingGroup {
	s.DataDisks = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDeploymentsetId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.DeploymentsetId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDesiredSize(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetDiskInit(v []*DiskInit) *CreateClusterNodePoolRequestScalingGroup {
	s.DiskInit = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetEnableHighDensityMode(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.EnableHighDensityMode = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetImageId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.ImageId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetImageType(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.ImageType = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstanceChargeType(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstanceMetadataOptions(v *InstanceMetadataOptions) *CreateClusterNodePoolRequestScalingGroup {
	s.InstanceMetadataOptions = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstancePatterns(v []*InstancePatterns) *CreateClusterNodePoolRequestScalingGroup {
	s.InstancePatterns = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInstanceTypes(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInternetChargeType(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.InternetChargeType = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetInternetMaxBandwidthOut(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetKeyPair(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetLoginAsNonRoot(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetLoginPassword(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetMultiAzPolicy(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.MultiAzPolicy = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetOnDemandBaseCapacity(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetPeriod(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.Period = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetPeriodUnit(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetPlatform(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.Platform = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetPrivatePoolOptions(v *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) *CreateClusterNodePoolRequestScalingGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetRamRoleName(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.RamRoleName = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetRdsInstances(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.RdsInstances = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetResourcePoolOptions(v *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) *CreateClusterNodePoolRequestScalingGroup {
	s.ResourcePoolOptions = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetScalingPolicy(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSecurityGroupId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSecurityGroupIds(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSecurityHardeningOs(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.SecurityHardeningOs = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSocEnabled(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.SocEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSpotInstancePools(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSpotInstanceRemedy(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSpotPriceLimit(v []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) *CreateClusterNodePoolRequestScalingGroup {
	s.SpotPriceLimit = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSpotStrategy(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SpotStrategy = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskBurstingEnabled(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskCategories(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategories = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskCategory(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskEncryptAlgorithm(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskEncrypted(v bool) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskKmsKeyId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskKmsKeyId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskPerformanceLevel(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskProvisionedIops(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskSize(v int64) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetSystemDiskSnapshotPolicyId(v string) *CreateClusterNodePoolRequestScalingGroup {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetTags(v []*CreateClusterNodePoolRequestScalingGroupTags) *CreateClusterNodePoolRequestScalingGroup {
	s.Tags = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetVswitchIds(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) Validate() error {
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DiskInit != nil {
		for _, item := range s.DiskInit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstanceMetadataOptions != nil {
		if err := s.InstanceMetadataOptions.Validate(); err != nil {
			return err
		}
	}
	if s.InstancePatterns != nil {
		for _, item := range s.InstancePatterns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcePoolOptions != nil {
		if err := s.ResourcePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SpotPriceLimit != nil {
		for _, item := range s.SpotPriceLimit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateClusterNodePoolRequestScalingGroupCpuOptions struct {
	// Specifies whether to enable nested virtualization. Valid values: disabled: disables nested virtualization. enabled: enables nested virtualization.
	//
	// example:
	//
	// enabled
	NestedVirtualization *string `json:"nested_virtualization,omitempty" xml:"nested_virtualization,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupCpuOptions) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupCpuOptions) GetNestedVirtualization() *string {
	return s.NestedVirtualization
}

func (s *CreateClusterNodePoolRequestScalingGroupCpuOptions) SetNestedVirtualization(v string) *CreateClusterNodePoolRequestScalingGroupCpuOptions {
	s.NestedVirtualization = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupCpuOptions) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions struct {
	// The ID of the private pool. When `match_criteria` is set to `Target`, you must specify the private pool ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The private pool capacity matching mode for instance startup. After an elasticity assurance or capacity reservation takes effect, a private pool is generated for you to select when you start instances. Valid values:
	//
	// - `Open`: open mode. The system automatically matches the capacity of an open private pool. If no eligible private pool capacity is available, the system uses public pool resources to start the instances.
	//
	// - `Target`: specified mode. The system uses the capacity of the specified private pool to start instances. If the specified private pool capacity is unavailable, the instances fail to start.
	//
	// - `None`: no private pool mode. The system does not use private pool capacity to start instances.
	//
	// example:
	//
	// Target
	MatchCriteria *string `json:"match_criteria,omitempty" xml:"match_criteria,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) SetId(v string) *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) SetMatchCriteria(v string) *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroupResourcePoolOptions struct {
	// The list of private pool IDs, which are the IDs of elasticity assurances or capacity reservations. You can specify only the IDs of private pools in Target mode. Valid values of N: 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool strategy used when creating instances. Resource pools include private pools generated after elasticity assurances or capacity reservations take effect, as well as public pools, which are available for selection during instance startup. Valid values:
	//
	// PrivatePoolFirst: Private pool first. When this strategy is selected and resouce_pool_options.private_pool_ids is specified, the specified private pools are used first. If no private pool is specified or the specified private pools have insufficient capacity, open private pools are automatically matched. If no eligible private pool is available, the public pool is used to create the instance.
	//
	// PrivatePoolOnly: Private pool only. When this strategy is selected, resouce_pool_options.private_pool_ids must be specified. If the specified private pools have insufficient capacity, the instance fails to start.
	//
	// None: No resource pool strategy is used.
	//
	// Default value: None.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) SetPrivatePoolIds(v []*string) *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) SetStrategy(v string) *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroupSpotPriceLimit struct {
	// The spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum price per instance.
	//
	// <props="china">Unit: CNY/hour.
	//
	//
	//
	// <props="intl">Unit: USD/hour.
	//
	// example:
	//
	// 0.39
	PriceLimit *string `json:"price_limit,omitempty" xml:"price_limit,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) SetInstanceType(v string) *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) SetPriceLimit(v string) *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupSpotPriceLimit) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroupTags struct {
	// The tag key.
	//
	// example:
	//
	// node-k-1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// node-v-1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateClusterNodePoolRequestScalingGroupTags) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestScalingGroupTags) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestScalingGroupTags) GetKey() *string {
	return s.Key
}

func (s *CreateClusterNodePoolRequestScalingGroupTags) GetValue() *string {
	return s.Value
}

func (s *CreateClusterNodePoolRequestScalingGroupTags) SetKey(v string) *CreateClusterNodePoolRequestScalingGroupTags {
	s.Key = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupTags) SetValue(v string) *CreateClusterNodePoolRequestScalingGroupTags {
	s.Value = &v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroupTags) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestTeeConfig struct {
	// Specifies whether to enable the confidential computing cluster. Valid values:
	//
	// - true: Enables confidential computing.
	//
	// - false: Does not enable confidential computing.
	//
	// example:
	//
	// true
	TeeEnable *bool `json:"tee_enable,omitempty" xml:"tee_enable,omitempty"`
}

func (s CreateClusterNodePoolRequestTeeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestTeeConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestTeeConfig) GetTeeEnable() *bool {
	return s.TeeEnable
}

func (s *CreateClusterNodePoolRequestTeeConfig) SetTeeEnable(v bool) *CreateClusterNodePoolRequestTeeConfig {
	s.TeeEnable = &v
	return s
}

func (s *CreateClusterNodePoolRequestTeeConfig) Validate() error {
	return dara.Validate(s)
}
