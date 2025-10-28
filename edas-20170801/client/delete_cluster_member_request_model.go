// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterMemberRequest
	GetClusterId() *string
	SetClusterMemberId(v string) *DeleteClusterMemberRequest
	GetClusterMemberId() *string
}

type DeleteClusterMemberRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 52984524-****-****-85f2-a34b0e5bb521
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The member ID of the ECS instance that you want to remove from the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zej4i2jdf3****jigng
	ClusterMemberId *string `json:"ClusterMemberId,omitempty" xml:"ClusterMemberId,omitempty"`
}

func (s DeleteClusterMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterMemberRequest) GetClusterMemberId() *string {
	return s.ClusterMemberId
}

func (s *DeleteClusterMemberRequest) SetClusterId(v string) *DeleteClusterMemberRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterMemberRequest) SetClusterMemberId(v string) *DeleteClusterMemberRequest {
	s.ClusterMemberId = &v
	return s
}

func (s *DeleteClusterMemberRequest) Validate() error {
	return dara.Validate(s)
}
