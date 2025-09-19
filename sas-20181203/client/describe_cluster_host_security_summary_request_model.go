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
	// The key of the condition that is used to query containers. Valid values:
	//
	// 	- **instanceId**: the instance ID
	//
	// 	- **appName**: the name of the application
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **regionId**: the region ID
	//
	// 	- **nodeName**: the name of the node
	//
	// 	- **namespace**: the namespace
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
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the condition that is used to query containers.
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
	// The type of the query. Valid values:
	//
	// 	- **containerId**
	//
	// 	- **uuid**
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
