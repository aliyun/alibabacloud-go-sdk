// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodepool interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScaling(v *NodepoolAutoScaling) *Nodepool
	GetAutoScaling() *NodepoolAutoScaling
	SetCount(v int64) *Nodepool
	GetCount() *int64
	SetInterconnectConfig(v *NodepoolInterconnectConfig) *Nodepool
	GetInterconnectConfig() *NodepoolInterconnectConfig
	SetInterconnectMode(v string) *Nodepool
	GetInterconnectMode() *string
	SetKubernetesConfig(v *NodepoolKubernetesConfig) *Nodepool
	GetKubernetesConfig() *NodepoolKubernetesConfig
	SetManagement(v *NodepoolManagement) *Nodepool
	GetManagement() *NodepoolManagement
	SetMaxNodes(v int64) *Nodepool
	GetMaxNodes() *int64
	SetNodeComponents(v []*NodepoolNodeComponents) *Nodepool
	GetNodeComponents() []*NodepoolNodeComponents
	SetNodeConfig(v *NodepoolNodeConfig) *Nodepool
	GetNodeConfig() *NodepoolNodeConfig
	SetNodepoolInfo(v *NodepoolNodepoolInfo) *Nodepool
	GetNodepoolInfo() *NodepoolNodepoolInfo
	SetScalingGroup(v *NodepoolScalingGroup) *Nodepool
	GetScalingGroup() *NodepoolScalingGroup
	SetTeeConfig(v *NodepoolTeeConfig) *Nodepool
	GetTeeConfig() *NodepoolTeeConfig
}

type Nodepool struct {
	// The auto scaling configurations for the node pool.
	AutoScaling *NodepoolAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// [This parameter is deprecated. Use desired_size instead.]
	//
	// The number of nodes in the node pool.
	//
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated.]
	//
	// The configurations of the edge node pool.
	InterconnectConfig *NodepoolInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network mode of the edge node pool. This parameter is valid only for edge node pools. Valid values:
	//
	// - `basic`: basic mode.
	//
	// - `private`: dedicated mode. This mode is supported in clusters of Kubernetes 1.22 and later.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// The Kubernetes configurations for the nodes.
	KubernetesConfig *NodepoolKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool.
	Management *NodepoolManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// 边缘节点池允许容纳的最大节点数量，该参数大于等于 0。0 表示无额外限制（仅受限于集群整体可以容纳的节点数，节点池本身无额外限制）。边缘节点池该参数值往往大于 0；ess 类型节点池和默认的 edge 类型节点池该参数值为 0。
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// 节点组件列表。
	NodeComponents []*NodepoolNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// 节点配置。
	NodeConfig *NodepoolNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The configurations of the node pool.
	NodepoolInfo *NodepoolNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group.
	ScalingGroup *NodepoolScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The configurations of the confidential computing node pool.
	TeeConfig *NodepoolTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
}

func (s Nodepool) String() string {
	return dara.Prettify(s)
}

func (s Nodepool) GoString() string {
	return s.String()
}

func (s *Nodepool) GetAutoScaling() *NodepoolAutoScaling {
	return s.AutoScaling
}

func (s *Nodepool) GetCount() *int64 {
	return s.Count
}

func (s *Nodepool) GetInterconnectConfig() *NodepoolInterconnectConfig {
	return s.InterconnectConfig
}

func (s *Nodepool) GetInterconnectMode() *string {
	return s.InterconnectMode
}

func (s *Nodepool) GetKubernetesConfig() *NodepoolKubernetesConfig {
	return s.KubernetesConfig
}

func (s *Nodepool) GetManagement() *NodepoolManagement {
	return s.Management
}

func (s *Nodepool) GetMaxNodes() *int64 {
	return s.MaxNodes
}

func (s *Nodepool) GetNodeComponents() []*NodepoolNodeComponents {
	return s.NodeComponents
}

func (s *Nodepool) GetNodeConfig() *NodepoolNodeConfig {
	return s.NodeConfig
}

func (s *Nodepool) GetNodepoolInfo() *NodepoolNodepoolInfo {
	return s.NodepoolInfo
}

func (s *Nodepool) GetScalingGroup() *NodepoolScalingGroup {
	return s.ScalingGroup
}

func (s *Nodepool) GetTeeConfig() *NodepoolTeeConfig {
	return s.TeeConfig
}

func (s *Nodepool) SetAutoScaling(v *NodepoolAutoScaling) *Nodepool {
	s.AutoScaling = v
	return s
}

func (s *Nodepool) SetCount(v int64) *Nodepool {
	s.Count = &v
	return s
}

func (s *Nodepool) SetInterconnectConfig(v *NodepoolInterconnectConfig) *Nodepool {
	s.InterconnectConfig = v
	return s
}

func (s *Nodepool) SetInterconnectMode(v string) *Nodepool {
	s.InterconnectMode = &v
	return s
}

func (s *Nodepool) SetKubernetesConfig(v *NodepoolKubernetesConfig) *Nodepool {
	s.KubernetesConfig = v
	return s
}

func (s *Nodepool) SetManagement(v *NodepoolManagement) *Nodepool {
	s.Management = v
	return s
}

func (s *Nodepool) SetMaxNodes(v int64) *Nodepool {
	s.MaxNodes = &v
	return s
}

func (s *Nodepool) SetNodeComponents(v []*NodepoolNodeComponents) *Nodepool {
	s.NodeComponents = v
	return s
}

func (s *Nodepool) SetNodeConfig(v *NodepoolNodeConfig) *Nodepool {
	s.NodeConfig = v
	return s
}

func (s *Nodepool) SetNodepoolInfo(v *NodepoolNodepoolInfo) *Nodepool {
	s.NodepoolInfo = v
	return s
}

func (s *Nodepool) SetScalingGroup(v *NodepoolScalingGroup) *Nodepool {
	s.ScalingGroup = v
	return s
}

func (s *Nodepool) SetTeeConfig(v *NodepoolTeeConfig) *Nodepool {
	s.TeeConfig = v
	return s
}

func (s *Nodepool) Validate() error {
	if s.AutoScaling != nil {
		if err := s.AutoScaling.Validate(); err != nil {
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

type NodepoolAutoScaling struct {
	// Deprecated
	//
	// [This parameter is deprecated.]
	//
	// The peak bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated.]
	//
	// The billing method for the EIP. Valid values:
	//
	// - `PayByBandwidth`: Pay-by-bandwidth.
	//
	// - `PayByTraffic`: Pay-by-traffic.
	//
	// Default value: PayByBandwidth.
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Whether to enable auto scaling.
	//
	// - `true`: enables auto scaling for the node pool.
	//
	// - `false`: Disables auto scaling. If you set this parameter to `false`, other parameters in the `auto_scaling` object do not take effect.
	//
	// Default value: `false`.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated.]
	//
	// Whether to associate an EIP with each node in the pool. Valid values:
	//
	// - `true`: associates an EIP.
	//
	// - `false`: does not associate an EIP.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances in the auto scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances in the auto scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The type of instances to which auto scaling applies. This determines the scaling behavior. Valid values:
	//
	// - `cpu`: Regular instances.
	//
	// - `gpu`: GPU-accelerated instances.
	//
	// - `gpushare`: Shared GPU-accelerated instances.
	//
	// - `spot`: Spot instances.
	//
	// Default value: `cpu`.
	//
	// example:
	//
	// cpu
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s NodepoolAutoScaling) String() string {
	return dara.Prettify(s)
}

func (s NodepoolAutoScaling) GoString() string {
	return s.String()
}

func (s *NodepoolAutoScaling) GetEipBandwidth() *int64 {
	return s.EipBandwidth
}

func (s *NodepoolAutoScaling) GetEipInternetChargeType() *string {
	return s.EipInternetChargeType
}

func (s *NodepoolAutoScaling) GetEnable() *bool {
	return s.Enable
}

func (s *NodepoolAutoScaling) GetIsBondEip() *bool {
	return s.IsBondEip
}

func (s *NodepoolAutoScaling) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *NodepoolAutoScaling) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *NodepoolAutoScaling) GetType() *string {
	return s.Type
}

func (s *NodepoolAutoScaling) SetEipBandwidth(v int64) *NodepoolAutoScaling {
	s.EipBandwidth = &v
	return s
}

func (s *NodepoolAutoScaling) SetEipInternetChargeType(v string) *NodepoolAutoScaling {
	s.EipInternetChargeType = &v
	return s
}

func (s *NodepoolAutoScaling) SetEnable(v bool) *NodepoolAutoScaling {
	s.Enable = &v
	return s
}

func (s *NodepoolAutoScaling) SetIsBondEip(v bool) *NodepoolAutoScaling {
	s.IsBondEip = &v
	return s
}

func (s *NodepoolAutoScaling) SetMaxInstances(v int64) *NodepoolAutoScaling {
	s.MaxInstances = &v
	return s
}

func (s *NodepoolAutoScaling) SetMinInstances(v int64) *NodepoolAutoScaling {
	s.MinInstances = &v
	return s
}

func (s *NodepoolAutoScaling) SetType(v string) *NodepoolAutoScaling {
	s.Type = &v
	return s
}

func (s *NodepoolAutoScaling) Validate() error {
	return dara.Validate(s)
}

type NodepoolInterconnectConfig struct {
	// Deprecated
	//
	// 【该字段已废弃】
	//
	// 边缘增强型节点池的网络带宽，单位：Mbps。
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// Deprecated
	//
	// 【该字段已废弃】
	//
	// 边缘增强型节点池绑定的云连接网实例 ID(CCNID)。
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// Deprecated
	//
	// 【该字段已废弃】
	//
	// 边缘增强型节点池绑定的云连接网实例所属的地域。
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// Deprecated
	//
	// [This parameter is deprecated.]
	//
	// The ID of the Cloud Enterprise Network (C
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// Deprecated
	//
	// 【该字段已废弃】
	//
	// 边缘增强型节点池的购买时长，单位：月。
	//
	// example:
	//
	// 1
	ImprovedPeriod *string `json:"improved_period,omitempty" xml:"improved_period,omitempty"`
}

func (s NodepoolInterconnectConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolInterconnectConfig) GoString() string {
	return s.String()
}

func (s *NodepoolInterconnectConfig) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *NodepoolInterconnectConfig) GetCcnId() *string {
	return s.CcnId
}

func (s *NodepoolInterconnectConfig) GetCcnRegionId() *string {
	return s.CcnRegionId
}

func (s *NodepoolInterconnectConfig) GetCenId() *string {
	return s.CenId
}

func (s *NodepoolInterconnectConfig) GetImprovedPeriod() *string {
	return s.ImprovedPeriod
}

func (s *NodepoolInterconnectConfig) SetBandwidth(v int64) *NodepoolInterconnectConfig {
	s.Bandwidth = &v
	return s
}

func (s *NodepoolInterconnectConfig) SetCcnId(v string) *NodepoolInterconnectConfig {
	s.CcnId = &v
	return s
}

func (s *NodepoolInterconnectConfig) SetCcnRegionId(v string) *NodepoolInterconnectConfig {
	s.CcnRegionId = &v
	return s
}

func (s *NodepoolInterconnectConfig) SetCenId(v string) *NodepoolInterconnectConfig {
	s.CenId = &v
	return s
}

func (s *NodepoolInterconnectConfig) SetImprovedPeriod(v string) *NodepoolInterconnectConfig {
	s.ImprovedPeriod = &v
	return s
}

func (s *NodepoolInterconnectConfig) Validate() error {
	return dara.Validate(s)
}

type NodepoolKubernetesConfig struct {
	// Whether to install CloudMonitor on the nodes. After installation, you can view monitoring information about the instances in the CloudMonitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: installs CloudMonitor on nodes.
	//
	// - `false`: does not install CloudMonitor on nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of the nodes. This parameter is available only for clusters of Kubernetes 1.12.6 and later. The following policies are supported:
	//
	// - `static`: allows pods with specific resource characteristics to be granted with enhanced CPU affinity and exclusivity on the node.
	//
	// - `none`: indicates that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels to add to the nodes in the node pool.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The naming convention of the node. A node name consists of a prefix, an IP address, and a suffix.
	//
	// - The prefix and suffix can consist of one or more parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). A node name must start and end with a lowercase letter or a digit.
	//
	// - The node IP address is the complete private IP address of the node.
	//
	// This parameter consists of four comma-separated parts. For example, if you set the parameter to `customized,aliyun,ip,com`, where `customized` and `ip` are fixed strings, `aliyun` is the prefix, and `com` is the suffix, the node name is in the format of `aliyun.192.168.xxx.xxx.com`.
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The container runtime. Valid values:
	//
	// - `containerd`: Recommended. This runtime is supported in all cluster versions.
	//
	// - `Sandboxed-Container.runv`: a sandboxed container that provides higher isolation. This runtime is supported in clusters of Kubernetes 1.24 and earlier.
	//
	// - `docker`: This runtime is supported in clusters of Kubernetes 1.22 and earlier.
	//
	// Default value: `containerd`
	//
	// This parameter is required.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The version of the container runtime.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.6.20
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// The taints.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// The user data of the node.
	//
	// example:
	//
	// MXM=
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s NodepoolKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolKubernetesConfig) GoString() string {
	return s.String()
}

func (s *NodepoolKubernetesConfig) GetCmsEnabled() *bool {
	return s.CmsEnabled
}

func (s *NodepoolKubernetesConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *NodepoolKubernetesConfig) GetLabels() []*Tag {
	return s.Labels
}

func (s *NodepoolKubernetesConfig) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *NodepoolKubernetesConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *NodepoolKubernetesConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *NodepoolKubernetesConfig) GetTaints() []*Taint {
	return s.Taints
}

func (s *NodepoolKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *NodepoolKubernetesConfig) SetCmsEnabled(v bool) *NodepoolKubernetesConfig {
	s.CmsEnabled = &v
	return s
}

func (s *NodepoolKubernetesConfig) SetCpuPolicy(v string) *NodepoolKubernetesConfig {
	s.CpuPolicy = &v
	return s
}

func (s *NodepoolKubernetesConfig) SetLabels(v []*Tag) *NodepoolKubernetesConfig {
	s.Labels = v
	return s
}

func (s *NodepoolKubernetesConfig) SetNodeNameMode(v string) *NodepoolKubernetesConfig {
	s.NodeNameMode = &v
	return s
}

func (s *NodepoolKubernetesConfig) SetRuntime(v string) *NodepoolKubernetesConfig {
	s.Runtime = &v
	return s
}

func (s *NodepoolKubernetesConfig) SetRuntimeVersion(v string) *NodepoolKubernetesConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *NodepoolKubernetesConfig) SetTaints(v []*Taint) *NodepoolKubernetesConfig {
	s.Taints = v
	return s
}

func (s *NodepoolKubernetesConfig) SetUserData(v string) *NodepoolKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *NodepoolKubernetesConfig) Validate() error {
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

type NodepoolManagement struct {
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to enable auto repair. This parameter takes effect only if `enable` is set to `true`.
	//
	// - `true`
	//
	// - `false`
	//
	// example:
	//
	// false
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto repair policy for nodes.
	AutoRepairPolicy *NodepoolManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to enable auto upgrade.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto upgrade policy.
	AutoUpgradePolicy *NodepoolManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically fix CVEs.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The policy for automatically fixing CVEs.
	AutoVulFixPolicy *NodepoolManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable the managed node pool feature. Valid values:
	//
	// - `true`
	//
	// - `false`: If you set this parameter to false, the other parameters in the `management` object are ignored.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// The auto upgrade configurations. This parameter takes effect only when `enable` is set to `true`.
	UpgradeConfig *NodepoolManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s NodepoolManagement) String() string {
	return dara.Prettify(s)
}

func (s NodepoolManagement) GoString() string {
	return s.String()
}

func (s *NodepoolManagement) GetAutoFaultDiagnosis() *bool {
	return s.AutoFaultDiagnosis
}

func (s *NodepoolManagement) GetAutoRepair() *bool {
	return s.AutoRepair
}

func (s *NodepoolManagement) GetAutoRepairPolicy() *NodepoolManagementAutoRepairPolicy {
	return s.AutoRepairPolicy
}

func (s *NodepoolManagement) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *NodepoolManagement) GetAutoUpgradePolicy() *NodepoolManagementAutoUpgradePolicy {
	return s.AutoUpgradePolicy
}

func (s *NodepoolManagement) GetAutoVulFix() *bool {
	return s.AutoVulFix
}

func (s *NodepoolManagement) GetAutoVulFixPolicy() *NodepoolManagementAutoVulFixPolicy {
	return s.AutoVulFixPolicy
}

func (s *NodepoolManagement) GetEnable() *bool {
	return s.Enable
}

func (s *NodepoolManagement) GetUpgradeConfig() *NodepoolManagementUpgradeConfig {
	return s.UpgradeConfig
}

func (s *NodepoolManagement) SetAutoFaultDiagnosis(v bool) *NodepoolManagement {
	s.AutoFaultDiagnosis = &v
	return s
}

func (s *NodepoolManagement) SetAutoRepair(v bool) *NodepoolManagement {
	s.AutoRepair = &v
	return s
}

func (s *NodepoolManagement) SetAutoRepairPolicy(v *NodepoolManagementAutoRepairPolicy) *NodepoolManagement {
	s.AutoRepairPolicy = v
	return s
}

func (s *NodepoolManagement) SetAutoUpgrade(v bool) *NodepoolManagement {
	s.AutoUpgrade = &v
	return s
}

func (s *NodepoolManagement) SetAutoUpgradePolicy(v *NodepoolManagementAutoUpgradePolicy) *NodepoolManagement {
	s.AutoUpgradePolicy = v
	return s
}

func (s *NodepoolManagement) SetAutoVulFix(v bool) *NodepoolManagement {
	s.AutoVulFix = &v
	return s
}

func (s *NodepoolManagement) SetAutoVulFixPolicy(v *NodepoolManagementAutoVulFixPolicy) *NodepoolManagement {
	s.AutoVulFixPolicy = v
	return s
}

func (s *NodepoolManagement) SetEnable(v bool) *NodepoolManagement {
	s.Enable = &v
	return s
}

func (s *NodepoolManagement) SetUpgradeConfig(v *NodepoolManagementUpgradeConfig) *NodepoolManagement {
	s.UpgradeConfig = v
	return s
}

func (s *NodepoolManagement) Validate() error {
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

type NodepoolManagementAutoRepairPolicy struct {
	// Specifies whether to allow the node to be restarted.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
}

func (s NodepoolManagementAutoRepairPolicy) String() string {
	return dara.Prettify(s)
}

func (s NodepoolManagementAutoRepairPolicy) GoString() string {
	return s.String()
}

func (s *NodepoolManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *NodepoolManagementAutoRepairPolicy) SetRestartNode(v bool) *NodepoolManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *NodepoolManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type NodepoolManagementAutoUpgradePolicy struct {
	// Specifies whether to allow the kubelet to be automatically upgraded.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
}

func (s NodepoolManagementAutoUpgradePolicy) String() string {
	return dara.Prettify(s)
}

func (s NodepoolManagementAutoUpgradePolicy) GoString() string {
	return s.String()
}

func (s *NodepoolManagementAutoUpgradePolicy) GetAutoUpgradeKubelet() *bool {
	return s.AutoUpgradeKubelet
}

func (s *NodepoolManagementAutoUpgradePolicy) SetAutoUpgradeKubelet(v bool) *NodepoolManagementAutoUpgradePolicy {
	s.AutoUpgradeKubelet = &v
	return s
}

func (s *NodepoolManagementAutoUpgradePolicy) Validate() error {
	return dara.Validate(s)
}

type NodepoolManagementAutoVulFixPolicy struct {
	// Specifies whether to allow the node to be restarted.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The CVE vulnerability levels to automatically fix. You can specify multiple levels separated by commas.
	//
	// example:
	//
	// asap,nntf
	VulLevel *string `json:"vul_level,omitempty" xml:"vul_level,omitempty"`
}

func (s NodepoolManagementAutoVulFixPolicy) String() string {
	return dara.Prettify(s)
}

func (s NodepoolManagementAutoVulFixPolicy) GoString() string {
	return s.String()
}

func (s *NodepoolManagementAutoVulFixPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *NodepoolManagementAutoVulFixPolicy) GetVulLevel() *string {
	return s.VulLevel
}

func (s *NodepoolManagementAutoVulFixPolicy) SetRestartNode(v bool) *NodepoolManagementAutoVulFixPolicy {
	s.RestartNode = &v
	return s
}

func (s *NodepoolManagementAutoVulFixPolicy) SetVulLevel(v string) *NodepoolManagementAutoVulFixPolicy {
	s.VulLevel = &v
	return s
}

func (s *NodepoolManagementAutoVulFixPolicy) Validate() error {
	return dara.Validate(s)
}

type NodepoolManagementUpgradeConfig struct {
	// Specifies whether to enable auto upgrade. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	// example:
	//
	// false
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes. Valid values: 1 to 1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 0
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes.
	//
	// example:
	//
	// 0
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You must specify this parameter or `surge`.
	//
	// example:
	//
	// 0
	SurgePercentage *int64 `json:"surge_percentage,omitempty" xml:"surge_percentage,omitempty"`
}

func (s NodepoolManagementUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolManagementUpgradeConfig) GoString() string {
	return s.String()
}

func (s *NodepoolManagementUpgradeConfig) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *NodepoolManagementUpgradeConfig) GetMaxUnavailable() *int64 {
	return s.MaxUnavailable
}

func (s *NodepoolManagementUpgradeConfig) GetSurge() *int64 {
	return s.Surge
}

func (s *NodepoolManagementUpgradeConfig) GetSurgePercentage() *int64 {
	return s.SurgePercentage
}

func (s *NodepoolManagementUpgradeConfig) SetAutoUpgrade(v bool) *NodepoolManagementUpgradeConfig {
	s.AutoUpgrade = &v
	return s
}

func (s *NodepoolManagementUpgradeConfig) SetMaxUnavailable(v int64) *NodepoolManagementUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *NodepoolManagementUpgradeConfig) SetSurge(v int64) *NodepoolManagementUpgradeConfig {
	s.Surge = &v
	return s
}

func (s *NodepoolManagementUpgradeConfig) SetSurgePercentage(v int64) *NodepoolManagementUpgradeConfig {
	s.SurgePercentage = &v
	return s
}

func (s *NodepoolManagementUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type NodepoolNodeComponents struct {
	// 节点组件配置。
	Config *NodepoolNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// 节点组件名称。
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点组件版本
	//
	// example:
	//
	// 1.33.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s NodepoolNodeComponents) String() string {
	return dara.Prettify(s)
}

func (s NodepoolNodeComponents) GoString() string {
	return s.String()
}

func (s *NodepoolNodeComponents) GetConfig() *NodepoolNodeComponentsConfig {
	return s.Config
}

func (s *NodepoolNodeComponents) GetName() *string {
	return s.Name
}

func (s *NodepoolNodeComponents) GetVersion() *string {
	return s.Version
}

func (s *NodepoolNodeComponents) SetConfig(v *NodepoolNodeComponentsConfig) *NodepoolNodeComponents {
	s.Config = v
	return s
}

func (s *NodepoolNodeComponents) SetName(v string) *NodepoolNodeComponents {
	s.Name = &v
	return s
}

func (s *NodepoolNodeComponents) SetVersion(v string) *NodepoolNodeComponents {
	s.Version = &v
	return s
}

func (s *NodepoolNodeComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NodepoolNodeComponentsConfig struct {
	// 节点组件自定义配置。
	CustomConfig map[string]*string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s NodepoolNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *NodepoolNodeComponentsConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *NodepoolNodeComponentsConfig) SetCustomConfig(v map[string]*string) *NodepoolNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *NodepoolNodeComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type NodepoolNodeConfig struct {
	// Kubelet 参数配置。
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
}

func (s NodepoolNodeConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolNodeConfig) GoString() string {
	return s.String()
}

func (s *NodepoolNodeConfig) GetKubeletConfiguration() *KubeletConfig {
	return s.KubeletConfiguration
}

func (s *NodepoolNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *NodepoolNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *NodepoolNodeConfig) Validate() error {
	if s.KubeletConfiguration != nil {
		if err := s.KubeletConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NodepoolNodepoolInfo struct {
	// The name of the node pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group to which the node pool belongs.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of the node pool. Valid values:
	//
	// - `ess`: a regular node pool.
	//
	// - `edge`: an edge node pool.
	//
	// example:
	//
	// ess
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s NodepoolNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s NodepoolNodepoolInfo) GoString() string {
	return s.String()
}

func (s *NodepoolNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *NodepoolNodepoolInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *NodepoolNodepoolInfo) GetType() *string {
	return s.Type
}

func (s *NodepoolNodepoolInfo) SetName(v string) *NodepoolNodepoolInfo {
	s.Name = &v
	return s
}

func (s *NodepoolNodepoolInfo) SetResourceGroupId(v string) *NodepoolNodepoolInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *NodepoolNodepoolInfo) SetType(v string) *NodepoolNodepoolInfo {
	s.Type = &v
	return s
}

func (s *NodepoolNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type NodepoolScalingGroup struct {
	// Whether to enable auto-renewal for the nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: enables auto-renewal.
	//
	// - `false`: disables auto-renewal.
	//
	// Default value: `true`.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period for the nodes. This parameter is required and takes effect only if `instance_charge_type` is set to `PrePaid`.
	//
	// If `PeriodUnit` is set to `Month`, the valid values are 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// example:
	//
	// 0
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the required number of ECS instances when `multi_az_policy` is set to `COST_OPTIMIZED` and preemptible instances cannot be created due to price or inventory constraints. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The configurations of the data disks that are attached to the nodes in the node pool.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The ID of the deployment set.
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
	// The configurations of block devices.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The ID of the custom image. If you do not set this parameter, the default system image is used.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200904.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the OS image. You must specify this parameter or `platform`. Valid values:
	//
	// - `AliyunLinux`: Alinux 2 image.
	//
	// - `AliyunLinux3`: Alinux 3 image.
	//
	// - `AliyunLinux3Arm64`: Alinux 3 image for ARM.
	//
	// - `AliyunLinuxUEFI`: Alinux 2 UEFI image.
	//
	// - `CentOS`: CentOS image.
	//
	// - `Windows`: Windows image.
	//
	// - `WindowsCore`: Windows Core image.
	//
	// - `ContainerOS`: Container-optimized image.
	//
	// example:
	//
	// AliyunLinux
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
	// The configurations of metadata access for the ECS instances.
	//
	// This feature is available only to allowlisted users. To use this feature, submit a ticket.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance types.
	//
	// This parameter is required.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method of the public IP address. Valid values:
	//
	// - `PayByBandwidth`: pay-by-bandwidth.
	//
	// - `PayByTraffic`: pay-by-traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound public bandwidth of a node. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair. You must specify this parameter or `login_password`.
	//
	// > If you create a managed node pool, you can specify only `key_pair`.
	//
	// example:
	//
	// np-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether to use a non-root user to log on to the ECS instances that are created.
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The password for SSH access. You must specify this parameter or `key_pair`. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for the ECS instances in a scaling group that spans multiple zones. Valid values:
	//
	// - `PRIORITY`: The system scales instances based on the priority of the vSwitches specified in `vswitch_ids`. The system preferentially scales instances in the zone where the vSwitch with the highest priority resides. If the scaling fails, the system scales instances in the zone where the vSwitch with the next highest priority resides.
	//
	// - `COST_OPTIMIZED`: The system scales instances based on the vCPU price from lowest to highest. The system preferentially scales instances that have the lowest vCPU price. If the scaling configuration includes multiple instance types and some of them are preemptible instance types, the system preferentially scales the preemptible instances. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when preemptible instances cannot be created due to reasons such as stock-outs.
	//
	//   > The `COST_OPTIMIZED` policy takes effect only when multiple instance types are specified or preemptible instances are selected in the scaling configuration.
	//
	// - `BALANCE`: The system evenly distributes ECS instances across the specified zones. If the distribution of instances becomes unbalanced due to stock-outs, you can call the [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) operation to rebalance the resources.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be provisioned in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than this value, the system preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the instances that exceed the minimum number specified by `on_demand_base_capacity`. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of the nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`. If `period_unit` is set to `Month`, the valid values for `period` are 1, 2, 3, 6, and 12.
	//
	// Default value: 1.
	//
	// example:
	//
	// 0
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the subscription nodes in the node pool. This parameter is required when `instance_charge_type` is set to `PrePaid`.
	//
	// `Month`: The billing cycle is measured in months.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
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
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The configurations of the private node pool.
	PrivatePoolOptions *NodepoolScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker RAM role.
	//
	// 	Notice: This parameter can be configured only when you create a node pool in an ACK managed cluster of Kubernetes 1.22 or later.
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// A list of ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy that are used to create instances. Note the following when you set this parameter:
	//
	// This parameter takes effect only when you create pay-as-you-go instances.
	//
	// You cannot specify this parameter together with `private_pool_options.match_criteria` or `private_pool_options.id`.
	ResourcePoolOptions *NodepoolScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling mode of the scaling group. Valid values:
	//
	// - `release`: the standard mode. In this mode, resources are scaled by creating and releasing ECS instances based on the resource usage.
	//
	// - `recycle`: the rapid mode. In this mode, resources are scaled by creating, stopping, and starting ECS instances. This speeds up the scaling of resources. When an instance is stopped, you are not charged for its computing resources but are charged for its storage resources. This does not apply to instances that have local disks.
	//
	// Default value: `release`.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The ID of the security group to which you want to add the nodes. You must specify this parameter or `security_group_ids`. We recommend that you specify `security_group_ids`.
	//
	// example:
	//
	// sg-2zeihch86ooz9io4****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// A list of security group IDs. You must specify this parameter or `security_group_id`. We recommend that you specify `security_group_ids`. If you specify both `security_group_id` and `security_group_ids`, `security_group_ids` takes precedence.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// The number of instance types. The scaling group creates preemptible instances of multiple instance types that have the lowest cost in a balanced manner. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable replenishment for preemptible instances. If this feature is enabled, the scaling group attempts to create a new instance to replace a preemptible instance that is about to be reclaimed. Valid values:
	//
	// - `true`
	//
	// - `false`
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price limits for specific spot instance types.
	SpotPriceLimit []*NodepoolScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The bidding policy for the pay-as-you-go instances. Valid values:
	//
	// - NoSpot: normal pay-as-you-go instances.
	//
	// - SpotWithPriceLimit: spot instances with a user-defined maximum hourly price.
	//
	// - SpotAsPriceGo: spot instances for which the system automatically bids based on the current market price.
	//
	// For more information, see [Preemptible instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// - true: enables the performance burst feature.
	//
	// - false: disables the performance burst feature.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// A prioritized list of system disk types. The system attempts to create a system disk of a disk type with a higher priority. If the disk type is unavailable, the system uses the next disk type to create the system disk. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system disk. Valid values:
	//
	// - `cloud_efficiency`: ultra disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: enhanced SSD (ESSD).
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
	// The encryption algorithm that is used for the system disk. Set the value to aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values: true and false.
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used to encrypt the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the ESSD that is used as the system disk. This parameter takes effect only for ESSDs.
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS of the system disk. The valid values are 0 to min{50000, 1000 × Capacity - Base IOPS}. The default Base IOPS is min{1800 + 50 × Capacity, 50000}.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// Valid values: 40 to 500.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The tags to add to the ECS instances in the node pool.
	//
	// A tag key must be unique and can be up to 128 characters in length. A tag key and a tag value cannot start with `aliyun` or `acs:` and cannot contain `https://` or `http://`.
	Tags []*NodepoolScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The IDs of the vSwitches to which the nodes can be added.
	//
	// This parameter is required.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s NodepoolScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroup) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroup) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *NodepoolScalingGroup) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *NodepoolScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *NodepoolScalingGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *NodepoolScalingGroup) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *NodepoolScalingGroup) GetDesiredSize() *int64 {
	return s.DesiredSize
}

func (s *NodepoolScalingGroup) GetDiskInit() []*DiskInit {
	return s.DiskInit
}

func (s *NodepoolScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *NodepoolScalingGroup) GetImageType() *string {
	return s.ImageType
}

func (s *NodepoolScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *NodepoolScalingGroup) GetInstanceMetadataOptions() *InstanceMetadataOptions {
	return s.InstanceMetadataOptions
}

func (s *NodepoolScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *NodepoolScalingGroup) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *NodepoolScalingGroup) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *NodepoolScalingGroup) GetKeyPair() *string {
	return s.KeyPair
}

func (s *NodepoolScalingGroup) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *NodepoolScalingGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *NodepoolScalingGroup) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *NodepoolScalingGroup) GetOnDemandBaseCapacity() *int64 {
	return s.OnDemandBaseCapacity
}

func (s *NodepoolScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int64 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *NodepoolScalingGroup) GetPeriod() *int64 {
	return s.Period
}

func (s *NodepoolScalingGroup) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *NodepoolScalingGroup) GetPlatform() *string {
	return s.Platform
}

func (s *NodepoolScalingGroup) GetPrivatePoolOptions() *NodepoolScalingGroupPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *NodepoolScalingGroup) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *NodepoolScalingGroup) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *NodepoolScalingGroup) GetResourcePoolOptions() *NodepoolScalingGroupResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *NodepoolScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *NodepoolScalingGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *NodepoolScalingGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *NodepoolScalingGroup) GetSpotInstancePools() *int64 {
	return s.SpotInstancePools
}

func (s *NodepoolScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *NodepoolScalingGroup) GetSpotPriceLimit() []*NodepoolScalingGroupSpotPriceLimit {
	return s.SpotPriceLimit
}

func (s *NodepoolScalingGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *NodepoolScalingGroup) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *NodepoolScalingGroup) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *NodepoolScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *NodepoolScalingGroup) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *NodepoolScalingGroup) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *NodepoolScalingGroup) GetSystemDiskKmsKeyId() *string {
	return s.SystemDiskKmsKeyId
}

func (s *NodepoolScalingGroup) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *NodepoolScalingGroup) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *NodepoolScalingGroup) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *NodepoolScalingGroup) GetTags() []*NodepoolScalingGroupTags {
	return s.Tags
}

func (s *NodepoolScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *NodepoolScalingGroup) SetAutoRenew(v bool) *NodepoolScalingGroup {
	s.AutoRenew = &v
	return s
}

func (s *NodepoolScalingGroup) SetAutoRenewPeriod(v int64) *NodepoolScalingGroup {
	s.AutoRenewPeriod = &v
	return s
}

func (s *NodepoolScalingGroup) SetCompensateWithOnDemand(v bool) *NodepoolScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *NodepoolScalingGroup) SetDataDisks(v []*DataDisk) *NodepoolScalingGroup {
	s.DataDisks = v
	return s
}

func (s *NodepoolScalingGroup) SetDeploymentsetId(v string) *NodepoolScalingGroup {
	s.DeploymentsetId = &v
	return s
}

func (s *NodepoolScalingGroup) SetDesiredSize(v int64) *NodepoolScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *NodepoolScalingGroup) SetDiskInit(v []*DiskInit) *NodepoolScalingGroup {
	s.DiskInit = v
	return s
}

func (s *NodepoolScalingGroup) SetImageId(v string) *NodepoolScalingGroup {
	s.ImageId = &v
	return s
}

func (s *NodepoolScalingGroup) SetImageType(v string) *NodepoolScalingGroup {
	s.ImageType = &v
	return s
}

func (s *NodepoolScalingGroup) SetInstanceChargeType(v string) *NodepoolScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *NodepoolScalingGroup) SetInstanceMetadataOptions(v *InstanceMetadataOptions) *NodepoolScalingGroup {
	s.InstanceMetadataOptions = v
	return s
}

func (s *NodepoolScalingGroup) SetInstanceTypes(v []*string) *NodepoolScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *NodepoolScalingGroup) SetInternetChargeType(v string) *NodepoolScalingGroup {
	s.InternetChargeType = &v
	return s
}

func (s *NodepoolScalingGroup) SetInternetMaxBandwidthOut(v int64) *NodepoolScalingGroup {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *NodepoolScalingGroup) SetKeyPair(v string) *NodepoolScalingGroup {
	s.KeyPair = &v
	return s
}

func (s *NodepoolScalingGroup) SetLoginAsNonRoot(v bool) *NodepoolScalingGroup {
	s.LoginAsNonRoot = &v
	return s
}

func (s *NodepoolScalingGroup) SetLoginPassword(v string) *NodepoolScalingGroup {
	s.LoginPassword = &v
	return s
}

func (s *NodepoolScalingGroup) SetMultiAzPolicy(v string) *NodepoolScalingGroup {
	s.MultiAzPolicy = &v
	return s
}

func (s *NodepoolScalingGroup) SetOnDemandBaseCapacity(v int64) *NodepoolScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *NodepoolScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int64) *NodepoolScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *NodepoolScalingGroup) SetPeriod(v int64) *NodepoolScalingGroup {
	s.Period = &v
	return s
}

func (s *NodepoolScalingGroup) SetPeriodUnit(v string) *NodepoolScalingGroup {
	s.PeriodUnit = &v
	return s
}

func (s *NodepoolScalingGroup) SetPlatform(v string) *NodepoolScalingGroup {
	s.Platform = &v
	return s
}

func (s *NodepoolScalingGroup) SetPrivatePoolOptions(v *NodepoolScalingGroupPrivatePoolOptions) *NodepoolScalingGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *NodepoolScalingGroup) SetRamRoleName(v string) *NodepoolScalingGroup {
	s.RamRoleName = &v
	return s
}

func (s *NodepoolScalingGroup) SetRdsInstances(v []*string) *NodepoolScalingGroup {
	s.RdsInstances = v
	return s
}

func (s *NodepoolScalingGroup) SetResourcePoolOptions(v *NodepoolScalingGroupResourcePoolOptions) *NodepoolScalingGroup {
	s.ResourcePoolOptions = v
	return s
}

func (s *NodepoolScalingGroup) SetScalingPolicy(v string) *NodepoolScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *NodepoolScalingGroup) SetSecurityGroupId(v string) *NodepoolScalingGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *NodepoolScalingGroup) SetSecurityGroupIds(v []*string) *NodepoolScalingGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *NodepoolScalingGroup) SetSpotInstancePools(v int64) *NodepoolScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *NodepoolScalingGroup) SetSpotInstanceRemedy(v bool) *NodepoolScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *NodepoolScalingGroup) SetSpotPriceLimit(v []*NodepoolScalingGroupSpotPriceLimit) *NodepoolScalingGroup {
	s.SpotPriceLimit = v
	return s
}

func (s *NodepoolScalingGroup) SetSpotStrategy(v string) *NodepoolScalingGroup {
	s.SpotStrategy = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskBurstingEnabled(v bool) *NodepoolScalingGroup {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskCategories(v []*string) *NodepoolScalingGroup {
	s.SystemDiskCategories = v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskCategory(v string) *NodepoolScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskEncryptAlgorithm(v string) *NodepoolScalingGroup {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskEncrypted(v bool) *NodepoolScalingGroup {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskKmsKeyId(v string) *NodepoolScalingGroup {
	s.SystemDiskKmsKeyId = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskPerformanceLevel(v string) *NodepoolScalingGroup {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskProvisionedIops(v int64) *NodepoolScalingGroup {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *NodepoolScalingGroup) SetSystemDiskSize(v int64) *NodepoolScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *NodepoolScalingGroup) SetTags(v []*NodepoolScalingGroupTags) *NodepoolScalingGroup {
	s.Tags = v
	return s
}

func (s *NodepoolScalingGroup) SetVswitchIds(v []*string) *NodepoolScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *NodepoolScalingGroup) Validate() error {
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

type NodepoolScalingGroupPrivatePoolOptions struct {
	// The ID of the private node pool.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private pool that is used to start instances. An elasticity assurance or a capacity reservation service generates a private pool after the service takes effect. You can select a private pool to start instances. Valid values:
	//
	// - `Open`: open mode. The system automatically matches an open private pool. If no matching private pool is found, the resources in the public pool are used.
	//
	// - `Target`: targeted mode. The instance is started from the specified private pool. If the private pool is unavailable, the instance fails to be started.
	//
	// - `None`: no mode. The instance is not started from a private pool.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"match_criteria,omitempty" xml:"match_criteria,omitempty"`
}

func (s NodepoolScalingGroupPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroupPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroupPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *NodepoolScalingGroupPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *NodepoolScalingGroupPrivatePoolOptions) SetId(v string) *NodepoolScalingGroupPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *NodepoolScalingGroupPrivatePoolOptions) SetMatchCriteria(v string) *NodepoolScalingGroupPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *NodepoolScalingGroupPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type NodepoolScalingGroupResourcePoolOptions struct {
	// The list of private pool IDs. The IDs are the IDs of elasticity assurances or capacity reservations. This parameter accepts only IDs of private pools in targeted mode. You can specify up to 20 IDs.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy that is used to create instances. Resource pools include private pools generated by an elasticity assurance or a capacity reservation service and public pools. You can select a resource pool to start an instance. Valid values:
	//
	// PrivatePoolFirst: The system prioritizes the use of private pools. If you specify `resource_pool_options.private_pool_ids`, the specified private pools are used first. If no private pool is specified or the capacity of the specified private pool is insufficient, the system automatically matches an open private pool. If no matching private pool is found, the resources in the public pool are used.
	//
	// PrivatePoolOnly: The instance can be created only from a private pool. You must specify `resource_pool_options.private_pool_ids`. If the capacity of the specified private pool is insufficient, the instance fails to be created.
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

func (s NodepoolScalingGroupResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroupResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroupResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *NodepoolScalingGroupResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *NodepoolScalingGroupResourcePoolOptions) SetPrivatePoolIds(v []*string) *NodepoolScalingGroupResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *NodepoolScalingGroupResourcePoolOptions) SetStrategy(v string) *NodepoolScalingGroupResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *NodepoolScalingGroupResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type NodepoolScalingGroupSpotPriceLimit struct {
	// The spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum price of a single instance.
	//
	// example:
	//
	// 0.39
	PriceLimit *string `json:"price_limit,omitempty" xml:"price_limit,omitempty"`
}

func (s NodepoolScalingGroupSpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroupSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroupSpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *NodepoolScalingGroupSpotPriceLimit) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *NodepoolScalingGroupSpotPriceLimit) SetInstanceType(v string) *NodepoolScalingGroupSpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *NodepoolScalingGroupSpotPriceLimit) SetPriceLimit(v string) *NodepoolScalingGroupSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *NodepoolScalingGroupSpotPriceLimit) Validate() error {
	return dara.Validate(s)
}

type NodepoolScalingGroupTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s NodepoolScalingGroupTags) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroupTags) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroupTags) GetKey() *string {
	return s.Key
}

func (s *NodepoolScalingGroupTags) GetValue() *string {
	return s.Value
}

func (s *NodepoolScalingGroupTags) SetKey(v string) *NodepoolScalingGroupTags {
	s.Key = &v
	return s
}

func (s *NodepoolScalingGroupTags) SetValue(v string) *NodepoolScalingGroupTags {
	s.Value = &v
	return s
}

func (s *NodepoolScalingGroupTags) Validate() error {
	return dara.Validate(s)
}

type NodepoolTeeConfig struct {
	// Specifies whether to create a confidential computing node pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	TeeEnable *bool `json:"tee_enable,omitempty" xml:"tee_enable,omitempty"`
}

func (s NodepoolTeeConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolTeeConfig) GoString() string {
	return s.String()
}

func (s *NodepoolTeeConfig) GetTeeEnable() *bool {
	return s.TeeEnable
}

func (s *NodepoolTeeConfig) SetTeeEnable(v bool) *NodepoolTeeConfig {
	s.TeeEnable = &v
	return s
}

func (s *NodepoolTeeConfig) Validate() error {
	return dara.Validate(s)
}
