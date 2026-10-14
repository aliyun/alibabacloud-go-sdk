// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *CreateClusterRequest
	GetClusterType() *string
	SetContainerCidr(v string) *CreateClusterRequest
	GetContainerCidr() *string
	SetControlPlaneConfig(v *CreateClusterRequestControlPlaneConfig) *CreateClusterRequest
	GetControlPlaneConfig() *CreateClusterRequestControlPlaneConfig
	SetEnsRegionId(v string) *CreateClusterRequest
	GetEnsRegionId() *string
	SetIsp(v string) *CreateClusterRequest
	GetIsp() *string
	SetKubernetesVersion(v string) *CreateClusterRequest
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *CreateClusterRequest
	GetLoadBalancerId() *string
	SetName(v string) *CreateClusterRequest
	GetName() *string
	SetNatGatewayId(v string) *CreateClusterRequest
	GetNatGatewayId() *string
	SetPodVswitchIds(v []*string) *CreateClusterRequest
	GetPodVswitchIds() []*string
	SetProfile(v string) *CreateClusterRequest
	GetProfile() *string
	SetPublicAccess(v bool) *CreateClusterRequest
	GetPublicAccess() *bool
	SetServiceCidr(v string) *CreateClusterRequest
	GetServiceCidr() *string
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
	SetVswitchIds(v []*string) *CreateClusterRequest
	GetVswitchIds() []*string
}

type CreateClusterRequest struct {
	// Kubernetes or ManagedKubernetes. The former indicates an ECK cluster, and the latter indicates a serverless VC.
	//
	// Must be used together with the profile parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The pod network CIDR block. The CIDR block must be a valid private CIDR block, which includes the following CIDR blocks and their subnets: 10.0.0.0/8, 172.16-31.0.0/12-16, and 192.168.0.0/16. The CIDR block cannot overlap with the CIDR blocks used by the VPC or existing Kubernetes clusters in the VPC. You cannot modify the CIDR block after the cluster is created.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	// The control plane configuration.
	ControlPlaneConfig *CreateClusterRequestControlPlaneConfig `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty" type:"Struct"`
	// The ENS Region ID of the cluster.
	//
	// example:
	//
	// cn-beijing
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The cluster version, which is consistent with the Kubernetes community baseline version. Select the latest version. If you do not specify this parameter, the latest version is used by default.
	//
	// example:
	//
	// 1.32.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The ID of the CLB instance used for API Server access. If you specify this parameter, automatic creation of a CLB instance for the API Server is skipped.
	//
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// mycluster-1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// Each pod vSwitch corresponds to a worker node vSwitch. The zone of the pod vSwitch must be the same as that of the worker node vSwitch. If you do not specify this parameter, vswitch_ids is used by default.
	PodVswitchIds []*string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty" type:"Repeated"`
	// Default or Serverless. The former indicates an ECK cluster, and the latter indicates a serverless VC.
	//
	// Must be used together with the cluster_type parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to enable public network access. The API Server is exposed through an EIP to enable public access to the cluster.
	//
	// ● true: Enable public network access.
	//
	// ● false: Disable public network access. If you select this option, you cannot access the cluster API Server from the Internet.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// The service network CIDR block. Valid ranges: 10.0.0.0/16-24, 172.16-31.0.0/16-24, and 192.168.0.0/16-24. The CIDR block cannot overlap with the VPC CIDR block 10.1.0.0/21 or the CIDR blocks used by existing Kubernetes clusters in the VPC. You cannot modify the CIDR block after the cluster is created.
	//
	// example:
	//
	// 172.19.0.0/20
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The VPC used by the cluster.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IDs of the vSwitches for cluster nodes. The ENS Region of the vswitch_ids must be the same as the ens_region_id value.
	//
	// example:
	//
	// ["vsw-xxx", "vsw-yyy"]
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *CreateClusterRequest) GetControlPlaneConfig() *CreateClusterRequestControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *CreateClusterRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateClusterRequest) GetIsp() *string {
	return s.Isp
}

func (s *CreateClusterRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *CreateClusterRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateClusterRequest) GetPodVswitchIds() []*string {
	return s.PodVswitchIds
}

func (s *CreateClusterRequest) GetProfile() *string {
	return s.Profile
}

func (s *CreateClusterRequest) GetPublicAccess() *bool {
	return s.PublicAccess
}

func (s *CreateClusterRequest) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetContainerCidr(v string) *CreateClusterRequest {
	s.ContainerCidr = &v
	return s
}

func (s *CreateClusterRequest) SetControlPlaneConfig(v *CreateClusterRequestControlPlaneConfig) *CreateClusterRequest {
	s.ControlPlaneConfig = v
	return s
}

func (s *CreateClusterRequest) SetEnsRegionId(v string) *CreateClusterRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateClusterRequest) SetIsp(v string) *CreateClusterRequest {
	s.Isp = &v
	return s
}

func (s *CreateClusterRequest) SetKubernetesVersion(v string) *CreateClusterRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *CreateClusterRequest) SetLoadBalancerId(v string) *CreateClusterRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetNatGatewayId(v string) *CreateClusterRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateClusterRequest) SetPodVswitchIds(v []*string) *CreateClusterRequest {
	s.PodVswitchIds = v
	return s
}

func (s *CreateClusterRequest) SetProfile(v string) *CreateClusterRequest {
	s.Profile = &v
	return s
}

func (s *CreateClusterRequest) SetPublicAccess(v bool) *CreateClusterRequest {
	s.PublicAccess = &v
	return s
}

func (s *CreateClusterRequest) SetServiceCidr(v string) *CreateClusterRequest {
	s.ServiceCidr = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetVswitchIds(v []*string) *CreateClusterRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	if s.ControlPlaneConfig != nil {
		if err := s.ControlPlaneConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestControlPlaneConfig struct {
	// The control plane virtual machine image ID.
	//
	// example:
	//
	// m-5ul335umat4e2y9ynwi84p3f9
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The node instance type.
	//
	// example:
	//
	// ens.esk.sn1.medium
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The logon password for the cluster control plane machines.
	//
	// example:
	//
	// ThisIsAPassword
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The node service port range.
	//
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"NodePortRange,omitempty" xml:"NodePortRange,omitempty"`
	// The control plane node runtime name.
	//
	// example:
	//
	// containerd
	Runtime *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	// The number of control plane nodes. Valid values: 3 and 5.
	//
	// example:
	//
	// 3
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The node system cloud disk type. Valid values: "local_ssd", "cloud_ssd", "cloud_efficiency", "local_hdd".
	//
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The node system cloud disk size.
	//
	// example:
	//
	// 120
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateClusterRequestControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterRequestControlPlaneConfig) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateClusterRequestControlPlaneConfig) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequestControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *CreateClusterRequestControlPlaneConfig) GetRuntime() *string {
	return s.Runtime
}

func (s *CreateClusterRequestControlPlaneConfig) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateClusterRequestControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *CreateClusterRequestControlPlaneConfig) SetImageId(v string) *CreateClusterRequestControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetInstanceSpec(v string) *CreateClusterRequestControlPlaneConfig {
	s.InstanceSpec = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetLoginPassword(v string) *CreateClusterRequestControlPlaneConfig {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetNodePortRange(v string) *CreateClusterRequestControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetRuntime(v string) *CreateClusterRequestControlPlaneConfig {
	s.Runtime = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSize(v int32) *CreateClusterRequestControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskCategory(v string) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) SetSystemDiskSize(v int64) *CreateClusterRequestControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequestControlPlaneConfig) Validate() error {
	return dara.Validate(s)
}
