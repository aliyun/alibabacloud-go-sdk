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
	// The cluster type. Valid values:
	//
	// - `Kubernetes`: ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: ACK managed cluster, including ACK Pro clusters, ACK basic clusters, ACK Serverless Pro clusters, ACK Serverless basic clusters, ACK Edge Pro clusters, and ACK Edge basic clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The Kubernetes version of the cluster. This version is consistent with the Kubernetes community baseline version. We recommend that you select the latest version. If you do not specify this parameter, the latest version is used by default.
	//
	// For more information about the Kubernetes versions supported by ACK, see [Overview of Kubernetes versions](https://help.aliyun.com/document_detail/185269.html).
	//
	// example:
	//
	// 1.32.1-aliyun.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	// The query mode. Valid values:
	//
	// - `supported`: Supported versions.
	//
	// - `creatable`: Versions available for cluster creation.
	//
	// If you specify KubernetesVersion, the system ignores the query mode.
	//
	// If you do not specify a query mode, the system returns versions available for cluster creation by default.
	//
	// example:
	//
	// supported
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The cluster profile. Valid values:
	//
	// - `Default`: cluster for non-edge scenarios.
	//
	// - `Edge`: cluster for edge scenarios.
	//
	// - `Serverless`: ASK cluster.
	//
	// Default value: `Default`.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to query the upgradable versions of the cluster version. This parameter takes effect only when the KubernetesVersion parameter is specified. Valid values:
	//
	// - true: Queries upgradable versions.
	//
	// - false: Does not query upgradable versions.
	//
	// example:
	//
	// 1.30.1-aliyun.1
	QueryUpgradableVersion *bool `json:"QueryUpgradableVersion,omitempty" xml:"QueryUpgradableVersion,omitempty"`
	// The ID of the region where the cluster is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The runtime type. You can specify a runtime type to filter system images supported by the runtime. Valid values:
	//
	// - `docker`: Docker runtime.
	//
	// - `containerd`: containerd runtime.
	//
	// - `Sandboxed-Container.runv`: sandboxed container.
	//
	// If you specify a runtime type, the system returns the image versions supported by the specified runtime.
	//
	// If you do not specify a runtime type, the system returns all images by default.
	//
	// example:
	//
	// containerd
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
