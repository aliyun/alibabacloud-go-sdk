// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeClassResponseBody
	GetDBClusterId() *string
	SetOrderId(v string) *ModifyDBNodeClassResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBNodeClassResponseBody
	GetRequestId() *string
}

type ModifyDBNodeClassResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 685F028C-4FCD-407D-A559-072D63******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeClassResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBNodeClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeClassResponseBody) SetDBClusterId(v string) *ModifyDBNodeClassResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassResponseBody) SetOrderId(v string) *ModifyDBNodeClassResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBNodeClassResponseBody) SetRequestId(v string) *ModifyDBNodeClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeClassResponseBody) Validate() error {
	return dara.Validate(s)
}
