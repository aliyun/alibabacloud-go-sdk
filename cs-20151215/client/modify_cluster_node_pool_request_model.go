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
	// The auto scaling configuration.
	AutoScaling *ModifyClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Whether to run concurrently.
	//
	// example:
	//
	// true
	Concurrency *bool `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The cluster-related configuration.
	KubernetesConfig *ModifyClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The managed node pool configuration.
	Management *ModifyClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// The node pool configuration.
	NodepoolInfo *ModifyClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// Node pool scaling group configuration.
	ScalingGroup *ModifyClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The configuration of the Kubernetes cluster for confidential computing.
	TeeConfig *ModifyClusterNodePoolRequestTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
	// Synchronously update node labels and taints.
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
	// This field is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The valid values are 1 to 100. The unit is Mbps.
	//
	// example:
	//
	// null
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// [Deprecated] Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The billing method for the EIP. Valid values:
	//
	// - `PayByBandwidth`: Pay by bandwidth.
	//
	// - `PayByTraffic`: Pay by traffic.
	//
	// Default value: `PayByBandwidth`.
	//
	// example:
	//
	// null
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Enable auto scaling. Valid values:
	//
	// - `true`: Enable auto scaling. When cluster capacity is insufficient for pod scheduling, ACK automatically scales nodes based on the minimum and maximum instance counts you configure. For clusters running Kubernetes 1.24 or later, instant elasticity is enabled by default. For earlier versions, auto scaling is enabled by default. For more information, see [Node scaling](https://help.aliyun.com/document_detail/2746785.html).
	//
	// - `false`: Disable auto scaling. ACK adjusts the number of nodes in the node pool to match the desired node count you specify.
	//
	// When set to false, other parameters under `auto_scaling` are ignored.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// [Deprecated] Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// - `true`: Bind an EIP.
	//
	// - `false`: Do not bind an EIP.
	//
	// Default value: `false`.
	//
	// example:
	//
	// null
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of scalable instances in the node pool, excluding existing instances. This parameter takes effect only when `enable=true`.
	//
	// Valid values: [min_instances, 2000]. Default value: 0.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of scalable instances in the node pool, excluding existing instances. This parameter takes effect only when `enable=true`.
	//
	// Valid values: [0, max_instances]. Default value: 0.
	//
	// > - If the minimum instance count is not zero, the system creates the specified number of ECS instances after the scaling group becomes active.
	//
	// >
	//
	// > - To avoid unintended scale-in, set the maximum instance count to at least the current number of nodes in the node pool.
	//
	// example:
	//
	// 2
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// Deprecated
	//
	// The auto scaling type, defined by instance type. Valid values:
	//
	// - `cpu`: Standard instance.
	//
	// - `gpu`: GPU instance.
	//
	// - `gpushare`: GPU shared instance.
	//
	// - `spot`: Spot instance.
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
	// Install Cloud Monitor on ECS nodes. After installation, you can view monitoring information for the created ECS instances in the Cloud Monitor console. We recommend enabling this. Valid values:
	//
	// - `true`: Install Cloud Monitor on ECS nodes.
	//
	// - `false`: Do not install Cloud Monitor on ECS nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The node CPU management policy. The following two policies are supported for clusters of version 1.12.6 and later:
	//
	// - `static`: Allows pods with certain resource characteristics on the node to have enhanced CPU affinity and exclusivity.
	//
	// - `none`: Enables the existing default CPU affinity scheme.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// Node labels. Add labels to Kubernetes cluster nodes. Label rules:
	//
	// - Labels are case-sensitive key-value pairs. You can set up to 20 labels.
	//
	// - Tag keys must be unique and up to 64 characters long. Tag values can be empty and up to 128 characters long. Neither tag keys nor tag values can start with `aliyun`, `acs:`, `https://`, or `http://`. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name parameter. A node name consists of three parts: prefix + node IP + suffix.
	//
	// The prefix and suffix can each consist of one or more parts separated by periods (.). Each part can contain lowercase letters, numbers, and hyphens (-). The node name must start and end with a lowercase letter or a number. The node IP is the complete private IP address of the node.
	//
	// The parameter is passed as four parts separated by commas. For example, if you pass the string "customized,aliyun,ip,com" (where "customized" and "ip" are fixed strings, "aliyun" is the prefix, and "com" is the suffix), the node name will be: aliyun.192.168.xxx.xxx.com.
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The instance pre-customization data. Before a node joins the cluster, it runs the specified pre-customization data script. For more information, see [User-Data script](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The container runtime name. ACK supports the following three container runtimes.
	//
	// - containerd: Recommended. Supports all cluster versions.
	//
	// - Sandboxed-Container.runv: Sandboxed container. Provides higher isolation. Supports clusters of version 1.31 and earlier.
	//
	// - docker: No longer maintained. Supports clusters of version 1.22 and earlier.
	//
	// Default value: containerd.
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
	// The node taint configuration.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Whether scaled-out nodes are unschedulable.
	//
	// - true: Unschedulable.
	//
	// - false: Schedulable.
	//
	// example:
	//
	// false
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After a node joins the cluster, it runs the specified user data script. For more information, see [User-Data script](https://help.aliyun.com/document_detail/49121.html).
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
	// Enable automatic node repair. This parameter takes effect only when `enable=true`.
	//
	// - `true`: Enable automatic repair.
	//
	// - `false`: Disable automatic repair.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The automatic node repair policy.
	AutoRepairPolicy *ModifyClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Enable automatic node upgrades. This parameter takes effect only when `enable=true`. Valid values:
	//
	// - `true`: Enable automatic upgrades.
	//
	// - `false`: Disable automatic upgrades.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The automatic upgrade policy.
	AutoUpgradePolicy *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Enable automatic CVE vulnerability fixes. This parameter takes effect only when `enable=true`. Valid values:
	//
	// - `true`: Enable automatic CVE fixes.
	//
	// - `false`: Disable automatic CVE fixes.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The automatic CVE fix policy.
	AutoVulFixPolicy *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Enable the managed node pool. Valid values:
	//
	// - `true`: Enable the managed node pool.
	//
	// - `false`: Disable the managed node pool. Other related configurations take effect only when `enable=true`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// [Deprecated] Use the top-level `auto_upgrade` parameter instead.
	//
	// The automatic upgrade configuration. This parameter takes effect only when `enable=true`.
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
	// Require manual approval for node repair.
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// The automatic node repair policy ID.
	//
	// example:
	//
	// r-xxxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// Allow node restart. This parameter takes effect only when `auto_repair=true`. Valid values:
	//
	// - `true`: Allow node restart.
	//
	// - `false`: Disallow node restart.
	//
	// Default value: `true`.
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
	// Enable automatic kubelet upgrades. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// - `true`: Enable automatic kubelet upgrades.
	//
	// - `false`: Disable automatic kubelet upgrades.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Enable automatic operating system upgrades. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// - `true`: Enable automatic OS upgrades.
	//
	// - `false`: Disable automatic OS upgrades.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Enable automatic runtime upgrades. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// - `true`: Enable automatic runtime upgrades.
	//
	// - `false`: Disable automatic runtime upgrades.
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
	// Packages to exclude during vulnerability fixes.
	//
	// Default value: `kernel`.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Allow node restart. This parameter takes effect only when `auto_vul_fix=true`. Valid values:
	//
	// - `true`: Allow node restart.
	//
	// - `false`: Disallow node restart.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The severity levels of vulnerabilities that can be fixed automatically, separated by commas. Example: `asap,later`. Supported levels:
	//
	// - `asap`: High
	//
	// - `later`: Medium
	//
	// - `nntf`: Low
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
	// [Deprecated] Use the top-level `auto_upgrade` parameter instead.
	//
	// Enable automatic upgrades. Valid values:
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes.
	//
	// Valid values: [1, 1000].
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of additional nodes. Choose either surge or `surge_percentage`.
	//
	// Nodes become unavailable during upgrades. Create extra nodes to maintain cluster load.
	//
	// > Do not create more extra nodes than the current number of nodes.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of additional nodes. Choose either surge or surge_percentage. The number of additional nodes equals `surge_percentage` × the number of nodes. For example, if `surge_percentage` is 50% and there are 6 nodes, then 3 additional nodes are created.
	//
	// The number of additional nodes = The percentage of additional nodes × The number of nodes in the node pool. For example, if the percentage of additional nodes is 50% and the number of nodes in the node pool is 6, the number of additional nodes is 3.
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
	// The node pool name.
	//
	// Name rules: Use only numbers, letters, or hyphens (-). The name must be 1 to 63 characters long and cannot start with a hyphen (-).
	//
	// example:
	//
	// default-nodepool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource group ID for the node pool. Nodes created from this node pool belong to this resource group.
	//
	// A resource can belong to only one resource group. Map resource groups to projects, applications, or organizations based on your business needs.
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
	// Enable automatic renewal for nodes. This parameter takes effect only when `instance_charge_type` is `PrePaid`. Valid values:
	//
	// - `true`: Enable automatic renewal.
	//
	// - `false`: Disable automatic renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The duration of each automatic renewal. Valid values:
	//
	// - When PeriodUnit=Week: 1, 2, 3.
	//
	// - When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// When `multi_az_policy` is set to `COST_OPTIMIZED`, this parameter specifies whether to automatically create on-demand instances to meet the required number of ECS instances if enough spot instances cannot be created due to price or inventory issues. Valid values:
	//
	// - `true`: Allow automatic creation of on-demand instances.
	//
	// - `false`: Disallow automatic creation of on-demand instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The data disk configuration for nodes. Valid values: [0, 10]. You can attach up to 10 data disks.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The deployment set to which the ECS instances created in the node pool belong. This setting applies only to new nodes. The deployment set of existing nodes is not changed.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The desired number of nodes in the node pool.
	//
	// This is the total number of nodes the node pool should maintain. Configure at least 2 nodes to ensure cluster components run properly. Adjust this number to scale the node pool in or out.
	//
	// If you do not need to create nodes, set this to 0 and adjust it manually later.
	//
	// example:
	//
	// 2
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// The custom image ID. Query supported images using `DescribeKubernetesVersionMetadata`. By default, the latest system image is used.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The operating system distribution type. We recommend using this field to specify the node OS. Valid values:
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
	// - `WindowsCore`: Windows Core image.
	//
	// - `ContainerOS`: Container-optimized image.
	//
	// - `AliyunLinux3ContainerOptimized`: Alibaba Cloud Linux 3 container-optimized image.
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method for nodes in the node pool. Valid values:
	//
	// - `PrePaid`: Subscription.
	//
	// - `PostPaid`: Pay-as-you-go.
	//
	// Default value: `PostPaid`.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The instance attribute configuration.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of instance types for the nodes. You can specify multiple instance types as fallback options. When creating a node, the system tries each instance type in order until it succeeds. The final instance type may vary depending on inventory.
	//
	// Valid values: [1, 10].
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method for public IP addresses. Valid values:
	//
	// - `PayByBandwidth`: Pay by bandwidth.
	//
	// - `PayByTraffic`: Pay by traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth for the public IP address of a node, measured in megabits per second (Mbps). The value must be an integer from 1 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The key pair name. Specify either key_pair or `login_password`. For managed node pools, only `key_pair` is supported.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH login password. Specify either login_password or `key_pair`. Password rules: 8 to 30 characters, including at least three of the following: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The scaling policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: Scales instances based on the virtual switches (VSwitchIds.N) you define. If an ECS instance cannot be created in the zone of a higher-priority virtual switch, the system automatically uses the next-priority virtual switch.
	//
	// - `COST_OPTIMIZED`: Tries to create instances starting with the lowest vCPU unit price. If multiple instance types are configured with a spot instance billing method, it prioritizes creating spot instances. Use the `CompensateWithOnDemand` parameter to specify whether to automatically create on-demand instances if spot instances cannot be created due to inventory or other reasons.
	//
	//   > `COST_OPTIMIZED` takes effect only when multiple instance types or spot instances are configured.
	//
	// - `BALANCE`: Evenly distributes ECS instances across the specified zones in the scaling group. If the distribution becomes unbalanced due to inventory shortages or other reasons, you can use the API `RebalanceInstances` to rebalance the resources. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html).
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// BALANCE
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of on-demand instances required in the scaling group. Valid values: [0, 1000]. When the number of on-demand instances is below this value, the system prioritizes creating on-demand instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of on-demand instances among the instances that exceed the minimum on-demand instance count (`on_demand_base_capacity`). Valid values: [0, 100].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration for nodes in the node pool. This parameter takes effect only when `instance_charge_type` is `PrePaid`, and is required in that case.
	//
	// - When `period_unit=Week`, valid values: {1, 2, 3, 4}.
	//
	// - When `period_unit=Month`, valid values: {1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60}.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle for nodes in the node pool. This parameter takes effect only when `instance_charge_type` is `PrePaid`, and is required in that case.
	//
	// - `Month`: Billed monthly.
	//
	// - `Week`: Billed weekly.
	//
	// Default value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// [Deprecated] Use the `image_type` parameter instead.
	//
	// The operating system platform. Valid values:
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
	// The private node pool configuration.
	PrivatePoolOptions *ModifyClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The list of RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and policy to use when creating instances. Note: This parameter is effective only when creating pay-as-you-go instances. It cannot be set at the same time as private_pool_options.match_criteria or private_pool_options.id.
	ResourcePoolOptions *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling group mode. Valid values:
	//
	// - `release`: Standard mode. Scales by creating and releasing ECS instances based on resource usage.
	//
	// - `recycle`: Fast mode. Scales by creating, stopping, and starting instances to improve scaling speed. You are not charged for compute resources when instances are stopped, only for storage fees (except for local disk models).
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The list of security group IDs.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// The number of available instance types. The scaling group creates spot instances in a balanced manner across the specified number of lowest-cost instance types. Valid values: [1, 10].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Enable spot instance replenishment. When enabled, the scaling group attempts to create a new instance to replace a spot instance that is about to be reclaimed. Valid values:
	//
	// - `true`: Enable spot instance replenishment.
	//
	// - `false`: Disable spot instance replenishment.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price range configuration for spot instances.
	SpotPriceLimit []*ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The spot instance type. Valid values:
	//
	// - `NoSpot`: On-demand instance.
	//
	// - `SpotWithPriceLimit`: Set a price limit for spot instances.
	//
	// - `SpotAsPriceGo`: Let the system automatically bid at the current market price.
	//
	// For more information, see [Spot instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Enable bursting for the system disk. Valid values:
	//
	// - true: Enable. During sudden spikes in read/write activity, the disk temporarily boosts performance until workloads return to normal.
	//
	// - false: Disable.
	//
	// You can set this parameter only when `system_disk_category` is `cloud_auto`. For more information, see [ESSD AutoPL disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The list of system disk types. If the highest-priority disk type is unavailable, the system tries the next priority disk type.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The system disk type for nodes. Valid values:
	//
	// - `cloud_efficiency`: Ultra disk.
	//
	// - `cloud_ssd`: Standard SSD.
	//
	// - `cloud_essd`: Enhanced SSD (ESSD).
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
	// The encryption algorithm for the system disk. Valid value: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Encrypt the system disk. Valid values:
	//
	// - true: Encrypt.
	//
	// - false: Do not encrypt.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The KMS key ID for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the system disk. This setting applies only to ESSD disks. Performance level depends on disk size. For more information, see [Enhanced SSD (ESSD)](https://help.aliyun.com/document_detail/122389.html).
	//
	// - PL0: Moderate I/O performance with stable read/write latency.
	//
	// - PL1: Moderate I/O performance with stable read/write latency.
	//
	// - PL2: High I/O performance with stable read/write latency.
	//
	// - PL3: Extremely high I/O performance with extremely stable read/write latency.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned IOPS for the system disk.
	//
	// Possible values: 0 to min{50,000, 1000 × capacity − baseline performance}. Baseline performance = min{1,800 + 50 × capacity, 50,000}.
	//
	// You can set this parameter only when `system_disk_category` is `cloud_auto`. For more information, see [ESSD AutoPL disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk for nodes. Unit: GiB.
	//
	// Valid values: [20, 2048].
	//
	// This value must be greater than or equal to max{20, ImageSize}.
	//
	// Default value: max{40, image size corresponding to ImageId}.
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
	// Add tags to ECS instances only.
	//
	// Tag keys must be unique and up to 128 characters long. Neither tag keys nor tag values can start with "aliyun", "acs:", or contain "https\\://" or "http\\://".
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of virtual switch IDs. Valid values: [1, 8].
	//
	// > To ensure high availability, select virtual switches from different zones.
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
	// The private node pool ID. This is required when `match_criteria` is set to `Target`.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The private node pool type, which determines the capacity option for instance startup. Elastic assurance services or capacity reservation services generate private pool capacity for instance startup. Valid values:
	//
	// - `Open`: Open mode. Automatically matches open-type private pool capacity. If no suitable private pool is found, public pool resources are used.
	//
	// - `Target`: Specified mode. Starts instances using the specified private pool capacity. If the capacity is unavailable, the instance startup fails.
	//
	// - `None`: No mode. The instance startup does not use private pool capacity.
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
	// The list of private pool IDs, which are the IDs of elastic assurance or capacity reservation services. You can only pass Target mode private pool IDs. The value of N can range from 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool policy for instance creation. Resource pools include private pools generated by elastic assurance or capacity reservation services, and public pools. Valid values: PrivatePoolFirst: Prioritizes private pools. If resouce_pool_options.private_pool_ids is specified, it uses the specified private pools first. If no private pool is specified or the specified capacity is insufficient, it automatically matches open-type private pools. If no suitable private pool is found, it uses the public pool. PrivatePoolOnly: Uses only private pools. You must specify resouce_pool_options.private_pool_ids. If the specified private pool capacity is insufficient, the instance startup fails. None: Does not use a resource pool policy. Default value: None.
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
	// The spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum hourly price for the instance.
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
	// Enable the Kubernetes cluster for confidential computing. Valid values:
	//
	// - `true`: Enable.
	//
	// - `false`: Disable.
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
