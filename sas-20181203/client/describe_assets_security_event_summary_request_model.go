// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsSecurityEventSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAssetsSecurityEventSummaryRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeAssetsSecurityEventSummaryRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeAssetsSecurityEventSummaryRequest
	GetContainerFieldValue() *string
	SetResourceOwnerId(v int64) *DescribeAssetsSecurityEventSummaryRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeAssetsSecurityEventSummaryRequest
	GetSourceIp() *string
}

type DescribeAssetsSecurityEventSummaryRequest struct {
	// The ID of the cluster to which the container belongs.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// cc58f827d893f4d7fb3e34b5d4395****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The key of the condition that is used to query on containers. Valid values:
	//
	// 	- **instanceId**: the ID of the container instance
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
	// 	- **ClusterType**: the type of the cluster
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
	// clusterName
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the condition that is used to query on containers.
	//
	// example:
	//
	// arms-prom-operator
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	ResourceOwnerId     *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 113.108.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAssetsSecurityEventSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsSecurityEventSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetsSecurityEventSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAssetsSecurityEventSummaryRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeAssetsSecurityEventSummaryRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeAssetsSecurityEventSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAssetsSecurityEventSummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAssetsSecurityEventSummaryRequest) SetClusterId(v string) *DescribeAssetsSecurityEventSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryRequest) SetContainerFieldName(v string) *DescribeAssetsSecurityEventSummaryRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryRequest) SetContainerFieldValue(v string) *DescribeAssetsSecurityEventSummaryRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryRequest) SetResourceOwnerId(v int64) *DescribeAssetsSecurityEventSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryRequest) SetSourceIp(v string) *DescribeAssetsSecurityEventSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryRequest) Validate() error {
	return dara.Validate(s)
}
