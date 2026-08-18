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
	// The intelligent managed configuration for the node pool.
	AutoMode *CreateClusterNodePoolRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The automatic scaling configuration.
	AutoScaling *CreateClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- Use desired_size instead.
	//
	// example:
	//
	// null
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The Lingjun node pool configuration.
	EfloNodeGroup *CreateClusterNodePoolRequestEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// Specifies whether the Pod network mode uses host network mode.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	InterconnectConfig *CreateClusterNodePoolRequestInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter takes effect only for node pools whose `type` is `edge`. Valid values:
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Specifies whether nodes in the edge node pool have Layer 3 network connectivity with each other.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The cluster-related configuration.
	KubernetesConfig *CreateClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configuration of the managed node pool feature.
	Management *CreateClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]**
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
	// Specifies whether to enable intelligent managed mode.
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
	// example:
	//
	// null
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable auto scaling. Valid values:
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated. Use internet_charge_type and internet_max_bandwidth_out instead.
	//
	// example:
	//
	// null
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of instances that can be scaled in the node pool, excluding your existing instances. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of instances that can be scaled in the node pool, excluding your existing instances. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The automatic scaling instance type. This parameter takes effect only when `enable=true`. Valid values:
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
	// Specifies whether to enable automatic addition for the Lingjun node pool.
	AutoAttachEnabled *bool `json:"auto_attach_enabled,omitempty" xml:"auto_attach_enabled,omitempty"`
	// The Lingjun cluster ID to associate when creating a Lingjun node pool.
	//
	// example:
	//
	// i1169130516633730****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The Lingjun group ID of the Lingjun cluster to associate when creating a Lingjun node pool.
	//
	// example:
	//
	// ng-ec3c96ff0aa****
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The Worker RAM role used by the Lingjun node pool.
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
	// example:
	//
	// null
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// null
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// null
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// null
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// **[Deprecated]**
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
	// false
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy for nodes. The following two policies are supported for clusters of version 1.12.6 or later:
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The node labels. You can add labels to nodes in the Kubernetes cluster.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The custom node name. After you customize the node name, the node name, ECS instance name, and ECS instance hostname are all changed accordingly.
	//
	// example:
	//
	// aliyun.com192.XX.YY.55test
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The instance pre-custom data. Before the node joins the cluster, the specified pre-custom data script is run. For more information, see [User-Data scripts](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The container runtime name. ACK supports the following three container runtimes.
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
	// Specifies whether the scaled-out nodes are unschedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The instance user data. After the node joins the cluster, the specified instance user data script is run. For more information, see [User-Data scripts](https://help.aliyun.com/document_detail/49121.html).
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
	// Specifies whether to enable ECS fault detection for node self-healing.
	AutoFaultDiagnosis *bool `json:"auto_fault_diagnosis,omitempty" xml:"auto_fault_diagnosis,omitempty"`
	// Specifies whether to enable automatic node repair. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The automatic node repair policy.
	AutoRepairPolicy *CreateClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to enable automatic node upgrade. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The automatic node upgrade policy.
	AutoUpgradePolicy *CreateClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to enable automatic CVE vulnerability fix. This parameter takes effect only when `enable=true`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The automatic CVE fix policy.
	AutoVulFixPolicy *CreateClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable node rotation. Only intelligent managed node pools support this feature, and it is enabled by default. Regular node pools do not support this feature.
	DriftEnabled *bool `json:"drift_enabled,omitempty" xml:"drift_enabled,omitempty"`
	// Specifies whether to enable the managed node pool feature. Valid values:
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `auto_upgrade` parameter at the upper level instead.
	//
	// The automatic upgrade configuration. This parameter takes effect only when `enable=true`.
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
	// Specifies whether to allow node restarts. This parameter takes effect only when `auto_repair=true`. Valid values:
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
	// Specifies whether to allow automatic kubelet upgrade. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow automatic operating system upgrade. This parameter takes effect only when `auto_upgrade=true`. Valid values:
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow automatic runtime upgrade. This parameter takes effect only when `auto_upgrade=true`. Valid values:
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
	// The packages to exclude during vulnerability fix.
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
	// Default value: `true`.
	//
	// example:
	//
	// false
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels allowed for automatic fix, separated by commas. Example: `asap,later`. Supported vulnerability levels:
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
	// **[Deprecated]*	- Use the `auto_upgrade` parameter at the upper level instead.
	//
	// example:
	//
	// null
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
	// 0
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of extra nodes. You can specify either this parameter or `surge`.
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
	// The node component configuration.
	Config *CreateClusterNodePoolRequestNodeComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The node component name.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node component version.
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

func (s *CreateClusterNodePoolRequestNodeComponentsConfig) SetCustomConfig(v map[string]interface{}) *CreateClusterNodePoolRequestNodeComponentsConfig {
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
	// The node pool name.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodepool-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource group ID of the node pool. Instances scaled out by the node pool belong to this resource group.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The node pool type. Valid values:
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
	// Deprecated
	//
	// **[Deprecated]*	- Use the security_hardening_os parameter instead.
	//
	// example:
	//
	// null
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Specifies whether to allow the automatic creation of pay-as-you-go instances to meet the required number of ECS instances when spot instances cannot be created due to cost or inventory reasons, if `multi_az_policy` is set to `COST_OPTIMIZED`. Valid values:
	//
	// - `true`: Allows the automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// - `false`: Does not allow the automatic creation of pay-as-you-go instances to meet the required number of ECS instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The CPU-related configuration options.
	CpuOptions *CreateClusterNodePoolRequestScalingGroupCpuOptions `json:"cpu_options,omitempty" xml:"cpu_options,omitempty" type:"Struct"`
	// The data cloud disk configurations for nodes in the node pool.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The deployment set ID. You can use a deployment set to distribute ECS instances scaled out by the node pool across different physical servers to ensure high availability and underlying disaster recovery capabilities. When ECS instances are created within a deployment set, they are launched in the specified region based on the preconfigured deployment strategy.
	//
	//
	// 	Notice: After you select a deployment set, the maximum number of nodes in the node pool is limited. By default, a deployment set supports a maximum of 20 × number of zones (the number of zones is determined by the vSwitches). Select a deployment set with caution and ensure that the quota in the deployment set is sufficient to avoid node scale-out failures.
	//
	// example:
	//
	// ds-bp1d19mmbsv3jf6xxxxx
	DeploymentsetId *string `json:"deploymentset_id,omitempty" xml:"deploymentset_id,omitempty"`
	// The desired number of nodes in the node pool.
	//
	// example:
	//
	// 0
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The block device initialization configuration.
	DiskInit []*DiskInit `json:"disk_init,omitempty" xml:"disk_init,omitempty" type:"Repeated"`
	// Specifies whether to enable high-density cloud disk mode. This is supported only when the node pool uses instance types. When enabled, the total number of system cloud disks and data cloud disks does not exceed the high-density cloud disk limit supported by the instance type.
	//
	// example:
	//
	// false
	EnableHighDensityMode *bool `json:"enable_high_density_mode,omitempty" xml:"enable_high_density_mode,omitempty"`
	// The custom image ID. The system-provided image is used by default.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of operating system image. Valid values:
	//
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of nodes in the node pool. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The ECS instance metadata access configuration.
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance attribute configuration.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The list of instance types for nodes in the node pool. When the node pool scales out, instances are created based on eligible instance types from this list.
	//
	// Number of supported instance types: [1, 10].
	//
	// > To ensure high availability, select multiple instance types.
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
	// The maximum outbound bandwidth for the public IP of nodes. Unit: Mbps (Mega bit per second). Valid values: [1,100\\].
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair for password-free logon. Use either this parameter or `login_password`.
	//
	// example:
	//
	// np-key-name
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether to log on to the scaled-out ECS instances as a non-root user.
	//
	// - true: Log on as a non-root user (ecs-user).
	//
	// - false: Log on as the root user.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The SSH logon password. Use either this parameter or `key_pair`. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// ****
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The multi-zone scaling policy for ECS instances in the scaling group. Valid values:
	//
	// - `PRIORITY`: Scales ECS instances based on the vSwitches (VSwitchIds.N) that you define. When ECS instances cannot be created in the zone of the vSwitch with the highest priority, the system automatically uses the vSwitch with the next highest priority to create ECS instances.
	//
	// - `COST_OPTIMIZED`: Attempts to create ECS instances in order of vCPU unit price from lowest to highest. When the scaling configuration specifies multiple instance types with the spot billing method, spot instances are created first. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically attempt to create pay-as-you-go instances when spot instances cannot be created due to insufficient inventory or other reasons.
	//
	//   >`COST_OPTIMIZED` takes effect only when the scaling configuration specifies multiple instance types or uses spot instances.
	//
	// - `BALANCE`: Evenly distributes ECS instances across the multiple zones specified in the scaling group. If the zones become unbalanced due to insufficient inventory or other reasons, you can call the [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) operation to rebalance resources.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. Valid values: [0,1000\\]. When the number of pay-as-you-go instances is less than this value, pay-as-you-go instances are preferentially created.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that exceed the minimum number of pay-as-you-go instances (`on_demand_base_capacity`) in the scaling group. Valid values: [0,100\\].
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of nodes in the node pool. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription period for nodes in the node pool. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use the `image_type` parameter instead.
	//
	// example:
	//
	// null
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The private node pool configuration.
	PrivatePoolOptions *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker RAM role.
	//
	// 	- If this parameter is left empty, the default worker RAM role created by the cluster is used.
	//
	// 	- If this parameter is specified, the specified RAM role must be a **normal service role**, and its **trusted service*	- must be configured as **Elastic Compute Service**. For more information, see [Create a normal service role](https://help.aliyun.com/document_detail/116800.html). If the specified RAM role is not the default worker RAM role created by the cluster, the role name cannot start with `KubernetesMasterRole-` or `KubernetesWorkerRole-`.
	//
	// 	Notice: This parameter is supported only by ACK managed clusters of version 1.22 or later.</notice>
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// The list of RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The resource pool and resource pool policy used when creating instances. After you set this parameter, note the following:
	ResourcePoolOptions *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling group mode. Valid values:
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// Deprecated
	//
	// The security group ID of the node pool. Use either this parameter or `security_group_ids`. We recommend that you use `security_group_ids`.
	//
	// example:
	//
	// sg-wz9a8g2mt6x5llu0****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The list of security group IDs. Use either this parameter or `security_group_id`. We recommend that you use `security_group_ids`. If both `security_group_id` and `security_group_ids` are specified, `security_group_ids` takes precedence.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// Alibaba Cloud OS security hardening. Valid values:
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Specifies whether to enable MLPS 2.0 security hardening. You can enable MLPS 2.0 security hardening for nodes only when the system image is Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3. Alibaba Cloud provides classified protection compliance baseline check standards and scanning programs for Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 MLPS 2.0 Level 3 images.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of available instance types. The scaling group creates spot instances of multiple types at the lowest cost. Valid values: [1,10\\].
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Specifies whether to enable supplementation of spot instances. If enabled, when the system receives a notification that a spot instance will be reclaimed, the scaling group attempts to create a new instance to replace the spot instance that will be reclaimed. Valid values:
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The price limit configurations for a single spot instance of the current instance type.
	SpotPriceLimit []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The type of spot instance. Valid values:
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
	// Multiple disk types for the system cloud disk. If a disk type with a higher priority is unavailable, the system automatically attempts to create the system cloud disk with the next priority disk type.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The type of the system cloud disk for nodes. Valid values:
	//
	// - `cloud_efficiency`: ultra cloud disk.
	//
	// - `cloud_ssd`: standard SSD.
	//
	// - `cloud_essd`: ESSD.
	//
	// - `cloud_auto`: ESSD AutoPL.
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
	// The performance level of the system cloud disk for nodes. This parameter takes effect only for ESSD disks. The performance level is related to the disk size. For more information, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The provisioned read/write IOPS of the node system cloud disk.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system cloud disk for nodes. Unit: GiB.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The system cloud disk snapshot policy.
	//
	// example:
	//
	// sp-0jl6xnmme8v7o935****
	SystemDiskSnapshotPolicyId *string `json:"system_disk_snapshot_policy_id,omitempty" xml:"system_disk_snapshot_policy_id,omitempty"`
	// Tags that are added only to ECS instances.
	Tags []*CreateClusterNodePoolRequestScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of vSwitch IDs. Valid values: [1,8\\].
	//
	// > To ensure high availability, select vSwitches in different zones.
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
	// The private node pool ID. When `match_criteria` is set to `Target`, you must specify the private pool ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the private node pool. The private pool option for instance startup. After an elasticity assurance or capacity reservation takes effect, a private pool is generated for instances to use during startup. Valid values:
	//
	// - `Open`: open mode. The system automatically matches open private pool capacity. If no matching private pool capacity is available, public pool resources are used to start the instance.
	//
	// - `Target`: targeted mode. The instance is started using the specified private pool capacity. If the specified private pool capacity is unavailable, the instance fails to start.
	//
	// - `None`: none mode. The instance does not use private pool capacity during startup.
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
	// The list of private pool IDs, which are elasticity assurance IDs or capacity reservation IDs. Only Target mode private pool IDs can be specified. Valid values of N: 1 to 20.
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
	// The resource pool strategy used when instances are created. Resource pools include private pools generated after elasticity assurance or capacity reservation takes effect, and public pools, which are available for instance startup. Valid values:
	//
	// - PrivatePoolFirst: private pool first. When this strategy is selected, if resouce_pool_options.private_pool_ids is specified, the specified private pools are used first. If no private pool is specified or the specified private pool has insufficient capacity, open-type private pools are automatically matched. If no eligible private pool is available, the public pool is used to create instances.
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
	// Specifies whether to enable the confidential computing cluster.
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
