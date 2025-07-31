// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
}

type CreateDBInstanceResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// dds-bp144a7f2db8****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 21077576248****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D8F1D721-6439-4257-A89C-F1E8E9C9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
