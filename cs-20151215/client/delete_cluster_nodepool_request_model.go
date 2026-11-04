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
	// Specifies whether to force delete the node pool. Valid values:
	//
	// - true: Force deletes the node pool. If the node pool contains existing nodes, the nodes are forcibly deleted.
	//
	// - false: Does not force delete the node pool. If the node pool contains existing nodes, the node pool cannot be deleted and an error is returned.
	//
	// Deleting a node pool does not automatically release subscription nodes. Automatic release subscription nodes, [change the billing method of the nodes from subscription to pay-as-you-go](https://www.alibabacloud.com/help/en/ecs/change-the-billing-method-of-an-instance-from-subscription-to-pay-as-you-go-1) and then log on to the ECS console automatic release the nodes.
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
