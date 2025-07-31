// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *CreateNodeBatchResponseBody
	GetNodeId() *string
	SetOrderId(v string) *CreateNodeBatchResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateNodeBatchResponseBody
	GetRequestId() *string
}

type CreateNodeBatchResponseBody struct {
	// The ID of the added mongos or shard node.
	//
	// example:
	//
	// d-bp18f7d6b6a7****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 50179021707****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 55D41A94-1ACE-55E8-8BC7-67D622E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeBatchResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNodeBatchResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateNodeBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeBatchResponseBody) SetNodeId(v string) *CreateNodeBatchResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateNodeBatchResponseBody) SetOrderId(v string) *CreateNodeBatchResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateNodeBatchResponseBody) SetRequestId(v string) *CreateNodeBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
