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
	// The cluster ID.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container search field. Valid values:
	//
	// - **instanceId**: container instance ID
	//
	// - **clusterId**: cluster ID
	//
	// - **regionId**: container region
	//
	// - **clusterName**: cluster name
	//
	// - **image**: image name
	//
	// - **imageRepoName**: image repository name
	//
	// - **imageRepoNamespace**: image repository namespace
	//
	// - **imageRepoTag**: image repository tag
	//
	// - **imageDigest**: image digest
	//
	// - **clusterType**: cluster type
	//
	// - **hostIp**: public IP address
	//
	// - **pod**: pod
	//
	// - **podIp**: pod IP address
	//
	// - **containerId**: container ID
	//
	// - **vulStatus**: whether the container has vulnerabilities
	//
	// - **alarmStatus**: whether the container has security alerts
	//
	// - **riskStatus**: whether the container has risks
	//
	// - **riskLevel**: container risk level
	//
	// - **containerScope**: container type.
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container search field.
	//
	// example:
	//
	// c2ac28b2d0c734df29a21d29f18ac****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The image digest.
	//
	// example:
	//
	// 402902de6480a020b9f29e7105e77b8a218bc1cccbc3935d3b38c8ea9ba2****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// repo
	ImageRepoName *string `json:"ImageRepoName,omitempty" xml:"ImageRepoName,omitempty"`
	// The image repository namespace.
	//
	// example:
	//
	// namespace
	ImageRepoNamespace *string `json:"ImageRepoNamespace,omitempty" xml:"ImageRepoNamespace,omitempty"`
	// The image tag.
	//
	// example:
	//
	// 3.54.0.1
	ImageTag        *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
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
