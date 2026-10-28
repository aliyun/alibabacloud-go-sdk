// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterHostSecuritySummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterHostSecuritySummaryRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeClusterHostSecuritySummaryRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeClusterHostSecuritySummaryRequest
	GetContainerFieldValue() *string
	SetResourceOwnerId(v int64) *DescribeClusterHostSecuritySummaryRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeClusterHostSecuritySummaryRequest
	GetSourceIp() *string
	SetTargetType(v string) *DescribeClusterHostSecuritySummaryRequest
	GetTargetType() *string
}

type DescribeClusterHostSecuritySummaryRequest struct {
	// The ID of the container cluster.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container search field. Valid values:
	//
	// - **instanceId**: The instance ID.
	//
	// - **appName**: The application name.
	//
	// - **clusterId**: The cluster ID.
	//
	// - **regionId**: The region ID.
	//
	// - **nodeName**: The node name.
	//
	// - **namespace**: The namespace.
	//
	// - **clusterName**: The cluster name.
	//
	// - **image**: The image name.
	//
	// - **imageRepoName**: The image repository name.
	//
	// - **imageRepoNamespace**: The image repository namespace.
	//
	// - **imageRepoTag**: The image tag.
	//
	// - **imageDigest**: The image digest.
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container search field.
	//
	// example:
	//
	// c1fdb5fd8d42e425d88fd73eec7be****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	ResourceOwnerId     *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 222.71.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The query type. Valid values:
	//
	// - **containerId**: The container ID.
	//
	// - **uuid**: The asset ID.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeClusterHostSecuritySummaryRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetClusterId(v string) *DescribeClusterHostSecuritySummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetContainerFieldName(v string) *DescribeClusterHostSecuritySummaryRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetContainerFieldValue(v string) *DescribeClusterHostSecuritySummaryRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetResourceOwnerId(v int64) *DescribeClusterHostSecuritySummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetSourceIp(v string) *DescribeClusterHostSecuritySummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) SetTargetType(v string) *DescribeClusterHostSecuritySummaryRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryRequest) Validate() error {
	return dara.Validate(s)
}
