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
	// The IDs of the ECS instances. Separate multiple IDs with a comma (,).
	//
	// - The instances must be in the same VPC as the target cluster.
	//
	// - An instance can belong to only one cluster at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze7s2v0b789k60p****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The logon password to set for the instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the target cluster.
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
