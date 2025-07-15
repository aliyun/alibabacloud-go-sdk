// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v []*string) *RenewDesktopGroupResponseBody
	GetOrderId() []*string
	SetRequestId(v string) *RenewDesktopGroupResponseBody
	GetRequestId() *string
}

type RenewDesktopGroupResponseBody struct {
	// The order IDs.
	OrderId []*string `json:"OrderId,omitempty" xml:"OrderId,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E55E6732-2028-52FA-AB06-EA29C36B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopGroupResponseBody) GetOrderId() []*string {
	return s.OrderId
}

func (s *RenewDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewDesktopGroupResponseBody) SetOrderId(v []*string) *RenewDesktopGroupResponseBody {
	s.OrderId = v
	return s
}

func (s *RenewDesktopGroupResponseBody) SetRequestId(v string) *RenewDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
