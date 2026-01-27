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
	SetKubernetesVersion(v string) *CreateClusterShrinkRequest
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *CreateClusterShrinkRequest
	GetLoadBalancerId() *string
	SetName(v string) *CreateClusterShrinkRequest
	GetName() *string
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
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr            *string `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	ControlPlaneConfigShrink *string `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty"`
	// example:
	//
	// cn-beijing
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// 1.32.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// example:
	//
	// lb-wz9t256gqa3vbouk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// mycluster-1
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PodVswitchIdsShrink *string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// example:
	//
	// 172.19.0.0/20
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateClusterShrinkRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *CreateClusterShrinkRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterShrinkRequest) GetName() *string {
	return s.Name
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
