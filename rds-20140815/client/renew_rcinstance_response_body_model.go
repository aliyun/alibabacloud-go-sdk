// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RenewRCInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *RenewRCInstanceResponseBody
	GetOrderId() *string
	SetOrderIds(v string) *RenewRCInstanceResponseBody
	GetOrderIds() *string
	SetRequestId(v string) *RenewRCInstanceResponseBody
	GetRequestId() *string
}

type RenewRCInstanceResponseBody struct {
	// The ID of the RDS Custom instance.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 23202700556****
	OrderIds *string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFFC7565-B3CF-5CFA-9E1F-164DD1E1F498
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewRCInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RenewRCInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewRCInstanceResponseBody) GetOrderIds() *string {
	return s.OrderIds
}

func (s *RenewRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewRCInstanceResponseBody) SetDBInstanceId(v string) *RenewRCInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *RenewRCInstanceResponseBody) SetOrderId(v string) *RenewRCInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewRCInstanceResponseBody) SetOrderIds(v string) *RenewRCInstanceResponseBody {
	s.OrderIds = &v
	return s
}

func (s *RenewRCInstanceResponseBody) SetRequestId(v string) *RenewRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
