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
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NodeIds      *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
