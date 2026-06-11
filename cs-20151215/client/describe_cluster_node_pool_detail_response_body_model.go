// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoMode(v *DescribeClusterNodePoolDetailResponseBodyAutoMode) *DescribeClusterNodePoolDetailResponseBody
	GetAutoMode() *DescribeClusterNodePoolDetailResponseBodyAutoMode
	SetAutoScaling(v *DescribeClusterNodePoolDetailResponseBodyAutoScaling) *DescribeClusterNodePoolDetailResponseBody
	GetAutoScaling() *DescribeClusterNodePoolDetailResponseBodyAutoScaling
	SetEfloNodeGroup(v *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) *DescribeClusterNodePoolDetailResponseBody
	GetEfloNodeGroup() *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup
	SetHostNetwork(v bool) *DescribeClusterNodePoolDetailResponseBody
	GetHostNetwork() *bool
	SetInterconnectConfig(v *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) *DescribeClusterNodePoolDetailResponseBody
	GetInterconnectConfig() *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig
	SetInterconnectMode(v string) *DescribeClusterNodePoolDetailResponseBody
	GetInterconnectMode() *string
	SetIntranet(v bool) *DescribeClusterNodePoolDetailResponseBody
	GetIntranet() *bool
	SetKubernetesConfig(v *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) *DescribeClusterNodePoolDetailResponseBody
	GetKubernetesConfig() *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig
	SetManagement(v *DescribeClusterNodePoolDetailResponseBodyManagement) *DescribeClusterNodePoolDetailResponseBody
	GetManagement() *DescribeClusterNodePoolDetailResponseBodyManagement
	SetMaxNodes(v int64) *DescribeClusterNodePoolDetailResponseBody
	GetMaxNodes() *int64
	SetNodeComponents(v []*DescribeClusterNodePoolDetailResponseBodyNodeComponents) *DescribeClusterNodePoolDetailResponseBody
	GetNodeComponents() []*DescribeClusterNodePoolDetailResponseBodyNodeComponents
	SetNodeConfig(v *DescribeClusterNodePoolDetailResponseBodyNodeConfig) *DescribeClusterNodePoolDetailResponseBody
	GetNodeConfig() *DescribeClusterNodePoolDetailResponseBodyNodeConfig
	SetNodepoolInfo(v *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) *DescribeClusterNodePoolDetailResponseBody
	GetNodepoolInfo() *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo
	SetScalingGroup(v *DescribeClusterNodePoolDetailResponseBodyScalingGroup) *DescribeClusterNodePoolDetailResponseBody
	GetScalingGroup() *DescribeClusterNodePoolDetailResponseBodyScalingGroup
	SetStatus(v *DescribeClusterNodePoolDetailResponseBodyStatus) *DescribeClusterNodePoolDetailResponseBody
	GetStatus() *DescribeClusterNodePoolDetailResponseBodyStatus
	SetTeeConfig(v *DescribeClusterNodePoolDetailResponseBodyTeeConfig) *DescribeClusterNodePoolDetailResponseBody
	GetTeeConfig() *DescribeClusterNodePoolDetailResponseBodyTeeConfig
}

type DescribeClusterNodePoolDetailResponseBody struct {
	// The smart hosting configurations.
	AutoMode *DescribeClusterNodePoolDetailResponseBodyAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The configurations of the node pool that is configured for automatic scaling.
	AutoScaling   *DescribeClusterNodePoolDetailResponseBodyAutoScaling   `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	EfloNodeGroup *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// Indicates whether the pod network uses the host network mode.
	//
	// - `true`: host network. Pods directly use the host\\"s network stack and share IP addresses and ports with the host.
	//
	// - `false`: container network. Pods have an independent network stack and do not use host network ports.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// [This parameter is deprecated]
	//
	// The network configurations of the edge node pool. This parameter is valid only for edge node pools.
	InterconnectConfig *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter is valid only for `edge` node pools. Valid values:
	//
	// - `basic`: public network. The nodes in the node pool interact with cloud nodes over the Internet. Applications in the node pool cannot directly access the VPC in the cloud.
	//
	// - `private`: dedicated network. The nodes in the node pool connect to the cloud network through leased lines, VPNs, or CEN. This provides higher communication quality between the cloud and the edge and offers more effective security.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Indicates whether nodes in the edge node pool have Layer 3 network connectivity.
	//
	// - `true`: connected. All nodes in this node pool have Layer 3 network connectivity.
	//
	// - `false`: not connected. All hosts in this node pool do not have Layer 3 network connectivity.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The cluster-related configurations.
	KubernetesConfig *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool.
	Management *DescribeClusterNodePoolDetailResponseBodyManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// [This parameter is deprecated]
	//
	// The maximum number of nodes that the edge node pool can contain.
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The list of node components.
	NodeComponents []*DescribeClusterNodePoolDetailResponseBodyNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// The node configurations.
	NodeConfig *DescribeClusterNodePoolDetailResponseBodyNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The node pool configurations.
	NodepoolInfo *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group for the node pool.
	ScalingGroup *DescribeClusterNodePoolDetailResponseBodyScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The status of the node pool.
	Status *DescribeClusterNodePoolDetailResponseBodyStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// The configurations of the confidential computing cluster.
	TeeConfig *DescribeClusterNodePoolDetailResponseBodyTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetAutoMode() *DescribeClusterNodePoolDetailResponseBodyAutoMode {
	return s.AutoMode
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetAutoScaling() *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	return s.AutoScaling
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetEfloNodeGroup() *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup {
	return s.EfloNodeGroup
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetHostNetwork() *bool {
	return s.HostNetwork
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetInterconnectConfig() *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	return s.InterconnectConfig
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetInterconnectMode() *string {
	return s.InterconnectMode
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetIntranet() *bool {
	return s.Intranet
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetKubernetesConfig() *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	return s.KubernetesConfig
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetManagement() *DescribeClusterNodePoolDetailResponseBodyManagement {
	return s.Management
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetMaxNodes() *int64 {
	return s.MaxNodes
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetNodeComponents() []*DescribeClusterNodePoolDetailResponseBodyNodeComponents {
	return s.NodeComponents
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetNodeConfig() *DescribeClusterNodePoolDetailResponseBodyNodeConfig {
	return s.NodeConfig
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetNodepoolInfo() *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	return s.NodepoolInfo
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetScalingGroup() *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	return s.ScalingGroup
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetStatus() *DescribeClusterNodePoolDetailResponseBodyStatus {
	return s.Status
}

func (s *DescribeClusterNodePoolDetailResponseBody) GetTeeConfig() *DescribeClusterNodePoolDetailResponseBodyTeeConfig {
	return s.TeeConfig
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetAutoMode(v *DescribeClusterNodePoolDetailResponseBodyAutoMode) *DescribeClusterNodePoolDetailResponseBody {
	s.AutoMode = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetAutoScaling(v *DescribeClusterNodePoolDetailResponseBodyAutoScaling) *DescribeClusterNodePoolDetailResponseBody {
	s.AutoScaling = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetEfloNodeGroup(v *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) *DescribeClusterNodePoolDetailResponseBody {
	s.EfloNodeGroup = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetHostNetwork(v bool) *DescribeClusterNodePoolDetailResponseBody {
	s.HostNetwork = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetInterconnectConfig(v *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) *DescribeClusterNodePoolDetailResponseBody {
	s.InterconnectConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetInterconnectMode(v string) *DescribeClusterNodePoolDetailResponseBody {
	s.InterconnectMode = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetIntranet(v bool) *DescribeClusterNodePoolDetailResponseBody {
	s.Intranet = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetKubernetesConfig(v *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) *DescribeClusterNodePoolDetailResponseBody {
	s.KubernetesConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetManagement(v *DescribeClusterNodePoolDetailResponseBodyManagement) *DescribeClusterNodePoolDetailResponseBody {
	s.Management = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetMaxNodes(v int64) *DescribeClusterNodePoolDetailResponseBody {
	s.MaxNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetNodeComponents(v []*DescribeClusterNodePoolDetailResponseBodyNodeComponents) *DescribeClusterNodePoolDetailResponseBody {
	s.NodeComponents = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetNodeConfig(v *DescribeClusterNodePoolDetailResponseBodyNodeConfig) *DescribeClusterNodePoolDetailResponseBody {
	s.NodeConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetNodepoolInfo(v *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) *DescribeClusterNodePoolDetailResponseBody {
	s.NodepoolInfo = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetScalingGroup(v *DescribeClusterNodePoolDetailResponseBodyScalingGroup) *DescribeClusterNodePoolDetailResponseBody {
	s.ScalingGroup = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetStatus(v *DescribeClusterNodePoolDetailResponseBodyStatus) *DescribeClusterNodePoolDetailResponseBody {
	s.Status = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) SetTeeConfig(v *DescribeClusterNodePoolDetailResponseBodyTeeConfig) *DescribeClusterNodePoolDetailResponseBody {
	s.TeeConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBody) Validate() error {
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
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
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

type DescribeClusterNodePoolDetailResponseBodyAutoMode struct {
	// Indicates whether to enable the feature.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyAutoMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyAutoMode) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoMode) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoMode) SetEnable(v bool) *DescribeClusterNodePoolDetailResponseBodyAutoMode {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoMode) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyAutoScaling struct {
	// The peak EIP bandwidth.
	//
	// Valid values: 1 to 100. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	// - `PayByBandwidth`: pay-by-bandwidth.
	//
	// - `PayByTraffic`: pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Indicates whether to enable automatic scaling. Valid values:
	//
	// - `true`: enables automatic scaling for the node pool. If the resources of the cluster cannot meet the scheduling requirements of pods, ACK automatically scales out or in nodes based on the configured minimum and maximum numbers of instances. For clusters of Kubernetes 1.24 or later, node elastic scaling is enabled by default. For clusters of a Kubernetes version earlier than 1.24, node autoscaling is enabled by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: disables automatic scaling. ACK adjusts the number of nodes in the node pool to the expected number of nodes. The number of nodes is always the same as the expected number of nodes.
	//
	// If this parameter is set to false, other parameters in auto_scaling do not take effect.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Indicates whether to associate an EIP with the node pool. Valid values:
	//
	// - `true`: Associates an EIP with the node pool.
	//
	// - `false`: Does not associate an EIP with the node pool.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances that can be created in the node pool. This value does not include the existing instances.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances that can be created in the node pool. This value does not include the existing instances.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The type of automatic scaling that is configured for the node pool. This parameter is specified based on the instance type for automatic scaling. Valid values:
	//
	// - `cpu`: regular instances.
	//
	// - `gpu`: GPU-accelerated instances.
	//
	// - `gpushare`: shared GPU-accelerated instances.
	//
	// - `spot`: spot instances.
	//
	// example:
	//
	// cpu
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyAutoScaling) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyAutoScaling) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetEipBandwidth() *int64 {
	return s.EipBandwidth
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetEipInternetChargeType() *string {
	return s.EipInternetChargeType
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetIsBondEip() *bool {
	return s.IsBondEip
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetEipBandwidth(v int64) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetEipInternetChargeType(v string) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.EipInternetChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetEnable(v bool) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetIsBondEip(v bool) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.IsBondEip = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetMaxInstances(v int64) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.MaxInstances = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetMinInstances(v int64) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.MinInstances = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) SetType(v string) *DescribeClusterNodePoolDetailResponseBodyAutoScaling {
	s.Type = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyAutoScaling) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup struct {
	// example:
	//
	// i113790071760688002461
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// i128147721760688002463
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) SetClusterId(v string) *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) SetGroupId(v string) *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup {
	s.GroupId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyEfloNodeGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyInterconnectConfig struct {
	// [This parameter is deprecated]
	//
	// The network bandwidth of the enhanced edge node pool. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the CCN instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The region where the CCN instance that is associated with the enhanced edge node pool resides.
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the CEN instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The subscription duration of the enhanced edge node pool. Unit: months.
	//
	// example:
	//
	// 1
	ImprovedPeriod *string `json:"improved_period,omitempty" xml:"improved_period,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GetCcnId() *string {
	return s.CcnId
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GetCcnRegionId() *string {
	return s.CcnRegionId
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GetCenId() *string {
	return s.CenId
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) GetImprovedPeriod() *string {
	return s.ImprovedPeriod
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) SetBandwidth(v int64) *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) SetCcnId(v string) *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	s.CcnId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) SetCcnRegionId(v string) *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	s.CcnRegionId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) SetCenId(v string) *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	s.CenId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) SetImprovedPeriod(v string) *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig {
	s.ImprovedPeriod = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyKubernetesConfig struct {
	// Indicates whether to install Cloud Monitor on the ECS nodes. After you install Cloud Monitor, you can view the monitoring information of the created ECS instances in the Cloud Monitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: Installs Cloud Monitor on the ECS nodes.
	//
	// - `false`: Does not install Cloud Monitor on the ECS nodes.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy for the nodes. The following policies are supported for clusters of Kubernetes 1.12.6 and later:
	//
	// - `static`: Allows pods with specific resource characteristics on a node to have enhanced CPU affinity and exclusivity.
	//
	// - `none`: Enables the default CPU affinity scheme.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The node labels.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name.
	//
	// A node name consists of a prefix, the IP address of the node, and a suffix:
	//
	// - The prefix and suffix can consist of one or more parts separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). The node name must start and end with a lowercase letter or a digit.
	//
	// - The IP address segment length indicates the number of digits to be truncated from the end of the node IP address. Valid values: 5 to 12.
	//
	// For example, if the node IP address is 192.168.0.55, the prefix is aliyun.com, the IP address segment length is 5, and the suffix is test, the node name is aliyun.com00055test.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-custom data of the node pool. The script is run before the node is initialized. For more information, see [Generate instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. ACK supports the following container runtimes.
	//
	// - containerd: recommended. It is supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: a sandboxed container that provides higher isolation. It is supported by clusters of Kubernetes 1.31 and earlier.
	//
	// - docker: no longer maintained. It is supported by clusters of Kubernetes 1.22 and earlier.
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
	// The node taints. Taints work with tolerations to prevent pods from being scheduled to unsuitable nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Indicates whether the scaled-out nodes are unschedulable.
	//
	// - true: The nodes are unschedulable.
	//
	// - false: The nodes are schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The custom data of the node pool. The script is run after the node is initialized. For more information, see [Generate instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvYmluL3NoCmVjaG8gIkhlbGxvIEFDSyEi
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetCmsEnabled() *bool {
	return s.CmsEnabled
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetLabels() []*Tag {
	return s.Labels
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetTaints() []*Taint {
	return s.Taints
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetUnschedulable() *bool {
	return s.Unschedulable
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetCmsEnabled(v bool) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.CmsEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetCpuPolicy(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.CpuPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetLabels(v []*Tag) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.Labels = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetNodeNameMode(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.NodeNameMode = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetPreUserData(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetRuntime(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.Runtime = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetRuntimeVersion(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetTaints(v []*Taint) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.Taints = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetUnschedulable(v bool) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.Unschedulable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) SetUserData(v string) *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig) Validate() error {
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

type DescribeClusterNodePoolDetailResponseBodyManagement struct {
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Indicates whether to enable auto repair. This parameter takes effect only if enable is set to true.
	//
	// - `true`: Auto repair is enabled.
	//
	// - `false`: Auto repair is disabled.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The policy for automatic node repair.
	AutoRepairPolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Indicates whether to enable automatic node upgrades. This parameter takes effect only if enable is set to true.
	//
	// - `true`: Automatic upgrades are enabled.
	//
	// - `false`: Automatic upgrades are disabled.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The policy for automatic upgrades.
	AutoUpgradePolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Indicates whether to automatically fix CVEs. This parameter takes effect only if enable is set to true.
	//
	// - `true`: CVEs are automatically fixed.
	//
	// - `false`: CVEs are not automatically fixed.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The policy for automatically fixing CVEs.
	AutoVulFixPolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Indicates whether to enable the managed node pool feature. Valid values:
	//
	// - `true`: Enables the managed node pool feature.
	//
	// - `false`: Disables the managed node pool feature. Other parameters in this section take effect only if this parameter is set to true.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The automatic upgrade configurations. This parameter takes effect only if enable is set to true.
	UpgradeConfig *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagement) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagement) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoFaultDiagnosis() *bool {
	return s.AutoFaultDiagnosis
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoRepair() *bool {
	return s.AutoRepair
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoRepairPolicy() *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy {
	return s.AutoRepairPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoUpgradePolicy() *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy {
	return s.AutoUpgradePolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoVulFix() *bool {
	return s.AutoVulFix
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetAutoVulFixPolicy() *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy {
	return s.AutoVulFixPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) GetUpgradeConfig() *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig {
	return s.UpgradeConfig
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoFaultDiagnosis(v bool) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoFaultDiagnosis = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoRepair(v bool) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoRepair = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoRepairPolicy(v *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoRepairPolicy = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoUpgrade(v bool) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoUpgrade = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoUpgradePolicy(v *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoUpgradePolicy = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoVulFix(v bool) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoVulFix = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetAutoVulFixPolicy(v *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.AutoVulFixPolicy = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetEnable(v bool) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) SetUpgradeConfig(v *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) *DescribeClusterNodePoolDetailResponseBodyManagement {
	s.UpgradeConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagement) Validate() error {
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

type DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy struct {
	// Indicates whether manual approval is required for node repair.
	//
	// example:
	//
	// false
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// The ID of the auto repair policy
	//
	// example:
	//
	// r-xxxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// Indicates whether to allow node restart. This parameter takes effect only if auto_repair is set to true.
	//
	// - `true`: Nodes can be restarted.
	//
	// - `false`: Nodes cannot be restarted.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) GetApprovalRequired() *bool {
	return s.ApprovalRequired
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) GetAutoRepairPolicyId() *string {
	return s.AutoRepairPolicyId
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) SetApprovalRequired(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) SetAutoRepairPolicyId(v string) *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy {
	s.AutoRepairPolicyId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) SetRestartNode(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy struct {
	// Indicates whether to allow automatic kubelet upgrades. This parameter takes effect only if auto_upgrade is set to true. Valid values:
	//
	// - `true`: Automatic kubelet upgrades are allowed.
	//
	// - `false`: Automatic kubelet upgrades are not allowed.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) GetAutoUpgradeKubelet() *bool {
	return s.AutoUpgradeKubelet
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) SetAutoUpgradeKubelet(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy {
	s.AutoUpgradeKubelet = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy struct {
	// The packages that should be excluded during CVE fixing.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Indicates whether to allow node restart. This parameter takes effect only if auto_vul_fix is set to true. Valid values:
	//
	// - `true`: Nodes can be restarted.
	//
	// - `false`: Nodes cannot be restarted.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The levels of CVEs that are allowed to be automatically fixed. The levels are separated by commas.
	//
	// - `asap`: high
	//
	// - `later`: medium
	//
	// - `nntf`: low
	//
	// example:
	//
	// asap,nntf
	VulLevel *string `json:"vul_level,omitempty" xml:"vul_level,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) GetExcludePackages() *string {
	return s.ExcludePackages
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) SetExcludePackages(v string) *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy {
	s.ExcludePackages = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) SetRestartNode(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy {
	s.RestartNode = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) SetVulLevel(v string) *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy {
	s.VulLevel = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig struct {
	// Indicates whether to enable automatic upgrades. Valid values:
	//
	// - `true`: Automatic upgrades are enabled.
	//
	// - `false`: Automatic upgrades are disabled.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes. Valid values: 1 to 1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes. You can specify only one of surge and surge_percentage.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You can specify only one of surge and surge_percentage.
	//
	// The number of extra nodes = Percentage of extra nodes × Number of nodes. For example, if you set the percentage of extra nodes to 50% and the number of existing nodes is 6, the number of extra nodes is 3 (50% × 6).
	//
	// example:
	//
	// 50
	SurgePercentage *int64 `json:"surge_percentage,omitempty" xml:"surge_percentage,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) GetMaxUnavailable() *int64 {
	return s.MaxUnavailable
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) GetSurge() *int64 {
	return s.Surge
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) GetSurgePercentage() *int64 {
	return s.SurgePercentage
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) SetAutoUpgrade(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig {
	s.AutoUpgrade = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) SetMaxUnavailable(v int64) *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) SetSurge(v int64) *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig {
	s.Surge = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) SetSurgePercentage(v int64) *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig {
	s.SurgePercentage = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyNodeComponents struct {
	// The configurations of the node component.
	Config *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s DescribeClusterNodePoolDetailResponseBodyNodeComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) GetConfig() *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig {
	return s.Config
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) SetConfig(v *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) *DescribeClusterNodePoolDetailResponseBodyNodeComponents {
	s.Config = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) SetName(v string) *DescribeClusterNodePoolDetailResponseBodyNodeComponents {
	s.Name = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) SetVersion(v string) *DescribeClusterNodePoolDetailResponseBodyNodeComponents {
	s.Version = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig struct {
	// The custom configurations of the node component.
	CustomConfig map[string]*string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) SetCustomConfig(v map[string]*string) *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyNodeConfig struct {
	// The Kubelet parameter settings.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
	// The node OS configurations.
	NodeOsConfig *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig `json:"node_os_config,omitempty" xml:"node_os_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) GetKubeletConfiguration() *KubeletConfig {
	return s.KubeletConfiguration
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) GetNodeOsConfig() *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig {
	return s.NodeOsConfig
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *DescribeClusterNodePoolDetailResponseBodyNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) SetNodeOsConfig(v *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) *DescribeClusterNodePoolDetailResponseBodyNodeConfig {
	s.NodeOsConfig = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) Validate() error {
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

type DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig struct {
	// The Hugepage configurations.
	Hugepage *Hugepage `json:"hugepage,omitempty" xml:"hugepage,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) GetHugepage() *Hugepage {
	return s.Hugepage
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) SetHugepage(v *Hugepage) *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig {
	s.Hugepage = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfigNodeOsConfig) Validate() error {
	if s.Hugepage != nil {
		if err := s.Hugepage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodePoolDetailResponseBodyNodepoolInfo struct {
	// The time when the node pool was created.
	//
	// example:
	//
	// 2025-04-10T14:25:37.285530433+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Indicates whether the node pool is the default node pool. A cluster usually has only one default node pool. Valid values:
	//
	// - `true`: the default node pool.
	//
	// - `false`: not the default node pool.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
	// The name of the node pool.
	//
	// example:
	//
	// default-nodepool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np615c0e0966124216a0412e10afe0****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of the node pool.
	//
	// - `ess`: a regular node pool. It includes the features of managed node pools and automatic scaling.
	//
	// - `edge`: an edge node pool.
	//
	// - `lingjun`: a Lingjun node pool.
	//
	// example:
	//
	// ess
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the node pool was last updated.
	//
	// example:
	//
	// 2025-04-15T15:39:45.41+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetCreated(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.Created = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetIsDefault(v bool) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.IsDefault = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetName(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetNodepoolId(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetRegionId(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetResourceGroupId(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetType(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.Type = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) SetUpdated(v string) *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo {
	s.Updated = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyScalingGroup struct {
	// Indicates whether to enable auto-renewal for the nodes. This parameter takes effect only if instance_charge_type is set to PrePaid. Valid values:
	//
	// - `true`: Auto-renewal is enabled.
	//
	// - `false`: Auto-renewal is disabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The duration of each auto-renewal. Valid values:
	//
	// - If PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated] Use the security_hardening_os parameter instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// If multi_az_policy is set to COST_OPTIMIZED, this parameter specifies whether to allow the system to automatically create on-demand instances to meet the required number of ECS instances when it is not possible to create a sufficient number of spot instances due to price or stock issues. Valid values:
	//
	// - `true`: Allows the system to automatically create on-demand instances to meet the required number of ECS instances.
	//
	// - `false`: Does not allow the system to automatically create on-demand instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The combination of the configurations, such as the type and size, of the data disks of the nodes.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The expected number of nodes in the node pool.
	//
	// example:
	//
	// 2
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The configurations for block device initialization.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The custom image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The OS image type.
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
	// - `ContainerOS`: Container-optimized image.
	//
	// - `AliyunLinux3ContainerOptimized`: Alinux3 container-optimized image.
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
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The configurations for accessing the metadata of ECS instances.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance attribute configurations.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of node instance types.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method for the public IP address of the nodes.
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound public bandwidth of the nodes. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair. You must specify either this parameter or login_password. When the node pool is a managed node pool, only key_pair is supported.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Indicates whether to log on to the created ECS instances as a non-root user.
	//
	// - true: Logs on as a non-root user (ecs-user).
	//
	// - false: Logs on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. You must specify either this parameter or key_pair. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// For security reasons, the password is encrypted in the query result.
	//
	// example:
	//
	// ********
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for the ECS instances in the multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: Scales instances based on the vSwitches that you define (VSwitchIds.N). If the ECS instances cannot be created in the zone of the vSwitch with a higher priority, the system automatically uses the vSwitch with the next priority to create the instances.
	//
	// - `COST_OPTIMIZED`: Attempts to create instances at the lowest vCPU unit price. If multiple instance types are specified for the scaling configuration and the preemption policy is configured, the system preferentially creates the corresponding spot instances. You can also use the `CompensateWithOnDemand` parameter to specify whether to automatically try to create on-demand instances when spot instances cannot be created due to reasons such as stock shortages.
	//
	//   > `COST_OPTIMIZED` takes effect only when multiple instance types are specified or spot instances are used for the scaling configuration.
	//
	// - `BALANCE`: Evenly distributes ECS instances across the specified zones of the scaling group. If the distribution of ECS instances becomes unbalanced between zones due to stock shortages, you can call the API RebalanceInstances operation to balance the resources. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) .
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// BALANCE
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of on-demand instances that the scaling group must contain. Valid values: 0 to 1000. If the number of on-demand instances is less than this value, on-demand instances are preferentially created.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of on-demand instances among the instances that exceed the minimum number of on-demand instances (on_demand_base_capacity). Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of the nodes. This parameter is required and takes effect only if instance_charge_type is set to PrePaid.
	//
	// - If period_unit is set to Week, the valid values of period are 1, 2, 3, and 4.
	//
	// - If period_unit is set to Month, the valid values of period are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 0
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the nodes. This parameter is required if instance_charge_type is set to PrePaid.
	//
	// - `Month`: The billing cycle is measured in months.
	//
	// - `Week`: The billing cycle is measured in weeks.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
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
	// example:
	//
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The private node pool configurations.
	PrivatePoolOptions *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// [This parameter is deprecated] Use ram_role_name instead.
	//
	// example:
	//
	// KubernetesWorkerRole-021dc54f-929b-437a-8ae0-34c24d3e****
	RamPolicy *string `json:"ram_policy,omitempty" xml:"ram_policy,omitempty"`
	// The name of the worker RAM role.
	//
	// example:
	//
	// KubernetesWorkerRole-4a4fa089-80c1-48a5-b3c6-9349311f****
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// If you specify a list of RDS instances, the ECS nodes of the cluster are automatically added to the RDS instance whitelist.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy used when creating instances.
	ResourcePoolOptions *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-2zeieod8giqmov7z****
	ScalingGroupId *string `json:"scaling_group_id,omitempty" xml:"scaling_group_id,omitempty"`
	// The scaling group mode. Valid values:
	//
	// - `release`: standard mode. Instances are created and released based on the resource usage.
	//
	// - `recycle`: fast mode. Instances are created, stopped, and started to accelerate scaling. Compute resources are not billed when instances are stopped, but storage resources are. This does not apply to instances with local disks.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The security group ID of the node pool. If the node pool is associated with multiple security groups, this is the first value in security_group_ids.
	//
	// example:
	//
	// sg-2ze60ockeekspl3d****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The list of security group IDs for the node pool.
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
	// Indicates whether to enable classified protection compliance. You can enable classified protection compliance for nodes only when you select Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 as the OS image. Alibaba Cloud provides baseline check standards and scanning programs for MLPS 2.0 Level 3-compliant Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 images.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of available instance types. The scaling group creates spot instances of multiple types that have the lowest costs in a balanced manner. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Indicates whether to enable the feature of supplementing spot instances. If this feature is enabled, the scaling group attempts to create a new instance to replace a spot instance that is reclaimed. Valid values:
	//
	// - `true`: Enables the feature of supplementing spot instances.
	//
	// - `false`: Disables the feature of supplementing spot instances.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The configurations of the price range for spot instances.
	SpotPriceLimit []*DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The preemption policy for the spot instances. Valid values:
	//
	// - NoSpot: The instances are not spot instances.
	//
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	//
	// - SpotAsPriceGo: The system automatically places bids based on the market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Indicates whether to enable performance burst for the system disk of the nodes. Valid values:
	//
	// - true: Enables performance burst. If you enable this feature, the cloud disk can temporarily improve its performance to handle sudden data read and write pressure when the business is unstable.
	//
	// - false: Disables performance burst.
	//
	// This parameter can be set only when system_disk_category is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The types of system disks. When a disk of a high-priority type is not available, the system automatically tries the next-priority disk type to create the system disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system disk of the nodes. Valid values:
	//
	// - `cloud_efficiency`: ultra disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: Enhanced SSD (ESSD).
	//
	// - `cloud_auto`: ESSD AutoPL disk.
	//
	// - `cloud_essd_entry`: ESSD Entry disk.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm that is used for the system disk. Valid value: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Indicates whether to encrypt the system disk. Valid values:
	//
	// - `true`: Encrypts the system disk.
	//
	// - `false`: Does not encrypt the system disk.
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
	// The performance level of the system disk of the nodes. This parameter is valid only for ESSDs. The disk performance level is related to the disk size. For more information, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: The I/O performance is moderate and the read/write latency is stable.
	//
	// - PL1: The I/O performance is moderate and the read/write latency is stable.
	//
	// - PL2: The I/O performance is high and the read/write latency is stable.
	//
	// - PL3: The I/O performance is very high and the read/write latency is very stable.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The pre-configured read and write IOPS of the system disk of the nodes.
	//
	// Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter can be set only when system_disk_category is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk of the nodes. Unit: GiB.
	//
	// Valid values: 20 to 2048.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The system disk snapshot policy
	//
	// example:
	//
	// sp-bp11g8z59rawcud9****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// The ECS instance tags.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of vSwitch IDs.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetCisEnabled() *bool {
	return s.CisEnabled
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetDesiredSize() *int64 {
	return s.DesiredSize
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetDiskInit() []*DiskInit {
	return s.DiskInit
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInstanceMetadataOptions() *InstanceMetadataOptions {
	return s.InstanceMetadataOptions
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInstancePatterns() []*InstancePatterns {
	return s.InstancePatterns
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetKeyPair() *string {
	return s.KeyPair
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetOnDemandBaseCapacity() *int64 {
	return s.OnDemandBaseCapacity
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int64 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetPrivatePoolOptions() *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetRamPolicy() *string {
	return s.RamPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetResourcePoolOptions() *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSpotInstancePools() *int64 {
	return s.SpotInstancePools
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSpotPriceLimit() []*DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit {
	return s.SpotPriceLimit
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskKmsKeyId() *string {
	return s.SystemDiskKmsKeyId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetTags() []*Tag {
	return s.Tags
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetAutoRenew(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetAutoRenewPeriod(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetCisEnabled(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.CisEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetCompensateWithOnDemand(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetDataDisks(v []*DataDisk) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.DataDisks = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetDeploymentsetId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.DeploymentsetId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetDesiredSize(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetDiskInit(v []*DiskInit) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.DiskInit = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetImageId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetImageType(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.ImageType = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInstanceChargeType(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInstanceMetadataOptions(v *InstanceMetadataOptions) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InstanceMetadataOptions = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInstancePatterns(v []*InstancePatterns) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InstancePatterns = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInstanceTypes(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInternetChargeType(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetInternetMaxBandwidthOut(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetKeyPair(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.KeyPair = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetLoginAsNonRoot(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.LoginAsNonRoot = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetLoginPassword(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.LoginPassword = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetMultiAzPolicy(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.MultiAzPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetOnDemandBaseCapacity(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetPeriod(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.Period = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetPeriodUnit(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetPlatform(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.Platform = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetPrivatePoolOptions(v *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetRamPolicy(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.RamPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetRamRoleName(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.RamRoleName = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetRdsInstances(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.RdsInstances = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetResourcePoolOptions(v *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.ResourcePoolOptions = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetScalingGroupId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetScalingPolicy(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSecurityGroupId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSecurityGroupIds(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSecurityHardeningOs(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SecurityHardeningOs = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSocEnabled(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SocEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSpotInstancePools(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSpotInstanceRemedy(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSpotPriceLimit(v []*DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SpotPriceLimit = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSpotStrategy(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskBurstingEnabled(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskCategories(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskCategory(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskEncryptAlgorithm(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskEncrypted(v bool) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskKmsKeyId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskKmsKeyId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskPerformanceLevel(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskProvisionedIops(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskSize(v int64) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetSystemDiskSnapshotPolicyId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetTags(v []*Tag) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.Tags = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetVswitchIds(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) Validate() error {
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

type DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions struct {
	// The private node pool ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. This parameter specifies the capacity option for the private pool that is used to start instances. After an elastic assurance service or a capacity reservation service takes effect, a private pool is generated. You can select the private pool to start instances. Valid values:
	//
	// - `Open`: The system automatically matches the capacity of an open private pool. If no matching private pool is found, the system uses public resources.
	//
	// - `Target`: The system uses the capacity of a specified private pool to start the instance. If the capacity of the specified private pool is unavailable, the instance fails to start.
	//
	// - `None`: The system does not use the capacity of a private pool.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"match_criteria,omitempty" xml:"match_criteria,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) SetId(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) SetMatchCriteria(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions struct {
	// The list of private pool IDs.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy used when creating instances. Valid values: PrivatePoolFirst: The private pool is used first. PrivatePoolOnly: Only the private pool is used. None: No resource pool policy is used.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) SetPrivatePoolIds(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) SetStrategy(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit struct {
	// The spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The price of a single instance.
	//
	// <props="china">
	//
	// Unit: CNY/hour.
	//
	//
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

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) SetInstanceType(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) SetPriceLimit(v string) *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyStatus struct {
	// The current status of the node pool. This parameter indicates the status of the node pool from different dimensions.
	Conditions []*DescribeClusterNodePoolDetailResponseBodyStatusConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The number of failed nodes.
	//
	// example:
	//
	// 0
	FailedNodes *int64 `json:"failed_nodes,omitempty" xml:"failed_nodes,omitempty"`
	// The number of healthy nodes.
	//
	// example:
	//
	// 3
	HealthyNodes *int64 `json:"healthy_nodes,omitempty" xml:"healthy_nodes,omitempty"`
	// The number of nodes that are being created.
	//
	// example:
	//
	// 0
	InitialNodes *int64 `json:"initial_nodes,omitempty" xml:"initial_nodes,omitempty"`
	// The number of offline nodes.
	//
	// example:
	//
	// 0
	OfflineNodes *int64 `json:"offline_nodes,omitempty" xml:"offline_nodes,omitempty"`
	// The number of nodes that are being removed.
	//
	// example:
	//
	// 0
	RemovingNodes *int64 `json:"removing_nodes,omitempty" xml:"removing_nodes,omitempty"`
	// The number of nodes that are in service.
	//
	// example:
	//
	// 3
	ServingNodes *int64 `json:"serving_nodes,omitempty" xml:"serving_nodes,omitempty"`
	// The state of the node pool. Valid values:
	//
	// - `active`: The node pool is active.
	//
	// - `scaling`: The node pool is being scaled.
	//
	// - `removing`: Nodes are being removed.
	//
	// - `deleting`: The node pool is being deleted.
	//
	// - `updating`: The node pool is being updated.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The total number of nodes in the node pool.
	//
	// example:
	//
	// 3
	TotalNodes *int64 `json:"total_nodes,omitempty" xml:"total_nodes,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetConditions() []*DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	return s.Conditions
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetFailedNodes() *int64 {
	return s.FailedNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetHealthyNodes() *int64 {
	return s.HealthyNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetInitialNodes() *int64 {
	return s.InitialNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetOfflineNodes() *int64 {
	return s.OfflineNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetRemovingNodes() *int64 {
	return s.RemovingNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetServingNodes() *int64 {
	return s.ServingNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) GetTotalNodes() *int64 {
	return s.TotalNodes
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetConditions(v []*DescribeClusterNodePoolDetailResponseBodyStatusConditions) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.Conditions = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetFailedNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.FailedNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetHealthyNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.HealthyNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetInitialNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.InitialNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetOfflineNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.OfflineNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetRemovingNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.RemovingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetServingNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.ServingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetState(v string) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.State = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) SetTotalNodes(v int64) *DescribeClusterNodePoolDetailResponseBodyStatus {
	s.TotalNodes = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatus) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterNodePoolDetailResponseBodyStatusConditions struct {
	// The time of the last status transition.
	//
	// example:
	//
	// 20**-**-30T10:39:00+08:00
	LastTransitionTime *string `json:"last_transition_time,omitempty" xml:"last_transition_time,omitempty"`
	// The detailed information.
	//
	// example:
	//
	// AutoUpgradeDisabled
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The reason.
	//
	// example:
	//
	// UpgradeDisabled
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The status.
	//
	// example:
	//
	// True
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type.
	//
	// example:
	//
	// ImageUpgradeReady
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyStatusConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyStatusConditions) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) GetMessage() *string {
	return s.Message
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) GetReason() *string {
	return s.Reason
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) SetLastTransitionTime(v string) *DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) SetMessage(v string) *DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	s.Message = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) SetReason(v string) *DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	s.Reason = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) SetStatus(v string) *DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	s.Status = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) SetType(v string) *DescribeClusterNodePoolDetailResponseBodyStatusConditions {
	s.Type = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyStatusConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyTeeConfig struct {
	// Indicates whether to enable the confidential computing cluster. Valid values:
	//
	// - `true`: Enables the confidential computing cluster.
	//
	// - `false`: Disables the confidential computing cluster.
	//
	// example:
	//
	// false
	TeeEnable *bool `json:"tee_enable,omitempty" xml:"tee_enable,omitempty"`
}

func (s DescribeClusterNodePoolDetailResponseBodyTeeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyTeeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolDetailResponseBodyTeeConfig) GetTeeEnable() *bool {
	return s.TeeEnable
}

func (s *DescribeClusterNodePoolDetailResponseBodyTeeConfig) SetTeeEnable(v bool) *DescribeClusterNodePoolDetailResponseBodyTeeConfig {
	s.TeeEnable = &v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyTeeConfig) Validate() error {
	return dara.Validate(s)
}
