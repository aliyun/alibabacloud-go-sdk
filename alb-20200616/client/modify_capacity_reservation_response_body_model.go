// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCapacityReservationResponseBody
	GetRequestId() *string
}

type ModifyCapacityReservationResponseBody struct {
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCapacityReservationResponseBody) SetRequestId(v string) *ModifyCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCapacityReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
