// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodesRequest
	GetClusterId() *string
	SetInstanceIds(v []*string) *DeleteNodesRequest
	GetInstanceIds() []*string
}

type DeleteNodesRequest struct {
	// The cluster ID. You can call the [listclusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs of the compute nodes that you want to delete.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DeleteNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeleteNodesRequest) SetClusterId(v string) *DeleteNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesRequest) SetInstanceIds(v []*string) *DeleteNodesRequest {
	s.InstanceIds = v
	return s
}

func (s *DeleteNodesRequest) Validate() error {
	return dara.Validate(s)
}
