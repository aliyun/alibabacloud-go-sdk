// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateStoragePlanResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *CreateStoragePlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateStoragePlanResponseBody
	GetRequestId() *string
}

type CreateStoragePlanResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// POLARDB-cn-**************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2035638*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStoragePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoragePlanResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateStoragePlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateStoragePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStoragePlanResponseBody) SetDBInstanceId(v string) *CreateStoragePlanResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateStoragePlanResponseBody) SetOrderId(v string) *CreateStoragePlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateStoragePlanResponseBody) SetRequestId(v string) *CreateStoragePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoragePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
