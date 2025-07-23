// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *JoinClusterRequest
	GetClusterId() *string
	SetInstanceId(v string) *JoinClusterRequest
	GetInstanceId() *string
}

type JoinClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-NZB9Oj5Yfd8Y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the HSM that you want to add to the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s JoinClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinClusterRequest) GoString() string {
	return s.String()
}

func (s *JoinClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *JoinClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *JoinClusterRequest) SetClusterId(v string) *JoinClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *JoinClusterRequest) SetInstanceId(v string) *JoinClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinClusterRequest) Validate() error {
	return dara.Validate(s)
}
