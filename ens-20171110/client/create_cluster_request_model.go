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
	SetClusterVersion(v string) *CreateClusterRequest
	GetClusterVersion() *string
	SetContainerCidr(v string) *CreateClusterRequest
	GetContainerCidr() *string
	SetControlPlaneConfig(v *CreateClusterRequestControlPlaneConfig) *CreateClusterRequest
	GetControlPlaneConfig() *CreateClusterRequestControlPlaneConfig
	SetEnsRegionId(v string) *CreateClusterRequest
	GetEnsRegionId() *string
	SetKubernetesVersion(v string) *CreateClusterRequest
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *CreateClusterRequest
	GetLoadBalancerId() *string
	SetName(v string) *CreateClusterRequest
	GetName() *string
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
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.18.8
	ClusterVersion *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr      *string                                 `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	ControlPlaneConfig *CreateClusterRequestControlPlaneConfig `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty" type:"Struct"`
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
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PodVswitchIds []*string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty" type:"Repeated"`
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

func (s *CreateClusterRequest) GetClusterVersion() *string {
	return s.ClusterVersion
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

func (s *CreateClusterRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *CreateClusterRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateClusterRequest) GetName() *string {
	return s.Name
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

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
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
	// example:
	//
	// m-5ul335umat4e2y9ynwi84p3f9
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// ens.esk.sn1.medium
	InstanceSpec  *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"NodePortRange,omitempty" xml:"NodePortRange,omitempty"`
	// example:
	//
	// containerd
	Runtime *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	// example:
	//
	// 3
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
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
