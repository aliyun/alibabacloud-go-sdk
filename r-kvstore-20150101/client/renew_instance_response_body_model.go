// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *RenewInstanceResponseBody
	GetEndTime() *string
	SetOrderId(v string) *RenewInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
}

type RenewInstanceResponseBody struct {
	// The end time of the order.
	//
	// example:
	//
	// 2019-02-19T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2222245-222A-4155-9349-E22222****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *RenewInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) SetEndTime(v string) *RenewInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *RenewInstanceResponseBody) SetOrderId(v string) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
