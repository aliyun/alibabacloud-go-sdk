// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKubernetesVersionMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *DescribeKubernetesVersionMetadataRequest
	GetClusterType() *string
	SetKubernetesVersion(v string) *DescribeKubernetesVersionMetadataRequest
	GetKubernetesVersion() *string
	SetMode(v string) *DescribeKubernetesVersionMetadataRequest
	GetMode() *string
	SetProfile(v string) *DescribeKubernetesVersionMetadataRequest
	GetProfile() *string
	SetQueryUpgradableVersion(v bool) *DescribeKubernetesVersionMetadataRequest
	GetQueryUpgradableVersion() *bool
	SetRegion(v string) *DescribeKubernetesVersionMetadataRequest
	GetRegion() *string
	SetRuntime(v string) *DescribeKubernetesVersionMetadataRequest
	GetRuntime() *string
}

type DescribeKubernetesVersionMetadataRequest struct {
	// The cluster type that you want to use. Valid values:
	//
	// 	- `Kubernetes`: ACK dedicated cluster.
	//
	// 	- `ManagedKubernetes`: ACK managed cluster. ACK managed clusters include ACK Pro clusters, ACK Basic clusters, ACK Serverless Pro clusters, ACK Serverless Basic clusters, ACK Edge Pro clusters, and ACK Edge Basic clusters.
	//
	// 	- `ExternalKubernetes`: registered cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The Kubernetes version of the cluster. The Kubernetes versions supported by ACK are the same as the Kubernetes versions supported by open source Kubernetes. We recommend that you specify the latest Kubernetes version. If you do not configure this parameter, the latest Kubernetes version is used.
	//
	// For more information about the Kubernetes versions supported by ACK, see [Release notes for Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.16.9-aliyun.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The query mode. Valid values:
	//
	// 	- `supported`: queries all supported Kubernetes versions.
	//
	// 	- `creatable`: queries only Kubernetes versions of clusters that you can create.
	//
	// If you specify `KubernetesVersion`, this parameter does not take effect.
	//
	// If you do not specify a query mode, Kubernetes versions of clusters that you can create are returned.
	//
	// example:
	//
	// supported
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The scenario where clusters are used. Valid values:
	//
	// 	- `Default`: non-edge computing scenarios
	//
	// 	- `Edge`: edge computing scenarios
	//
	// 	- `Serverless`: serverless scenarios.
	//
	// Default value: `Default`.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to query the Kubernetes versions available for updates. This parameter takes effect only when the KubernetesVersion parameter is specified.
	//
	// 	- true: queries the Kubernetes versions available for updates.
	//
	// 	- false: does not query the Kubernetes versions available for updates.
	//
	// example:
	//
	// 1.30.1-aliyun.1
	QueryUpgradableVersion *bool `json:"QueryUpgradableVersion,omitempty" xml:"QueryUpgradableVersion,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The container runtime type that you want to use. You can specify a runtime type to query only OS images that support the runtime type. Valid values:
	//
	// 	- `docker`: Docker
	//
	// 	- `containerd`: containerd
	//
	// 	- `Sandboxed-Container.runv`: Sandboxed-Container
	//
	// If you specify a runtime type, only the OS images that support the specified runtime type are returned.
	//
	// Otherwise, all OS images are returned.
	//
	// example:
	//
	// docker
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
}

func (s DescribeKubernetesVersionMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKubernetesVersionMetadataRequest) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesVersionMetadataRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeKubernetesVersionMetadataRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *DescribeKubernetesVersionMetadataRequest) GetMode() *string {
	return s.Mode
}

func (s *DescribeKubernetesVersionMetadataRequest) GetProfile() *string {
	return s.Profile
}

func (s *DescribeKubernetesVersionMetadataRequest) GetQueryUpgradableVersion() *bool {
	return s.QueryUpgradableVersion
}

func (s *DescribeKubernetesVersionMetadataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeKubernetesVersionMetadataRequest) GetRuntime() *string {
	return s.Runtime
}

func (s *DescribeKubernetesVersionMetadataRequest) SetClusterType(v string) *DescribeKubernetesVersionMetadataRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetKubernetesVersion(v string) *DescribeKubernetesVersionMetadataRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetMode(v string) *DescribeKubernetesVersionMetadataRequest {
	s.Mode = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetProfile(v string) *DescribeKubernetesVersionMetadataRequest {
	s.Profile = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetQueryUpgradableVersion(v bool) *DescribeKubernetesVersionMetadataRequest {
	s.QueryUpgradableVersion = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetRegion(v string) *DescribeKubernetesVersionMetadataRequest {
	s.Region = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) SetRuntime(v string) *DescribeKubernetesVersionMetadataRequest {
	s.Runtime = &v
	return s
}

func (s *DescribeKubernetesVersionMetadataRequest) Validate() error {
	return dara.Validate(s)
}
