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
	// Kubernetes or ManagedKubernetes. The former indicates an ECK cluster, and the latter indicates a serverless VC. This parameter must be used together with the Profile parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The CIDR block for pods. It must be a valid private CIDR block, such as 10.0.0.0/8, 172.16.0.0/12 to 172.31.0.0/16, or 192.168.0.0/16, and their subnets. The CIDR block cannot overlap with the CIDR block of the VPC or the CIDR blocks used by existing Kubernetes clusters in the VPC. The CIDR block cannot be modified after the cluster is created.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	// The control plane configuration.
	ControlPlaneConfigShrink *string `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty"`
	// The ID of the Edge Node Service (ENS) region where the cluster is deployed.
	//
	// example:
	//
	// cn-beijing
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The Internet service provider (ISP) information. Valid values:
	//
	// - **cmcc**: China Mobile.
	//
	// - **unicom**: China Unicom.
	//
	// - **telecom**: China Telecom.
	//
	// example:
	//
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The Kubernetes version of the cluster. The version aligns with the Kubernetes community baseline version. We recommend that you select the latest version. If you do not specify this parameter, the latest version is used by default.
	//
	// example:
	//
	// 1.32.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance used for API server access. If you specify this parameter, automatic creation of a CLB instance for the API server is skipped.
	//
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// mycluster-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway used by the cluster. If you specify this parameter, automatic creation of a NAT gateway is skipped.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The IDs of the pod vSwitches. Each pod vSwitch corresponds to a worker node vSwitch. The pod vSwitch and the worker node vSwitch must be in the same zone. If you do not specify this parameter, the vSwitches specified by the VswitchIds parameter are used by default.
	PodVswitchIdsShrink *string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty"`
	// Default or Serverless. The former indicates an ECK cluster, and the latter indicates a serverless VC. This parameter must be used together with the ClusterType parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to enable public network access. You can expose the API server by using an elastic IP address (EIP) to enable public network access to the cluster. Valid values:
	//
	// - true: enables public network access.
	//
	// - false: disables public network access. If you select false, you cannot access the API server of the cluster over the Internet.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// The CIDR block for Services. Valid ranges include 10.0.0.0/16 to 10.0.0.0/24, 172.16.0.0/16 to 172.31.0.0/24, and 192.168.0.0/16 to 192.168.0.0/24. The CIDR block cannot overlap with the VPC CIDR block 10.1.0.0/21 or the CIDR blocks used by existing Kubernetes clusters in the VPC. The CIDR block cannot be modified after the cluster is created.
	//
	// example:
	//
	// 172.19.0.0/20
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The ID of the virtual private cloud (VPC) used by the cluster.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IDs of the vSwitches where the cluster nodes are deployed. The ENS region of the vSwitches must be the same as the ENS region specified by the EnsRegionId parameter.
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
