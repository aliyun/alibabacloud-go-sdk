// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodesShrinkRequest
	GetClusterId() *string
	SetInstanceIdsShrink(v string) *DeleteNodesShrinkRequest
	GetInstanceIdsShrink() *string
}

type DeleteNodesShrinkRequest struct {
	// The cluster ID. You can call the [listclusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs of the compute nodes that you want to delete.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DeleteNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DeleteNodesShrinkRequest) SetClusterId(v string) *DeleteNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesShrinkRequest) SetInstanceIdsShrink(v string) *DeleteNodesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DeleteNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
