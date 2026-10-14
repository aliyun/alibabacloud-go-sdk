// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersV1ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeClustersV1ResponseBodyClusters) *DescribeClustersV1ResponseBody
	GetClusters() []*DescribeClustersV1ResponseBodyClusters
	SetRequestId(v string) *DescribeClustersV1ResponseBody
	GetRequestId() *string
}

type DescribeClustersV1ResponseBody struct {
	// The list of clusters.
	Clusters []*DescribeClustersV1ResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClustersV1ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBody) GetClusters() []*DescribeClustersV1ResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClustersV1ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClustersV1ResponseBody) SetClusters(v []*DescribeClustersV1ResponseBodyClusters) *DescribeClustersV1ResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersV1ResponseBody) SetRequestId(v string) *DescribeClustersV1ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersV1ResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClustersV1ResponseBodyClusters struct {
	// The user ID.
	//
	// example:
	//
	// 1375383353108460
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The cluster instance ID.
	//
	// example:
	//
	// eck-xxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster kubeconfig.
	//
	// example:
	//
	// {
	//
	//         "kind": "Config",
	//
	//         "apiVersion": "v1",
	//
	//         "preferences": {},
	//
	//         "clusters": [
	//
	//           {
	//
	//             "name": "kubernetes",
	//
	//             "cluster": {
	//
	//               "server": "https://000.000.000.000:6443",
	//
	//               "certificate-authority-data": ""
	//
	//             }
	//
	//           }
	//
	//         ],
	//
	//         "users": [
	//
	//           {
	//
	//             "name": "kubernetes-admin",
	//
	//             "user": {
	//
	//               "client-certificate-data": "",
	//
	//               "client-key-data": ""
	//
	//             }
	//
	//           }
	//
	//         ],
	//
	//         "contexts": [
	//
	//           {
	//
	//             "name": "kubernetes-admin@kubernetes",
	//
	//             "context": {
	//
	//               "cluster": "kubernetes",
	//
	//               "user": "kubernetes-admin"
	//
	//             }
	//
	//           }
	//
	//         ],
	//
	//         "current-context": "kubernetes-admin@kubernetes"
	//
	//       }
	Config interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The pod network CIDR block. The value must be a valid private CIDR block, which includes the following CIDR blocks and their subnets: 10.0.0.0/8, 172.16-31.0.0/12-16, and 192.168.0.0/16. The CIDR block cannot overlap with the VPC CIDR block or the CIDR blocks used by existing Kubernetes clusters in the VPC. You cannot modify this value after the cluster is created.
	//
	// example:
	//
	// 10.0.0.0/8
	ContainerCidr *string `json:"ContainerCidr,omitempty" xml:"ContainerCidr,omitempty"`
	// The control plane configuration.
	ControlPlaneConfig *DescribeClustersV1ResponseBodyClustersControlPlaneConfig `json:"ControlPlaneConfig,omitempty" xml:"ControlPlaneConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2025-04-07T09:57:26+08:00
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The ENS node ID.
	//
	// example:
	//
	// cn-hangzhou-55
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The authentication token for cluster access.
	//
	// example:
	//
	// xxxxxxxxxx.yyyyyyy
	JoinToken *string `json:"JoinToken,omitempty" xml:"JoinToken,omitempty"`
	// The cluster version, which is consistent with the Kubernetes community baseline version. Use the latest version. If you do not specify this parameter, the latest version is used by default.
	//
	// example:
	//
	// 1.31.9-aliyunedge.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The load balancing instance ID.
	//
	// example:
	//
	// lb-58dngw0fyimzzvwljfec7hy0z
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test-eck-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The vSwitch IDs for pods. Each pod vSwitch corresponds to a worker node vSwitch, and the zones of the pod vSwitch and the worker node vSwitch must be the same. If you leave this parameter empty, the value of vswitch_ids is used by default.
	PodVswitchIds []*string `json:"PodVswitchIds,omitempty" xml:"PodVswitchIds,omitempty" type:"Repeated"`
	// Indicates whether public network access is enabled. The API Server is exposed through an EIP to enable public access to the cluster. Valid values:
	//
	// - true: Public network access is enabled.
	//
	// - false: Public network access is disabled. If you select this option, the cluster API Server cannot be accessed from the Internet.
	//
	// example:
	//
	// true
	PublicAccess *bool `json:"PublicAccess,omitempty" xml:"PublicAccess,omitempty"`
	// The service network CIDR block. Valid ranges: 10.0.0.0/16-24, 172.16-31.0.0/16-24, and 192.168.0.0/16-24. The CIDR block cannot overlap with the VPC CIDR block 10.1.0.0/21 or the CIDR blocks used by existing Kubernetes clusters in the VPC. You cannot modify this value after the cluster is created.
	//
	// example:
	//
	// 192.168.0.0/16
	ServiceCidr *string `json:"ServiceCidr,omitempty" xml:"ServiceCidr,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The VPC used by the cluster.
	//
	// example:
	//
	// n-5wsgr3xeolb2ist303wp3cscp
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch IDs of the cluster nodes. The ENS region of the vSwitch IDs must be the same as the value of ens_region_id.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s DescribeClustersV1ResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClusters) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeClustersV1ResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetConfig() interface{} {
	return s.Config
}

func (s *DescribeClustersV1ResponseBodyClusters) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *DescribeClustersV1ResponseBodyClusters) GetControlPlaneConfig() *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	return s.ControlPlaneConfig
}

func (s *DescribeClustersV1ResponseBodyClusters) GetCreated() *string {
	return s.Created
}

func (s *DescribeClustersV1ResponseBodyClusters) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetIsp() *string {
	return s.Isp
}

func (s *DescribeClustersV1ResponseBodyClusters) GetJoinToken() *string {
	return s.JoinToken
}

func (s *DescribeClustersV1ResponseBodyClusters) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *DescribeClustersV1ResponseBodyClusters) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeClustersV1ResponseBodyClusters) GetPodVswitchIds() []*string {
	return s.PodVswitchIds
}

func (s *DescribeClustersV1ResponseBodyClusters) GetPublicAccess() *bool {
	return s.PublicAccess
}

func (s *DescribeClustersV1ResponseBodyClusters) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *DescribeClustersV1ResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeClustersV1ResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClustersV1ResponseBodyClusters) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClustersV1ResponseBodyClusters) SetAliUid(v string) *DescribeClustersV1ResponseBodyClusters {
	s.AliUid = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetClusterId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetConfig(v interface{}) *DescribeClustersV1ResponseBodyClusters {
	s.Config = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetContainerCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ContainerCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetControlPlaneConfig(v *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) *DescribeClustersV1ResponseBodyClusters {
	s.ControlPlaneConfig = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetCreated(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Created = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetEnsRegionId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetIsp(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Isp = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetJoinToken(v string) *DescribeClustersV1ResponseBodyClusters {
	s.JoinToken = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetKubernetesVersion(v string) *DescribeClustersV1ResponseBodyClusters {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetLoadBalancerId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetName(v string) *DescribeClustersV1ResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetPodVswitchIds(v []*string) *DescribeClustersV1ResponseBodyClusters {
	s.PodVswitchIds = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetPublicAccess(v bool) *DescribeClustersV1ResponseBodyClusters {
	s.PublicAccess = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetServiceCidr(v string) *DescribeClustersV1ResponseBodyClusters {
	s.ServiceCidr = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetState(v string) *DescribeClustersV1ResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVpcId(v string) *DescribeClustersV1ResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) SetVswitchIds(v []*string) *DescribeClustersV1ResponseBodyClusters {
	s.VswitchIds = v
	return s
}

func (s *DescribeClustersV1ResponseBodyClusters) Validate() error {
	if s.ControlPlaneConfig != nil {
		if err := s.ControlPlaneConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersV1ResponseBodyClustersControlPlaneConfig struct {
	// The container runtime type.
	//
	// example:
	//
	// containerd
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" xml:"ContainerRuntime,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-68be8cb9f71fhyvjekxa23qsf
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ens.sn1.medium
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
	// 5
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The system cloud disk type of the node. Valid values: local_ssd, cloud_ssd, cloud_efficiency, and local_hdd.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The system cloud disk size. Unit: GB.
	//
	// example:
	//
	// 100
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeClustersV1ResponseBodyClustersControlPlaneConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetContainerRuntime() *string {
	return s.ContainerRuntime
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetNodePortRange() *string {
	return s.NodePortRange
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetContainerRuntime(v string) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.ContainerRuntime = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetImageId(v string) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.ImageId = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetInstanceSpec(v string) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetNodePortRange(v string) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.NodePortRange = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetSize(v int64) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.Size = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetSystemDiskCategory(v string) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) SetSystemDiskSize(v int64) *DescribeClustersV1ResponseBodyClustersControlPlaneConfig {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeClustersV1ResponseBodyClustersControlPlaneConfig) Validate() error {
	return dara.Validate(s)
}
