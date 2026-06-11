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
	// The intelligent managed node pool configurations.
	AutoMode *CreateClusterNodePoolRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The auto scaling configurations.
	AutoScaling *CreateClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// [This parameter is deprecated] Use desired_size instead.
	//
	// The number of nodes in the node pool.
	//
	// example:
	//
	// null
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The Lingjun node pool configurations.
	EfloNodeGroup *CreateClusterNodePoolRequestEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// Specifies whether to use the host network for the pod network.
	//
	// - `true`: host network. Pods directly use the network stack of the host and share the IP address and ports with the host.
	//
	// - `false`: container network. Pods have an independent network stack and do not occupy host network ports.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated]
	//
	// The configurations of the edge node pool.
	InterconnectConfig *CreateClusterNodePoolRequestInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter takes effect only for node pools of the \\`edge\\` type. Valid values:
	//
	// - `basic`: public network. Nodes in the node pool interact with cloud nodes over the Internet. Applications in the node pool cannot directly access the VPC in the cloud.
	//
	// - `private`: private network. Nodes in the node pool connect to the cloud over a leased line, VPN, or CEN to achieve higher communication quality and better security.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Specifies whether nodes in the edge node pool can communicate with each other at Layer 3.
	//
	// - `true`: All nodes in the node pool can communicate with each other at Layer 3.
	//
	// - `false`: All hosts in the node pool cannot communicate with each other at Layer 3.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The cluster-related configurations.
	KubernetesConfig *CreateClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool feature.
	Management *CreateClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// Deprecated
	//
	// [This parameter is deprecated]
	//
	// The maximum number of nodes that the edge node pool can contain.
	//
	// example:
	//
	// null
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// A list of node components.
	NodeComponents []*CreateClusterNodePoolRequestNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// The node configurations.
	NodeConfig *CreateClusterNodePoolRequestNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The node pool configurations.
	NodepoolInfo *CreateClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group for the node pool.
	ScalingGroup *CreateClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The configurations of the confidential computing cluster.
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
	// Specifies whether to enable the intelligent managed mode.
	//
	// Valid values:
	//
	// - true: Enables the intelligent managed mode. You can enable this mode only when the intelligent managed mode is enabled for the cluster.
	//
	// - false: Disables the intelligent managed mode.
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
	// [This parameter is deprecated] Use \\`internet_charge_type\\` and \\`internet_max_bandwidth_out\\` instead.
	//
	// The peak bandwidth of the EIP. Unit: Mbps.
	//
	// example:
	//
	// null
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated] Use \\`internet_charge_type\\` and \\`internet_max_bandwidth_out\\` instead.
	//
	// The billing method of the EIP. Valid values:
	//
	// - `PayByBandwidth`: pay-by-bandwidth.
	//
	// - `PayByTraffic`: pay-by-traffic.
	//
	// Default value: `PayByBandwidth`.
	//
	// example:
	//
	// null
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable auto scaling. Valid values:
	//
	// - `true`: Enables auto scaling for the node pool. If the resources planned for the cluster cannot meet the scheduling requirements of pods, Container Service for Kubernetes (ACK) automatically scales out or scales in nodes based on the configured minimum and maximum numbers of instances. For clusters of Kubernetes 1.24 or later, instant scaling is enabled by default. For clusters of a Kubernetes version earlier than 1.24, node autoscaling is enabled by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: Disables auto scaling. ACK adjusts the number of nodes in the node pool based on the value of \\`desired_size\\` to maintain a specific number of nodes.
	//
	// If you set this parameter to false, other parameters in \\`auto_scaling\\` do not take effect.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated] This parameter is deprecated. Use \\`internet_charge_type\\` and \\`internet_max_bandwidth_out\\` instead.
	//
	// Specifies whether to associate an EIP with the node. Valid values:
	//
	// - `true`: associates an EIP with the node.
	//
	// - `false`: does not associate an EIP with the node.
	//
	// Default value: `false`.
	//
	// example:
	//
	// null
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances that can be created in the node pool. This does not include existing instances. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
	//
	// The value must be in the range of [\\`min_instances\\`, 2000]. Default value: 0.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances that can be created in the node pool. This does not include existing instances. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
	//
	// The value must be in the range of [0, \\`max_instances\\`]. Default value: 0.
	//
	// > - If the minimum number of instances is not 0, the specified number of ECS instances are automatically created after the scaling group is created.
	//
	// >
	//
	// > - We recommend that you set the maximum number of instances to a value that is not smaller than the current number of nodes in the node pool. Otherwise, nodes in the node pool are scaled in after auto scaling is enabled.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The type of instances that are automatically scaled. This parameter takes effect only if \\`enable\\` is set to \\`true\\`. Valid values:
	//
	// - `cpu`: regular instance.
	//
	// - `gpu`: GPU-accelerated instance.
	//
	// - `gpushare`: shared GPU-accelerated instance.
	//
	// - `spot`: spot instance.
	//
	// Default value: `cpu`.
	//
	// 	Notice: You cannot change the value of this parameter after the node pool is created.
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
	// The ID of the Lingjun cluster that you want to associate with the Lingjun node pool when you create the node pool.
	//
	// example:
	//
	// i1169130516633730****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The ID of the Lingjun group in the Lingjun cluster that you want to associate with the Lingjun node pool when you create the node pool.
	//
	// example:
	//
	// ng-ec3c96ff0aa****
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s CreateClusterNodePoolRequestEfloNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestEfloNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetClusterId(v string) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) SetGroupId(v string) *CreateClusterNodePoolRequestEfloNodeGroup {
	s.GroupId = &v
	return s
}

func (s *CreateClusterNodePoolRequestEfloNodeGroup) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestInterconnectConfig struct {
	// [This parameter is deprecated]
	//
	// The network bandwidth of the enhanced edge node pool. Unit: Mbps.
	//
	// example:
	//
	// null
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the CCN instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// null
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The region of the CCN instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// null
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the CEN instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// null
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The subscription duration of the enhanced edge node pool. Unit: months.
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
	// Specifies whether to install Cloud Monitor on the ECS nodes. After Cloud Monitor is installed, you can view the monitoring information of the created ECS instances in the Cloud Monitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: Installs Cloud Monitor on the ECS nodes.
	//
	// - `false`: Does not install Cloud Monitor on the ECS nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of the node. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// - `static`: Allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// - `none`: Enables the default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels that you want to add to the nodes in the Kubernetes cluster.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name. If you customize the node name, the node name, ECS instance name, and ECS instance hostname are changed.
	//
	// > For Windows instances for which custom node names are enabled, the hostname is fixed as the IP address. Hyphens (-) are used to replace periods (.) in the IP address. The hostname does not contain a prefix or a suffix.
	//
	// A node name consists of a prefix, the node IP address, and a suffix.
	//
	// - The total length must be 2 to 64 characters. The node name must start and end with a lowercase letter or a digit.
	//
	// - The prefix and suffix can contain uppercase letters, lowercase letters, digits, hyphens (-), and periods (.). They must start with an uppercase or lowercase letter. They cannot start or end with a hyphen (-) or a period (.). You cannot use consecutive hyphens (-)or periods (.). You cannot use consecutive hyphens (-) or periods (.).
	//
	// - The prefix is required (due to an ECS limit). The suffix is optional.
	//
	// - The node IP address is the complete private IP address of the node.
	//
	// For example, if the node IP address is 192.XX.YY.55, the prefix is aliyun.com, and the suffix is test.
	//
	// - If the node is a Linux node, the node name, ECS instance name, and ECS instance hostname are all aliyun.com192.XX.YY.55test.
	//
	// - If the node is a Windows node, the ECS instance hostname is 192-XX-YY-55, and the node name and ECS instance name are both aliyun.com192.XX.YY.55test.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-join instance user data. The specified user data script is run before the node joins the cluster. For more information, see [User-Data scripts](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. ACK supports the following three container runtimes.
	//
	// - containerd: We recommend that you use this runtime. It is supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: a sandboxed container that provides higher isolation. It is supported by clusters of Kubernetes 1.31 or earlier.
	//
	// - docker: no longer maintained. It is supported by clusters of Kubernetes 1.22 or earlier.
	//
	// Default value: containerd.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The version of the container runtime.
	//
	// example:
	//
	// 1.6.38
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// The taint configurations.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Specifies whether the scaled-out nodes are unschedulable.
	//
	// - true: The nodes are unschedulable.
	//
	// - false: The nodes are schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After the node joins the cluster, the specified user data script is run. For more information, see [User-Data scripts](https://help.aliyun.com/document_detail/49121.html).
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
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to enable auto node repair. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
	//
	// - `true`: Enables auto node repair.
	//
	// - `false`: Disables auto node repair.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto node repair policy.
	AutoRepairPolicy *CreateClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to enable auto node upgrade. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
	//
	// - `true`: Enables auto node upgrade.
	//
	// - `false`: Disables auto node upgrade.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto node upgrade policy.
	AutoUpgradePolicy *CreateClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically fix CVE vulnerabilities. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
	//
	// - `true`: Automatically fixes CVE vulnerabilities.
	//
	// - `false`: Does not automatically fix CVE vulnerabilities.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The policy for automatically fixing CVE vulnerabilities.
	AutoVulFixPolicy *CreateClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable the managed node pool feature. Valid values:
	//
	// - `true`: Enables the managed node pool feature.
	//
	// - `false`: Disables the managed node pool feature. If you set this parameter to \\`false\\`, the other parameters of \\`management\\` do not take effect.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated] Use the \\`auto_upgrade\\` parameter instead.
	//
	// The auto upgrade configurations. This parameter takes effect only if \\`enable\\` is set to \\`true\\`.
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
	// Specifies whether to allow node restart. This parameter takes effect only if \\`auto_repair\\` is set to \\`true\\`. Valid values:
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Disallows node restart.
	//
	// Default value: `true`
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

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *CreateClusterNodePoolRequestManagementAutoRepairPolicy) SetApprovalRequired(v bool) *CreateClusterNodePoolRequestManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
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
	// Specifies whether to allow auto kubelet upgrade. This parameter takes effect only if \\`auto_upgrade\\` is set to \\`true\\`. Valid values:
	//
	// - `true`: Allows auto kubelet upgrade.
	//
	// - `false`: Disallows auto kubelet upgrade.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow auto OS upgrade. This parameter takes effect only if \\`auto_upgrade\\` is set to \\`true\\`. Valid values:
	//
	// - `true`: Allows auto OS upgrade.
	//
	// - `false`: Disallows auto OS upgrade.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow auto runtime upgrade. This parameter takes effect only if \\`auto_upgrade\\` is set to \\`true\\`. Valid values:
	//
	// - `true`: Allows auto runtime upgrade.
	//
	// - `false`: Disallows auto runtime upgrade.
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
	// The packages that should be excluded from vulnerability fixing.
	//
	// Default value: `kernel`.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only if \\`auto_vul_fix\\` is set to \\`true\\`. Valid values:
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Disallows node restart.
	//
	// Default value: `true`
	//
	// example:
	//
	// false
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels that are allowed to be automatically fixed. Separate multiple levels with commas. Example: `asap,later`. Supported vulnerability levels:
	//
	// - `asap`: high
	//
	// - `later`: medium
	//
	// - `nntf`: low
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
	// [This parameter is deprecated] Use the \\`auto_upgrade\\` parameter instead.
	//
	// Specifies whether to enable auto upgrade. Valid values:
	//
	// - `true`: enables auto upgrade.
	//
	// - `false`: disables auto upgrade.
	//
	// example:
	//
	// null
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes.
	//
	// The value must be in the range of [1, 1000].
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes. You can specify only one of \\`surge\\` and \\`surge_percentage\\`.
	//
	// Nodes become unavailable during an upgrade. You can create extra nodes to compensate for the load on the cluster.
	//
	// > We recommend that the number of extra nodes does not exceed the current number of nodes.
	//
	// example:
	//
	// 0
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You can specify only one of \\`surge\\` and \\`surge_percentage\\`.
	//
	// Number of extra nodes = Percentage of extra nodes × Number of nodes. For example, if you set the percentage of extra nodes to 50% and the number of existing nodes is 6, three extra nodes are created.
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
	// The configurations of the node component.
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
	// The custom configurations of the node component.
	CustomConfig map[string]*string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s CreateClusterNodePoolRequestNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) SetCustomConfig(v map[string]*string) *CreateClusterNodePoolRequestNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestNodeConfig struct {
	// The Kubelet parameter settings.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
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

func (s *CreateClusterNodePoolRequestNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *CreateClusterNodePoolRequestNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *CreateClusterNodePoolRequestNodeConfig) Validate() error {
	if s.KubeletConfiguration != nil {
		if err := s.KubeletConfiguration.Validate(); err != nil {
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
	// The ID of the resource group. Instances that are created in the node pool belong to this resource group.
	//
	// A resource can belong to only one resource group. You can map resource groups to concepts such as projects, applications, or organizations based on your business scenarios.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of the node pool. Valid values:
	//
	// - `ess`: a regular node pool. This type of node pool provides managed features and auto scaling.
	//
	// - `edge`: an edge node pool.
	//
	// - `lingjun`: a Lingjun node pool.
	//
	// - `hybrid`: a hybrid cloud node pool.
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
	// Specifies whether to enable auto-renewal for the nodes in the node pool. This parameter takes effect only if \\`instance_charge_type\\` is set to \\`PrePaid\\`. Valid values:
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
	// The auto-renewal period. Valid values:
	//
	// - If \\`PeriodUnit\\` is set to \\`Week\\`: 1, 2, and 3.
	//
	// - If \\`PeriodUnit\\` is set to \\`Month\\`: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [Deprecated] Use the \\`security_hardening_os\\` parameter instead.
	//
	// example:
	//
	// null
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the instance quantity requirement when \\`multi_az_policy\\` is set to \\`COST_OPTIMIZED\\` and spot instances cannot be created due to issues such as price or insufficient inventory. Valid values:
	//
	// - `true`: Allows the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: Does not allow the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The data disk configurations of the nodes in the node pool.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The ID of the deployment set. You can use a deployment set to distribute the ECS instances created in the node pool across different physical servers to ensure high availability and underlying disaster recovery. When you create ECS instances in a deployment set, the instances are launched in the specified region based on the deployment strategy that you set.
	//
	// 	Notice:
	//
	// After you select a deployment set, the maximum number of nodes in the node pool is limited. By default, a deployment set supports a maximum of 20 × Number of zones (the number of zones is determined by the vSwitches) nodes. Select a deployment set with sufficient quota to prevent node creation failures.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The expected number of nodes in the node pool.
	//
	// The total number of nodes that the node pool should maintain. We recommend that you configure at least two nodes to ensure that cluster components run as expected. You can adjust the expected number of nodes to scale out or scale in the node pool.
	//
	// If you do not need to create nodes, set this parameter to 0. You can manually adjust the number of nodes later.
	//
	// example:
	//
	// 0
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configurations.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The ID of the custom image. By default, the system-provided image is used.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The OS image type. Valid values:
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
	// - `WindowsCore`: WindowsCore image.
	//
	// - `ContainerOS`: Container-Optimized OS.
	//
	// - `AliyunLinux3ContainerOptimized`: Alinux3 Container-Optimized OS.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of the nodes in the node pool. Valid values:
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
	// The ECS instance metadata access configurations.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance property configurations.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// A list of instance types for the nodes in the node pool. When the system creates nodes in the node pool, it selects an instance type from the list that meets the requirements.
	//
	// The number of instance types must be in the range of [1, 10].
	//
	// > For high availability, we recommend that you select multiple instance types.
	//
	// This parameter is required.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method of the public IP address. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound public bandwidth of the node. Unit: Mbit/s. The value must be in the range of [1, 100].
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair that is used for passwordless logon. You must specify one of \\`key_pair\\` and \\`login_password\\`.
	//
	// > If you select Container-Optimized OS for the node pool, you can use only \\`key_pair\\`.
	//
	// example:
	//
	// np-key-name
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether to log on to the created ECS instances as a non-root user.
	//
	// - true: Logs on as the ecs-user.
	//
	// - false: Logs on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. You must specify one of \\`key_pair\\` and \\`login_password\\`. The password must be 8 to 30 characters in length, and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// ****
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for the ECS instances in the multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: Scales instances based on the vSwitch priority. The system scales instances based on the order of vSwitches that you specify in \\`VSwitchIds.N\\`. If the system fails to create an ECS instance in the zone where the vSwitch with the highest priority resides, it automatically uses the vSwitch with the next highest priority to create the instance.
	//
	// - `COST_OPTIMIZED`: Creates instances based on the vCPU unit price in ascending order. When multiple instance types are specified and the preemptible instance policy is configured, the system gives priority to creating the lowest-cost instance type. You can also use the \\`CompensateWithOnDemand\\` parameter to specify whether to automatically create pay-as-you-go instances when preemptible instances cannot be created due to insufficient inventory.
	//
	//   > `COST_OPTIMIZED` takes effect only when multiple instance types are specified or the preemptible instance policy is configured.
	//
	// - `BALANCE`: Evenly distributes ECS instances across the specified zones. If the distribution of ECS instances becomes unbalanced due to insufficient inventory, you can call the API [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) operation to balance the resource distribution.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be included in the scaling group. The value must be in the range of [0, 1000]. If the number of pay-as-you-go instances is less than this value, the system gives priority to creating pay-as-you-go instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that are created after the minimum number of pay-as-you-go instances (\\`on_demand_base_capacity\\`) is met. The value must be in the range of [0, 100].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of the nodes in the node pool. This parameter is required and takes effect only if \\`instance_charge_type\\` is set to \\`PrePaid\\`.
	//
	// - If \\`period_unit\\` is set to \\`Week\\`, the valid values of \\`period\\` are 1, 2, 3, and 4.
	//
	// - If \\`period_unit\\` is set to \\`Month\\`, the valid values of \\`period\\` are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the nodes in the node pool. This parameter is required and takes effect only if \\`instance_charge_type\\` is set to \\`PrePaid\\`.
	//
	// - `Month`: The billing cycle is measured in months.
	//
	// - `Week`: The billing cycle is measured in weeks.
	//
	// Default value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated] Use the \\`image_type\\` parameter instead.
	//
	// The OS distribution. Valid values:
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
	// The private pool configurations.
	PrivatePoolOptions *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker RAM role.
	//
	// - If this parameter is left empty, the default worker RAM role of the cluster is used.
	//
	// - If this parameter is not empty, the specified RAM role must be a **service role**, and its **trusted service*	- must be **Elastic Compute Service**. For more information, see [Create a service role](https://help.aliyun.com/document_detail/116800.html). If the specified RAM role is not the default worker RAM role of the cluster, the name of the role cannot start with \\`KubernetesMasterRole-\\` or \\`KubernetesWorkerRole-\\`.
	//
	// 	Notice:
	//
	// This parameter is supported only by ACK managed clusters of Kubernetes 1.22 or later.
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// A list of RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and policy used to create instances. Note:
	//
	// This parameter takes effect only when you create pay-as-you-go instances.
	//
	// This parameter cannot be set at the same time as \\`private_pool_options.match_criteria\\` and \\`private_pool_options.id\\`.
	ResourcePoolOptions *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling mode of the scaling group. Valid values:
	//
	// - `release`: standard mode. The system creates and releases ECS instances to scale the group.
	//
	// - `recycle`: accelerated mode. The system creates, stops, and starts ECS instances to scale the group. This improves the scaling speed. When an instance is stopped, its computing resources are not billed, but its storage resources are. This does not apply to instances with local disks.
	//
	// Default value: `release`.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// Deprecated
	//
	// The ID of the security group for the node pool. You must specify one of \\`security_group_ids\\` and \\`security_group_id\\`. We recommend that you specify \\`security_group_ids\\`.
	//
	// example:
	//
	// sg-wz9a8g2mt6x5llu0****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// A list of security group IDs. You must specify one of \\`security_group_ids\\` and \\`security_group_id\\`. We recommend that you specify \\`security_group_ids\\`. If you specify both \\`security_group_id\\` and \\`security_group_ids\\`, \\`security_group_ids\\` takes precedence.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// Alibaba Cloud OS security hardening. Valid values:
	//
	// - `true`: Enables Alibaba Cloud OS security hardening.
	//
	// - `false`: Disables Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. You can enable MLPS 2.0 security hardening for nodes only when you select Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 for the OS image. Alibaba Cloud provides baseline check standards and scanning programs for Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 Level 3 of MLPS 2.0 to ensure classified protection compliance.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of instance types that you can specify. The scaling group creates preemptible instances of multiple instance types that are available at the lowest cost. The value must be in the range of [1, 10].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable the instance reclaim mode. After this mode is enabled, when the system receives a message that a spot instance is about to be reclaimed, the scaling group attempts to create a new instance to replace the instance that is about to be reclaimed. Valid values:
	//
	// - `true`: Enables the instance reclaim mode.
	//
	// - `false`: Disables the instance reclaim mode.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The configurations of the price range for a single spot instance.
	SpotPriceLimit []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The bidding policy for the spot instances. Valid values:
	//
	// - `NoSpot`: The instance is not a spot instance.
	//
	// - `SpotWithPriceLimit`: Sets the maximum bid price for the spot instance.
	//
	// - `SpotAsPriceGo`: The system automatically bids based on the current market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/165053.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk of the node. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// This parameter can be set only when \\`system_disk_category\\` is set to \\`cloud_auto\\`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The types of system disks. If a disk of a high-priority type is unavailable, the system automatically uses a disk of the next priority type to create the system disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system disk of the node. Valid values:
	//
	// - `cloud_efficiency`: ultra disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud_auto`: ESSD AutoPL disk.
	//
	// - `cloud_essd_entry`: ESSD Entry disk.
	//
	// Default value: `cloud_efficiency`.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm that is used to encrypt the system disk. Valid value: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: encrypts the system disk.
	//
	// - false: does not encrypt the system disk.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The ID of the KMS key that is used to encrypt the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the system disk for each node. This parameter applies only to Enhanced SSD (ESSD) disks. The performance level of an ESSD is determined by its size. For more information, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: A moderate maximum concurrent I/O performance and a relatively stable read/write latency.
	//
	// - PL1: A moderate maximum concurrent I/O performance and a relatively stable read/write latency.
	//
	// - PL2: A high maximum concurrent I/O performance and a stable read/write latency.
	//
	// - PL3: A very high maximum concurrent I/O performance and a very stable read/write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS of the system disk of the node.
	//
	// Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter can be set only when \\`system_disk_category\\` is set to \\`cloud_auto\\`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk of the node. Unit: GiB.
	//
	// The value must be in the range of [20, 2048].
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
	// The tags that you want to add only to ECS instances.
	//
	// A tag key cannot be repeated. The tag key can be up to 128 characters in length. The tag key and the tag value cannot start with “aliyun” or “acs:”, and cannot contain “https\\://” or “http\\://”.
	Tags []*CreateClusterNodePoolRequestScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// A list of vSwitch IDs. The number of vSwitch IDs must be in the range of [1, 8].
	//
	// > For high availability, we recommend that you select vSwitches in different zones.
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

type CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions struct {
	// The ID of the private pool. You must specify the private pool ID when \\`match_criteria\\` is set to \\`Target\\`.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private pool. This parameter specifies the capacity of the private pool that you want to use to launch instances. An elastic assurance service or a capacity reservation service takes effect after it is generated. You can select a capacity type when you launch an instance. Valid values:
	//
	// - `Open`: Open mode. The system automatically matches the capacity of open private pools. If no matching private pool is found, the instance is launched using public pool resources.
	//
	// - `Target`: Target mode. The instance is launched using the capacity of the specified private pool. If the capacity of the private pool is unavailable, the instance fails to be launched.
	//
	// - `None`: The instance is launched without using the capacity of a private pool.
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
	// A list of private pool IDs. These are the IDs of elastic assurance services or capacity reservation services. You can only specify the IDs of private pools in Target mode. The number of IDs must be in the range of 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy for instance creation. Resource pools include private pools (generated by elastic assurance or capacity reservation services) and public pools. Valid values:
	//
	// PrivatePoolFirst: Prioritizes private pools. If you specify \\`resouce_pool_options.private_pool_ids\\`, the specified private pools are used first. If no private pool is specified or the specified pools have insufficient capacity, the system automatically tries to use open private pools. If no suitable private pool is found, the public pool is used.
	//
	// PrivatePoolOnly: Uses only private pools. You must specify \\`resouce_pool_options.private_pool_ids\\`. If the specified private pools have insufficient capacity, the instance fails to launch.
	//
	// None: No resource pool policy is used.
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
	// The instance type of the spot instance.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum bid price for a single instance.
	//
	// <props="intl">
	//
	// Unit: USD/hour.
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
	// The name of the tag.
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
	// Specifies whether to enable the confidential computing cluster.
	//
	// - true: Enables confidential computing.
	//
	// - false: Disables confidential computing.
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
