// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterBasicInfoRequest
	GetClusterId() *string
	SetTargetType(v string) *DescribeClusterBasicInfoRequest
	GetTargetType() *string
	SetType(v string) *DescribeClusterBasicInfoRequest
	GetType() *string
}

type DescribeClusterBasicInfoRequest struct {
	// The ID of the cluster that you want to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// c870ec78ecbcb41d2a35c679823ef****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The dimension from which you want to configure the feature. Valid values:
	//
	// 	- **Cluster**: the ID of the cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- **containerNetwork**: container network topology
	//
	// 	- **interceptionSwitch**: cluster microsegmentation
	//
	// This parameter is required.
	//
	// example:
	//
	// containerNetwork
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeClusterBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterBasicInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterBasicInfoRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeClusterBasicInfoRequest) GetType() *string {
	return s.Type
}

func (s *DescribeClusterBasicInfoRequest) SetClusterId(v string) *DescribeClusterBasicInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterBasicInfoRequest) SetTargetType(v string) *DescribeClusterBasicInfoRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeClusterBasicInfoRequest) SetType(v string) *DescribeClusterBasicInfoRequest {
	s.Type = &v
	return s
}

func (s *DescribeClusterBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
