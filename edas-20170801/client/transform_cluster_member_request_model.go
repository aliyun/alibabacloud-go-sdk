// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformClusterMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *TransformClusterMemberRequest
	GetInstanceIds() *string
	SetPassword(v string) *TransformClusterMemberRequest
	GetPassword() *string
	SetTargetClusterId(v string) *TransformClusterMemberRequest
	GetTargetClusterId() *string
}

type TransformClusterMemberRequest struct {
	// The ID of the instance that you want to import or migrate. Separate multiple IDs with commas (,).
	//
	// 	- An instance may not belong to a cluster, but an instance can belong to only one cluster at most.
	//
	// 	- The ECS instances and the destination cluster must be in the same virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze7s2v0b789k60p****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The logon password of the ECS instance that you want to import or migrate to the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the destination cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3e3f77b-462e-****-****-bec8727a****
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
}

func (s TransformClusterMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *TransformClusterMemberRequest) GetPassword() *string {
	return s.Password
}

func (s *TransformClusterMemberRequest) GetTargetClusterId() *string {
	return s.TargetClusterId
}

func (s *TransformClusterMemberRequest) SetInstanceIds(v string) *TransformClusterMemberRequest {
	s.InstanceIds = &v
	return s
}

func (s *TransformClusterMemberRequest) SetPassword(v string) *TransformClusterMemberRequest {
	s.Password = &v
	return s
}

func (s *TransformClusterMemberRequest) SetTargetClusterId(v string) *TransformClusterMemberRequest {
	s.TargetClusterId = &v
	return s
}

func (s *TransformClusterMemberRequest) Validate() error {
	return dara.Validate(s)
}
