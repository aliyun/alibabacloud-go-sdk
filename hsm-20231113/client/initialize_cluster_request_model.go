// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InitializeClusterRequest
	GetClusterId() *string
}

type InitializeClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-p94y1dud9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InitializeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeClusterRequest) GoString() string {
	return s.String()
}

func (s *InitializeClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InitializeClusterRequest) SetClusterId(v string) *InitializeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *InitializeClusterRequest) Validate() error {
	return dara.Validate(s)
}
