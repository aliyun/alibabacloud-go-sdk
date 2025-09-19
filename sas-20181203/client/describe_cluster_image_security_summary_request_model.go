// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterImageSecuritySummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterImageSecuritySummaryRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeClusterImageSecuritySummaryRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeClusterImageSecuritySummaryRequest
	GetContainerFieldValue() *string
	SetImageDigest(v string) *DescribeClusterImageSecuritySummaryRequest
	GetImageDigest() *string
	SetImageRepoName(v string) *DescribeClusterImageSecuritySummaryRequest
	GetImageRepoName() *string
	SetImageRepoNamespace(v string) *DescribeClusterImageSecuritySummaryRequest
	GetImageRepoNamespace() *string
	SetImageTag(v string) *DescribeClusterImageSecuritySummaryRequest
	GetImageTag() *string
	SetResourceOwnerId(v int64) *DescribeClusterImageSecuritySummaryRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeClusterImageSecuritySummaryRequest
	GetSourceIp() *string
}

type DescribeClusterImageSecuritySummaryRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The key of the condition that is used to query containers. Valid values:
	//
	// 	- **instanceId**: the instance ID of the container
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **regionId**: the region ID of the container
	//
	// 	- **clusterName**: the name of the cluster
	//
	// 	- **image**: the name of the image
	//
	// 	- **imageRepoName**: the name of the image repository
	//
	// 	- **imageRepoNamespace**: the namespace to which the image repository belongs
	//
	// 	- **imageRepoTag**: the tag that is added to the image repository
	//
	// 	- **imageDigest**: the digest of the image
	//
	// 	- **clusterType**: the type of the cluster
	//
	// 	- **hostIp**: the public IP address
	//
	// 	- **pod**: the pod
	//
	// 	- **podIp**: the IP address of the pod
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **vulStatus**: whether vulnerabilities are detected on the container
	//
	// 	- **alarmStatus**: whether alerts are generated for the container
	//
	// 	- **riskStatus**: whether risks are detected on the container
	//
	// 	- **riskLevel**: the risk level of the container
	//
	// 	- **containerScope**: the type of the container
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the condition that is used to query containers.
	//
	// example:
	//
	// c2ac28b2d0c734df29a21d29f18ac****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 402902de6480a020b9f29e7105e77b8a218bc1cccbc3935d3b38c8ea9ba2****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// repo
	ImageRepoName *string `json:"ImageRepoName,omitempty" xml:"ImageRepoName,omitempty"`
	// The namespace of the image repository.
	//
	// example:
	//
	// namespace
	ImageRepoNamespace *string `json:"ImageRepoNamespace,omitempty" xml:"ImageRepoNamespace,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// 3.54.0.1
	ImageTag        *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 60.190.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetImageRepoName() *string {
	return s.ImageRepoName
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetImageRepoNamespace() *string {
	return s.ImageRepoNamespace
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterImageSecuritySummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetClusterId(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetContainerFieldName(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetContainerFieldValue(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetImageDigest(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetImageRepoName(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ImageRepoName = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetImageRepoNamespace(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ImageRepoNamespace = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetImageTag(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetResourceOwnerId(v int64) *DescribeClusterImageSecuritySummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) SetSourceIp(v string) *DescribeClusterImageSecuritySummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryRequest) Validate() error {
	return dara.Validate(s)
}
