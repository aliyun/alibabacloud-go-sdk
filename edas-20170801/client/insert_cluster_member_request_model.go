// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InsertClusterMemberRequest
	GetClusterId() *string
	SetInstanceIds(v string) *InsertClusterMemberRequest
	GetInstanceIds() *string
	SetPassword(v string) *InsertClusterMemberRequest
	GetPassword() *string
}

type InsertClusterMemberRequest struct {
	// The ID of the cluster into which you want to import ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3e3f77b-462e-****-****-bec8727a4dc8
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The ID of the ECS instance that you want to import into the cluster. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze7s2v0b789k60p****
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty"`
	// The logon password of the ECS instance that you want to import into the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// YourPassword
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s InsertClusterMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InsertClusterMemberRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *InsertClusterMemberRequest) GetPassword() *string {
	return s.Password
}

func (s *InsertClusterMemberRequest) SetClusterId(v string) *InsertClusterMemberRequest {
	s.ClusterId = &v
	return s
}

func (s *InsertClusterMemberRequest) SetInstanceIds(v string) *InsertClusterMemberRequest {
	s.InstanceIds = &v
	return s
}

func (s *InsertClusterMemberRequest) SetPassword(v string) *InsertClusterMemberRequest {
	s.Password = &v
	return s
}

func (s *InsertClusterMemberRequest) Validate() error {
	return dara.Validate(s)
}
