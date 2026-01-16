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
	SetEnsRegionId(v string) *DescribeClusterDetailResponseBody
	GetEnsRegionId() *string
	SetJoinToken(v string) *DescribeClusterDetailResponseBody
	GetJoinToken() *string
	SetKubernetesVersion(v string) *DescribeClusterDetailResponseBody
	GetKubernetesVersion() *string
	SetLoadBalancerId(v string) *DescribeClusterDetailResponseBody
	GetLoadBalancerId() *string
	SetName(v string) *DescribeClusterDetailResponseBody
	GetName() *string
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
	// example:
	//
	// 1574790082031102
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// eck-11111111
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr      *string                                              `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	ControlPlaneConfig *DescribeClusterDetailResponseBodyControlPlaneConfig `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty" type:"Struct"`
	// example:
	//
	// cn-fuzhou-23
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// xxxx.yyy
	JoinToken *string `json:"JoinToken,omitempty" xml:"JoinToken,omitempty"`
	// example:
	//
	// 1.32.1.aliyunedge.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// example:
	//
	// your-cluster-name
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PodVswitchIds []*string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 172.19.0.0/20
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// n-xxxxxxxxx
	VpcId      *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeClusterDetailResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
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

func (s *DescribeClusterDetailResponseBody) SetEnsRegionId(v string) *DescribeClusterDetailResponseBody {
	s.EnsRegionId = &v
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
	return nil
}

type DescribeClusterDetailResponseBodyControlPlaneConfig struct {
	// example:
	//
	// containerd
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" xml:"ContainerRuntime,omitempty"`
	// example:
	//
	// m-5ul335umat4e2y9ynwi84p3f9
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// ens.esk.sn1.medium
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// example:
	//
	// 30000-32767
	NodePortRange *string `json:"NodePortRange,omitempty" xml:"NodePortRange,omitempty"`
	// example:
	//
	// 3
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// cloud_efficiency
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
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
