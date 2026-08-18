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
	SetEfloNodeGroup(v *ModifyClusterNodePoolRequestEfloNodeGroup) *ModifyClusterNodePoolRequest
	GetEfloNodeGroup() *ModifyClusterNodePoolRequestEfloNodeGroup
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
	// Specifies whether to enable concurrency.
	//
	// example:
	//
	// true
	Concurrency *bool `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The Lingjun node pool configuration. (Not effective)
	EfloNodeGroup *ModifyClusterNodePoolRequestEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// The cluster-related configuration.
	KubernetesConfig *ModifyClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The managed node pool configuration.
	Management *ModifyClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// The node pool configuration.
	NodepoolInfo *ModifyClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The scaling group configuration of the node pool.
	ScalingGroup *ModifyClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The confidential computing cluster configuration.
	TeeConfig *ModifyClusterNodePoolRequestTeeConfig `json:"tee_config,omitempty" xml:"tee_config,omitempty" type:"Struct"`
	// Specifies whether to synchronously update node labels and taints.
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

func (s *ModifyClusterNodePoolRequest) GetEfloNodeGroup() *ModifyClusterNodePoolRequestEfloNodeGroup {
	return s.EfloNodeGroup
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

func (s *ModifyClusterNodePoolRequest) SetEfloNodeGroup(v *ModifyClusterNodePoolRequestEfloNodeGroup) *ModifyClusterNodePoolRequest {
	s.EfloNodeGroup = v
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
	if s.EfloNodeGroup != nil {
		if err := s.EfloNodeGroup.Validate(); err != nil {
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
	// **[Deprecated]*	- Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// The peak bandwidth of the EIP.
	//
	// Valid values: [1,100]. Unit: Mbit/s.
	//
	// example:
	//
	// null
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// example:
	//
	// null
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable auto scaling. Valid values:
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// example:
	//
	// null
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances that can be scaled out in the node pool, excluding your existing instances. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of scalable instances in the node pool, excluding your existing instances. This parameter takes effect only when `enable=true`.
	//
	// Valid values: [0, max_instances]. Default value: 0.
	//
	// > - If the minimum number of instances is not 0, the scaling group automatically creates the corresponding number of ECS instances after it takes effect.
	//
	// > - Set the maximum number of instances to a value that is not less than the current number of nodes in the node pool. Otherwise, nodes in the node pool are scaled in after the auto scaling feature takes effect.
	//
	// example:
	//
	// 2
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// Deprecated
	//
	// The auto scaling type, classified by instance type. Valid values:
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

type ModifyClusterNodePoolRequestEfloNodeGroup struct {
	// Specifies whether to enable automatic addition for the Lingjun node pool. (Not effective)
	AutoAttachEnabled *string `json:"auto_attach_enabled,omitempty" xml:"auto_attach_enabled,omitempty"`
}

func (s ModifyClusterNodePoolRequestEfloNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolRequestEfloNodeGroup) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolRequestEfloNodeGroup) GetAutoAttachEnabled() *string {
	return s.AutoAttachEnabled
}

func (s *ModifyClusterNodePoolRequestEfloNodeGroup) SetAutoAttachEnabled(v string) *ModifyClusterNodePoolRequestEfloNodeGroup {
	s.AutoAttachEnabled = &v
	return s
}

func (s *ModifyClusterNodePoolRequestEfloNodeGroup) Validate() error {
	return dara.Validate(s)
}

type ModifyClusterNodePoolRequestKubernetesConfig struct {
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
	// The CPU management policy for nodes. The following two policies are supported for clusters of version 1.12.6 or later:
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The node labels. You can add labels to the nodes in the Kubernetes cluster. Label definition rules:
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name parameter. A node name consists of three parts: prefix + node IP address + suffix.
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The pre-user data of the instance. Before a node joins the cluster, the specified pre-user data script is run. For more information, see [User data scripts](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// IyEvdXNyL2Jpbi9iYXNoCmVjaG8gIkhlbGxvIEFDSyEi
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The container runtime name. ACK supports the following three container runtimes.
	//
	// - containerd: Recommended. Supports all cluster versions.
	//
	// - Sandboxed-Container.runv: Sandboxed container runtime that provides higher isolation. Supports clusters of version 1.31 and earlier.
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
	// Specifies whether the nodes added during the scale-out are unschedulable.
	//
	// example:
	//
	// false
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After a node joins the cluster, the specified user data script is run. For more information, see [User data scripts](https://help.aliyun.com/document_detail/49121.html).
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
	// Specifies whether to enable ECS fault detection for node self-healing.
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to automatically repair nodes. This parameter takes effect only when `enable=true`.
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
	// The auto repair node policy.
	AutoRepairPolicy *ModifyClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically upgrade nodes. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto upgrade policy.
	AutoUpgradePolicy *ModifyClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to automatically fix CVE vulnerabilities. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The auto CVE fix policy.
	AutoVulFixPolicy *ModifyClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable node rotation. Only intelligent managed node pools support this feature, and it is enabled by default. Common node pools do not support this feature.
	DriftEnabled *bool `json:"drift_enabled,omitempty" xml:"drift_enabled,omitempty"`
	// Specifies whether to enable the managed node pool. Valid values:
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `auto_upgrade` parameter at the upper level instead.
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

func (s *ModifyClusterNodePoolRequestManagement) GetDriftEnabled() *bool {
	return s.DriftEnabled
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

func (s *ModifyClusterNodePoolRequestManagement) SetDriftEnabled(v bool) *ModifyClusterNodePoolRequestManagement {
	s.DriftEnabled = &v
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
	// Specifies whether node repair requires manual approval.
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// The auto repair policy ID.
	//
	// example:
	//
	// r-xxxxxxxxxx
	AutoRepairPolicyId *string `json:"auto_repair_policy_id,omitempty" xml:"auto_repair_policy_id,omitempty"`
	// The maximum number of nodes that can be repaired in parallel. When a large number of abnormal nodes exist in the node pool, this parameter specifies the maximum number or percentage of nodes that can be repaired simultaneously. You can specify a number (such as 5, valid range: 1 to 100000) or a percentage (such as 10%, valid range: 1% to 100%). Default value: 1.
	//
	// example:
	//
	// 5
	MaxParallelRepairingNodes *string `json:"max_parallel_repairing_nodes,omitempty" xml:"max_parallel_repairing_nodes,omitempty"`
	// The circuit breaker condition for self-healing. When the number or percentage of faulty nodes exceeds this threshold, self-healing enters a circuit breaker state and stops initiating new repair actions. You can specify a number (such as 10, valid range: 1 to 100000) or a percentage (such as 20%, valid range: 1% to 100%). Default value: 100%.
	//
	// example:
	//
	// 20%
	MaxUnhealthyNodesThreshold *string `json:"max_unhealthy_nodes_threshold,omitempty" xml:"max_unhealthy_nodes_threshold,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only when `auto_repair=true`. Valid values:
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

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GetMaxParallelRepairingNodes() *string {
	return s.MaxParallelRepairingNodes
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) GetMaxUnhealthyNodesThreshold() *string {
	return s.MaxUnhealthyNodesThreshold
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

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) SetMaxParallelRepairingNodes(v string) *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	s.MaxParallelRepairingNodes = &v
	return s
}

func (s *ModifyClusterNodePoolRequestManagementAutoRepairPolicy) SetMaxUnhealthyNodesThreshold(v string) *ModifyClusterNodePoolRequestManagementAutoRepairPolicy {
	s.MaxUnhealthyNodesThreshold = &v
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
	// Specifies whether to allow auto upgrade of kubelet. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow auto upgrade of the operating system. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow auto upgrade of the runtime. This parameter takes effect only when `auto_upgrade=true`. Valid values:
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
	// Specifies the packages to exclude during vulnerability fix.
	//
	// example:
	//
	// kernel
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Specifies whether to allow node restarts. This parameter takes effect only when `auto_vul_fix=true`. Valid values:
	//
	// - `true`: Node restarts are allowed.
	//
	// - `false`: Node restarts are not allowed.
	//
	// Default value: `true`
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels that are allowed for auto fix, separated by commas. Example: `asap,later`. Supported vulnerability levels:
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
	// **[Deprecated]*	- Use the `auto_upgrade` parameter at the upper level instead.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of unavailable nodes.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of extra nodes. You can specify either this parameter or `surge_percentage`.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You can specify either this parameter or `surge`.
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
	// Naming rules: The name must be 1 to 63 characters in length and can contain digits, Chinese characters, letters, and hyphens (-). It cannot start with a hyphen (-).
	//
	// example:
	//
	// default-nodepool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource group ID of the node pool. Instances scaled out by the node pool belong to this resource group.
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
	// Specifies whether to enable auto-renewal for nodes. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The duration of a single auto-renewal cycle. Valid values:
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Specifies whether to allow the automatic creation of pay-as-you-go instances to meet the required number of ECS instances when spot instances cannot be created due to cost or inventory reasons, if `multi_az_policy` is set to `COST_OPTIMIZED`. Valid values:
	//
	// - `true`: allows the automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: does not allow the automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The data cloud disk configurations for nodes. Valid values: 0 to 10. You can add up to 10 data cloud disks.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The deployment set to which the ECS instances created by the node pool belong. This setting takes effect only on new nodes. The deployment set of existing nodes is not changed.
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
	// Specifies whether to enable high-density cloud disk mode. This is supported only when the node pool uses instance types. When enabled, the total number of system cloud disks and data cloud disks does not exceed the high-density cloud disk count supported by the instance type.
	//
	// example:
	//
	// false
	EnableHighDensityMode *bool `json:"enable_high_density_mode,omitempty" xml:"enable_high_density_mode,omitempty"`
	// The custom image ID. You can call `DescribeKubernetesVersionMetadata` to query the images supported by the system. The latest system image is used by default.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The operating system distribution type. We recommend that you use this field to specify the node operating system. Valid values:
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
	// The billing method of nodes in the node pool. Valid values:
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The instance attribute configurations.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of node instance types. You can specify multiple instance types as alternatives. During node creation, the system attempts to purchase instances starting from the first specification until the creation succeeds. The final purchased instance type may vary depending on inventory availability.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method for public IP addresses. Valid values:
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth for the node public IP address. Unit: Mbps (Mega bit per second). Valid values: [1,100\\].
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The key pair name. Specify either this parameter or `login_password`. When the node pool is a managed node pool, only `key_pair` is supported.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// The SSH logon password. Specify either this parameter or `key_pair`. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The multi-zone scaling policy for ECS instances in the scaling group. Valid values:
	//
	// example:
	//
	// BALANCE
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances required by the scaling group. Valid values: [0,1000\\]. Pay-as-you-go instances are created with priority when the number of pay-as-you-go instances is less than this value.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the instances that exceed the minimum number of pay-as-you-go instances (`on_demand_base_capacity`) in the scaling group. Valid values: [0,100\\].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid` and is required in this case.
	//
	// - When `period_unit=Week`, valid values of `period`: 1, 2, 3, and 4.
	//
	// - When `period_unit=Month`, valid values of `period`: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of nodes in the node pool. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `image_type` parameter instead.
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
	// The resource pools and resource pool policies used when creating instances. After you set this parameter, note the following:
	ResourcePoolOptions *ModifyClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling group mode. Valid values:
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The list of security group IDs.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// The number of available instance types. The scaling group creates spot instances in a cost-optimized manner across multiple instance types. Valid values: [1,10\\].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable spot instance supplementation. If enabled, when the scaling group receives a system message that a spot instance will be reclaimed, it attempts to create a new instance to replace the spot instance that is about to be reclaimed. Valid values:
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price limit configurations for spot instances.
	SpotPriceLimit []*ModifyClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The type of spot instance. Valid values:
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable burst (I/O performance burst) for the node system cloud disk. Valid values:
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The multiple cloud disk types for the system cloud disk. If a cloud disk type with a higher priority is unavailable, the system automatically attempts the next priority cloud disk type to create the system cloud disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system cloud disk for nodes. Valid values:
	//
	// - `cloud_efficiency`: ultra cloud disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud_auto`: ESSD AutoPL cloud disk.
	//
	// - `cloud_essd_entry`: ESSD Entry disk.
	//
	// Default value: `cloud_efficiency`.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm used by the system cloud disk. Valid values: aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system cloud disk. Valid values:
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// The KMS key ID used by the system cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level of the node system cloud disk. This parameter takes effect only for ESSD cloud disks. The performance level is related to the cloud disk size. For more information, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS for the node system cloud disk.
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
	// The snapshot policy for the system cloud disk.
	//
	// example:
	//
	// sp-0jl6xnmme8v7o935****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// Adds tags only to ECS instances.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of vSwitch IDs. Valid values: 1 to 8.
	//
	// > To ensure high availability, select vSwitches in different zones.
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

func (s *ModifyClusterNodePoolRequestScalingGroup) GetEnableHighDensityMode() *bool {
	return s.EnableHighDensityMode
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

func (s *ModifyClusterNodePoolRequestScalingGroup) SetEnableHighDensityMode(v bool) *ModifyClusterNodePoolRequestScalingGroup {
	s.EnableHighDensityMode = &v
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
	// The private node pool ID. When `match_criteria` is set to `Target`, you must specify the private pool ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. The private pool option for instance launch. After an elasticity assurance or capacity reservation takes effect, a private pool is generated for instances to use during launch. Valid values:
	//
	// - `Open`: open mode. The system automatically matches open private pool capacity. If no matching private pool capacity is available, public pool resources are used to launch the instance.
	//
	// - `Target`: targeted mode. The instance is launched using the specified private pool capacity. If the specified private pool capacity is unavailable, the instance fails to launch.
	//
	// - `None`: none mode. The instance launch does not use private pool capacity.
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
	// The list of private pool IDs, which are elasticity assurance IDs or capacity reservation IDs. Only Target mode private pool IDs can be specified. Valid values of N: 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool strategy used when instances are created. Resource pools include private pools generated after elasticity assurance or capacity reservation takes effect, and public pools, which are available for instance startup. Valid values:
	//
	// - PrivatePoolFirst: private pool preferred. When this strategy is selected, if resouce_pool_options.private_pool_ids is specified, the specified private pools are used first. If no private pool is specified or the specified private pool has insufficient capacity, open-type private pools are automatically matched. If no eligible private pool is available, the public pool is used to create instances.
	//
	// - PrivatePoolOnly: private pool only. When this strategy is selected, you must specify resouce_pool_options.private_pool_ids. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - None: no resource pool strategy is used.
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
	// The spot instance type.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The maximum price of a single instance.
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
	// - `true`: Enabled.
	//
	// - `false`: Not enabled.
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
