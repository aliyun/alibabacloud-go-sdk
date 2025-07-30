// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateTCInstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *CreateTCInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateTCInstanceResponseBody
	GetRequestId() *string
}

type CreateTCInstanceResponseBody struct {
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22179******0904
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTCInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTCInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTCInstanceResponseBody) SetInstanceId(v string) *CreateTCInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTCInstanceResponseBody) SetOrderId(v int64) *CreateTCInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTCInstanceResponseBody) SetRequestId(v string) *CreateTCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
