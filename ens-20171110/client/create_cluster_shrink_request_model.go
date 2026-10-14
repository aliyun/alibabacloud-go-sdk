// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *CreateClusterShrinkRequest
	GetClusterType() *string
	SetContainerCidr(v string) *CreateClusterShrinkRequest
	GetContainerCidr() *string
	SetControlPlaneConfigShrink(v string) *CreateClusterShrinkRequest
	GetControlPlaneConfigShrink() *string
	SetEnsRegionId(v string) *CreateClusterShrinkRequest
	GetEnsRegionId() *string
	SetIsp(v string) *CreateClusterShrinkRequest
	GetIsp() *string
	SetKubernetesVersion(v string) *CreateClusterShrinkRequest
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *CreateClusterShrinkRequest
	GetLoadBalancerId() *string
	SetName(v string) *CreateClusterShrinkRequest
	GetName() *string
	SetNatGatewayId(v string) *CreateClusterShrinkRequest
	GetNatGatewayId() *string
	SetPodVswitchIdsShrink(v string) *CreateClusterShrinkRequest
	GetPodVswitchIdsShrink() *string
	SetProfile(v string) *CreateClusterShrinkRequest
	GetProfile() *string
	SetPublicAccess(v bool) *CreateClusterShrinkRequest
	GetPublicAccess() *bool
	SetServiceCidr(v string) *CreateClusterShrinkRequest
	GetServiceCidr() *string
	SetVpcId(v string) *CreateClusterShrinkRequest
	GetVpcId() *string
	SetVswitchIdsShrink(v string) *CreateClusterShrinkRequest
	GetVswitchIdsShrink() *string
}

type CreateClusterShrinkRequest struct {
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
	ControlPlaneConfigShrink *string `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty"`
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
	PodVswitchIdsShrink *string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty"`
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
	VswitchIdsShrink *string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterShrinkRequest) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *CreateClusterShrinkRequest) GetControlPlaneConfigShrink() *string {
	return s.ControlPlaneConfigShrink
}

func (s *CreateClusterShrinkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateClusterShrinkRequest) GetIsp() *string {
	return s.Isp
}

func (s *CreateClusterShrinkRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *CreateClusterShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateClusterShrinkRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateClusterShrinkRequest) GetPodVswitchIdsShrink() *string {
	return s.PodVswitchIdsShrink
}

func (s *CreateClusterShrinkRequest) GetProfile() *string {
	return s.Profile
}

func (s *CreateClusterShrinkRequest) GetPublicAccess() *bool {
	return s.PublicAccess
}

func (s *CreateClusterShrinkRequest) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *CreateClusterShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterShrinkRequest) GetVswitchIdsShrink() *string {
	return s.VswitchIdsShrink
}

func (s *CreateClusterShrinkRequest) SetClusterType(v string) *CreateClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetContainerCidr(v string) *CreateClusterShrinkRequest {
	s.ContainerCidr = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetControlPlaneConfigShrink(v string) *CreateClusterShrinkRequest {
	s.ControlPlaneConfigShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetEnsRegionId(v string) *CreateClusterShrinkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIsp(v string) *CreateClusterShrinkRequest {
	s.Isp = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetKubernetesVersion(v string) *CreateClusterShrinkRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetLoadBalancerId(v string) *CreateClusterShrinkRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetName(v string) *CreateClusterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNatGatewayId(v string) *CreateClusterShrinkRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetPodVswitchIdsShrink(v string) *CreateClusterShrinkRequest {
	s.PodVswitchIdsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetProfile(v string) *CreateClusterShrinkRequest {
	s.Profile = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetPublicAccess(v bool) *CreateClusterShrinkRequest {
	s.PublicAccess = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetServiceCidr(v string) *CreateClusterShrinkRequest {
	s.ServiceCidr = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVpcId(v string) *CreateClusterShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVswitchIdsShrink(v string) *CreateClusterShrinkRequest {
	s.VswitchIdsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
