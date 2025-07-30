// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShardingNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v []*string) *AddShardingNodeResponseBody
	GetNodeIds() []*string
	SetOrderId(v int64) *AddShardingNodeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *AddShardingNodeResponseBody
	GetRequestId() *string
}

type AddShardingNodeResponseBody struct {
	// The IDs of the data shards.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 20741011111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B79C1A90-495B-4E99-A2AA-A4DB13B8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddShardingNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddShardingNodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddShardingNodeResponseBody) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *AddShardingNodeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *AddShardingNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddShardingNodeResponseBody) SetNodeIds(v []*string) *AddShardingNodeResponseBody {
	s.NodeIds = v
	return s
}

func (s *AddShardingNodeResponseBody) SetOrderId(v int64) *AddShardingNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *AddShardingNodeResponseBody) SetRequestId(v string) *AddShardingNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShardingNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
