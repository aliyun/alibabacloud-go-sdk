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
	// List of node pool instances.
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
	// Intelligent management configuration.
	AutoMode *DescribeClusterNodePoolsResponseBodyNodepoolsAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// Automatic scaling configuration.
	AutoScaling *DescribeClusterNodePoolsResponseBodyNodepoolsAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Lingjun node group information.
	EfloNodeGroup *DescribeClusterNodePoolsResponseBodyNodepoolsEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// [This field is deprecated.]
	//
	// Network configuration for edge node pools. This parameter applies only to edge-type node pools.
	InterconnectConfig *DescribeClusterNodePoolsResponseBodyNodepoolsInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// Network type for edge node pools. This parameter applies only to node pools where `type` is `edge`. Valid values:
	//
	// - `basic`: Public network. Nodes in the node pool interact with cloud nodes over the public network and cannot directly access the VPC intranet.
	//
	// - `private`: Private network. Nodes in the node pool connect to the cloud through leased lines, VPN, or CEN, providing higher cloud-edge communication quality and better security.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Cluster-related configuration.
	KubernetesConfig *DescribeClusterNodePoolsResponseBodyNodepoolsKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// Managed node pool configuration. This parameter takes effect only for professional managed clusters.
	Management *DescribeClusterNodePoolsResponseBodyNodepoolsManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// Maximum number of nodes allowed in an edge node pool. This parameter is greater than or equal to 0. A value of 0 means no additional limit (only limited by the overall cluster capacity). This parameter is usually greater than 0 for edge node pools, and 0 for ess-type node pools and default edge-type node pools.
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// List of node components.
	NodeComponents []*DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// Node configuration.
	NodeConfig *DescribeClusterNodePoolsResponseBodyNodepoolsNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// Node pool information.
	NodepoolInfo *DescribeClusterNodePoolsResponseBodyNodepoolsNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// Scaling group configuration for the node pool.
	ScalingGroup *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// Node pool status.
	Status *DescribeClusterNodePoolsResponseBodyNodepoolsStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// Confidential computing configuration.
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
	// Indicates whether intelligent management is enabled.
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
	// Peak bandwidth of the EIP.
	//
	// Valid values: [1,100]. Unit: Mbps.
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// EIP billing method. Valid values:
	//
	// - `PayByBandwidth`: Pay-by-bandwidth.
	//
	// - `PayByTraffic`: Pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Indicates whether automatic scaling is enabled. Valid values:
	//
	// - `true`: Enables automatic scaling for the node pool. When the cluster capacity cannot meet application pod scheduling requirements, ACK automatically scales node resources based on the configured minimum and maximum instance counts. Clusters of version 1.24 or later enable instant elasticity by default. Clusters earlier than version 1.24 enable node autoscaling by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: Disables automatic scaling. ACK adjusts the number of nodes in the node pool to match the desired node count and maintains this count.
	//
	// If this parameter is set to false, other parameters in `auto_scaling` do not take effect.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Indicates whether an EIP is bound. Valid values:
	//
	// - `true`: Binds an EIP.
	//
	// - `false`: Does not bind an EIP.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// Maximum number of scalable instances in the node pool, excluding existing instances.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// Minimum number of scalable instances in the node pool, excluding existing instances.
	//
	// example:
	//
	// 2
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// Type of automatic scaling, categorized by instance type. Valid values:
	//
	// - `cpu`: Standard instance.
	//
	// - `gpu`: GPU instance.
	//
	// - `gpushare`: Shared GPU instance.
	//
	// - `spot`: Spot instance.
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
	// Lingjun cluster ID.
	//
	// example:
	//
	// i113790071760688002461
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Lingjun group ID.
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
	// [This field is deprecated.]
	//
	// Network bandwidth for enhanced edge node pools. Unit: Mbps.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// [This field is deprecated.]
	//
	// CCN instance ID bound to enhanced edge node pools.
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// [This field is deprecated.]
	//
	// Region of the CCN instance bound to enhanced edge node pools.
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// [This field is deprecated.]
	//
	// CEN instance ID bound to enhanced edge node pools.
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// [This field is deprecated.]
	//
	// Subscription duration for enhanced edge node pools. Unit: months.
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
	// Indicates whether Cloud Monitor is installed on ECS nodes. After installation, you can view monitoring information for created ECS instances in the Cloud Monitor console. We recommend enabling this feature. Valid values:
	//
	// - `true`: Installs Cloud Monitor on ECS nodes.
	//
	// - `false`: Does not install Cloud Monitor on ECS nodes.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// Node CPU management policy. Clusters of version 1.12.6 or later support the following policies:
	//
	// - `static`: Enhances CPU affinity and exclusivity for pods with specific resource characteristics on the node.
	//
	// - `none`: Uses the default CPU affinity scheme.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Node labels.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// Custom node name.
	//
	// A node name consists of three parts: prefix + IP address substring + suffix:
	//
	// - Both prefix and suffix can consist of one or more parts separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). The node name must start and end with a lowercase letter or digit.
	//
	// - The IP address substring length specifies the number of trailing digits to extract from the node IP address. Valid values: 5 to 12.
	//
	// Example: If the node IP address is 192.168.0.55, the prefix is aliyun.com, the IP address substring length is 5, and the suffix is test, the node name is aliyun.com00055test.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// Pre-custom data for the node pool. This is a script that runs before node initialization. For more information, see [Generate instance custom data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvYmluL3NoCmVjaG8gIkhlbGxvIEFD
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// Container runtime name. ACK supports the following container runtimes:
	//
	// - containerd: Recommended. Supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: Sandboxed container that provides higher isolation. Supported by clusters of version 1.31 or earlier.
	//
	// - docker: No longer maintained. Supported by clusters of version 1.22 or earlier.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// Container runtime version.
	//
	// example:
	//
	// 1.6.38
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// Node taint information. Taints and tolerations work together to prevent pods from being scheduled onto unsuitable nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Indicates whether newly scaled nodes are unschedulable.
	//
	// - true: Unschedulable.
	//
	// - false: Schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// Custom data for the node pool. This is a script that runs after node initialization. For more information, see [Generate instance custom data](https://help.aliyun.com/document_detail/49121.html).
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
	// Automatic repair. This parameter takes effect only when `enable=true`.
	//
	// - `true`: Enables automatic repair.
	//
	// - `false`: Disables automatic repair.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// Automatic node repair policy.
	AutoRepairPolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Indicates whether node auto-upgrade is enabled. This parameter takes effect only when `enable=true`.
	//
	// - `true`: Enables auto-upgrade.
	//
	// - `false`: Disables auto-upgrade.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// Auto-upgrade policy.
	AutoUpgradePolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Indicates whether CVEs are automatically fixed. This parameter takes effect only when `enable=true`.
	//
	// - `true`: Enables automatic CVE fixing.
	//
	// - `false`: Disables automatic CVE fixing.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// CVE automatic fix policy.
	AutoVulFixPolicy *DescribeClusterNodePoolsResponseBodyNodepoolsManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Indicates whether the managed node pool feature is enabled. Valid values:
	//
	// - `true`: Enables the managed node pool.
	//
	// - `false`: Disables the managed node pool. Other related configurations take effect only when `enable=true`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Auto-upgrade configuration. This parameter takes effect only when `enable=true`.
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
	// ID of the automatic repair policy.
	//
	// example:
	//
	// r-xxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// Indicates whether node restart is allowed. This parameter takes effect only when `auto_repair=true`.
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Disallows node restart.
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
	// Indicates whether kubelet auto-upgrade is allowed. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// - `true`: Allows kubelet auto-upgrade.
	//
	// - `false`: Disallows kubelet auto-upgrade.
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
	// Packages excluded from vulnerability fixes.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Indicates whether node restart is allowed. This parameter takes effect only when `auto_vul_fix=true`. Valid values:
	//
	// - `true`: Allows node restart.
	//
	// - `false`: Disallows node restart.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// Vulnerability levels that can be automatically fixed, separated by commas.
	//
	// - `asap`: High
	//
	// - `later`: Medium
	//
	// - `nntf`: Low
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
	// Indicates whether auto-upgrade is enabled. Valid values:
	//
	// - `true`: Enables auto-upgrade.
	//
	// - `false`: Disables auto-upgrade.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// Maximum number of unavailable nodes. Valid values: [1,1000].
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// Number of extra nodes. Specify either this parameter or `surge_percentage`.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// Percentage of extra nodes. Specify either this parameter or `surge`.
	//
	// Number of extra nodes = surge percentage × number of nodes. For example, if you set the surge percentage to 50% and the current number of nodes is 6, the number of extra nodes is 50% × 6 = 3.
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
	// Node component configuration.
	Config *DescribeClusterNodePoolsResponseBodyNodepoolsNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Node component name.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Node component version.
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
	// Custom configuration for node components.
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
	// Kubelet parameter configuration.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
	// Node operating system configuration.
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
	// Hugepage configuration.
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
	// Time when the node pool was created.
	//
	// example:
	//
	// 2025-04-15T16:33:29.362888807+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Indicates whether the node pool is the default node pool. A cluster usually has only one default node pool. Valid values:
	//
	// - `true`: Default node pool.
	//
	// - `false`: Non-default node pool.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
	// Node pool name.
	//
	// example:
	//
	// default-nodepool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Node pool ID.
	//
	// example:
	//
	// np615c0e0966124216a0412e10afe0****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// Node pool type. Valid values:
	//
	// - `ess`: Standard node pool (supports managed features and automatic elastic scaling).
	//
	// - `edge`: Edge node pool.
	//
	// - `lingjun`: Lingjun node pool.
	//
	// example:
	//
	// ess
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Time when the node pool was last updated.
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
	// Indicates whether auto-renewal is enabled for nodes. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// Duration of each auto-renewal cycle. Valid values:
	//
	// - When PeriodUnit=Week: 1, 2, 3.
	//
	// - When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, 60.
	//
	// example:
	//
	// 0
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [This field is deprecated.]
	//
	// Use the security_hardening_os parameter instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// When `multi_az_policy` is set to `COST_OPTIMIZED`, indicates whether to automatically attempt creating pay-as-you-go instances if sufficient spot instances cannot be created due to price or inventory issues. Valid values:
	//
	// - `true`: Allows automatic attempts to create pay-as-you-go instances to meet the required ECS instance count.
	//
	// - `false`: Disallows automatic attempts to create pay-as-you-go instances to meet the required ECS instance count.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// Combination of data disk types, sizes, and other configurations for nodes.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// Deployment set ID.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// Desired number of nodes in the node pool.
	//
	// example:
	//
	// 2
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// Block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// Custom image ID. You can query supported images using `DescribeKubernetesVersionMetadata`.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Operating system image type.
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
	// Billing method for nodes in the node pool. Valid values:
	//
	// - `PrePaid`: Subscription.
	//
	// - `PostPaid`: Pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// Instance attribute configuration.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// List of node instance types. You can select multiple instance types as alternatives. When creating a node, the system attempts to purchase instances starting from the first type until successful. The actual purchased instance type may vary due to inventory availability.
	//
	// example:
	//
	// ecs.n4.large
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// Billing method for public network bandwidth of nodes.
	//
	// - PayByBandwidth: Pay-by-bandwidth.
	//
	// - PayByTraffic: Pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// Maximum outbound public bandwidth for nodes. Unit: Mbps. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// Key pair name. Specify either this parameter or `login_password`.
	//
	// For managed node pools, only `key_pair` is supported.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Indicates whether to log on to the ECS instance as a non-root user.
	//
	// - true: Log on as the non-root user (ecs-user).
	//
	// - false: Log on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// SSH logon password. Specify either this parameter or `key_pair`. The password must be 8 to 30 characters long and contain at least three of the following: uppercase letters, lowercase letters, digits, and special characters.
	//
	// For security reasons, the returned password is encrypted.
	//
	// example:
	//
	// ******
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// ECS instance scaling policy for multi-zone scaling groups. Valid values:
	//
	// - `PRIORITY`: Scales based on the virtual switches (VSwitchIds.N) you define. If ECS instances cannot be created in the zone of a higher-priority virtual switch, the system automatically uses the next priority virtual switch.
	//
	// - `COST_OPTIMIZED`: Attempts to create instances from lowest to highest vCPU price. When multiple instance types or spot billing are configured, spot instances are prioritized. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically attempt pay-as-you-go instance creation if spot instances cannot be created due to inventory or other reasons.
	//
	//   > `COST_OPTIMIZED` takes effect only when multiple instance types are configured or spot instances are selected.
	//
	// - `BALANCE`: Distributes ECS instances evenly across the specified zones. If zones become unbalanced due to inventory shortages, you can use the `RebalanceInstances` API to rebalance resources. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html).
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// Minimum number of pay-as-you-go instances required in the scaling group. Valid values: [0,1000]. If the number of pay-as-you-go instances is less than this value, pay-as-you-go instances are prioritized for creation.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// Percentage of pay-as-you-go instances among instances exceeding the minimum pay-as-you-go instance count (`on_demand_base_capacity`). Valid values: [0,100].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// Subscription duration for nodes. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// - When `period_unit=Week`, valid values for `period` are {1, 2, 3, 4}.
	//
	// - When `period_unit=Month`, valid values for `period` are {1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60}.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// Billing cycle for nodes. Specify this parameter when `instance_charge_type` is set to `PrePaid`.
	//
	// - `Month`: Billed monthly.
	//
	// - `Week`: Billed weekly.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// [This field is deprecated.]
	//
	// Operating system distribution. Valid values:
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
	// Private pool options.
	PrivatePoolOptions *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// This field is deprecated. Use ram_role_name instead.
	//
	// example:
	//
	// KubernetesWorkerRole-021dc54f-929b-437a-8ae0-34c24d3e****
	RamPolicy *string `json:"ram_policy,omitempty" xml:"ram_policy,omitempty"`
	// Worker RAM role name.
	//
	// example:
	//
	// KubernetesWorkerRole-4a4fa089-80c1-48a5-b3c6-9349311f****
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// If RDS instances are specified, ECS nodes in the cluster are automatically added to the RDS access whitelist.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// Resource pool and strategy used when creating instances.
	ResourcePoolOptions *DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// Scaling group ID.
	//
	// example:
	//
	// asg-2ze8n5qw4atggut8****
	ScalingGroupId *string `json:"scaling_group_id,omitempty" xml:"scaling_group_id,omitempty"`
	// Scaling group mode. Valid values:
	//
	// - `release`: Standard mode. Scales by creating or releasing ECS instances based on resource usage.
	//
	// - `recycle`: Fast mode. Scales by creating, stopping, or starting instances to speed up subsequent scaling operations. (Stopped instances incur no compute charges, only storage fees, except for local-disk instance types.)
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// [This field is deprecated.]
	//
	// Security group ID for the node pool. When multiple security groups are bound to the node pool, this value is the first value in `security_group_ids`.
	//
	// example:
	//
	// sg-2ze1iuk12m2sb4c4****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// List of security group IDs for the node pool.
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
	// Indicates whether classified protection compliance hardening is enabled. You can enable this feature for nodes only when Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 is selected as the OS image. Alibaba Cloud provides baseline check standards and scanning programs for MLPS 2.0 Level-3 compliance for Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 images.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// Number of available instance types. The scaling group creates spot instances evenly across the lowest-cost instance types. Valid values: [1,10].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Indicates whether to replenish spot instances. When enabled, if the system receives a notification that a spot instance will be reclaimed, the scaling group attempts to create a new instance to replace the instance to be reclaimed. Valid values:
	//
	// - `true`: Enables replenishment.
	//
	// - `false`: Disables replenishment.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// Market price range configuration for spot instances.
	SpotPriceLimit []*DescribeClusterNodePoolsResponseBodyNodepoolsScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// Spot instance type. Valid values:
	//
	// - NoSpot: Regular instance.
	//
	// - SpotWithPriceLimit: Sets a maximum price for spot instances.
	//
	// - SpotAsPriceGo: Uses the current market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Indicates whether burst performance is enabled for the system disk. Valid values:
	//
	// - true: Enabled. When facing sudden read/write pressure from fluctuating workloads, the disk temporarily boosts performance based on actual workload conditions until the workload stabilizes.
	//
	// - false: Disabled.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// Multiple system disk types. If the system cannot use a higher-priority disk type, it automatically tries the next priority disk type to create the system disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// System disk type for nodes. Valid values:
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: Enterprise SSD.
	//
	// - `cloud_auto`: ESSD AutoPL disk.
	//
	// - `cloud_essd_entry`: ESSD Entry disk.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// Encryption algorithm used for the system disk. Valid values: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Indicates whether the system disk is encrypted. Valid values:
	//
	// - true: Encrypted.
	//
	// - false: Not encrypted.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// KMS key ID used for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// Disk performance level for the system disk. This parameter applies only to ESSD disks. Disk performance levels depend on disk size. For more information, see [ESSD disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: Moderate I/O performance with stable read/write latency.
	//
	// - PL1: Moderate I/O performance with stable read/write latency.
	//
	// - PL2: High I/O performance with stable read/write latency.
	//
	// - PL3: Extremely high I/O performance with highly stable read/write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// Provisioned read/write IOPS for the system disk. Configure this parameter when the disk type is cloud_auto.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// System disk size for nodes. Unit: GiB.
	//
	// Valid values: [20,2048].
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// System disk snapshot policy.
	//
	// example:
	//
	// sp-0jl6xnmme8v7o935****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// ECS instance tags.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// List of virtual switch IDs.
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
	// Private pool ID. This is the ID of an elasticity assurance service or capacity reservation service.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Private pool type. Specifies how instance startup uses private pool capacity. After an elasticity assurance service or capacity reservation service takes effect, it generates private pool capacity for instance startup. Valid values:
	//
	// - `Open`: Open mode. Automatically matches open-type private pool capacity. If no matching private pool capacity is available, public pool resources are used.
	//
	// - `Target`: Target mode. Uses the specified private pool capacity to start instances. If the private pool capacity is unavailable, instance startup fails.
	//
	// - `None`: None mode. Instance startup does not use private pool capacity.
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
	// List of private pool IDs.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// Resource pool strategy used when creating instances. Valid values:
	//
	// PrivatePoolFirst: Private pool first.
	//
	// PrivatePoolOnly: Private pool only.
	//
	// None: Do not use resource pool strategy.
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
	// Spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// Market price range per instance.
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
	// Number of failed instances.
	//
	// example:
	//
	// 0
	FailedNodes *int64 `json:"failed_nodes,omitempty" xml:"failed_nodes,omitempty"`
	// Number of healthy instances.
	//
	// example:
	//
	// 3
	HealthyNodes *int64 `json:"healthy_nodes,omitempty" xml:"healthy_nodes,omitempty"`
	// Number of nodes being created.
	//
	// example:
	//
	// 0
	InitialNodes *int64 `json:"initial_nodes,omitempty" xml:"initial_nodes,omitempty"`
	// Number of offline nodes.
	//
	// example:
	//
	// 0
	OfflineNodes *int64 `json:"offline_nodes,omitempty" xml:"offline_nodes,omitempty"`
	// Number of nodes being removed.
	//
	// example:
	//
	// 0
	RemovingNodes *int64 `json:"removing_nodes,omitempty" xml:"removing_nodes,omitempty"`
	// Number of active nodes.
	//
	// example:
	//
	// 3
	ServingNodes *int64 `json:"serving_nodes,omitempty" xml:"serving_nodes,omitempty"`
	// Node pool state. Valid values:
	//
	// - `active`: Activated.
	//
	// - `scaling`: Scaling.
	//
	// - `removing`: Removing nodes.
	//
	// - `deleting`: Deleting.
	//
	// - `updating`: Updating.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Total number of nodes in the node pool.
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
	// Indicates whether confidential computing is enabled. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
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
