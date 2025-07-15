// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *CreateDesktopsResponseBody
	GetDesktopId() []*string
	SetOrderId(v string) *CreateDesktopsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDesktopsResponseBody
	GetRequestId() *string
}

type CreateDesktopsResponseBody struct {
	// The IDs of the cloud computers that are created. If multiple cloud computers are created, multiple IDs are returned.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// > This parameter is returned only when you set the ChargeType parameter to PrePaid.
	//
	// example:
	//
	// 123456789
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponseBody) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *CreateDesktopsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDesktopsResponseBody) SetDesktopId(v []*string) *CreateDesktopsResponseBody {
	s.DesktopId = v
	return s
}

func (s *CreateDesktopsResponseBody) SetOrderId(v string) *CreateDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsResponseBody) SetRequestId(v string) *CreateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
