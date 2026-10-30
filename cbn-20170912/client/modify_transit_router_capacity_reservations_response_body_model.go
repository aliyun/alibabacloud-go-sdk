// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCapacityReservationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTransitRouterCapacityReservationsResponseBody
	GetRequestId() *string
}

type ModifyTransitRouterCapacityReservationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 004E99FB-E996-5777-888E-BA1D8F215407
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTransitRouterCapacityReservationsResponseBody) SetRequestId(v string) *ModifyTransitRouterCapacityReservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsResponseBody) Validate() error {
	return dara.Validate(s)
}
