// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ScaleClusterNodePoolRequest
	GetCount() *int64
}

type ScaleClusterNodePoolRequest struct {
	// The number of nodes to add. For example, if the node pool currently has 2 nodes and you add 2 more, the node pool will have 4 nodes. Due to the node quota limit of the current cluster, you can add up to 500 nodes in a single operation.
	//
	// example:
	//
	// 2
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s ScaleClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolRequest) GetCount() *int64 {
	return s.Count
}

func (s *ScaleClusterNodePoolRequest) SetCount(v int64) *ScaleClusterNodePoolRequest {
	s.Count = &v
	return s
}

func (s *ScaleClusterNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
