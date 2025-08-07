// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBNodesResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *DeleteDBNodesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DeleteDBNodesResponseBody
	GetRequestId() *string
}

type DeleteDBNodesResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6566B2E6-3157-4B57-A693-AFB751******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBNodesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBNodesResponseBody) SetDBClusterId(v string) *DeleteDBNodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) SetOrderId(v string) *DeleteDBNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) SetRequestId(v string) *DeleteDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
