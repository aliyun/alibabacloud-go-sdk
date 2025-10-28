// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterRequest
	GetClusterId() *string
	SetMode(v int32) *DeleteClusterRequest
	GetMode() *int32
}

type DeleteClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8b96ade0-0a07-****-af9d-5ed83640d076
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster ID. Valid values:
	//
	// 	- 0: specifies the ID of the cluster in Enterprise Distributed Application Service (EDAS).
	//
	// 	- 1: specifies the ID of the ACK cluster.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterRequest) GetMode() *int32 {
	return s.Mode
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetMode(v int32) *DeleteClusterRequest {
	s.Mode = &v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	return dara.Validate(s)
}
