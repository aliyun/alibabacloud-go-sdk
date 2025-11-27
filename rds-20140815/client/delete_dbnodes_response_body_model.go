// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBNodesResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *DeleteDBNodesResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *DeleteDBNodesResponseBody
	GetRequestId() *string
}

type DeleteDBNodesResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 100780000000000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBNodesResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DeleteDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBNodesResponseBody) SetDBInstanceId(v string) *DeleteDBNodesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) SetOrderId(v int64) *DeleteDBNodesResponseBody {
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
