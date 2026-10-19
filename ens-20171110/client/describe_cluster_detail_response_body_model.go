// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *DescribeClusterDetailResponseBody
	GetAliUid() *string
	SetClusterId(v string) *DescribeClusterDetailResponseBody
	GetClusterId() *string
	SetConfig(v interface{}) *DescribeClusterDetailResponseBody
	GetConfig() interface{}
	SetContainerCidr(v string) *DescribeClusterDetailResponseBody
	GetContainerCidr() *string
	SetControlPlaneConfig(v *DescribeClusterDetailResponseBodyControlPlaneConfig) *DescribeClusterDetailResponseBody
	GetControlPlaneConfig() *DescribeClusterDetailResponseBodyControlPlaneConfig
	SetControlPlaneMeta(v *DescribeClusterDetailResponseBodyControlPlaneMeta) *DescribeClusterDetailResponseBody
	GetControlPlaneMeta() *DescribeClusterDetailResponseBodyControlPlaneMeta
	SetCreated(v string) *DescribeClusterDetailResponseBody
	GetCreated() *string
	SetEnsRegionId(v string) *DescribeClusterDetailResponseBody
	GetEnsRegionId() *string
	SetIsp(v string) *DescribeClusterDetailResponseBody
	GetIsp() *string
	SetJoinToken(v string) *DescribeClusterDetailResponseBody
	GetJoinToken() *string
	SetKubernetesVersion(v string) *DescribeClusterDetailResponseBody
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *DescribeClusterDetailResponseBody
	GetLoadBalancerId() *string
	SetName(v string) *DescribeClusterDetailResponseBody
	GetName() *string
	SetNatGatewayId(v string) *DescribeClusterDetailResponseBody
	GetNatGatewayId() *string
	SetNodeCount(v *DescribeClusterDetailResponseBodyNodeCount) *DescribeClusterDetailResponseBody
	GetNodeCount() *DescribeClusterDetailResponseBodyNodeCount
	SetPodStatus(v *DescribeClusterDetailResponseBodyPodStatus) *DescribeClusterDetailResponseBody
	GetPodStatus() *DescribeClusterDetailResponseBodyPodStatus
	SetPodVswitchIds(v []*string) *DescribeClusterDetailResponseBody
	GetPodVswitchIds() []*string
	SetPublicAccess(v bool) *DescribeClusterDetailResponseBody
	GetPublicAccess() *bool
	SetRequestId(v string) *DescribeClusterDetailResponseBody
	GetRequestId() *string
	SetServiceCidr(v string) *DescribeClusterDetailResponseBody
	GetServiceCidr() *string
	SetState(v string) *DescribeClusterDetailResponseBody
	GetState() *string
	SetVpcId(v string) *DescribeClusterDetailResponseBody
	GetVpcId() *string
	SetVswitchIds(v []*string) *DescribeClusterDetailResponseBody
	GetVswitchIds() []*string
}

type DescribeClusterDetailResponseBody struct {
	// The user ID.
	//
	// example:
	//
	// 1574790082031102
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// eck-11111111
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster kubeconfig.
	//
	// example:
	//
	// apiVersion: v1
	//
	// clusters:
	//
	// - cluster:
	//
	//     certificate-authority-data: x
	//
	//     server: https://111.111.111.111:6443
	//
	//   name: kubernetes
	//
	// contexts:
	//
	// - context:
	//
	//     cluster: kubernetes
	//
	//     user: user
	//
	//   name: eck-xxxxx
	//
	// current-context: eck-xxxx
	//
	// kind: Config
	//
	// preferences: {}
	//
	// users:
	//
	// - name: user
	//
	//   user:
	//
	//     client-certificate-data: x
	//
	//     client-key-data: x
	Config interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The pod CIDR block.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	// The control plane configuration.
	ControlPlaneConfig *DescribeClusterDetailResponseBodyControlPlaneConfig `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty" type:"Struct"`
	// The control plane metadata.
	ControlPlaneMeta *DescribeClusterDetailResponseBodyControlPlaneMeta `json:"ControlPlaneMeta,omitempty" xml:"ControlPlaneMeta,omitempty" type:"Struct"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2025-04-07T09:57:26+08:00
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The ENS edge zone ID.
	//
	// example:
	//
	// cn-fuzhou-23
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ISP information.
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The authentication token for joining the cluster.
	//
	// example:
	//
	// xxxx.yyy
	JoinToken *string `json:"JoinToken,omitempty" xml:"JoinToken,omitempty"`
	// The cluster version, which aligns with the Kubernetes community baseline version. Select the latest version. If not specified, the latest version is used by default.
	//
	// example:
	//
	// 1.32.1.aliyunedge.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance for the API server.
	//
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// your-cluster-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5tawjw5j7sgd2deujxuk0****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The number of nodes in the cluster.
	NodeCount *DescribeClusterDetailResponseBodyNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	// The pod status information of the cluster.
	PodStatus *DescribeClusterDetailResponseBodyPodStatus `json:"PodStatus,omitempty" xml:"PodStatus,omitempty" type:"Struct"`
	// The pod vSwitch IDs.
	PodVswitchIds []*string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty" type:"Repeated"`
	// Indicates whether public network access is enabled.
	//
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service CIDR block.
	//
	// example:
	//
	// 172.19.0.0/20
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The cluster state. Valid values:
	//
	// - **initial**: Creating.
	//
	// - **failed**: Creation failed.
	//
	// - **running**: Running.
	//
	// - **deleting**: Releasing.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The VPC used by the cluster.
	//
	// example:
	//
	// n-xxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch IDs where the cluster nodes reside.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s DescribeClusterDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeClusterDetailResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterDetailResponseBody) GetConfig() interface{} {
	return s.Config
}

func (s *DescribeClusterDetailResponseBody) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *DescribeClusterDetailResponseBody) GetControlPlaneConfig() *DescribeClusterDetailResponseBodyControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *DescribeClusterDetailResponseBody) GetControlPlaneMeta() *DescribeClusterDetailResponseBodyControlPlaneMeta {
	return s.ControlPlaneMeta
}

func (s *DescribeClusterDetailResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterDetailResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClusterDetailResponseBody) GetIsp() *string {
	return s.Isp
}

func (s *DescribeClusterDetailResponseBody) GetJoinToken() *string {
	return s.JoinToken
}

func (s *DescribeClusterDetailResponseBody) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *DescribeClusterDetailResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeClusterDetailResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeClusterDetailResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeClusterDetailResponseBody) GetNodeCount() *DescribeClusterDetailResponseBodyNodeCount {
	return s.NodeCount
}

func (s *DescribeClusterDetailResponseBody) GetPodStatus() *DescribeClusterDetailResponseBodyPodStatus {
	return s.PodStatus
}

func (s *DescribeClusterDetailResponseBody) GetPodVswitchIds() []*string {
	return s.PodVswitchIds
}

func (s *DescribeClusterDetailResponseBody) GetPublicAccess() *bool {
	return s.PublicAccess
}

func (s *DescribeClusterDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterDetailResponseBody) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *DescribeClusterDetailResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeClusterDetailResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterDetailResponseBody) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClusterDetailResponseBody) SetAliUid(v string) *DescribeClusterDetailResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterId(v string) *DescribeClusterDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetConfig(v interface{}) *DescribeClusterDetailResponseBody {
	s.Config = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetContainerCidr(v string) *DescribeClusterDetailResponseBody {
	s.ContainerCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetControlPlaneConfig(v *DescribeClusterDetailResponseBodyControlPlaneConfig) *DescribeClusterDetailResponseBody {
	s.ControlPlaneConfig = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetControlPlaneMeta(v *DescribeClusterDetailResponseBodyControlPlaneMeta) *DescribeClusterDetailResponseBody {
	s.ControlPlaneMeta = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCreated(v string) *DescribeClusterDetailResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetEnsRegionId(v string) *DescribeClusterDetailResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetIsp(v string) *DescribeClusterDetailResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetJoinToken(v string) *DescribeClusterDetailResponseBody {
	s.JoinToken = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetKubernetesVersion(v string) *DescribeClusterDetailResponseBody {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetLoadBalancerId(v string) *DescribeClusterDetailResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetName(v string) *DescribeClusterDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNatGatewayId(v string) *DescribeClusterDetailResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNodeCount(v *DescribeClusterDetailResponseBodyNodeCount) *DescribeClusterDetailResponseBody {
	s.NodeCount = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPodStatus(v *DescribeClusterDetailResponseBodyPodStatus) *DescribeClusterDetailResponseBody {
	s.PodStatus = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPodVswitchIds(v []*string) *DescribeClusterDetailResponseBody {
	s.PodVswitchIds = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPublicAccess(v bool) *DescribeClusterDetailResponseBody {
	s.PublicAccess = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRequestId(v string) *DescribeClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetServiceCidr(v string) *DescribeClusterDetailResponseBody {
	s.ServiceCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetState(v string) *DescribeClusterDetailResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVpcId(v string) *DescribeClusterDetailResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchIds(v []*string) *DescribeClusterDetailResponseBody {
	s.VswitchIds = v
	return s
}

func (s *DescribeClusterDetailResponseBody) Validate() error {
	if s.ControlPlaneConfig != nil {
		if err := s.ControlPlaneConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ControlPlaneMeta != nil {
		if err := s.ControlPlaneMeta.Validate(); err != nil {
			return err
		}
	}
	if s.NodeCount != nil {
		if err := s.NodeCount.Validate(); err != nil {
			return err
		}
	}
	if s.PodStatus != nil {
		if err := s.PodStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyControlPlaneConfig struct {
	// The runtime type.
	//
	// example:
	//
	// containerd
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" xml:"ContainerRuntime,omitempty"`
	// The ID of the control plane virtual machine image.
	//
	// example:
	//
	// m-5ul335umat4e2y9ynwi84p3f9
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type of the node.
	//
	// example:
	//
	// ens.esk.sn1.medium
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The node service port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"NodePortRange,omitempty" xml:"NodePortRange,omitempty"`
	// The number of control plane nodes.
	//
	// example:
	//
	// 3
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The system cloud disk category of the node.
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The system cloud disk size of the node.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetContainerRuntime() *string {
	return s.ContainerRuntime
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetContainerRuntime(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.ContainerRuntime = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetImageId(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetInstanceSpec(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetNodePortRange(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSize(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskCategory(v string) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) SetSystemDiskSize(v int64) *DescribeClusterDetailResponseBodyControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyControlPlaneMeta struct {
	// The master node information.
	MasterInstances []*DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances `json:"MasterInstances,omitempty" xml:"MasterInstances,omitempty" type:"Repeated"`
	// The API server address.
	MasterUrl *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl `json:"MasterUrl,omitempty" xml:"MasterUrl,omitempty" type:"Struct"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneMeta) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMeta) GetMasterInstances() []*DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances {
	return s.MasterInstances
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMeta) GetMasterUrl() *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl {
	return s.MasterUrl
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMeta) SetMasterInstances(v []*DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) *DescribeClusterDetailResponseBodyControlPlaneMeta {
	s.MasterInstances = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMeta) SetMasterUrl(v *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) *DescribeClusterDetailResponseBodyControlPlaneMeta {
	s.MasterUrl = v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMeta) Validate() error {
	if s.MasterInstances != nil {
		for _, item := range s.MasterInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MasterUrl != nil {
		if err := s.MasterUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances struct {
	// The instance ID.
	//
	// example:
	//
	// i-573utia8wbecrbgsm6o8fx132
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IP address.
	//
	// example:
	//
	// 10.10.0.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The node name.
	//
	// example:
	//
	// i-573utia8wbecrbgsm6o8fx132
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) GetIp() *string {
	return s.Ip
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) GetName() *string {
	return s.Name
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) SetInstanceId(v string) *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) SetIp(v string) *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances {
	s.Ip = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) SetName(v string) *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances {
	s.Name = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl struct {
	// The private endpoint of the API server.
	//
	// example:
	//
	// https://10.0.0.1:6443
	PrivateUrl *string `json:"PrivateUrl,omitempty" xml:"PrivateUrl,omitempty"`
	// The public address of the API server.
	//
	// example:
	//
	// https://0.0.0.0:6443
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
}

func (s DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) GetPrivateUrl() *string {
	return s.PrivateUrl
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) SetPrivateUrl(v string) *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl {
	s.PrivateUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) SetPublicUrl(v string) *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl {
	s.PublicUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyControlPlaneMetaMasterUrl) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyNodeCount struct {
	// The number of ready nodes.
	//
	// example:
	//
	// 3
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The total number of nodes.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeClusterDetailResponseBodyNodeCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyNodeCount) GetReady() *int32 {
	return s.Ready
}

func (s *DescribeClusterDetailResponseBodyNodeCount) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeClusterDetailResponseBodyNodeCount) SetReady(v int32) *DescribeClusterDetailResponseBodyNodeCount {
	s.Ready = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyNodeCount) SetTotal(v int32) *DescribeClusterDetailResponseBodyNodeCount {
	s.Total = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyNodeCount) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterDetailResponseBodyPodStatus struct {
	// The number of pods in the Running state.
	//
	// example:
	//
	// 10
	Running *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	// The total number of pods.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeClusterDetailResponseBodyPodStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyPodStatus) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyPodStatus) GetRunning() *int32 {
	return s.Running
}

func (s *DescribeClusterDetailResponseBodyPodStatus) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeClusterDetailResponseBodyPodStatus) SetRunning(v int32) *DescribeClusterDetailResponseBodyPodStatus {
	s.Running = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyPodStatus) SetTotal(v int32) *DescribeClusterDetailResponseBodyPodStatus {
	s.Total = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyPodStatus) Validate() error {
	return dara.Validate(s)
}
