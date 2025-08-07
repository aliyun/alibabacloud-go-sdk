// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodesClassResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBNodesClassResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBNodesClassResponseBody
	GetRequestId() *string
}

type ModifyDBNodesClassResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 685F028C-4FCD-407D-A559-072D63******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodesClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesClassResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodesClassResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBNodesClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodesClassResponseBody) SetDBClusterId(v string) *ModifyDBNodesClassResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodesClassResponseBody) SetOrderId(v string) *ModifyDBNodesClassResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBNodesClassResponseBody) SetRequestId(v string) *ModifyDBNodesClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodesClassResponseBody) Validate() error {
	return dara.Validate(s)
}
