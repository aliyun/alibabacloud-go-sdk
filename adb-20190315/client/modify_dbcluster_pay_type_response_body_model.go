// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterPayTypeResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBClusterPayTypeResponseBody
	GetOrderId() *string
	SetPayType(v string) *ModifyDBClusterPayTypeResponseBody
	GetPayType() *string
	SetRequestId(v string) *ModifyDBClusterPayTypeResponseBody
	GetRequestId() *string
}

type ModifyDBClusterPayTypeResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2023532****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterPayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterPayTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBClusterPayTypeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBClusterPayTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterPayTypeResponseBody) SetDBClusterId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetOrderId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetPayType(v string) *ModifyDBClusterPayTypeResponseBody {
	s.PayType = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetRequestId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
