// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodepoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteClusterNodepoolRequest
	GetForce() *bool
}

type DeleteClusterNodepoolRequest struct {
	// Specifies whether to forcefully delete the node pool.
	//
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteClusterNodepoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodepoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodepoolRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteClusterNodepoolRequest) SetForce(v bool) *DeleteClusterNodepoolRequest {
	s.Force = &v
	return s
}

func (s *DeleteClusterNodepoolRequest) Validate() error {
	return dara.Validate(s)
}
