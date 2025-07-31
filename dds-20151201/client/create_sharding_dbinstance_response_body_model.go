// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardingDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateShardingDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateShardingDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateShardingDBInstanceResponseBody
	GetRequestId() *string
}

type CreateShardingDBInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// dds-bp114f14849d****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21010996721****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D8F1D721-6439-4257-A89C-F1E8E9C9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateShardingDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateShardingDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateShardingDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateShardingDBInstanceResponseBody) SetDBInstanceId(v string) *CreateShardingDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateShardingDBInstanceResponseBody) SetOrderId(v string) *CreateShardingDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateShardingDBInstanceResponseBody) SetRequestId(v string) *CreateShardingDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateShardingDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
