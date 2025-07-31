// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *CreateNodeResponseBody
	GetNodeId() *string
	SetOrderId(v string) *CreateNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateNodeResponseBody
	GetRequestId() *string
}

type CreateNodeResponseBody struct {
	// The node ID.
	//
	// example:
	//
	// d-bp1b234bf7a4****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20951063702****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7D48FB19-20CA-4725-A870-3D8F5CE6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeResponseBody) SetNodeId(v string) *CreateNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateNodeResponseBody) SetOrderId(v string) *CreateNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
