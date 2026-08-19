// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDDoSSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateDDoSSpecResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *UpdateDDoSSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateDDoSSpecResponseBody
	GetRequestId() *string
}

type UpdateDDoSSpecResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// esa-ddos-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 31223****11
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDDoSSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDDoSSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDDoSSpecResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDDoSSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateDDoSSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDDoSSpecResponseBody) SetInstanceId(v string) *UpdateDDoSSpecResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateDDoSSpecResponseBody) SetOrderId(v string) *UpdateDDoSSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateDDoSSpecResponseBody) SetRequestId(v string) *UpdateDDoSSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDDoSSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
