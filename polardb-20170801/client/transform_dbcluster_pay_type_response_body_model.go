// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBClusterPayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *TransformDBClusterPayTypeResponseBody
	GetChargeType() *string
	SetDBClusterId(v string) *TransformDBClusterPayTypeResponseBody
	GetDBClusterId() *string
	SetExpiredTime(v string) *TransformDBClusterPayTypeResponseBody
	GetExpiredTime() *string
	SetOrderId(v string) *TransformDBClusterPayTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *TransformDBClusterPayTypeResponseBody
	GetRequestId() *string
}

type TransformDBClusterPayTypeResponseBody struct {
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the cluster expires.
	//
	// >  This parameter is returned if you set the **PayType*	- parameter to **Prepaid**.
	//
	// example:
	//
	// 2020-04-20T10:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20515760028****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5E71541A-6007-4DCC-A38A-F872C31FEB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformDBClusterPayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformDBClusterPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransformDBClusterPayTypeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TransformDBClusterPayTypeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *TransformDBClusterPayTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *TransformDBClusterPayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformDBClusterPayTypeResponseBody) SetChargeType(v string) *TransformDBClusterPayTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetDBClusterId(v string) *TransformDBClusterPayTypeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetExpiredTime(v string) *TransformDBClusterPayTypeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetOrderId(v string) *TransformDBClusterPayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetRequestId(v string) *TransformDBClusterPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
