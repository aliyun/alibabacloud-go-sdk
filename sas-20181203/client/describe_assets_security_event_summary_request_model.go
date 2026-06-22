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
	// The ID of the container cluster that you want to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// cc58f827d893f4d7fb3e34b5d4395****
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
	// clusterName
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container search field.
	//
	// example:
	//
	// arms-prom-operator
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	ResourceOwnerId     *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
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
