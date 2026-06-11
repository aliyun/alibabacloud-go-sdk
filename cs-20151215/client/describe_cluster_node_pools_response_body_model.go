// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodepools(v []*DescribeClusterNodePoolsResponseBodyNodepools) *DescribeClusterNodePoolsResponseBody
	GetNodepools() []*DescribeClusterNodePoolsResponseBodyNodepools
}

type DescribeClusterNodePoolsResponseBody struct {
	// The list of node pools.
	Nodepools []*DescribeClusterNodePoolsResponseBodyNodepools `json:"nodepools,omitempty" xml:"nodepools,omitempty" type:"Repeated"`
}

func (s DescribeClusterNodePoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBody) GetNodepools() []*DescribeClusterNodePoolsResponseBodyNodepools {
	return s.Nodepools
}

func (s *DescribeClusterNodePoolsResponseBody) SetNodepools(v []*DescribeClusterNodePoolsResponseBodyNodepools) *DescribeClusterNodePoolsResponseBody {
	s.Nodepools = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBody) Validate() error {
	if s.Nodepools != nil {
		for _, item := range s.Nodepools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterNodePoolsResponseBodyNodepools struct {
	// The intelligent hosting configurations.
	AutoMode *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The auto scaling configuration.
	AutoScaling *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// The information about the Lingjun node group.
	EfloNodeGroup *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// [This parameter is deprecated]
	//
	// The network configurations of the edge node pool. This parameter is valid only for edge node pools.
	InterconnectConfig *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter is valid only for node pools of the `edge` type. Valid values:
	//
	// - `basic`: public network. The nodes in the node pool interact with cloud nodes over the Internet. Applications in the node pool cannot directly access the VPC on the cloud.
	//
	// - `private`: private network. The nodes in the node pool are connected to the cloud over a leased line, a VPN connection, or a CEN instance. This provides higher cloud-to-edge communication quality and enhanced security.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// The cluster-related configurations.
	KubernetesConfig *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool. This parameter takes effect only in ACK Pro clusters.
	Management *DescribeClusterNodePoolsResponseBodyNodepoolsManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// The maximum number of nodes that the edge node pool can contain. The value of this parameter must be greater than or equal to 0. A value of 0 indicates that no limit is imposed on the number of nodes in the node pool, except for the limit on the total number of nodes in the cluster. The value of this parameter is usually greater than 0 for edge node pools. The value is 0 for ess node pools and default edge node pools
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The list of node components.
	NodeComponents []*DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// The node configurations.
	NodeConfig *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The information about the node pool.
	NodepoolInfo *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group for the node pool.
	ScalingGroup *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The status of the node pool.
	Status *DescribeClusterNodePoolsResponseBodyNodepoolsStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// The confidential computing configurations.
	TeeConfig *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepools) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepools) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetAutoMode() *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode {
	return s.AutoMode
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetAutoScaling() *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	return s.AutoScaling
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetEfloNodeGroup() *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup {
	return s.EfloNodeGroup
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetInterconnectConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	return s.InterconnectConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetInterconnectMode() *string {
	return s.InterconnectMode
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetKubernetesConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	return s.KubernetesConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetManagement() *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	return s.Management
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetMaxNodes() *int64 {
	return s.MaxNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetNodeComponents() []*DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents {
	return s.NodeComponents
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetNodeConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig {
	return s.NodeConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetNodepoolInfo() *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	return s.NodepoolInfo
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetScalingGroup() *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	return s.ScalingGroup
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetStatus() *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	return s.Status
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) GetTeeConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig {
	return s.TeeConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetAutoMode(v *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.AutoMode = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetAutoScaling(v *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.AutoScaling = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetEfloNodeGroup(v *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.EfloNodeGroup = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetInterconnectConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.InterconnectConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetInterconnectMode(v string) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.InterconnectMode = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetKubernetesConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.KubernetesConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetManagement(v *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.Management = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetMaxNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.MaxNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetNodeComponents(v []*DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.NodeComponents = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetNodeConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.NodeConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetNodepoolInfo(v *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.NodepoolInfo = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetScalingGroup(v *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.ScalingGroup = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetStatus(v *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.Status = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) SetTeeConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) *DescribeClusterNodePoolsResponseBodyNodepools {
	s.TeeConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepools) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode struct {
	// Indicates whether to enable intelligent hosting.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) SetEnable(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling struct {
	// The maximum bandwidth of the EIP.
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
	// Indicates whether auto scaling is enabled. Valid values:
	//
	// - `true`: Enables auto scaling for the node pool. When cluster resources are insufficient for pod scheduling, ACK automatically scales nodes based on the configured minimum and maximum number of instances. For clusters running Kubernetes 1.24 or later, on-demand node scaling is enabled by default. For clusters running Kubernetes versions earlier than 1.24, node autoscaling is enabled by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: Auto scaling is disabled. ACK adjusts the number of nodes in the node pool to the expected number of nodes.
	//
	// If this parameter is set to false, other parameters in the auto_scaling object do not take effect.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Indicates whether to associate an EIP with the node. Valid values:
	//
	// - `true`: Associates an EIP with the node.
	//
	// - `false`: Does not associate an EIP with the node.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances that can be created in the node pool. This value does not include existing instances.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances that must be kept in the node pool. This value does not include existing instances.
	//
	// example:
	//
	// 2
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The type of auto scaling. This parameter specifies the type of instances that are used for auto scaling. Valid values:
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetEipBandwidth() *int64 {
	return s.EipBandwidth
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetEipInternetChargeType() *string {
	return s.EipInternetChargeType
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetIsBondEip() *bool {
	return s.IsBondEip
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetEipBandwidth(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetEipInternetChargeType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.EipInternetChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetEnable(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetIsBondEip(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.IsBondEip = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetMaxInstances(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.MaxInstances = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetMinInstances(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.MinInstances = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) SetType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling {
	s.Type = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup struct {
	// The ID of the Lingjun cluster.
	//
	// example:
	//
	// i113790071760688002461
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The ID of the Lingjun group.
	//
	// example:
	//
	// i128147721760688002463
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) SetClusterId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) SetGroupId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup {
	s.GroupId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig struct {
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
	// The ID of the Cloud Connect Network (CCN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The region where the CCN instance associated with the enhanced edge node pool resides.
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the Cloud Enterprise Network (CEN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// [This parameter is deprecated]
	//
	// The subscription duration of the enhanced edge node pool. Unit: month.
	//
	// example:
	//
	// 1
	ImprovedPeriod *string `json:"improved_period,omitempty" xml:"improved_period,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GetCcnId() *string {
	return s.CcnId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GetCcnRegionId() *string {
	return s.CcnRegionId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GetCenId() *string {
	return s.CenId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) GetImprovedPeriod() *string {
	return s.ImprovedPeriod
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) SetBandwidth(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) SetCcnId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	s.CcnId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) SetCcnRegionId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	s.CcnRegionId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) SetCenId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	s.CenId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) SetImprovedPeriod(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig {
	s.ImprovedPeriod = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig struct {
	// Indicates whether to install CloudMonitor on the ECS nodes. After CloudMonitor is installed, you can view the monitoring information of the created ECS instances in the CloudMonitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: Installs CloudMonitor on the ECS nodes.
	//
	// - `false`: Does not install CloudMonitor on the ECS nodes
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of the node. The following policies are supported for clusters of Kubernetes 1.12.6 and later:
	//
	// - `static`: allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// - `none`: indicates that the default CPU affinity scheme is enabled.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The node labels.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name.
	//
	// A node name consists of three parts: a prefix, the middle part of an IP address, and a suffix:
	//
	// - The prefix and suffix can contain one or more parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). The node name must start and end with a lowercase letter or a digit.
	//
	// - The IP address segment length specifies the number of digits to be truncated from the end of the node IP address. Valid values: 5 to 12.
	//
	// For example, if the node IP address is 192.1.168.0.55, the prefix is aliyun.com, the IP address segment length is 5, and the suffix is test, the node name is aliyun.com00055test.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-custom data of the node pool. The script is run before the node is initialized. For more information, see [Generate instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvYmluL3NoCmVjaG8gIkhlbGxvIEFD
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. ACK supports the following container runtimes.
	//
	// - containerd: recommended. This runtime is supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: a sandboxed container that provides higher isolation. This runtime is supported by clusters of Kubernetes 1.31 and earlier.
	//
	// - docker: no longer maintained. This runtime is supported by clusters of Kubernetes 1.22 and earlier.
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
	// The node taints. Taints and tolerations work together to prevent pods from being scheduled to unsuitable nodes. For more information, see [Taints and Tolerations](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Indicates whether the scaled-out nodes are unschedulable.
	//
	// - true: The scaled-out nodes are unschedulable.
	//
	// - false: The scaled-out nodes are schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The custom data of the node pool. The script is run after the node is initialized. For more information, see [Generate instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvYmluL3NoCmVjaG8gIkhlbGxvIEFD****
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetCmsEnabled() *bool {
	return s.CmsEnabled
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetLabels() []*Tag {
	return s.Labels
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetTaints() []*Taint {
	return s.Taints
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetUnschedulable() *bool {
	return s.Unschedulable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetCmsEnabled(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.CmsEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetCpuPolicy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.CpuPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetLabels(v []*Tag) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.Labels = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetNodeNameMode(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.NodeNameMode = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetPreUserData(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetRuntime(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.Runtime = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetRuntimeVersion(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetTaints(v []*Taint) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.Taints = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetUnschedulable(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.Unschedulable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) SetUserData(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodepoolsManagement struct {
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Indicates whether to enable auto repair. This parameter takes effect only when enable is set to true.
	//
	// - `true`: Enables auto repair.
	//
	// - `false`: Disables auto repair.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto repair policy for nodes.
	AutoRepairPolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Indicates whether to enable auto update. This parameter takes effect only when enable is set to true.
	//
	// - `true`: Enables auto update.
	//
	// - `false`: Disables auto update.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto update policy.
	AutoUpgradePolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Indicates whether to automatically fix CVEs. This parameter takes effect only when enable is set to true.
	//
	// - `true`: Allows CVEs to be automatically fixed.
	//
	// - `false`: Does not allow CVEs to be automatically fixed.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The policy for automatically fixing CVEs.
	AutoVulFixPolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Indicates whether to enable the managed node pool feature. Valid values:
	//
	// - `true`: Enables the managed node pool feature.
	//
	// - `false`: The managed node pool feature is disabled. If you set this parameter to true, the other parameters take effect.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The auto update configurations. This parameter takes effect only when enable is set to true.
	UpgradeConfig *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagement) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoFaultDiagnosis() *bool {
	return s.AutoFaultDiagnosis
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoRepair() *bool {
	return s.AutoRepair
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoRepairPolicy() *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy {
	return s.AutoRepairPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoUpgradePolicy() *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy {
	return s.AutoUpgradePolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoVulFix() *bool {
	return s.AutoVulFix
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetAutoVulFixPolicy() *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy {
	return s.AutoVulFixPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) GetUpgradeConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig {
	return s.UpgradeConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoFaultDiagnosis(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoFaultDiagnosis = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoRepair(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoRepair = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoRepairPolicy(v *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoRepairPolicy = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoUpgrade(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoUpgrade = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoUpgradePolicy(v *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoUpgradePolicy = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoVulFix(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoVulFix = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetAutoVulFixPolicy(v *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.AutoVulFixPolicy = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetEnable(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.Enable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) SetUpgradeConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) *DescribeClusterNodePoolsResponseBodyNodepoolsManagement {
	s.UpgradeConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagement) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy struct {
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
	// r-xxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// Indicates whether to allow node restart. This parameter takes effect only when auto_repair is set to true.
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Does not allow node restart.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) GetApprovalRequired() *bool {
	return s.ApprovalRequired
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) GetAutoRepairPolicyId() *string {
	return s.AutoRepairPolicyId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) SetApprovalRequired(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) SetAutoRepairPolicyId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy {
	s.AutoRepairPolicyId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) SetRestartNode(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy struct {
	// Indicates whether to allow auto update of the kubelet. This parameter takes effect only when auto_upgrade is set to true. Valid values:
	//
	// - `true`: Allows auto update of the kubelet.
	//
	// - `false`: Does not allow auto update of the kubelet.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) GetAutoUpgradeKubelet() *bool {
	return s.AutoUpgradeKubelet
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) SetAutoUpgradeKubelet(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy {
	s.AutoUpgradeKubelet = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy struct {
	// The packages that should be excluded during CVE fixing.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Indicates whether to allow node restart. This parameter takes effect only when auto_vul_fix is set to true. Valid values:
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Does not allow node restart.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The CVE levels that are allowed to be automatically fixed. Separate multiple CVE levels with commas.
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) GetExcludePackages() *string {
	return s.ExcludePackages
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) SetExcludePackages(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy {
	s.ExcludePackages = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) SetRestartNode(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy {
	s.RestartNode = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) SetVulLevel(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy {
	s.VulLevel = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig struct {
	// Indicates whether to enable auto update. Valid values:
	//
	// - `true`: Enables auto update.
	//
	// - `false`: Disables auto update.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes. Valid values: 1 to 1000
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
	// The number of extra nodes = Percentage of extra nodes × Number of nodes. For example, if you set the percentage of extra nodes to 50% and the number of existing nodes is 6, three extra nodes are created.
	//
	// example:
	//
	// 50
	SurgePercentage *int64 `json:"surge_percentage,omitempty" xml:"surge_percentage,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) GetMaxUnavailable() *int64 {
	return s.MaxUnavailable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) GetSurge() *int64 {
	return s.Surge
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) GetSurgePercentage() *int64 {
	return s.SurgePercentage
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) SetAutoUpgrade(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig {
	s.AutoUpgrade = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) SetMaxUnavailable(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) SetSurge(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig {
	s.Surge = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) SetSurgePercentage(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig {
	s.SurgePercentage = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsManagementUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents struct {
	// The configurations of the node component.
	Config *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) GetConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig {
	return s.Config
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) SetConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents {
	s.Config = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) SetName(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents {
	s.Name = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) SetVersion(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents {
	s.Version = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig struct {
	// The custom configurations of the node component.
	CustomConfig map[string]*string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) SetCustomConfig(v map[string]*string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig struct {
	// The Kubelet parameter configurations.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
	// The node OS configurations.
	NodeOsConfig *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig `json:"node_os_config,omitempty" xml:"node_os_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) GetKubeletConfiguration() *KubeletConfig {
	return s.KubeletConfiguration
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) GetNodeOsConfig() *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig {
	return s.NodeOsConfig
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) SetNodeOsConfig(v *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig {
	s.NodeOsConfig = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig struct {
	// The Hugepage configuration.
	Hugepage *Hugepage `json:"hugepage,omitempty" xml:"hugepage,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) GetHugepage() *Hugepage {
	return s.Hugepage
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) SetHugepage(v *Hugepage) *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig {
	s.Hugepage = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfigNodeOsConfig) Validate() error {
	if s.Hugepage != nil {
		if err := s.Hugepage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo struct {
	// The time when the node pool was created.
	//
	// example:
	//
	// 2025-04-15T16:33:29.362888807+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Indicates whether the node pool is the default node pool. A cluster has only one default node pool. Valid values:
	//
	// - `true`: The node pool is the default node pool.
	//
	// - `false`: The node pool is not the default node pool.
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
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of the node pool. Valid values:
	//
	// - `ess`: a regular node pool. This type of node pool provides managed features and supports automatic scaling.
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
	// 2025-04-15T16:33:32.823+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetCreated(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.Created = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetIsDefault(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.IsDefault = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetName(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetNodepoolId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetRegionId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetResourceGroupId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.Type = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) SetUpdated(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo {
	s.Updated = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup struct {
	// Indicates whether to enable auto-renewal for the nodes. This parameter takes effect only when instance_charge_type is set to PrePaid. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
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
	// 0
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated]
	//
	// Use the security_hardening_os parameter instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// If multi_az_policy is set to `COST_OPTIMIZED`, this parameter specifies whether to allow the system to automatically create pay-as-you-go instances to meet the required number of ECS instances when preemptible instances cannot be created due to reasons such as price and inventory. Valid values:
	//
	// - `true`: Allows the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: Does not allow the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The combination of the configurations of the data disks of the node, such as the disk type and size.
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
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The custom image ID. You can call the `DescribeKubernetesVersionMetadata` operation to query the images supported by the system.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The OS image type.
	//
	// - `AliyunLinux`: Alibaba Cloud Linux 2 image.
	//
	// - `AliyunLinuxSecurity`: Alibaba Cloud Linux 2 UEFI image.
	//
	// - `AliyunLinux3`: Alibaba Cloud Linux 3 image.
	//
	// - `AliyunLinux3Arm64`: Alibaba Cloud Linux 3 ARM image.
	//
	// - `AliyunLinux3Security`: Alibaba Cloud Linux 3 UEFI image.
	//
	// - `CentOS`: CentOS image.
	//
	// - `Windows`: Windows image.
	//
	// - `WindowsCore`: WindowsCore image.
	//
	// - `ContainerOS`: Container-optimized image.
	//
	// - `AliyunLinux3ContainerOptimized`: Alibaba Cloud Linux 3 container-optimized image.
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
	// The instance attribute configurations.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of node instance types. You can select multiple instance types as alternatives. When a node is created, the system starts from the first instance type until the node is created. The instance type that is used to create the node may vary based on the inventory.
	//
	// example:
	//
	// ecs.n4.large
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method of the public IP address of the node.
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth of the public IP address of the node. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair. You must set one of key_pair and login_password.
	//
	// You can set only `key_pair` for managed node pools.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Indicates whether to log on to the created ECS instance as a non-root user.
	//
	// - true: Log on as a non-root user (ecs-user).
	//
	// - false: Log on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. You must set one of key_pair and login_password. The password must be 8 to 30 characters in length, and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// For security reasons, the password is encrypted.
	//
	// example:
	//
	// ******
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for the ECS instances in the multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: The system scales ECS instances based on the vSwitches that you specify (VSwitchIds.N). If an ECS instance cannot be created in the zone where the vSwitch with a higher priority resides, the system uses the vSwitch with the next priority to create the ECS instance.
	//
	// - `COST_OPTIMIZED`: The system creates ECS instances at the lowest vCPU price. If multiple instance types are specified in the scaling configuration and the preemptible instance feature is enabled, the system preferentially creates preemptible instances. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when preemptible instances cannot be created due to reasons such as insufficient inventory.
	//
	//   > `COST_OPTIMIZED` takes effect only if the scaling configuration uses multiple instance types or spot instances.
	//
	// - `BALANCE`: The system evenly distributes ECS instances across the specified zones of the scaling group. If the distribution of ECS instances becomes unbalanced due to insufficient inventory, you can call the API `RebalanceInstances` operation to balance the resource distribution. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) .
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be contained in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than this value, the system preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that are created after the minimum number of pay-as-you-go instances (`on_demand_base_capacity`) is met. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of the nodes. This parameter is required and takes effect only when instance_charge_type is set to PrePaid.
	//
	// - If period_unit is set to Week, the valid values of period are 1, 2, 3, and 4.
	//
	// - If period_unit is set to Month, the valid values of period are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the nodes. This parameter is required when instance_charge_type is set to PrePaid.
	//
	// - `Month`: The billing cycle is measured in months.
	//
	// - `Week`: The billing cycle is measured in weeks.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// [This parameter is deprecated]
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
	// example:
	//
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The private pool options.
	PrivatePoolOptions *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// This parameter is deprecated. Use ram_role_name instead.
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
	// If you specify a list of RDS instances, the ECS instances of the cluster nodes are automatically added to the RDS instance whitelist.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy that are used when an instance is created.
	ResourcePoolOptions *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling group ID.
	//
	// example:
	//
	// asg-2ze8n5qw4atggut8****
	ScalingGroupId *string `json:"scaling_group_id,omitempty" xml:"scaling_group_id,omitempty"`
	// The scaling mode of the scaling group. Valid values:
	//
	// - `release`: standard mode. The system creates and releases ECS instances to scale resources based on the resource usage.
	//
	// - `recycle`: fast mode. The system creates, stops, and starts ECS instances to scale resources. This improves the scaling speed. When an instance is stopped, its computing resources are not billed, but its storage resources are. This does not apply to instance types with local disks.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// [This parameter is deprecated]
	//
	// The ID of the security group for the node pool. If the node pool is associated with multiple security groups, this parameter returns the first security group ID in the `security_group_ids` list.
	//
	// example:
	//
	// sg-2ze1iuk12m2sb4c4****
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
	// Indicates whether to enable MLPS 2.0 security hardening. You can enable MLPS 2.0 security hardening for nodes only when you select Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 as the OS image. Alibaba Cloud provides classified protection compliance baseline check standards and scanning programs for MLPS 2.0 Level 3-compliant versions of Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of available instance types. The scaling group creates preemptible instances of multiple instance types that are available at the lowest cost. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Indicates whether to enable the feature of supplementing preemptible instances. If this feature is enabled, the scaling group attempts to create a new instance to replace a preemptible instance when the scaling group receives a system message that the preemptible instance is to be reclaimed. Valid values:
	//
	// - `true`: Enables the feature of supplementing preemptible instances.
	//
	// - `false`: Disables the feature of supplementing preemptible instances.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The configurations of the price range for the spot instances.
	SpotPriceLimit []*DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The bidding policy for the spot instances. Valid values:
	//
	// - NoSpot: The instances are not spot instances.
	//
	// - SpotWithPriceLimit: The maximum bid price is specified for the spot instances.
	//
	// - SpotAsPriceGo: The system automatically places bids based on the market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Indicates whether to enable bursting for the system disk of the node. Valid values:
	//
	// - true: enables bursting. After bursting is enabled, the performance of the cloud disk is temporarily improved to handle sudden data read and write pressure when the business is unstable. The performance of the cloud disk is restored after the business becomes stable.
	//
	// - false: disables bursting.
	//
	// This parameter can be set only when system_disk_category is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The multi-disk type of the system disk. When a disk of a higher priority is unavailable, the system automatically uses a disk of a lower priority to create the system disk.
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
	// Indicates whether to encrypt the system disk. Valid values:
	//
	// - true: Encrypts the system disk.
	//
	// - false: Does not encrypt the system disk.
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
	// The performance level of the system disk of the node. This parameter is valid only for ESSDs. The performance level of the disk is related to the disk size. For more information, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: The disk has medium concurrent I/O performance and stable read and write latency.
	//
	// - PL1: The disk has medium concurrent I/O performance and stable read and write latency.
	//
	// - PL2: The disk has high concurrent I/O performance and stable read and write latency.
	//
	// - PL3: The disk has ultra-high concurrent I/O performance and ultra-stable read and write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The pre-configured read and write IOPS of the system disk of the node. This parameter is configured when the disk type is cloud_auto.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk of the node. Unit: GiB.
	//
	// Valid values: 20 to 2048.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The snapshot policy for the system disk
	//
	// example:
	//
	// sp-0jl6xnmme8v7o935****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// The tags of the ECS instances.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of vSwitch IDs.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetCisEnabled() *bool {
	return s.CisEnabled
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetDesiredSize() *int64 {
	return s.DesiredSize
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetDiskInit() []*DiskInit {
	return s.DiskInit
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetInstancePatterns() []*InstancePatterns {
	return s.InstancePatterns
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetKeyPair() *string {
	return s.KeyPair
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetOnDemandBaseCapacity() *int64 {
	return s.OnDemandBaseCapacity
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int64 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetPrivatePoolOptions() *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetRamPolicy() *string {
	return s.RamPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetResourcePoolOptions() *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSecurityHardeningOs() *bool {
	return s.SecurityHardeningOs
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSocEnabled() *bool {
	return s.SocEnabled
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSpotInstancePools() *int64 {
	return s.SpotInstancePools
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSpotPriceLimit() []*DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit {
	return s.SpotPriceLimit
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskKmsKeyId() *string {
	return s.SystemDiskKmsKeyId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetTags() []*Tag {
	return s.Tags
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetAutoRenew(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetAutoRenewPeriod(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetCisEnabled(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.CisEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetCompensateWithOnDemand(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetDataDisks(v []*DataDisk) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.DataDisks = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetDeploymentsetId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.DeploymentsetId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetDesiredSize(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetDiskInit(v []*DiskInit) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.DiskInit = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetImageId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetImageType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.ImageType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetInstanceChargeType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetInstancePatterns(v []*InstancePatterns) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.InstancePatterns = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetInstanceTypes(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetInternetChargeType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetInternetMaxBandwidthOut(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetKeyPair(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.KeyPair = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetLoginAsNonRoot(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.LoginAsNonRoot = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetLoginPassword(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.LoginPassword = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetMultiAzPolicy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.MultiAzPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetOnDemandBaseCapacity(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetPeriod(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.Period = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetPeriodUnit(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetPlatform(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.Platform = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetPrivatePoolOptions(v *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetRamPolicy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.RamPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetRamRoleName(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.RamRoleName = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetRdsInstances(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.RdsInstances = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetResourcePoolOptions(v *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.ResourcePoolOptions = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetScalingGroupId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetScalingPolicy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSecurityGroupId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSecurityGroupIds(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSecurityHardeningOs(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SecurityHardeningOs = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSocEnabled(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SocEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSpotInstancePools(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSpotInstanceRemedy(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSpotPriceLimit(v []*DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SpotPriceLimit = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSpotStrategy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskBurstingEnabled(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskCategories(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskCategory(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskEncryptAlgorithm(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskEncrypted(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskKmsKeyId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskKmsKeyId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskPerformanceLevel(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskProvisionedIops(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskSize(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetSystemDiskSnapshotPolicyId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetTags(v []*Tag) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.Tags = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) SetVswitchIds(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup) Validate() error {
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

type DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions struct {
	// The private pool ID. This is the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. This parameter specifies the capacity of the private pool that is used to start an instance. The capacity of a private pool is generated after an elasticity assurance or a capacity reservation takes effect. You can select a capacity option when you start an instance. Valid values:
	//
	// - `Open`: The system automatically matches the capacity of an open private pool. If no matching private pool is found, the resources in the public pool are used.
	//
	// - `Target`: The system uses the capacity of the specified private pool to start the instance. If the capacity of the private pool is unavailable, the instance fails to be started.
	//
	// - `None`: The instance is started without using the capacity of a private pool.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"match_criteria,omitempty" xml:"match_criteria,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) SetId(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) SetMatchCriteria(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions struct {
	// The list of private pool IDs.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy that is used when an instance is created. Valid values:
	//
	// PrivatePoolFirst: The private pool is used first.
	//
	// PrivatePoolOnly: Only the private pool is used.
	//
	// None: No resource pool policy is used.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) SetPrivatePoolIds(v []*string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) SetStrategy(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit struct {
	// The instance type of the spot instance.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The price range for a single instance.
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) SetInstanceType(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) SetPriceLimit(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsStatus struct {
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
	// The number of running nodes.
	//
	// example:
	//
	// 3
	ServingNodes *int64 `json:"serving_nodes,omitempty" xml:"serving_nodes,omitempty"`
	// The status of the node pool. Valid values:
	//
	// - `active`: The node pool is active.
	//
	// - `scaling`: The node pool is being scaled.
	//
	// - `removing`: Nodes are being removed from the node pool.
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetFailedNodes() *int64 {
	return s.FailedNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetHealthyNodes() *int64 {
	return s.HealthyNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetInitialNodes() *int64 {
	return s.InitialNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetOfflineNodes() *int64 {
	return s.OfflineNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetRemovingNodes() *int64 {
	return s.RemovingNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetServingNodes() *int64 {
	return s.ServingNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) GetTotalNodes() *int64 {
	return s.TotalNodes
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetFailedNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.FailedNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetHealthyNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.HealthyNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetInitialNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.InitialNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetOfflineNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.OfflineNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetRemovingNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.RemovingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetServingNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.ServingNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetState(v string) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.State = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) SetTotalNodes(v int64) *DescribeClusterNodePoolsResponseBodyNodepoolsStatus {
	s.TotalNodes = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig struct {
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

func (s DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) GetTeeEnable() *bool {
	return s.TeeEnable
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) SetTeeEnable(v bool) *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig {
	s.TeeEnable = &v
	return s
}

func (s *DescribeClusterNodePoolsResponseBodyNodepoolsTeeConfig) Validate() error {
	return dara.Validate(s)
}
