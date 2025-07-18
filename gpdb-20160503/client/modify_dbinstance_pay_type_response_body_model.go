// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstancePayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstancePayTypeResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *ModifyDBInstancePayTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBInstancePayTypeResponseBody
	GetRequestId() *string
}

type ModifyDBInstancePayTypeResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 50578361067****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstancePayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstancePayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstancePayTypeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstancePayTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBInstancePayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstancePayTypeResponseBody) SetDBInstanceId(v string) *ModifyDBInstancePayTypeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstancePayTypeResponseBody) SetOrderId(v string) *ModifyDBInstancePayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstancePayTypeResponseBody) SetRequestId(v string) *ModifyDBInstancePayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstancePayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
