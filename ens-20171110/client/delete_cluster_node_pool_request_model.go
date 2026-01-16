// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterNodePoolRequest
	GetClusterId() *string
	SetNodepoolId(v string) *DeleteClusterNodePoolRequest
	GetNodepoolId() *string
}

type DeleteClusterNodePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s DeleteClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterNodePoolRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DeleteClusterNodePoolRequest) SetClusterId(v string) *DeleteClusterNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterNodePoolRequest) SetNodepoolId(v string) *DeleteClusterNodePoolRequest {
	s.NodepoolId = &v
	return s
}

func (s *DeleteClusterNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
