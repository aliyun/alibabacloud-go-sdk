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
	AutoMode *DescribeClusterNodePoolDetailResponseBodyAutoMode `json:"auto_mode,omitempty" xml:"auto_mode,omitempty" type:"Struct"`
	// The auto scaling configuration of the node pool.
	AutoScaling *DescribeClusterNodePoolDetailResponseBodyAutoScaling `json:"auto_scaling,omitempty" xml:"auto_scaling,omitempty" type:"Struct"`
	// Indicates whether the pods in the edge node pool can use the host network.
	//
	// `true`: sets to host network.
	//
	// `false`: sets to container network.
	//
	// example:
	//
	// true
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network,omitempty"`
	// The network configuration of the edge node pool. This parameter takes effect only for edge node pools.
	InterconnectConfig *DescribeClusterNodePoolDetailResponseBodyInterconnectConfig `json:"interconnect_config,omitempty" xml:"interconnect_config,omitempty" type:"Struct"`
	// The network type of the edge node pool. This parameter takes effect only if you set the type parameter of the node pool to edge. Valid values:
	//
	// `basic`: Internet.
	//
	// `private`: private network.
	//
	// example:
	//
	// improved
	InterconnectMode *string `json:"interconnect_mode,omitempty" xml:"interconnect_mode,omitempty"`
	// Specifies whether all nodes in the edge node pool can communicate with each other at Layer 3.
	//
	// `true`: The nodes in the edge node pool can communicate with each other at Layer 3.
	//
	// `false`: The nodes in the edge node pool cannot communicate with each other at Layer 3.
	//
	// example:
	//
	// true
	Intranet *bool `json:"intranet,omitempty" xml:"intranet,omitempty"`
	// The configurations of the cluster in which the node pool is deployed.
	KubernetesConfig *DescribeClusterNodePoolDetailResponseBodyKubernetesConfig `json:"kubernetes_config,omitempty" xml:"kubernetes_config,omitempty" type:"Struct"`
	// The configuration of the managed node pool feature.
	Management *DescribeClusterNodePoolDetailResponseBodyManagement `json:"management,omitempty" xml:"management,omitempty" type:"Struct"`
	// This parameter is deprecated.
	//
	// The maximum number of nodes allowed in an edge node pool.
	//
	// example:
	//
	// 10
	MaxNodes *int64 `json:"max_nodes,omitempty" xml:"max_nodes,omitempty"`
	// The node configurations.
	NodeConfig *DescribeClusterNodePoolDetailResponseBodyNodeConfig `json:"node_config,omitempty" xml:"node_config,omitempty" type:"Struct"`
	// The configuration of the node pool.
	NodepoolInfo *DescribeClusterNodePoolDetailResponseBodyNodepoolInfo `json:"nodepool_info,omitempty" xml:"nodepool_info,omitempty" type:"Struct"`
	// The configurations of the scaling group that is used by the node pool.
	ScalingGroup *DescribeClusterNodePoolDetailResponseBodyScalingGroup `json:"scaling_group,omitempty" xml:"scaling_group,omitempty" type:"Struct"`
	// The status details about the node pool.
	Status *DescribeClusterNodePoolDetailResponseBodyStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// The configuration of confidential computing.
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
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyAutoMode struct {
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
	// The maximum bandwidth of the elastic IP address (EIP).
	//
	// example:
	//
	// 5
	EipBandwidth *int64 `json:"eip_bandwidth,omitempty" xml:"eip_bandwidth,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- `PayByBandwidth`: pay-by-bandwidth.
	//
	// 	- `PayByTraffic`: pay-by-data-transfer.
	//
	// example:
	//
	// PayByBandwidth
	EipInternetChargeType *string `json:"eip_internet_charge_type,omitempty" xml:"eip_internet_charge_type,omitempty"`
	// Indicates whether auto scaling is enabled. Valid values:
	//
	// 	- `true`: auto scaling is enabled.
	//
	// 	- `false`: auto scaling is disabled. If this parameter is set to false, other parameters in the `auto_scaling` section do not take effect.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Indicates whether an EIP is associated with the node pool. Valid values:
	//
	// 	- `true`: An EIP is associated with the node pool.
	//
	// 	- `false`: No EIP is associated with the node pool.
	//
	// example:
	//
	// true
	IsBondEip *bool `json:"is_bond_eip,omitempty" xml:"is_bond_eip,omitempty"`
	// The maximum number of Elastic Compute Service (ECS) instances that can be created in the node pool.
	//
	// example:
	//
	// 10
	MaxInstances *int64 `json:"max_instances,omitempty" xml:"max_instances,omitempty"`
	// The minimum number of ECS instances that must be kept in the node pool.
	//
	// example:
	//
	// 1
	MinInstances *int64 `json:"min_instances,omitempty" xml:"min_instances,omitempty"`
	// The instance types that can be used for the auto scaling of the node pool. Valid values:
	//
	// 	- `cpu`: regular instance.
	//
	// 	- `gpu`: GPU-accelerated instance.
	//
	// 	- `gpushare`: shared GPU-accelerated instance.
	//
	// 	- `spot`: preemptible instance.
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

type DescribeClusterNodePoolDetailResponseBodyInterconnectConfig struct {
	// The bandwidth of the enhanced edge node pool. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// The ID of the Cloud Connect Network (CCN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// ccn-qm5i0i0q9yi*******
	CcnId *string `json:"ccn_id,omitempty" xml:"ccn_id,omitempty"`
	// The region to which the CCN instance that is associated with the enhanced edge node pool belongs.
	//
	// example:
	//
	// cn-shanghai
	CcnRegionId *string `json:"ccn_region_id,omitempty" xml:"ccn_region_id,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance that is associated with the enhanced edge node pool.
	//
	// example:
	//
	// cen-ey9k9nfhz0f*******
	CenId *string `json:"cen_id,omitempty" xml:"cen_id,omitempty"`
	// The subscription duration of the enhanced edge node pool. The duration is measured in months.
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
	// Indicates whether the CloudMonitor agent is installed on ECS nodes in the cluster. After the CloudMonitor agent is installed, you can view monitoring information about the ECS instances in the CloudMonitor console. Installation is recommended. Valid values:
	//
	// 	- `true`: The CloudMonitor agent is installed on ECS nodes.
	//
	// 	- `false`: The CloudMonitor agent is not installed on ECS nodes.
	//
	// example:
	//
	// true
	CmsEnabled *bool `json:"cms_enabled,omitempty" xml:"cms_enabled,omitempty"`
	// The CPU management policy of the nodes in the node pool. The following policies are supported if the Kubernetes version of the cluster is 1.12.6 or later.
	//
	// 	- `static`: allows pods with specific resource characteristics on the node to be granted enhanced CPU affinity and exclusivity.
	//
	// 	- `none`: indicates that the default CPU affinity is used.
	//
	// example:
	//
	// none
	CpuPolicy *string `json:"cpu_policy,omitempty" xml:"cpu_policy,omitempty"`
	// The labels that you want to add to the nodes in the cluster. You must add labels based on the following rules:
	//
	// 	- A label is a case-sensitive key-value pair. You can add up to 20 labels.
	//
	// 	- The key must be unique and cannot exceed 64 characters in length. The value can be empty and cannot exceed 128 characters in length. Keys and values cannot start with `aliyun`, `acs:`, `https://`, or `http://`. For more information, see [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set).
	Labels []*Tag `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// A custom node name consists of a prefix, an IP substring, and a suffix.
	//
	// 	- The prefix and suffix can contain multiple parts that are separated by periods (.). Each part can contain lowercase letters, digits, and hyphens (-). A custom node name must start and end with a digit or lowercase letter.
	//
	// 	- The IP substring length specifies the number of digits to be truncated from the end of the node IP address. The IP substring length ranges from 5 to 12.
	//
	// For example, if the node IP address is 192.168.0.55, the prefix is aliyun.com, the IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.
	//
	// example:
	//
	// customized,test.,5,.com
	NodeNameMode *string `json:"node_name_mode,omitempty" xml:"node_name_mode,omitempty"`
	// The user-defined script that is executed before nodes are initialized. For more information, see [Generate user-defined data](https://help.aliyun.com/document_detail/49121.html).
	//
	// example:
	//
	// dGhpcyBpcyBhIGV4YW1wbGU
	PreUserData *string `json:"pre_user_data,omitempty" xml:"pre_user_data,omitempty"`
	// The name of the container runtime.
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
	// The taints that you want to add to nodes. Taints can be used together with tolerations to prevent pods from being scheduled to specific nodes. For more information, see [taint-and-toleration](https://kubernetes.io/zh/docs/concepts/scheduling-eviction/taint-and-toleration/).
	Taints []*Taint `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
	// Whether the expanded node is schedulable.
	//
	// example:
	//
	// true
	Unschedulable *bool `json:"unschedulable,omitempty" xml:"unschedulable,omitempty"`
	// The custom script to be executed after nodes in the node pool are initialized. For more information, see [Generate user-defined data](https://help.aliyun.com/document_detail/49121.html).
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
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyManagement struct {
	// Indicates whether auto repair is enabled. This parameter takes effect only when `enable=true` is specified. Valid values:
	//
	// 	- `true`: Auto repair is enabled.
	//
	// 	- `false`: Auto repair is disabled.
	//
	// example:
	//
	// true
	AutoRepair *bool `json:"auto_repair,omitempty" xml:"auto_repair,omitempty"`
	// Automatic repair node policy.
	AutoRepairPolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy `json:"auto_repair_policy,omitempty" xml:"auto_repair_policy,omitempty" type:"Struct"`
	// Whether to automatically upgrade.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// Automatic upgrade policy.
	AutoUpgradePolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoUpgradePolicy `json:"auto_upgrade_policy,omitempty" xml:"auto_upgrade_policy,omitempty" type:"Struct"`
	// Whether to automatically fix CVEs.
	//
	// example:
	//
	// true
	AutoVulFix *bool `json:"auto_vul_fix,omitempty" xml:"auto_vul_fix,omitempty"`
	// Automatically repair CVE policies.
	AutoVulFixPolicy *DescribeClusterNodePoolDetailResponseBodyManagementAutoVulFixPolicy `json:"auto_vul_fix_policy,omitempty" xml:"auto_vul_fix_policy,omitempty" type:"Struct"`
	// Indicates whether the managed node pool feature is enabled. Valid values:
	//
	// 	- `true`: The managed node pool feature is enabled.
	//
	// 	- `false`: The managed node pool feature is disabled. Other parameters in this section take effect only when `enable=true` is specified.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The configuration of auto update. The configuration takes effect only when `enable=true` is specified.
	UpgradeConfig *DescribeClusterNodePoolDetailResponseBodyManagementUpgradeConfig `json:"upgrade_config,omitempty" xml:"upgrade_config,omitempty" type:"Struct"`
}

func (s DescribeClusterNodePoolDetailResponseBodyManagement) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolDetailResponseBodyManagement) GoString() string {
	return s.String()
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
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy struct {
	ApprovalRequired *bool `json:"approval_required,omitempty" xml:"approval_required,omitempty"`
	// Whether to allow restarting nodes.
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

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) GetRestartNode() *bool {
	return s.RestartNode
}

func (s *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy) SetApprovalRequired(v bool) *DescribeClusterNodePoolDetailResponseBodyManagementAutoRepairPolicy {
	s.ApprovalRequired = &v
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
	// Whether to allow automatic upgrading of kubelet.
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
	ExcludePackages *string `json:"exclude_packages,omitempty" xml:"exclude_packages,omitempty"`
	// Whether to allow restarting nodes.
	//
	// example:
	//
	// true
	RestartNode *bool `json:"restart_node,omitempty" xml:"restart_node,omitempty"`
	// The vulnerability levels allowed for auto-fixing, separated by commas.
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
	// Indicates whether auto update is enabled. Valid values:
	//
	// 	- `true`: Auto update is enabled.
	//
	// 	- `false`: Auto update is disabled.
	//
	// example:
	//
	// true
	AutoUpgrade *bool `json:"auto_upgrade,omitempty" xml:"auto_upgrade,omitempty"`
	// The maximum number of nodes that can be in the Unavailable state. Valid values: 1 to 1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	MaxUnavailable *int64 `json:"max_unavailable,omitempty" xml:"max_unavailable,omitempty"`
	// The number of additional nodes.
	//
	// example:
	//
	// 5
	Surge *int64 `json:"surge,omitempty" xml:"surge,omitempty"`
	// The percentage of additional nodes to the nodes in the node pool. You must set this parameter or `surge`.
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

type DescribeClusterNodePoolDetailResponseBodyNodeConfig struct {
	// The configurations of the kubelet.
	KubeletConfiguration *KubeletConfig `json:"kubelet_configuration,omitempty" xml:"kubelet_configuration,omitempty"`
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

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) SetKubeletConfiguration(v *KubeletConfig) *DescribeClusterNodePoolDetailResponseBodyNodeConfig {
	s.KubeletConfiguration = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyNodeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyNodepoolInfo struct {
	// The time when the node pool was created.
	//
	// example:
	//
	// 2020-09-27T19:14:09.156823496+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// Indicates whether the node pool is a default node pool. A Container Service for Kubernetes (ACK) cluster usually has only one default node pool. Valid values: `true`: The node pool is a default node pool. `false`: The node pool is not a default node pool.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
	// The name of the node pool.
	//
	// The name must be 1 to 63 characters in length, and can contain digits, letters, and hyphens (-). It cannot start with a hyphen (-).
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyvw3wjmb****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The type of node pool.
	//
	// example:
	//
	// ess
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the node pool was last updated.
	//
	// example:
	//
	// 2020-09-27T20:37:46+08:00
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
	// Indicates whether auto-renewal is enabled for the nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:
	//
	// 	- `true`: Auto-renewal is enabled.
	//
	// 	- `false`: Auto-renewal is disabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// The duration of the auto-renewal. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// If you specify `PeriodUnit=Month`, the valid values are 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"auto_renew_period,omitempty" xml:"auto_renew_period,omitempty"`
	// Deprecated
	//
	// [**Deprecated**] Please use the parameter security_hardening_os instead.
	//
	// example:
	//
	// false
	CisEnabled *bool `json:"cis_enabled,omitempty" xml:"cis_enabled,omitempty"`
	// Indicates whether pay-as-you-go instances are automatically created to meet the required number of ECS instances if preemptible instances cannot be created due to reasons such as cost or insufficient inventory. This parameter takes effect when `multi_az_policy` is set to `COST_OPTIMIZED`. Valid values:
	//
	// 	- `true`: Pay-as-you-go instances are automatically created to meet the required number of ECS instances if preemptible instances cannot be created.
	//
	// 	- `false`: Pay-as-you-go instances are not automatically created to meet the required number of ECS instances if preemptible instances cannot be created.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"compensate_with_on_demand,omitempty" xml:"compensate_with_on_demand,omitempty"`
	// The configurations of the data disks that are attached to the nodes in the node pool. The configurations include the disk category and disk size.
	DataDisks []*DataDisk `json:"data_disks,omitempty" xml:"data_disks,omitempty" type:"Repeated"`
	// The ID of the deployment set to which the ECS instances in the node pool belong.
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
	// The ID of the custom image. You can call the `DescribeKubernetesVersionMetadata` operation to query the images supported by ACK.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Operating system image type.
	//
	// example:
	//
	// AliyunLinux
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// The billing method of the nodes in the node pool. Valid values:
	//
	// 	- `PrePaid`: the subscription billing method.
	//
	// 	- `PostPaid`: the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType      *string                  `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	InstanceMetadataOptions *InstanceMetadataOptions `json:"instance_metadata_options,omitempty" xml:"instance_metadata_options,omitempty"`
	// The instance properties.
	InstancePatterns []*InstancePatterns `json:"instance_patterns,omitempty" xml:"instance_patterns,omitempty" type:"Repeated"`
	// A list of instance types. You can select multiple instance types. When the system needs to create a node, it starts from the first instance type until the node is created. The instance type that is used to create the node varies based on the actual instance stock.
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
	// The billing method of the public IP address of the node.
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
	// The name of the key pair. You must set this parameter or the `login_password` parameter. You must set `key_pair` if the node pool is a managed node pool.
	//
	// example:
	//
	// pro-nodepool
	KeyPair *string `json:"key_pair,omitempty" xml:"key_pair,omitempty"`
	// Whether the popped ECS instance uses a non-root user for login.
	//
	// example:
	//
	// true
	LoginAsNonRoot *bool `json:"login_as_non_root,omitempty" xml:"login_as_non_root,omitempty"`
	// The password for SSH logon. You must set this parameter or the `key_pair` parameter. The password must be 8 to 30 characters in length, and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// For security purposes, the returned password is encrypted.
	//
	// example:
	//
	// ********
	LoginPassword *string `json:"login_password,omitempty" xml:"login_password,omitempty"`
	// The ECS instance scaling policy for a multi-zone scaling group. Valid values:
	//
	// 	- `PRIORITY`: the scaling group is scaled based on the VSwitchIds.N parameter. If an ECS instance cannot be created in the zone where the vSwitch that has the highest priority resides, Auto Scaling creates the ECS instance in the zone where the vSwitch that has the next highest priority resides.
	//
	// 	- `COST_OPTIMIZED`: ECS instances are created based on the vCPU unit price in ascending order. Preemptible instances are preferably created when preemptible instance types are specified in the scaling configuration. You can set the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when preemptible instances cannot be created due to insufficient resources.
	//
	//     **
	//
	//     **Note**The `COST_OPTIMIZED` setting takes effect only when multiple instance types are specified or at least one instance type is specified for preemptible instances.
	//
	// 	- `BALANCE`: ECS instances are evenly distributed across multiple zones specified by the scaling group. If ECS instances become imbalanced among multiple zones due to insufficient inventory, you can call the RebalanceInstances operation of Auto Scaling to balance the instance distribution among zones. For more information, see [RebalanceInstances](https://help.aliyun.com/document_detail/71516.html).
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// BALANCE
	MultiAzPolicy *string `json:"multi_az_policy,omitempty" xml:"multi_az_policy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be kept in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than the value of this parameter, Auto Scaling preferably creates pay-as-you-go instances.
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
	// The subscription duration of worker nodes. This parameter takes effect and is required only when `instance_charge_type` is set to `PrePaid`.
	//
	// If `PeriodUnit=Month` is specified, the valid values are 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 0
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// The billing cycle of the nodes. This parameter is required if `instance_charge_type` is set to `PrePaid`.
	//
	// Valid value: `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// The release version of the operating system. Valid values:
	//
	// 	- `CentOS`
	//
	// 	- `AliyunLinux`
	//
	// 	- `Windows`
	//
	// 	- `WindowsCore`
	//
	// example:
	//
	// AliyunLinux
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The configuration of the private node pool.
	PrivatePoolOptions *DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions `json:"private_pool_options,omitempty" xml:"private_pool_options,omitempty" type:"Struct"`
	// The name of the worker Resource Access Management (RAM) role. The RAM role is assigned to the worker nodes of the cluster to allow the worker nodes to manage ECS instances.
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
	// After you specify the list of RDS instances, the ECS instances in the cluster are automatically added to the whitelist of the RDS instances.
	RdsInstances []*string `json:"rds_instances,omitempty" xml:"rds_instances,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-2zeieod8giqmov7z****
	ScalingGroupId *string `json:"scaling_group_id,omitempty" xml:"scaling_group_id,omitempty"`
	// The scaling mode of the scaling group. Valid values:
	//
	// 	- `release`: the standard mode. ECS instances are created and released based on resource usage.
	//
	// 	- `recycle`: the swift mode. ECS instances are created, stopped, or started during scaling events. This reduces the time required for the next scale-out event. When the instance is stopped, you are charged only for the storage service. This does not apply to ECS instances that are attached with local disks.
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"scaling_policy,omitempty" xml:"scaling_policy,omitempty"`
	// The ID of the security group to which the node pool is added. If the node pool is added to multiple security groups, the first ID in the value of `security_group_ids` is returned.
	//
	// example:
	//
	// sg-2ze60ockeekspl3d****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The IDs of the security groups to which the node pool is added.
	SecurityGroupIds []*string `json:"security_group_ids,omitempty" xml:"security_group_ids,omitempty" type:"Repeated"`
	// Alibaba Cloud OS security hardening. Values:
	//
	// - `true`: Enable Alibaba Cloud OS security hardening.
	//
	// - `false`: Do not enable Alibaba Cloud OS security hardening.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SecurityHardeningOs *bool `json:"security_hardening_os,omitempty" xml:"security_hardening_os,omitempty"`
	// Indicates whether to enable security reinforcement compliant with the hardening standards. This option is available only when the system image is set to Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3. Alibaba Cloud provides baseline check standards and scanning programs compliant with Grade 3, Version 2.0 of the hardening standards for both Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 images.
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
	// Indicates whether preemptible instances are supplemented when the number of preemptible instances drops below the specified minimum number. If this parameter is set to true, when the scaling group receives a system message that a preemptible instance is to be reclaimed, the scaling group attempts to create a new instance to replace this instance. Valid values: Valid values:
	//
	// 	- `true`: Supplementation of preemptible instances is enabled.
	//
	// 	- `false`: Supplementation of preemptible instances is disabled.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"spot_instance_remedy,omitempty" xml:"spot_instance_remedy,omitempty"`
	// The bid configurations of preemptible instances.
	SpotPriceLimit []*DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit `json:"spot_price_limit,omitempty" xml:"spot_price_limit,omitempty" type:"Repeated"`
	// The type of preemptible instance. Valid values:
	//
	// 	- NoSpot: a non-preemptible instance.
	//
	// 	- SpotWithPriceLimit: a preemptible instance that is configured with the highest bid price.
	//
	// 	- SpotAsPriceGo: a preemptible instance for which the system automatically bids based on the current market price.
	//
	// For more information, see [Preemptible instances](https://help.aliyun.com/document_detail/157759.html).
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// Indicates whether to enable the burst feature for the system disk. Valid values:
	//
	// 	- true: enables the burst feature for the system disk. The performance burst feature allows ESSD AutoPL disks to burst their performance when spikes in read/write workloads occur and reduce the performance to the baseline level at the end of workload spikes.
	//
	// 	- false: does not enable the burst feature for the system disk.
	//
	// This parameter is effective only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	SystemDiskBurstingEnabled *bool `json:"system_disk_bursting_enabled,omitempty" xml:"system_disk_bursting_enabled,omitempty"`
	// The categories of the system disk for nodes. The system attempts to create system disks of a disk category with a lower priority if the disk category with a higher priority is unavailable. Valid values: Valid values:
	//
	// 	- `cloud`: basic disk.
	//
	// 	- `cloud_efficiency`: ultra disk.
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: Enterprise SSD (ESSD).
	//
	// 	- `cloud_auto`: ESSD AutoPL disk.
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk.
	//
	// Default value: `cloud_efficiency`.
	SystemDiskCategories []*string `json:"system_disk_categories,omitempty" xml:"system_disk_categories,omitempty" type:"Repeated"`
	// The system disk type. Valid values:
	//
	// 	- `cloud`: basic disk
	//
	// 	- `cloud_efficiency`: ultra disk
	//
	// 	- `cloud_ssd`: standard SSD
	//
	// 	- `cloud_essd`: Enterprise SSD (ESSD)
	//
	// 	- `cloud_auto`: ESSD AutoPL disk
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk
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
	// Specifies whether to encrypt the system disk. Valid values: Valid values:
	//
	// 	- `true`: encrypts the system disk.
	//
	// 	- `false`: does not encrypt the system disk.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"system_disk_encrypted,omitempty" xml:"system_disk_encrypted,omitempty"`
	// System disk\\"s KMS key ID.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKmsKeyId *string `json:"system_disk_kms_key_id,omitempty" xml:"system_disk_kms_key_id,omitempty"`
	// The performance level (PL) of the system disk that you want to use for the node. This parameter takes effect only for enhanced SSDs (ESSDs).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"system_disk_performance_level,omitempty" xml:"system_disk_performance_level,omitempty"`
	// Pre-configured read and write IOPS for the system disk of the node, configured when the disk type is cloud_auto.
	//
	// example:
	//
	// 1000
	SystemDiskProvisionedIops *int64 `json:"system_disk_provisioned_iops,omitempty" xml:"system_disk_provisioned_iops,omitempty"`
	// The system disk size of a node. Unit: GiB.
	//
	// Valid values: 20 to 500.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"system_disk_size,omitempty" xml:"system_disk_size,omitempty"`
	// The labels that you want to add only to ECS instances.
	//
	// The label key must be unique and cannot exceed 128 characters in length. The label key and value cannot start with aliyun or acs: or contain https:// or http://.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The IDs of vSwitches. You can specify 1 to 20 vSwitches.
	//
	// > We recommend that you select vSwitches in different zones to ensure high availability.
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

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetTags(v []*Tag) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.Tags = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) SetVswitchIds(v []*string) *DescribeClusterNodePoolDetailResponseBodyScalingGroup {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterNodePoolDetailResponseBodyScalingGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyScalingGroupPrivatePoolOptions struct {
	// The ID of the private node pool.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of private node pool. This parameter specifies the type of private node pool that you want to use to create instances. A private node pool is generated when an elasticity assurance or a capacity reservation service takes effect. The system selects a private node pool to launch instances. Valid values:
	//
	// 	- `Open`: open private pool. The system selects an open private node pool to launch instances. If no matching open private node pool is available, the resources in the public node pool are used.
	//
	// 	- `Target`: specific private pool. The system uses the resources of the specified private node pool to launch instances. If the specified private node pool is unavailable, instances cannot be launched.
	//
	// 	- `None`: no private node pool is used. The resources of private node pools are not used to launch the instances.
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

type DescribeClusterNodePoolDetailResponseBodyScalingGroupSpotPriceLimit struct {
	// The instance type of the preemptible instances.
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The price cap of a preemptible instance of the type.
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
	// 	- `active`: The node pool is active.
	//
	// 	- `scaling`: The node pool is being scaled.
	//
	// 	- `removing`: Nodes are being removed from the node pool.
	//
	// 	- `deleting`: The node pool is being deleted.
	//
	// 	- `updating`: The node pool is being updated.
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
	return dara.Validate(s)
}

type DescribeClusterNodePoolDetailResponseBodyTeeConfig struct {
	// Indicates whether confidential computing is enabled. Valid values:
	//
	// 	- `true`: Confidential computing is enabled.
	//
	// 	- `false`: Confidential computing is disabled.
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
