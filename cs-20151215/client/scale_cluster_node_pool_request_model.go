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
	// The number of worker nodes that you want to add. For example, the current node pool contains two nodes. After the two node is scaled out, the node pool contains four nodes. Due to the limit of the node quota, you can add at most 500 nodes in each request.
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
