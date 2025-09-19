// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTarget(v string) *DescribeClusterInfoListRequest
	GetTarget() *string
	SetTargetType(v string) *DescribeClusterInfoListRequest
	GetTargetType() *string
	SetType(v string) *DescribeClusterInfoListRequest
	GetType() *string
}

type DescribeClusterInfoListRequest struct {
	// The operation value. The value specifies the ID of the cluster.
	//
	// example:
	//
	// c23551de6149343e8a54e69fbefe6****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The dimension based on which you want to configure the feature. Valid values:
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
	// 	- **containerNetwork**: container network
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

func (s DescribeClusterInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterInfoListRequest) GetTarget() *string {
	return s.Target
}

func (s *DescribeClusterInfoListRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeClusterInfoListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeClusterInfoListRequest) SetTarget(v string) *DescribeClusterInfoListRequest {
	s.Target = &v
	return s
}

func (s *DescribeClusterInfoListRequest) SetTargetType(v string) *DescribeClusterInfoListRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeClusterInfoListRequest) SetType(v string) *DescribeClusterInfoListRequest {
	s.Type = &v
	return s
}

func (s *DescribeClusterInfoListRequest) Validate() error {
	return dara.Validate(s)
}
