// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBInstancePayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *TransformDBInstancePayTypeResponseBody
	GetChargeType() *string
	SetDBInstanceId(v string) *TransformDBInstancePayTypeResponseBody
	GetDBInstanceId() *string
	SetExpiredTime(v string) *TransformDBInstancePayTypeResponseBody
	GetExpiredTime() *string
	SetOrderId(v int64) *TransformDBInstancePayTypeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *TransformDBInstancePayTypeResponseBody
	GetRequestId() *string
}

type TransformDBInstancePayTypeResponseBody struct {
	// The billing method. Valid values:
	//
	// - POSTPAY: pay-as-you-go
	//
	// - PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The expiration time.
	//
	// > This parameter is not returned if the billing method is changed to pay-as-you-go.
	//
	// example:
	//
	// 2020-04-20T10:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20515760028****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5E6E09DE-5B12-4BFF-A55E-1C86EDE06D9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformDBInstancePayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformDBInstancePayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransformDBInstancePayTypeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransformDBInstancePayTypeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TransformDBInstancePayTypeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *TransformDBInstancePayTypeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TransformDBInstancePayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformDBInstancePayTypeResponseBody) SetChargeType(v string) *TransformDBInstancePayTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransformDBInstancePayTypeResponseBody) SetDBInstanceId(v string) *TransformDBInstancePayTypeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *TransformDBInstancePayTypeResponseBody) SetExpiredTime(v string) *TransformDBInstancePayTypeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *TransformDBInstancePayTypeResponseBody) SetOrderId(v int64) *TransformDBInstancePayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *TransformDBInstancePayTypeResponseBody) SetRequestId(v string) *TransformDBInstancePayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformDBInstancePayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
