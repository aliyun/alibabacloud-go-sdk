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
	AutoMode *CreateClusterNodePoolRequestAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The configurations of auto scaling.
	AutoScaling *CreateClusterNodePoolRequestAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Deprecated
	//
	// This parameter is deprecated. Use the desired_size parameter instead.
	//
	// The number of nodes in the node pool.
	//
	// example:
	//
	// 1
	Count         *int64                                     `json:"count,omitempty" xml:"count,omitempty"`
	EfloNodeGroup *CreateClusterNodePoolRequestEfloNodeGroup `json:"eflo_node_group,omitempty" xml:"eflo_node_group,omitempty" type:"Struct"`
	// Specifies whether to set the network type of the pod to host network.
	//
	// 	- `true`: sets to host network.
	//
	// 	- `false`: sets to container network.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// The configurations of the edge node pool.
	InterconnectConfig *CreateClusterNodePoolRequestInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter takes effect only when the `type` of the node pool is set to `edge`. Valid values:
	//
	// 	- `basic`: Internet.
	//
	// 	- `private`: private network.
	//
	// example:
	//
	// basic
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Specifies whether all nodes in the edge node pool can communicate with each other at Layer 3.
	//
	// 	- `true`: The nodes in the edge node pool can communicate with each other at Layer 3.
	//
	// 	- `false`: The nodes in the edge node pool cannot communicate with each other at Layer 3.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The configurations of the cluster.
	KubernetesConfig *CreateClusterNodePoolRequestKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configurations of the managed node pool feature.
	Management *CreateClusterNodePoolRequestManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// The maximum number of nodes that can be contained in the edge node pool.
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The node configurations.
	NodeConfig *CreateClusterNodePoolRequestNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The configurations of the node pool.
	NodepoolInfo *CreateClusterNodePoolRequestNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group that is used by the node pool.
	ScalingGroup *CreateClusterNodePoolRequestScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The configurations of confidential computing for the cluster.
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
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestAutoMode struct {
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
	// This parameter is deprecated.
	//
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// **
	//
	// **Important*	- This parameter is deprecated. Use the internet_charge_type and internet_max_bandwidth_out parameters instead.
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// The metering method of the EIP. Valid values:
	//
	// 	- `PayByBandwidth`: pay-by-bandwidth.
	//
	// 	- `PayByTraffic`: pay-by-data-transfer.
	//
	// Default value: `PayByBandwidth`.
	//
	// **
	//
	// **Important*	- This parameter is deprecated. Use the internet_charge_type and internet_max_bandwidth_out parameters instead.
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Specifies whether to enable auto scaling for the node pool. Valid values:
	//
	// 	- `true`: enables auto scaling.
	//
	// 	- `false`: disables auto scaling. If you set this parameter to false, other parameters in the `auto_scaling` section do not take effect.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// Specifies whether to associate an elastic IP address (EIP) with the node pool. Valid values:
	//
	// 	- `true`: associates an EIP with the node pool.
	//
	// 	- `false`: does not associate an EIP with the node pool.
	//
	// Default value: `false`.
	//
	// **
	//
	// **Important*	- This parameter is deprecated. Use the internet_charge_type and internet_max_bandwidth_out parameters instead.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number to which the Elastic Compute Service (ECS) instances in the node pool can be scaled. The number of nodes in the node pool cannot be greater than this value. This parameter takes effect only when `enable` is set to true. Valid values: [min_instances, 2000]. Default value: 0.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number to which the ECS instances in the node pool can be scaled. The number of nodes in the node pool cannot be smaller than this value. This parameter takes effect only when `enable` is set to true. Valid values: [0, max_instances]. Default value: 0.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The instance type that is used for auto scaling. This parameter takes effect only when `enable` is set to true. Valid values:
	//
	// 	- `cpu`: regular instance.
	//
	// 	- `gpu`: GPU-accelerated instance.
	//
	// 	- `gpushare`: shared GPU-accelerated instance.
	//
	// 	- `spot`: preemptible instance.
	//
	// Default value: `cpu`.
	//
	// >  You cannot modify this parameter after the node pool is created.
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
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	GroupId   *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
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
	// This parameter is deprecated.
	//
	// The bandwidth of the enhanced edge node pool. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// This parameter is deprecated.
	//
	// The ID of the Cloud Connect Network (CCN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// This parameter is deprecated.
	//
	// The region to which the CCN instance that is associated with the enhanced edge node pool belongs.
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// This parameter is deprecated.
	//
	// The ID of the Cloud Enterprise Network (CEN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// This parameter is deprecated.
	//
	// The subscription duration of the enhanced edge node pool. The duration is measured in months.
	//
	// example:
	//
	// 1
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
	// Specifies whether to install the CloudMonitor agent on ECS nodes. After the CloudMonitor agent is installed on ECS nodes, you can view monitoring information about the instances in the CloudMonitor console. We recommend that you install the CloudMonitor agent. Valid values:
	//
	// 	- `true`: installs the CloudMonitor agent on ECS nodes.
	//
	// 	- `false`: does not install the CloudMonitor agent on ECS nodes.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of nodes in the node pool. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later:
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted with enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: specifies that the default CPU affinity is used.
	//
	// Default value: `none`.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels that you want to add to the nodes in the cluster.
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// A custom node name consists of a prefix, a node IP address, and a suffix.
	//
	// 	- The prefix and the suffix can contain multiple parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). A custom node name must start and end with a digit or lowercase letter.
	//
	// 	- The node IP address in a custom node name is the private IP address of the node.
	//
	// Set the parameter to a value in the customized,aliyun,ip,com format. The value consists of four parts that are separated by commas (,). customized and ip are fixed content. aliyun is the prefix and com is the suffix. Example: aliyun.192.168.xxx.xxx.com.
	//
	// example:
	//
	// customized,aliyun,ip,com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The user-defined data of nodes. You can specify custom scripts that are automatically executed before the nodes are initialized.
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime. The following types of runtime are supported by ACK:
	//
	// 	- containerd: containerd is the recommended runtime and supports all Kubernetes versions.
	//
	// 	- Sandboxed-Container.runv: The Sandbox-Container runtime provides improved isolation and supports Kubernetes 1.24 and earlier.
	//
	// 	- docker: The Docker runtime supports Kubernetes 1.22 and earlier.
	//
	// Default value: containerd.
	//
	// example:
	//
	// docker
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The version of the container runtime.
	//
	// example:
	//
	// 19.03.5
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// The configurations of the taints.
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Specifies whether the nodes are schedulable after a scale-out operation is performed.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The user-defined data of nodes. You can specify custom scripts that are automatically executed after the nodes are initialized.
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
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestManagement struct {
	// Specifies whether to enable auto node repair. This parameter takes effect only when `enable` is set to true.
	//
	// 	- `true`: enables auto node repair.
	//
	// 	- `false`: disables auto node repair.
	//
	// If `enable` is set to true, the default value of this parameter is `true`. If `enable` is set to false, the default value of this parameter is `false`.
	//
	// example:
	//
	// false
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// The auto node repair policy.
	AutoRepairPolicy *CreateClusterNodePoolRequestManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Specifies whether to enable auto node upgrade. This parameter takes effect only when `enable` is set to true.
	//
	// 	- `true`: enables auto node upgrade.
	//
	// 	- `false`: disables auto node upgrade.
	//
	// If `enable` is set to true, the default value of this parameter is `true`. If `enable` is set to false, the default value of this parameter is `false`.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The auto node upgrade policy.
	AutoUpgradePolicy *CreateClusterNodePoolRequestManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Specifies whether to enable auto Common Vulnerabilities and Exposures (CVE) patching. This parameter takes effect only when `enable` is set to true.
	//
	// 	- `true`: enables auto CVE patching.
	//
	// 	- `false`: disables auto CVE patching.
	//
	// If `enable` is set to true, the default value of this parameter is `true`. If `enable` is set to false, the default value of this parameter is `false`.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// The auto CVE patching policy.
	AutoVulFixPolicy *CreateClusterNodePoolRequestManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Specifies whether to enable the managed node pool feature. Valid values:
	//
	// 	- `true`: enables the managed node pool feature.
	//
	// 	- `false`: disables the managed node pool feature. Other parameters in this section take effect only when enable is set to true.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Deprecated
	//
	// The configurations of auto upgrade. The configurations take effect only when `enable` is set to true.
	UpgradeConfig *CreateClusterNodePoolRequestManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s CreateClusterNodePoolRequestManagement) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolRequestManagement) GoString() string {
	return s.String()
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
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestManagementAutoRepairPolicy struct {
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only when `auto_repair` is set to true. Valid values:
	//
	// 	- `true`: allows node restart.
	//
	// 	- `false`: does not allow node restart.
	//
	// If `auto_repair` is set to true, the default value of this parameter is `true`. If `auto_repair` is set to false, the default value of this parameter is `false`.
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
	// Specifies whether to allow the auto upgrade of the kubelet. This parameter takes effect only when `auto_upgrade` is set to true. Valid values:
	//
	// 	- `true`: allows the auto upgrade of the kubelet.
	//
	// 	- `false`: does not allow the auto upgrade of the kubelet.
	//
	// If `auto_upgrade` is set to true, the default value of this parameter is `true`. If `auto_upgrade` is set to false, the default value of this parameter is `false`.
	//
	// example:
	//
	// true
	AutoUpgradeKubelet *bool `json:"auto_upgrade_kubelet,omitempty" xml:"auto_upgrade_kubelet,omitempty"`
	// Specifies whether to allow the auto upgrade of the OS. This parameter takes effect only when `auto_upgrade` is set to true. Valid values:
	//
	// 	- `true`: allows the auto upgrade of the OS.
	//
	// 	- `false`: does not allow the auto upgrade of the OS.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AutoUpgradeOs *bool `json:"auto_upgrade_os,omitempty" xml:"auto_upgrade_os,omitempty"`
	// Specifies whether to allow the auto upgrade of the runtime. This parameter takes effect only when `auto_upgrade` is set to true. Valid values:
	//
	// 	- `true`: allows the auto upgrade of the runtime.
	//
	// 	- `false`: does not allow the auto upgrade of the runtime.
	//
	// Default value: `false`.
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
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Specifies whether to allow node restart. This parameter takes effect only when `auto_vul_fix` is set to true. Valid values:
	//
	// 	- `true`: allows node restart.
	//
	// 	- `false`: does not allow node restart. If `auto_vul_fix` is set to true, the default value of this parameter is `false`. If `auto_vul_fix` is set to false, the default value of this parameter is `false`.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The severity levels of CVEs that can be automatically patched. Separate multiple levels with commas (,). Example: `asap,later`. Valid values:
	//
	// 	- `asap`: high.
	//
	// 	- `later`: medium.
	//
	// 	- `nntf`: low.
	//
	// If `auto_vul_fix` is set to true, the default value of this parameter is `asap`.
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
	// Specifies whether to enable auto upgrade. Valid values:
	//
	// 	- `true`: enables auto OS upgrade.
	//
	// 	- `false`: disables auto OS upgrade.
	//
	// **
	//
	// **Caution*	- This parameter is deprecated. Use the preceding auto_upgrade parameter instead.
	//
	// example:
	//
	// false
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of nodes that can be in the Unavailable state. Valid values: 1 to 1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of nodes that are temporarily added to the node pool during an auto upgrade.
	//
	// example:
	//
	// 0
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of additional nodes that are temporarily added to the node pool during an auto upgrade. You must set this parameter or `surge`.
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

type CreateClusterNodePoolRequestNodeConfig struct {
	// The configurations of the kubelet.
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
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestNodepoolInfo struct {
	// The name of the node pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the resource group to which the node pool belongs. Instances that are added to the node pool belong to this resource group.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of node pool. Valid values:
	//
	// 	- `ess`: regular node pool, which supports the managed node pool feature and the auto scaling feature.
	//
	// 	- `edge`: edge node pool.
	//
	// 	- `lingjun`: Lingjun node pool.
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
	// Specifies whether to enable auto-renewal for nodes in the node pool. This parameter takes effect only when you set `instance_charge_type` to `PrePaid`. Valid values:
	//
	// 	- `true`: enables auto-renewal.
	//
	// 	- `false`: disables auto-renewal.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// 	- Valid values when PeriodUnit is set to Week: 1, 2, and 3.
	//
	// 	- Valid values when PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use security_hardening_os instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the required number of ECS instances if preemptible instances cannot be created due to reasons such as the price or insufficient inventory. This parameter takes effect when you set `multi_az_policy` to `COST_OPTIMIZED`. Valid values:
	//
	// 	- `true`: automatically creates pay-as-you-go instances to meet the required number of ECS instances if preemptible instances cannot be created.
	//
	// 	- `false`: does not create pay-as-you-go instances to meet the required number of ECS instances if preemptible instances cannot be created.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The configurations of the data disks that are attached to nodes in the node pool.
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
	// 0
	DesiredSize *int64 `json:"desired_size,omitempty" xml:"desired_size,omitempty"`
	// The custom image ID. By default, the image provided by Container Service for Kubernetes (ACK) is used.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The type of the OS image. You must specify this parameter or `platform`. Valid values:
	//
	// 	- `AliyunLinux`: Alibaba Cloud Linux 2.
	//
	// 	- `AliyunLinuxSecurity`: Alibaba Cloud Linux 2 (UEFI).
	//
	// 	- `AliyunLinux3`: Alibaba Cloud Linux 3
	//
	// 	- `AliyunLinux3Arm64`: Alibaba Cloud Linux 3 for ARM.
	//
	// 	- `AliyunLinux3Security`: Alibaba Cloud Linux 3 for ARM.
	//
	// 	- `CentOS`: CentOS.
	//
	// 	- `Windows`: Windows.
	//
	// 	- `WindowsCore`: Windows Core.
	//
	// 	- `ContainerOS`: ContainerOS.
	//
	// 	- `AliyunLinux3ContainerOptimized`: Alibaba Cloud Linux 3 Container-optimized.
	//
	// example:
	//
	// AliyunLinux
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of nodes in the node pool. Valid values:
	//
	// 	- `PrePaid`: subscription.
	//
	// 	- `PostPaid`: pay-as-you-go.
	//
	// Default value: `PostPaid`
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType      *string                  `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance attributes.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// The instance types of nodes in the node pool. When the system adds a node to the node pool, the system selects the most appropriate one from the specified instance types for the node. You can specify 1 to 10 instance types.
	//
	// >  To ensure high availability, we recommend that you specify multiple instance types.
	//
	// This parameter is required.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The metering method of the public IP address. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth.
	//
	// 	- PayByTraffic: pay-by-data-transfer.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"internet_charge_type,omitempty" xml:"internet_charge_type,omitempty"`
	// The maximum outbound bandwidth of the public IP address. Unit: Mbit/s. Valid values: 1 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int64 `json:"internet_max_bandwidth_out,omitempty" xml:"internet_max_bandwidth_out,omitempty"`
	// The name of the key pair used to log on to nodes in the node pool. You must set this parameter or `login_password`.
	//
	// >  If you select ContainerOS as the OS of nodes in the node pool, you must specify `key_pair`.
	//
	// example:
	//
	// np-key-name
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Specifies whether to allow a non-root user to log on to an ECS instance that is added to the node pool.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The password for SSH logon. You must specify this parameter or the `key_pair` parameter. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// example:
	//
	// Hello1234
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The ECS instance scaling policy for the multi-zone scaling group. Valid values:
	//
	// 	- `PRIORITY`: ECS instances are scaled based on the value of VSwitchIds.N. If an ECS instance cannot be created in the zone where the vSwitch that has the highest priority resides, the system creates the ECS instance in the zone where the vSwitch that has the next highest priority resides.
	//
	// 	- `COST_OPTIMIZED`: ECS instances are created based on the vCPU unit price in ascending order. Preemptible instances are preferably created if preemptible instance types are specified in the scaling configurations. You can set the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when preemptible instances cannot be created due to insufficient inventory.
	//
	//     **
	//
	//     **Note*	- `COST_OPTIMIZED` takes effect only when multiple instance types are specified or at least one preemptible instance type is specified.
	//
	// 	- `BALANCE`: ECS instances are evenly distributed across multiple zones for the scaling group. If ECS instances become imbalanced among multiple zones due to insufficient inventory, you can call the [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html) operation of Auto Scaling to evenly distribute the ECS instances among zones.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// COST_OPTIMIZED
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be kept in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than the value of this parameter, the system preferably creates pay-as-you-go instances.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int64 `json:"on_demand_base_capacity,omitempty" xml:"on_demand_base_capacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that exceed the number specified by `on_demand_base_capacity`. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int64 `json:"on_demand_percentage_above_base_capacity,omitempty" xml:"on_demand_percentage_above_base_capacity,omitempty"`
	// The subscription duration of nodes in the node pool. This parameter takes effect and is required if you set `instance_charge_type` to `PrePaid`.
	//
	// 	- If `period_unit` is set to Week, the valid values of `period` are 1, 2, 3, and 4.
	//
	// 	- If `period_unit` is set to Month, the valid values of `period` are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of nodes in the node pool. This parameter takes effect and is required only when you set `instance_charge_type` to `PrePaid`. Valid values:
	//
	// 	- `Month`: The subscription duration is measured in months.
	//
	// 	- `Week`: The subscription duration is measured in weeks.
	//
	// Default value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// Deprecated
	//
	// The operating system distribution. Valid values:
	//
	// 	- `CentOS`
	//
	// 	- `AliyunLinux`
	//
	// 	- `Windows`
	//
	// 	- `WindowsCore`
	//
	// Default value: `AliyunLinux`.
	//
	// example:
	//
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The configurations of the private node pool.
	PrivatePoolOptions *CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker RAM role.
	//
	// 	- If you do not specify this parameter, the default worker RAM role created by the cluster is used.
	//
	// 	- The specified RAM role must be a **regular service role*	- and the **Select Trusted Service*	- parameter must be set to **Elastic Compute Service**. For more information, see [Create a normal service role](https://help.aliyun.com/document_detail/116800.html). If the specified RAM role is not the default worker RAM role created by the cluster, the name of the RAM role cannot start with `KubernetesMasterRole-` or `KubernetesWorkerRole-`.
	//
	// >  This parameter is available only for ACK managed clusters that run Kubernetes 1.22 or later.
	//
	// example:
	//
	// example-role
	RamRoleName *string `json:"ram_role_name,omitempty" xml:"ram_role_name,omitempty"`
	// The IDs of ApsaraDB RDS instances.
	RdsInstances        []*string                                                    `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	ResourcePoolOptions *CreateClusterNodePoolRequestScalingGroupResourcePoolOptions `json:"resource_pool_options,omitempty" xml:"resource_pool_options,omitempty" type:"Struct"`
	// The scaling mode of the scaling group. Valid values:
	//
	// 	- `release`: the standard mode. ECS instances are created and released based on resource usage.
	//
	// 	- `recycle`: the swift mode. ECS instances are created, stopped, or started during scaling events. This reduces the time required for the next scale-out event. When the instance is stopped, you are charged only for the storage service. This does not apply to ECS instances that are attached with local disks.
	//
	// Default value: `release`.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// Deprecated
	//
	// The ID of the security group to which you want to add the node pool. You must specify this parameter or the `security_group_ids` parameter. We recommend that you specify `security_group_ids`.
	//
	// example:
	//
	// sg-wz9a8g2mt6x5llu0****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The IDs of security groups. You must specify this parameter or `security_group_id`. We recommend that you specify `security_group_ids`. If you specify both `security_group_id` and `security_group_ids`, `security_group_ids` is used.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// Indicates whether Alibaba Cloud Linux Security Hardening is enabled. Valid values:
	//
	// 	- `true`: enables Alibaba Cloud Linux Security Hardening.
	//
	// 	- `false`: disables Alibaba Cloud Linux Security Hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Specifies whether to enable MLPS Security Hardening. You can enable security hardening based on Multi-Level Protection Scheme (MLPS) only when Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 is installed on nodes. Alibaba Cloud provides standards for baseline checks and a scanner to ensure the compliance of Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 images with the level 3 standards of MLPS 2.0.
	//
	// example:
	//
	// false
	SocEnabled *bool `json:"soc_enabled,omitempty" xml:"soc_enabled,omitempty"`
	// The number of instance types that are available for creating preemptible instances. Auto Scaling creates preemptible instances of multiple instance types that are available at the lowest cost. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int64 `json:"spot_instance_pools,omitempty" xml:"spot_instance_pools,omitempty"`
	// Indicates whether preemptible instances can be supplemented. If you set this parameter to true, when the scaling group receives a system message indicating that a preemptible instance is to be reclaimed, the scaling group attempts to create a new instance to replace this instance. Valid values:
	//
	// 	- `true`: enables the supplementation of preemptible instances.
	//
	// 	- `false`: disables the supplementation of preemptible instances.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The instance type of preemptible instances and the price cap for the instance type.
	SpotPriceLimit []*CreateClusterNodePoolRequestScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The bidding policy of preemptible instances. Valid values:
	//
	// 	- `NoSpot`: non-preemptible.
	//
	// 	- `SpotWithPriceLimit`: specifies the highest bid for the preemptible instance.
	//
	// 	- `SpotAsPriceGo`: automatically submits bids based on the up-to-date market price.
	//
	// For more information, see [Use preemptible instances](https://help.aliyun.com/document_detail/165053.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Specifies whether to enable the burst feature for the system disk. Valid values:
	//
	// 	- true: enables the burst feature.
	//
	// 	- false: disables the burst feature.
	//
	// This parameter is available only when `SystemDiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The categories of the system disk for nodes. The system attempts to create system disks of a disk category with a lower priority if the disk category with a higher priority is unavailable. Valid values:
	//
	// 	- `cloud`: basic disk.
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: ESSD.
	//
	// 	- `cloud_auto`: ESSD AutoPL disk.
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The category of the system disk. Valid values:
	//
	// 	- `cloud`: basic disk.
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: ESSD.
	//
	// 	- `cloud_auto`: ESSD AutoPL disk.
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk.
	//
	// Default value: `cloud_efficiency`.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"system_disk_category,omitempty" xml:"system_disk_category,omitempty"`
	// The encryption algorithm that is used to encrypt the system disk. Set the value to aes-256.
	//
	// example:
	//
	// aes-256
	SystemDiskEncryptAlgorithm *string `json:"system_disk_encrypt_algorithm,omitempty" xml:"system_disk_encrypt_algorithm,omitempty"`
	// Specifies whether to encrypt the system disk. true: encrypts the system disk. false: does not encrypt the system disk.
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
	// The PL of the system disk. This parameter takes effect only for an ESSD.
	//
	// 	- PL0: moderate maximum concurrent I/O performance and low I/O latency.
	//
	// 	- PL1: moderate maximum concurrent I/O performance and low I/O latency.
	//
	// 	- PL2: high maximum concurrent I/O performance and low I/O latency.
	//
	// 	- PL3: ultra-high maximum concurrent I/O performance and ultra-low I/O latency.
	//
	// >  Alibaba Cloud disks support the preceding PLs. However, when you create a disk, the available PLs vary based on the ECS instance type that you selected. For more information, see [Overview of ECS instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// The preset IOPS of the system disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is supported only when `SystemDiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// Valid values: 20 to 20248.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The tags that you want to add only to ECS instances.
	//
	// The tag key must be unique and cannot exceed 128 characters in length. The tag key and value cannot start with aliyun or acs: or contain https:// or http://.
	Tags []*CreateClusterNodePoolRequestScalingGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The vSwitch IDs. You can specify one to eight vSwitch IDs.
	//
	// >  To ensure high availability, we recommend that you select vSwitches that reside in different zones.
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

func (s *CreateClusterNodePoolRequestScalingGroup) SetTags(v []*CreateClusterNodePoolRequestScalingGroupTags) *CreateClusterNodePoolRequestScalingGroup {
	s.Tags = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) SetVswitchIds(v []*string) *CreateClusterNodePoolRequestScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterNodePoolRequestScalingGroup) Validate() error {
	return dara.Validate(s)
}

type CreateClusterNodePoolRequestScalingGroupPrivatePoolOptions struct {
	// The ID of the private node pool.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of private node pool. This parameter specifies the type of private pool that you want to use to create instances. A private pool is generated when an elasticity assurance or a capacity reservation takes effect. The system selects a private pool to start instances. Valid values:
	//
	// 	- `Open`: uses an open private pool. The system selects an open private pool to start instances. If no matching open private pools are available, the resources in the public pool are used.
	//
	// 	- `Target`: uses a specified private pool. The system uses the resources of the specified private pool to start instances. If the specified private pool is unavailable, instances cannot be started.
	//
	// 	- `None`: no private pool is used. The resources of private pools are not used to launch the instances.
	//
	// example:
	//
	// Open
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
	PrivatePoolIds []*string `json:"private_pool_ids,omitempty" xml:"private_pool_ids,omitempty" type:"Repeated"`
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
	// The instance type of preemptible instances.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The price cap of a preemptible instance of the type.
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
	// Specifies whether to enable confidential computing for the cluster.
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
