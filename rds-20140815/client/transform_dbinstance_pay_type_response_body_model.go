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
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
