// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBNodesResponseBody
	GetDBInstanceId() *string
	SetNodeIds(v string) *CreateDBNodesResponseBody
	GetNodeIds() *string
	SetOrderId(v int64) *CreateDBNodesResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateDBNodesResponseBody
	GetRequestId() *string
}

type CreateDBNodesResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-2ze450g4ctg6t****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the created node. The value is a string. Multiple values are separated by commas (`,`).
	//
	// example:
	//
	// rn-abcd2*****
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2133400000*****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7A41C147-C8D0-4DAE-A1A2-17EBCD60DFA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBNodesResponseBody) GetNodeIds() *string {
	return s.NodeIds
}

func (s *CreateDBNodesResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBNodesResponseBody) SetDBInstanceId(v string) *CreateDBNodesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetNodeIds(v string) *CreateDBNodesResponseBody {
	s.NodeIds = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetOrderId(v int64) *CreateDBNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetRequestId(v string) *CreateDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
