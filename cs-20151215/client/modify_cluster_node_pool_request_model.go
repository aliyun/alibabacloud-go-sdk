// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScaling(v *ModifyClusterNodePoolRequestAutoScaling) *ModifyClusterNodePoolRequest
	GetAutoScaling() *ModifyClusterNodePoolRequestAutoScaling
	SetConcurrency(v bool) *ModifyClusterNodePoolRequest
	GetConcurrency() *bool
	SetKubernetesConfig(v *ModifyClusterNodePoolRequestKubernetesConfig) *ModifyClusterNodePoolRequest
	GetKubernetesConfig() *ModifyClusterNodePoolRequestKubernetesConfig
	SetManagement(v *ModifyClusterNodePoolRequestManagement) *ModifyClusterNodePoolRequest
	GetManagement() *ModifyClusterNodePoolRequestManagement
	SetNodepoolInfo(v *ModifyClusterNodePoolRequestNodepoolInfo) *ModifyClusterNodePoolRequest
	GetNodepoolInfo() *ModifyClusterNodePoolRequestNodepoolInfo
	SetScalingGroup(v *ModifyClusterNodePoolRequestScalingGroup) *ModifyClusterNodePoolRequest
	GetScalingGroup() *ModifyClusterNodePoolRequestScalingGroup
	SetTeeConfig(v *ModifyClusterNodePoolRequestTeeConfig) *ModifyClusterNodePoolRequest
	GetTeeConfig() *ModifyClusterNodePoolRequestTeeConfig
	SetUpdateNodes(v bool) *ModifyClusterNodePoolRequest
	GetUpdateNodes() *bool
}

type ModifyClusterNodePoolRequest struct {
	// The auto scaling configurations.
	AutoScaling *ModifyClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Specifies whether to run the task in parallel.
	//
	// example:
	//
	// true
	Concurrency *bool `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The Kubernetes-related configurations.
	KubernetesConfig *ModifyClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool.
	Management *ModifyClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// The node pool configurations.
	NodepoolInfo *ModifyClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the node pool scaling group.
	ScalingGroup *ModifyClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The configurations of the Kubernetes cluster for confidential computing.
	TeeConfig *ModifyClusterNodePoolRequestTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
	// Synchronously updates node labels and taints.
	//
	// example:
	//
	// true
	UpdateNodes *bool `json:"update_nodes,omitempty" xml:"update_nodes,omitempty"`
}

func (s ModifyClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequest) GetAutoScaling() *ModifyClusterNodePoolRequestAutoScaling {
	return s.AutoScaling
}

func (s *ModifyClusterNodePoolRequest) GetConcurrency() *bool {
	return s.Concurrency
}

func (s *ModifyClusterNodePoolRequest) GetKubernetesConfig() *ModifyClusterNodePoolRequestKubernetesConfig {
	return s.KubernetesConfig
}

func (s *ModifyClusterNodePoolRequest) GetManagement() *ModifyClusterNodePoolRequestManagement {
	return s.Management
}

func (s *ModifyClusterNodePoolRequest) GetNodepoolInfo() *ModifyClusterNodePoolRequestNodepoolInfo {
	return s.NodepoolInfo
}

func (s *ModifyClusterNodePoolRequest) GetScalingGroup() *ModifyClusterNodePoolRequestScalingGroup {
	return s.ScalingGroup
}

func (s *ModifyClusterNodePoolRequest) GetTeeConfig() *ModifyClusterNodePoolRequestTeeConfig {
	return s.TeeConfig
}

func (s *ModifyClusterNodePoolRequest) GetUpdateNodes() *bool {
	return s.UpdateNodes
}

func (s *ModifyClusterNodePoolRequest) SetAutoScaling(v *ModifyClusterNodePoolRequestAutoScaling) *ModifyClusterNodePoolRequest {
	s.AutoScaling = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetConcurrency(v bool) *ModifyClusterNodePoolRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetKubernetesConfig(v *ModifyClusterNodePoolRequestKubernetesConfig) *ModifyClusterNodePoolRequest {
	s.KubernetesConfig = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetManagement(v *ModifyClusterNodePoolRequestManagement) *ModifyClusterNodePoolRequest {
	s.Management = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetNodepoolInfo(v *ModifyClusterNodePoolRequestNodepoolInfo) *ModifyClusterNodePoolRequest {
	s.NodepoolInfo = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetScalingGroup(v *ModifyClusterNodePoolRequestScalingGroup) *ModifyClusterNodePoolRequest {
	s.ScalingGroup = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetTeeConfig(v *ModifyClusterNodePoolRequestTeeConfig) *ModifyClusterNodePoolRequest {
	s.TeeConfig = v
	return s
}

func (s *ModifyClusterNodePoolRequest) SetUpdateNodes(v bool) *ModifyClusterNodePoolRequest {
	s.UpdateNodes = &v
	return s
}

func (s *ModifyClusterNodePoolRequest) Validate() error {
	if s.AutoScaling != nil {
		if err := s.AutoScaling.Validate(); err != nil {
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

type ModifyClusterNodePoolRequestAutoScaling struct {
	// Deprecated
	//
	// This parameter is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The peak bandwidth of the EIP.
	//
	// Valid values: [1, 100]. Unit: Mbit/s.
	//
	// example:
	//
	// null
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
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
	// - `true`: Enables auto scaling for the node pool. If the resources in the cluster do not meet the scheduling requirements of application pods, ACK automatically scales the nodes based on the configured minimum and maximum numbers of instances. For clusters of Kubernetes 1.24 or later, instant elasticity is enabled by default. For clusters of a Kubernetes version earlier than 1.24, node autoscaling is enabled by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: Disables auto scaling. ACK adjusts the number of nodes in the node pool to the value of \\`desired_size\\` and keeps the number of nodes unchanged.
	//
	// If this parameter is set to false, other parameters in `auto_scaling` do not take effect.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
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
	// The maximum number of instances that can be created in the node pool. This parameter does not include existing instances. This parameter takes effect only when `enable=true`.
	//
	// Valid values: [min_instances, 2000]. Default value: 0.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances that can be created in the node pool. This parameter does not include existing instances. This parameter takes effect only when `enable=true`.
	//
	// Valid values: [0, max_instances]. Default value: 0.
	//
	// > - If the minimum number of instances is not 0, the specified number of ECS instances are automatically created after the scaling group is created.
	//
	// >
	//
	// > - Set the maximum number of instances to a value that is not smaller than the current number of nodes in the node pool. Otherwise, a scale-in event is triggered after auto scaling is enabled.
	//
	// example:
	//
	// 2
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// Deprecated
	//
	// The type of auto scaling. This parameter is specified based on the instance type. Valid values:
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
	// example:
	//
	// cpu
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyClusterNodePoolRequestAutoScaling) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestAutoScaling) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetEipBandwidth() *int64 {
	return s.EipBandwidth
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetEipInternetChargeType() *string {
	return s.EipInternetChargeType
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetIsBondEip() *bool {
	return s.IsBondEip
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetMaxInstances() *int64 {
	return s.MaxInstances
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *ModifyClusterNodePoolRequestAutoScaling) GetType() *string {
	return s.Type
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetEipBandwidth(v int64) *ModifyClusterNodePoolRequestAutoScaling {
	s.EipBandwidth = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetEipInternetChargeType(v string) *ModifyClusterNodePoolRequestAutoScaling {
	s.EipInternetChargeType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetEnable(v bool) *ModifyClusterNodePoolRequestAutoScaling {
	s.Enable = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetIsBondEip(v bool) *ModifyClusterNodePoolRequestAutoScaling {
	s.IsBondEip = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetMaxInstances(v int64) *ModifyClusterNodePoolRequestAutoScaling {
	s.MaxInstances = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetMinInstances(v int64) *ModifyClusterNodePoolRequestAutoScaling {
	s.MinInstances = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) SetType(v string) *ModifyClusterNodePoolRequestAutoScaling {
	s.Type = &v
	return s
}

func (s *ModifyClusterNodePoolRequestAutoScaling) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestKubernetesConfig struct {
	// Specifies whether to install Cloud Monitor on the ECS nodes. After Cloud Monitor is installed, you can view the monitoring information of the created ECS instances in the Cloud Monitor console. We recommend that you enable this feature. Valid values:
	//
	// - `true`: installs Cloud Monitor on ECS nodes.
	//
	// - `false`: does not install Cloud Monitor on ECS nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of the node. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// - `static`: allows pods with specific resource characteristics on the node to be granted with enhanced CPU affinity and exclusivity.
	//
	// - `none`: indicates that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels that you want to add to the nodes. The following rules apply:
	//
	// - A label is a case-sensitive key-value pair. You can add up to 20 labels.
	//
	// - The key must be unique and can be up to 64 characters in length. The value can be empty and can be up to 128 characters in length. The key and the value cannot start with `aliyun`, `acs:`, `https://`, or `http://`. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name parameter. A node name consists of three parts: a prefix, the node IP address, and a suffix.
	//
	// The prefix and suffix can contain one or more parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). The node name must start and end with a lowercase letter or a digit. The node IP address is the complete private IP address of the node.
	//
	// The parameter consists of four parts that are separated by commas. For example, if you pass the "customized,aliyun,ip,com" string (where "customized" and "ip" are fixed strings, "aliyun" is the prefix, and "com" is the suffix), the node name is aliyun.192.168.xxx.xxx.com.
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-customized instance data. Before a node is added to the cluster, the specified pre-customized instance data script is run. For more information, see [User data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. ACK supports the following three container runtimes.
	//
	// - containerd: We recommend that you use this runtime. It is supported by all cluster versions.
	//
	// - Sandboxed-Container.runv: a sandboxed container that provides higher isolation. It is supported by clusters of Kubernetes 1.31 and earlier.
	//
	// - docker: This runtime is no longer maintained. It is supported by clusters of Kubernetes 1.22 and earlier.
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
	// The node taint configurations.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Specifies whether the scaled-out nodes are unschedulable.
	//
	// - true: The nodes are unschedulable.
	//
	// - false: The nodes are schedulable.
	//
	// example:
	//
	// false
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After a node is added to the cluster, the specified user data script is run. For more information, see [User data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	UserData *string `json:"user_data,omitempty" xml:"user_data,omitempty"`
}

func (s ModifyClusterNodePoolRequestKubernetesConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestKubernetesConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetCmsEnabled() *bool {
	return s.CmsEnabled
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetCpuPolicy() *string {
	return s.CpuPolicy
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetLabels() []*Tag {
	return s.Labels
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetNodeNameMode() *string {
	return s.NodeNameMode
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetPreUserData() *string {
	return s.PreUserData
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetTaints() []*Taint {
	return s.Taints
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetUnschedulable() *bool {
	return s.Unschedulable
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) GetUserData() *string {
	return s.UserData
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetCmsEnabled(v bool) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.CmsEnabled = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetCpuPolicy(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.CpuPolicy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetLabels(v []*Tag) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.Labels = v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetNodeNameMode(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.NodeNameMode = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetPreUserData(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.PreUserData = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetRuntime(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.Runtime = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetRuntimeVersion(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetTaints(v []*Taint) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.Taints = v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetUnschedulable(v bool) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.Unschedulable = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) SetUserData(v string) *ModifyClusterNodePoolRequestKubernetesConfig {
	s.UserData = &v
	return s
}

func (s *ModifyClusterNodePoolRequestKubernetesConfig) Validate() error {
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

type ModifyClusterNodePoolRequestManagement struct {
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to enable auto node repair. This parameter takes effect only when enable is set to `true`.
	//
	// - `true`: Auto repair is enabled.
	//
	// - `false`: Auto repair is disabled.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto node repair policy.
	AutoRepairPolicy *ModifyClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to enable auto node upgrade. This parameter takes effect only when enable is set to `true`.
	//
	// - `true`: enables auto upgrade.
	//
	// - `false`: disables auto upgrade.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto upgrade policy.
	AutoUpgradePolicy *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically fix CVE vulnerabilities. This parameter takes effect only when enable is set to `true`.
	//
	// - `true`: allows automatic CVE fixing.
	//
	// - `false`: disallows automatic CVE fixing.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The policy for automatically fixing CVE vulnerabilities.
	AutoVulFixPolicy *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable the managed node pool. Valid values:
	//
	// - `true`: Enables the managed node pool.
	//
	// - `false`: Disables the managed node pool. Other related configurations are ignored.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `auto_upgrade` parameter at the upper level instead.
	//
	// The auto upgrade configurations. This parameter takes effect only when enable is set to `true`.
	UpgradeConfig *ModifyClusterNodePoolRequestManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s ModifyClusterNodePoolRequestManagement) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestManagement) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoFaultDiagnosis() *bool {
	return s.AutoFaultDiagnosis
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoRepair() *bool {
	return s.AutoRepair
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoRepairPolicy() *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	return s.AutoRepairPolicy
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoUpgradePolicy() *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy {
	return s.AutoUpgradePolicy
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoVulFix() *bool {
	return s.AutoVulFix
}

func (s *ModifyClusterNodePoolRequestManagement) GetAutoVulFixPolicy() *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy {
	return s.AutoVulFixPolicy
}

func (s *ModifyClusterNodePoolRequestManagement) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyClusterNodePoolRequestManagement) GetUpgradeConfig() *ModifyClusterNodePoolRequestManagementUpgradeConfig {
	return s.UpgradeConfig
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoFaultDiagnosis(v bool) *ModifyClusterNodePoolRequestManagement {
	s.AutoFaultDiagnosis = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoRepair(v bool) *ModifyClusterNodePoolRequestManagement {
	s.AutoRepair = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoRepairPolicy(v *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) *ModifyClusterNodePoolRequestManagement {
	s.AutoRepairPolicy = v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoUpgrade(v bool) *ModifyClusterNodePoolRequestManagement {
	s.AutoUpgrade = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoUpgradePolicy(v *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) *ModifyClusterNodePoolRequestManagement {
	s.AutoUpgradePolicy = v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoVulFix(v bool) *ModifyClusterNodePoolRequestManagement {
	s.AutoVulFix = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetAutoVulFixPolicy(v *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) *ModifyClusterNodePoolRequestManagement {
	s.AutoVulFixPolicy = v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetEnable(v bool) *ModifyClusterNodePoolRequestManagement {
	s.Enable = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) SetUpgradeConfig(v *ModifyClusterNodePoolRequestManagementUpgradeConfig) *ModifyClusterNodePoolRequestManagement {
	s.UpgradeConfig = v
	return s
}

func (s *ModifyClusterNodePoolRequestManagement) Validate() error {
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

type ModifyClusterNodePoolRequestManagementAutoRepairPolicy struct {
	// Specifies whether manual approval is required for node repair.
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// The ID of the auto repair policy.
	//
	// example:
	//
	// r-xxxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only when auto_repair is set to `true`. Valid values:
	//
	// - `true`: allows node restart.
	//
	// - `false`: disallows node restart.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
}

func (s ModifyClusterNodePoolRequestManagementAutoRepairPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GetApprovalRequired() *bool {
	return s.ApprovalRequired
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GetAutoRepairPolicyId() *string {
	return s.AutoRepairPolicyId
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) SetApprovalRequired(v bool) *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) SetAutoRepairPolicyId(v string) *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	s.AutoRepairPolicyId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) SetRestartNode(v bool) *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	s.RestartNode = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestManagementAutoUpgradePolicy struct {
	// Specifies whether to allow auto kubelet upgrade. This parameter takes effect only when auto_upgrade is set to `true`. Valid values:
	//
	// - `true`: allows auto kubelet upgrade.
	//
	// - `false`: disallows auto kubelet upgrade.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow auto operating system upgrade. This parameter takes effect only when auto_upgrade is set to `true`. Valid values:
	//
	// - `true`: allows auto operating system upgrade.
	//
	// - `false`: disallows auto operating system upgrade.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow auto runtime upgrade. This parameter takes effect only when auto_upgrade is set to `true`. Valid values:
	//
	// - `true`: allows auto runtime upgrade.
	//
	// - `false`: disallows auto runtime upgrade.
	//
	// Default value: `true`.
	//
	// example:
	//
	// false
	AutoUpgradeRuntime *bool `json:"auto_upgrade_runtime,omitempty" xml:"auto_upgrade_runtime,omitempty"`
}

func (s ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeKubelet() *bool {
	return s.AutoUpgradeKubelet
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeOs() *bool {
	return s.AutoUpgradeOs
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) GetAutoUpgradeRuntime() *bool {
	return s.AutoUpgradeRuntime
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeKubelet(v bool) *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeKubelet = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeOs(v bool) *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeOs = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) SetAutoUpgradeRuntime(v bool) *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy {
	s.AutoUpgradeRuntime = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestManagementAutoVulFixPolicy struct {
	// The packages that should be excluded during vulnerability fixing.
	//
	// Default value: `kernel`.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only when auto_vul_fix is set to `true`. Valid values:
	//
	// - `true`: allows node restart.
	//
	// - `false`: disallows node restart.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels that are allowed to be automatically fixed. The value is a comma-separated list. Example: `asap,later`. Supported vulnerability levels:
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

func (s ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) GetExcludePackages() *string {
	return s.ExcludePackages
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) GetVulLevel() *string {
	return s.VulLevel
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) SetExcludePackages(v string) *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.ExcludePackages = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) SetRestartNode(v bool) *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.RestartNode = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) SetVulLevel(v string) *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy {
	s.VulLevel = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestManagementUpgradeConfig struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the `auto_upgrade` parameter at the upper level instead.
	//
	// Specifies whether to enable auto upgrade:
	//
	// - true: enables auto upgrade.
	//
	// - false: disables auto upgrade.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes.
	//
	// Valid values: [1, 1000]
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes. You can specify only one of surge and `surge_percentage`.
	//
	// Nodes may become unavailable during an upgrade. You can create extra nodes to ensure service continuity.
	//
	// > The number of extra nodes must not exceed the current number of nodes.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You can specify only one of surge and `surge_percentage`.
	//
	// Number of extra nodes = Percentage of extra nodes × Number of nodes. For example, if you set the percentage of extra nodes to 50% and the number of existing nodes is 6, three extra nodes are created.
	//
	// example:
	//
	// 0
	SurgePercentage *int64 `json:"surge_percentage,omitempty" xml:"surge_percentage,omitempty"`
}

func (s ModifyClusterNodePoolRequestManagementUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestManagementUpgradeConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) GetAutoUpgrade() *bool {
	return s.AutoUpgrade
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) GetMaxUnavailable() *int64 {
	return s.MaxUnavailable
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) GetSurge() *int64 {
	return s.Surge
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) GetSurgePercentage() *int64 {
	return s.SurgePercentage
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) SetAutoUpgrade(v bool) *ModifyClusterNodePoolRequestManagementUpgradeConfig {
	s.AutoUpgrade = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) SetMaxUnavailable(v int64) *ModifyClusterNodePoolRequestManagementUpgradeConfig {
	s.MaxUnavailable = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) SetSurge(v int64) *ModifyClusterNodePoolRequestManagementUpgradeConfig {
	s.Surge = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) SetSurgePercentage(v int64) *ModifyClusterNodePoolRequestManagementUpgradeConfig {
	s.SurgePercentage = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestNodepoolInfo struct {
	// The name of the node pool.
	//
	// The name must be 1 to 63 characters in length, and can contain digits, letters, and hyphens (-). It cannot start with a hyphen (-).
	//
	// example:
	//
	// default-nodepool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group for the node pool. Instances created in the node pool belong to this resource group.
	//
	// A resource can belong to only one resource group. You can use resource groups to categorize resources by project, application, or organization.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
}

func (s ModifyClusterNodePoolRequestNodepoolInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestNodepoolInfo) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) GetName() *string {
	return s.Name
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) SetName(v string) *ModifyClusterNodePoolRequestNodepoolInfo {
	s.Name = &v
	return s
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) SetResourceGroupId(v string) *ModifyClusterNodePoolRequestNodepoolInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestNodepoolInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestScalingGroup struct {
	// Specifies whether to enable auto-renewal for the nodes. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// - `true`: Auto-renewal is enabled.
	//
	// - `false`: Auto-renewal is disabled.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// - If PeriodUnit=Week, valid values are 1, 2, and 3.
	//
	// - If PeriodUnit=Month, valid values are 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// If `multi_az_policy` is set to `COST_OPTIMIZED`, this parameter specifies whether to allow the system to automatically create pay-as-you-go instances to meet the required number of ECS instances when spot instances cannot be created due to reasons such as price and stock. Valid values:
	//
	// - `true`: allows the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: does not allow the system to automatically create pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The data disk configurations of the node. You can specify 0 to 10 data disks.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The ID of the deployment set to which the ECS instances in the node pool belong. This parameter is valid only for incremental nodes. The deployment sets of existing nodes are not changed.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The expected number of nodes in the node pool.
	//
	// The total number of nodes that the node pool should maintain. We recommend that you configure at least two nodes to ensure that the cluster components run as expected. You can adjust the expected number of nodes to scale in or scale out the node pool.
	//
	// If you do not need to create nodes, set this parameter to 0. You can manually adjust the number of nodes later.
	//
	// example:
	//
	// 2
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The ID of the custom image. You can call `DescribeKubernetesVersionMetadata` to query the images that are supported by the system. By default, the latest image is used.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The distribution of the operating system. We recommend that you use this parameter to specify the operating system of the nodes. Valid values:
	//
	// - `AliyunLinux`: Alinux2 image.
	//
	// - `AliyunLinuxSecurity`: Alinux2 image with UEFI.
	//
	// - `AliyunLinux3`: Alinux3 image.
	//
	// - `AliyunLinux3Arm64`: Alinux3 image for ARM.
	//
	// - `AliyunLinux3Security`: Alinux3 image with UEFI.
	//
	// - `CentOS`: CentOS image.
	//
	// - `Windows`: Windows image.
	//
	// - `WindowsCore`: WindowsCore image.
	//
	// - `ContainerOS`: container-optimized image.
	//
	// - `AliyunLinux3ContainerOptimized`: container-optimized Alinux3 image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of the nodes in the node pool. Valid values:
	//
	// - `PrePaid`: subscription
	//
	// - `PostPaid`: pay-as-you-go
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The instance attribute configurations.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// A list of node instance types. You can specify multiple instance types as alternatives. When a node is created, the system attempts to purchase the instance types in the order they are specified until one is successfully purchased. The final instance type that is purchased may vary depending on the stock.
	//
	// You can specify 1 to 10 instance types.
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
	// The maximum outbound bandwidth of the public IP address of the node. Unit: Mbit/s. Valid values: [1, 100].
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair. You must specify key_pair or `login_password`. For managed node pools, you can specify only `key_pair`.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. You must specify key_pair or `login_password`. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: The system scales ECS instances based on the vSwitch priority. The vSwitch priority is specified by the VSwitchIds.N parameter. If the system fails to create an ECS instance in the zone where the vSwitch with the highest priority resides, it attempts to create the ECS instance in the zone where the vSwitch with the next highest priority resides.
	//
	// - `COST_OPTIMIZED`: The system creates ECS instances of the instance type that has the lowest vCPU price. When the scaling configuration is set to create multiple instance types and the billing method is set to preemptible, the system preferentially creates preemptible instances. You can also use the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when the system fails to create preemptible instances due to insufficient stock.
	//
	//   > The `COST_OPTIMIZED` policy is valid only when multiple instance types are specified or preemptible instances are selected in the scaling configuration.
	//
	// - `BALANCE`: The system evenly distributes ECS instances across the zones that are specified in the scaling group. If the distribution of ECS instances becomes unbalanced due to insufficient stock, you can call the `RebalanceInstances` operation to rebalance the distribution. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) .
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// BALANCE
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of on-demand instances that must be contained in the scaling group. Valid values: [0, 1000]. When the number of on-demand instances is less than this value, on-demand instances are preferentially created.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of on-demand instances among the instances that exceed the minimum number of on-demand instances (`on_demand_base_capacity`). Valid values: [0, 100].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of the nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`.
	//
	// - If `period_unit=Week`, valid values of `period` are {1, 2, 3, 4}.
	//
	// - If `period_unit=Month`, valid values of `period` are {1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60}.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`.
	//
	// - `Month`: billed by month.
	//
	// - `Week`: billed by week.
	//
	// Default value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `image_type` parameter instead.
	//
	// The OS platform. Valid values:
	//
	// - `AliyunLinux`
	//
	// - `CentOS`
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
	PrivatePoolOptions *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// A list of ApsaraDB RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy that are used when you create an instance. If you specify this parameter, note the following points:
	//
	// This parameter is valid only when you create pay-as-you-go instances.
	//
	// This parameter cannot be specified at the same time as private_pool_options.match_criteria and private_pool_options.id.
	ResourcePoolOptions *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling mode of the scaling group. Valid values:
	//
	// - `release`: standard mode. This mode creates and releases ECS instances to perform scaling.
	//
	// - `recycle`: fast mode. This mode creates, stops, and starts ECS instances to perform scaling. This improves scaling speed. When an instance is stopped, its computing resources are not billed, but its storage resources are. This does not apply to instance types with local disks.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// A list of security group IDs.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// The number of available instance types. The scaling group creates spot instances of multiple instance types that are provided at the lowest cost. Valid values: [1, 10].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable the feature of supplementing spot instances. If this feature is enabled, the scaling group attempts to create a new instance to replace a spot instance that is reclaimed. Valid values:
	//
	// - `true`: enables the feature of supplementing spot instances.
	//
	// - `false`: disables the feature of supplementing spot instances.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price range for the spot instance.
	SpotPriceLimit []*ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The bidding policy for the spot instance. Valid values:
	//
	// - `NoSpot`: The instance is not a spot instance.
	//
	// - `SpotWithPriceLimit`: The instance is a spot instance for which you can specify the maximum hourly price.
	//
	// - `SpotAsPriceGo`: The system automatically bids for the instance. The bid is based on the market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// - true: enables the performance burst feature. If your business fluctuates and is subject to unexpected data read and write pressure, the cloud disk temporarily improves its performance until your business returns to a stable state.
	//
	// - false: disables the performance burst feature.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The multi-disk type of the system disk. If a disk of a higher priority disk type cannot be used, the system automatically tries the next priority disk type to create the system disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system disk. Valid values:
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
	// The encryption algorithm that is used for the system disk. Valid value: aes-256.
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
	// The performance level of the system disk. This parameter is valid only for ESSD disks. The performance level of a disk is related to its size. For more information, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: The maximum concurrent I/O performance is moderate and the read/write latency is relatively stable.
	//
	// - PL1: The maximum concurrent I/O performance is moderate and the read/write latency is relatively stable.
	//
	// - PL2: The maximum concurrent I/O performance is high and the read/write latency is stable.
	//
	// - PL3: The maximum concurrent I/O performance is very high and the read/write latency is very stable.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The pre-configured read/write IOPS of the system disk.
	//
	// Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is supported only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// Valid values: [20, 2048].
	//
	// The value of this parameter must be greater than or equal to max{20, ImageSize}.
	//
	// Default value: max{40, The size of the image that corresponds to the ImageId parameter}.
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
	// Tag keys cannot be repeated. A tag key can be up to 128 characters in length. Tag keys and tag values cannot start with "aliyun" or "acs:" and cannot contain "https\\://" or "http\\://".
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// A list of vSwitch IDs. You can specify 1 to 8 vSwitch IDs.
	//
	// > For high availability, select vSwitches in different zones.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s ModifyClusterNodePoolRequestScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroup) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetDeploymentsetId() *string {
	return s.DeploymentsetId
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetDesiredSize() *int64 {
	return s.DesiredSize
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetDiskInit() []*DiskInit {
	return s.DiskInit
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetImageType() *string {
	return s.ImageType
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInstancePatterns() []*InstancePatterns {
	return s.InstancePatterns
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetKeyPair() *string {
	return s.KeyPair
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetOnDemandBaseCapacity() *int64 {
	return s.OnDemandBaseCapacity
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int64 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetPeriod() *int64 {
	return s.Period
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetPlatform() *string {
	return s.Platform
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetPrivatePoolOptions() *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetRdsInstances() []*string {
	return s.RdsInstances
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetResourcePoolOptions() *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSpotInstancePools() *int64 {
	return s.SpotInstancePools
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSpotPriceLimit() []*ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit {
	return s.SpotPriceLimit
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskKmsKeyId() *string {
	return s.SystemDiskKmsKeyId
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetSystemDiskSnapshotPolicyId() *string {
	return s.SystemDiskSnapshotPolicyId
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetTags() []*Tag {
	return s.Tags
}

func (s *ModifyClusterNodePoolRequestScalingGroup) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetAutoRenew(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.AutoRenew = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetAutoRenewPeriod(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.AutoRenewPeriod = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetCompensateWithOnDemand(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetDataDisks(v []*DataDisk) *ModifyClusterNodePoolRequestScalingGroup {
	s.DataDisks = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetDeploymentsetId(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.DeploymentsetId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetDesiredSize(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.DesiredSize = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetDiskInit(v []*DiskInit) *ModifyClusterNodePoolRequestScalingGroup {
	s.DiskInit = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetImageId(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.ImageId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetImageType(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.ImageType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInstanceChargeType(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInstancePatterns(v []*InstancePatterns) *ModifyClusterNodePoolRequestScalingGroup {
	s.InstancePatterns = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInstanceTypes(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.InstanceTypes = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInternetChargeType(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetInternetMaxBandwidthOut(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetKeyPair(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.KeyPair = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetLoginPassword(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.LoginPassword = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetMultiAzPolicy(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.MultiAzPolicy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetOnDemandBaseCapacity(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetPeriod(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.Period = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetPeriodUnit(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetPlatform(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.Platform = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetPrivatePoolOptions(v *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) *ModifyClusterNodePoolRequestScalingGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetRdsInstances(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.RdsInstances = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetResourcePoolOptions(v *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) *ModifyClusterNodePoolRequestScalingGroup {
	s.ResourcePoolOptions = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetScalingPolicy(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSecurityGroupIds(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSpotInstancePools(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSpotInstanceRemedy(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSpotPriceLimit(v []*ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) *ModifyClusterNodePoolRequestScalingGroup {
	s.SpotPriceLimit = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSpotStrategy(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskBurstingEnabled(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskCategories(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategories = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskCategory(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskCategory = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskEncryptAlgorithm(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskEncrypted(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskKmsKeyId(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskKmsKeyId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskPerformanceLevel(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskProvisionedIops(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskSize(v int64) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetSystemDiskSnapshotPolicyId(v string) *ModifyClusterNodePoolRequestScalingGroup {
	s.SystemDiskSnapshotPolicyId = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetTags(v []*Tag) *ModifyClusterNodePoolRequestScalingGroup {
	s.Tags = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) SetVswitchIds(v []*string) *ModifyClusterNodePoolRequestScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroup) Validate() error {
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

type ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions struct {
	// The ID of the private node pool. If `match_criteria` is set to `Target`, you must specify the ID of the private pool.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. This parameter specifies the private pool capacity option for instance startup. After an elastic assurance service or a capacity reservation service takes effect, it generates a private pool of capacity for instance startup. Valid values:
	//
	// - `Open`: Open mode. The system automatically matches the capacity of private pools in Open mode. If no matching private pool is found, the instance is started using public pool resources.
	//
	// - `Target`: Specified mode. The instance is started using the capacity of a specified private pool. If the capacity of the specified private pool is unavailable, the instance fails to be started.
	//
	// - `None`: The instance is started without using the capacity of a private pool.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"match_criteria,omitempty" xml:"match_criteria,omitempty"`
}

func (s ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) SetId(v string) *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) SetMatchCriteria(v string) *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions struct {
	// A list of private pool IDs. The IDs of elastic assurance services or capacity reservation services. You can specify only the IDs of private pools in Target mode. You can specify 1 to 20 IDs.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy that is used when you create an instance. Resource pools include private pools that are generated after elastic assurance services or capacity reservation services take effect, and public pools. You can select a resource pool when you start an instance. Valid values:
	//
	// PrivatePoolFirst: Private pool first. If you select this policy and specify resouce_pool_options.private_pool_ids, the specified private pool is used first. If you do not specify a private pool or the capacity of the specified private pool is insufficient, the system automatically matches a private pool in Open mode. If no matching private pool is found, a public pool is used to create the instance.
	//
	// PrivatePoolOnly: Private pool only. If you select this policy, you must specify resouce_pool_options.private_pool_ids. If the capacity of the specified private pool is insufficient, the instance fails to be started.
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

func (s ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) SetPrivatePoolIds(v []*string) *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) SetStrategy(v string) *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit struct {
	// The instance type of the spot instance.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum price of a single instance.
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

func (s ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) SetInstanceType(v string) *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) SetPriceLimit(v string) *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestTeeConfig struct {
	// Specifies whether to enable the confidential computing cluster. Valid values:
	//
	// - `true`: enables the cluster.
	//
	// - `false`: disables the cluster.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	TeeEnable *bool `json:"tee_enable,omitempty" xml:"tee_enable,omitempty"`
}

func (s ModifyClusterNodePoolRequestTeeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestTeeConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestTeeConfig) GetTeeEnable() *bool {
	return s.TeeEnable
}

func (s *ModifyClusterNodePoolRequestTeeConfig) SetTeeEnable(v bool) *ModifyClusterNodePoolRequestTeeConfig {
	s.TeeEnable = &v
	return s
}

func (s *ModifyClusterNodePoolRequestTeeConfig) Validate() error {
	return dara.Validate(s)
}
