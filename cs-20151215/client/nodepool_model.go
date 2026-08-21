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
	// The auto scaling node pool configuration.
	AutoScaling *NodepoolAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- Use desired_size instead.
	//
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	InterconnectConfig *NodepoolInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This value is meaningful only for node pools whose `type` is `edge`. Valid values:
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// The cluster-related configuration.
	KubernetesConfig *NodepoolKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The managed node pool configuration.
	Management *NodepoolManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// The maximum number of nodes allowed in the edge node pool. This parameter must be greater than or equal to 0. A value of 0 indicates no additional limit (limited only by the maximum number of nodes the cluster can contain, with no additional limit on the node pool itself). Edge node pools typically have a value greater than 0. ESS type node pools and default edge type node pools have a value of 0.
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The node component list.
	NodeComponents []*NodepoolNodeComponents `json:"node_components,omitempty" xml:"node_components,omitempty" type:"Repeated"`
	// The node configuration.
	NodeConfig *NodepoolNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The node pool configuration.
	NodepoolInfo *NodepoolNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The node pool scaling group configuration.
	ScalingGroup *NodepoolScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The confidential computing node pool configuration.
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
	// **[Deprecated]**
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable automatic scaling.
	//
	// - `true`: Enables the automatic scaling feature for the node pool.
	//
	// - `false`: Disables automatic scaling. If this parameter is set to `false`, other configuration parameters in `auto_scaling` do not take effect.
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
	// **[Deprecated]**
	//
	// Specifies whether to associate an Elastic IP Address (EIP). Valid values:
	//
	// - `true`: Associates an EIP.
	//
	// - `false`: Does not associate an EIP.
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
	// The auto scaling type, classified by auto scaling instance type. Valid values:
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
	// **[Deprecated]**
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
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
	// Specifies whether to install the CloudMonitor agent on ECS nodes. After installation, you can view monitoring information of the created ECS instances in the CloudMonitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: Installs the CloudMonitor agent on ECS nodes.
	//
	// - `false`: Does not install the CloudMonitor agent on ECS nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy for nodes. The following two policies are supported when the cluster version is 1.12.6 or later:
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The node labels. Adds labels to Kubernetes cluster nodes.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The node name consists of three parts: prefix + node IP + suffix:
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The container runtime. Valid values:
	//
	// - `containerd`: Recommended. Supported by all cluster versions.
	//
	// - `Sandboxed-Container.runv`: Sandboxed container that provides higher isolation. Supported by clusters of version 1.24 and earlier.
	//
	// - `docker`: Supported by clusters of version 1.22 and earlier.
	//
	// Default value: `containerd`
	//
	// This parameter is required.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The container runtime version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.6.20
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// The taint configuration.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// The node custom data.
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
	// Specifies whether to enable self-healing ECS fault detection for nodes.
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to enable auto repair. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// false
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto repair node policy.
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
	// The auto CVE fix policy.
	AutoVulFixPolicy *NodepoolManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable node rotation. Only intelligent managed node pools support this feature, and it is enabled by default. Common node pools do not support this feature.
	DriftEnabled *bool `json:"drift_enabled,omitempty" xml:"drift_enabled,omitempty"`
	// Specifies whether to enable the managed node pool. Valid values:
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// The auto upgrade configuration. This parameter takes effect only when `enable=true`.
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

func (s *NodepoolManagement) GetDriftEnabled() *bool {
	return s.DriftEnabled
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

func (s *NodepoolManagement) SetDriftEnabled(v bool) *NodepoolManagement {
	s.DriftEnabled = &v
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
	// The maximum number of parallel repairs. When a large number of abnormal nodes exist in the node pool, this specifies the maximum number or percentage of nodes that can be repaired simultaneously. Supports a number (such as 5, valid range: 1 to 100000) or a percentage (such as 10%, valid range: 1% to 100%). Default value: 1.
	//
	// example:
	//
	// 5
	MaxParallelRepairingNodes *string `json:"max_parallel_repairing_nodes,omitempty" xml:"max_parallel_repairing_nodes,omitempty"`
	// The self-healing circuit breaker threshold. When the number or percentage of faulty nodes exceeds this threshold, self-healing enters a circuit breaker state and stops initiating new repair actions. Supports a number (such as 10, valid range: 1 to 100000) or a percentage (such as 20%, valid range: 1% to 100%). Default value: 100%.
	//
	// example:
	//
	// 20%
	MaxUnhealthyNodesThreshold *string `json:"max_unhealthy_nodes_threshold,omitempty" xml:"max_unhealthy_nodes_threshold,omitempty"`
	// Specifies whether to allow node restart.
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

func (s *NodepoolManagementAutoRepairPolicy) GetMaxParallelRepairingNodes() *string {
	return s.MaxParallelRepairingNodes
}

func (s *NodepoolManagementAutoRepairPolicy) GetMaxUnhealthyNodesThreshold() *string {
	return s.MaxUnhealthyNodesThreshold
}

func (s *NodepoolManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *NodepoolManagementAutoRepairPolicy) SetMaxParallelRepairingNodes(v string) *NodepoolManagementAutoRepairPolicy {
	s.MaxParallelRepairingNodes = &v
	return s
}

func (s *NodepoolManagementAutoRepairPolicy) SetMaxUnhealthyNodesThreshold(v string) *NodepoolManagementAutoRepairPolicy {
	s.MaxUnhealthyNodesThreshold = &v
	return s
}

func (s *NodepoolManagementAutoRepairPolicy) SetRestartNode(v bool) *NodepoolManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *NodepoolManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type NodepoolManagementAutoUpgradePolicy struct {
	// Specifies whether to allow auto upgrade of kubelet.
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
	// Specifies whether to allow node restart.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels allowed for auto fix, separated by commas.
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
	// example:
	//
	// false
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes. Valid values: [1,1000\\].
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
	// The percentage of extra nodes. This parameter is mutually exclusive with `surge`.
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
	// The node component configuration.
	Config *NodepoolNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The node component name.
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
	// The custom configuration of the node component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s NodepoolNodeComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s NodepoolNodeComponentsConfig) GoString() string {
	return s.String()
}

func (s *NodepoolNodeComponentsConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *NodepoolNodeComponentsConfig) SetCustomConfig(v map[string]interface{}) *NodepoolNodeComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *NodepoolNodeComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type NodepoolNodeConfig struct {
	// The kubelet parameter settings.
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
	// The node pool name.
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
	// The node pool type. Valid values:
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
	// Specifies whether to enable auto-renewal for the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period for the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`, and is required in that case.
	//
	// example:
	//
	// 0
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the required number of ECS instances when spot instances cannot be created due to cost or inventory reasons, if `multi_az_policy` is set to `COST_OPTIMIZED`. Valid values:
	//
	// - `true`: Allows automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: Does not allow automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The CPU configuration options.
	CpuOptions *NodepoolScalingGroupCpuOptions `json:"cpu_options,omitempty" xml:"cpu_options,omitempty" type:"Struct"`
	// The data cloud disk configuration for node pool nodes.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The desired number of nodes in the node pool.
	//
	// example:
	//
	// 2
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The custom image ID. The system-provided image is used by default.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200904.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The operating system image type. You can specify either this parameter or the platform parameter. Valid values:
	//
	// - `AliyunLinux`: Alinux2 image.
	//
	// - `AliyunLinux3`: Alinux3 image.
	//
	// - `AliyunLinux3Arm64`: Alinux3 image for ARM.
	//
	// - `AliyunLinuxUEFI`: Alinux2 UEFI image.
	//
	// - `CentOS`: CentOS image.
	//
	// - `Windows`: Windows image.
	//
	// - `WindowsCore`: WindowsCore image.
	//
	// - `ContainerOS`: container-optimized image.
	//
	// example:
	//
	// AliyunLinux
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method for node pool nodes. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The ECS instance metadata access configuration.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance attributes.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The instance types.
	//
	// This parameter is required.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing type for public IP addresses. Valid values:
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth for node public IP addresses. Unit: Mbps (Mega bit per second). Valid values: [1,100].
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The key pair name. Choose either this or `login_password`.
	//
	// example:
	//
	// np-key
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether scaled-out ECS instances use non-root user logon.
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. Choose either this or `key_pair`. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The multi-zone scaling policy for ECS instances in the scaling group. Valid values:
	//
	// - `PRIORITY`: Scales ECS instances based on the virtual switches (VSwitchIds.N) that you define. When ECS instances cannot be created in the zone of the highest-priority vSwitch, the system automatically uses the next-priority vSwitch to create ECS instances.
	//
	// - `COST_OPTIMIZED`: Attempts to create ECS instances in ascending order of vCPU unit price. When the scaling configuration specifies multiple instance types with the spot billing method, spot instances are created first. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically attempt to create pay-as-you-go instances when spot instances cannot be created due to insufficient inventory.
	//
	//   >`COST_OPTIMIZED` takes effect only when multiple instance types are specified in the scaling configuration or spot instances are selected.
	//
	// - `BALANCE`: Evenly distributes ECS instances across the multiple zones specified in the scaling group. If zones become unbalanced due to insufficient inventory, you can call the RebalanceInstances operation to rebalance resources. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html).
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. Valid values: [0,1000]. If the number of pay-as-you-go instances is less than this value, pay-as-you-go instances are created first.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the instances that exceed the minimum number of pay-as-you-go instances (`on_demand_base_capacity`) in the scaling group. Valid values: [0,100].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration for node pool nodes. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`, and is required in that case. Valid values: when `period_unit` is set to Month, valid values for `period` are: {1, 2, 3, 6, 12}.
	//
	// example:
	//
	// 0
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing period unit for node pool nodes. This parameter must be specified when `instance_charge_type` is set to `PrePaid`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// The operating system distribution. Valid values:
	//
	// example:
	//
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The private node pool configuration.
	PrivatePoolOptions *NodepoolScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The RAM role name for worker nodes.
	//
	// 	Notice: This parameter is supported only when you create a node pool for ACK managed clusters of version 1.22 or later.
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// The list of RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy used when creating instances. After you set this parameter, note the following:
	ResourcePoolOptions *NodepoolScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling mode of the scaling group. Valid values:
	//
	// - `release`: standard mode. The scaling group scales in or out by creating or releasing ECS instances based on resource usage.
	//
	// - `recycle`: swift mode. The scaling group scales in or out by creating, stopping, or starting ECS instances, which improves the speed of subsequent scaling operations. Stopped instances are not charged for compute resources but are still charged for storage resources, except for instances that use local disks.
	//
	// Default value: `release`.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The security group ID of the node pool. Use either this parameter or `security_group_ids`. We recommend that you use `security_group_ids`.
	//
	// example:
	//
	// sg-2zeihch86ooz9io4****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The list of security group IDs. Use either this parameter or `security_group_id`. We recommend that you use `security_group_ids`. If both `security_group_id` and `security_group_ids` are specified, `security_group_ids` takes precedence.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// The number of available instance types. The scaling group creates spot instances of multiple types at the lowest cost in a balanced manner. Valid values: [1,10].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable supplementing spot instances. If enabled, when the system receives a notification that a spot instance will be reclaimed, the scaling group attempts to create a new instance to replace the spot instance that will be reclaimed. Valid values:
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price limit configurations for a single spot instance of the current instance type.
	SpotPriceLimit []*NodepoolScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The type of the spot instance. Valid values:
	//
	// - NoSpot: non-spot instance.
	//
	// - SpotWithPriceLimit: sets a maximum price for the spot instance.
	//
	// - SpotAsPriceGo: the system automatically bids at the current market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable burst (performance burst) for the node system cloud disk. Valid values:
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The multiple cloud disk types for the system cloud disk. If the highest-priority cloud disk type is unavailable, the system automatically attempts the next-priority cloud disk type to create the system cloud disk. Valid values:
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the node system cloud disk. Valid values:
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm used for the node system cloud disk. Valid values: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system cloud disk. Valid values: true: Encrypt. false: Do not encrypt.
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The KMS key ID used for the node system cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the system cloud disk for nodes. This parameter takes effect only for ESSDs.
	//
	// - PL0: moderate maximum concurrent I/O performance with relatively stable read/write latency.
	//
	// - PL1: moderate maximum concurrent I/O performance with relatively stable read/write latency.
	//
	// - PL2: high maximum concurrent I/O performance with stable read/write latency.
	//
	// - PL3: ultra-high maximum concurrent I/O performance with extremely stable read/write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS for the node system cloud disk. Valid values: 0~min{50,000, 1000×capacity-baseline performance}. Baseline performance=min{1,800+50×capacity, 50000}.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the node system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// Tags added only to ECS instances.
	Tags []*NodepoolScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The vSwitch IDs.
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

func (s *NodepoolScalingGroup) GetCpuOptions() *NodepoolScalingGroupCpuOptions {
	return s.CpuOptions
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

func (s *NodepoolScalingGroup) GetInstancePatterns() []*InstancePatterns {
	return s.InstancePatterns
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

func (s *NodepoolScalingGroup) SetCpuOptions(v *NodepoolScalingGroupCpuOptions) *NodepoolScalingGroup {
	s.CpuOptions = v
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

func (s *NodepoolScalingGroup) SetInstancePatterns(v []*InstancePatterns) *NodepoolScalingGroup {
	s.InstancePatterns = v
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

type NodepoolScalingGroupCpuOptions struct {
	// Specifies whether to enable nested virtualization. Valid values:
	//
	// example:
	//
	// enabled
	NestedVirtualization *string `json:"nested_virtualization,omitempty" xml:"nested_virtualization,omitempty"`
}

func (s NodepoolScalingGroupCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s NodepoolScalingGroupCpuOptions) GoString() string {
	return s.String()
}

func (s *NodepoolScalingGroupCpuOptions) GetNestedVirtualization() *string {
	return s.NestedVirtualization
}

func (s *NodepoolScalingGroupCpuOptions) SetNestedVirtualization(v string) *NodepoolScalingGroupCpuOptions {
	s.NestedVirtualization = &v
	return s
}

func (s *NodepoolScalingGroupCpuOptions) Validate() error {
	return dara.Validate(s)
}

type NodepoolScalingGroupPrivatePoolOptions struct {
	// The private node pool ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. The private pool capacity option for instance startup. After an elasticity assurance or capacity reservation takes effect, a private pool capacity is generated for instances to select during startup. Valid values:
	//
	// - `Open`: open mode. The system automatically matches open private pool capacity. If no matching private pool capacity is available, public pool resources are used to start the instance.
	//
	// - `Target`: targeted mode. The instance is started using the specified private pool capacity. If the specified private pool capacity is unavailable, the instance fails to start.
	//
	// - `None`: none mode. The instance does not use private pool capacity during startup.
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
	// The list of private pool IDs, which are elasticity assurance IDs or capacity reservation IDs. Only Target mode private pool IDs can be specified. Valid values of N: 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool strategy used when instances are created. Resource pools include private pools generated after Elasticity Assurance or Capacity Reservation takes effect, as well as public pools, for instances to select from during startup. Valid values:
	//
	// PrivatePoolFirst: private pool first. When this strategy is selected, if resource_pool_options.private_pool_ids is specified, the specified private pools are used preferentially. If no private pools are specified or the specified private pools have insufficient capacity, open-type private pools are automatically matched. If no eligible private pools are available, instances are created from the public pool.
	//
	// PrivatePoolOnly: private pool only. When this strategy is selected, resource_pool_options.private_pool_ids must be specified. If the specified private pools have insufficient capacity, instance startup fails.
	//
	// None: no resource pool strategy is used.
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
	// The maximum price for a single instance.
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
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
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
	// Specifies whether this is a confidential computing node pool.
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
