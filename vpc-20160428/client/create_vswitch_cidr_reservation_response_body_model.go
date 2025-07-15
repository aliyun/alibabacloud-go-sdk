// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchCidrReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVSwitchCidrReservationResponseBody
	GetRequestId() *string
	SetVSwitchCidrReservationId(v string) *CreateVSwitchCidrReservationResponseBody
	GetVSwitchCidrReservationId() *string
}

type CreateVSwitchCidrReservationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the reserved CIDR block.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
}

func (s CreateVSwitchCidrReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchCidrReservationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVSwitchCidrReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVSwitchCidrReservationResponseBody) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *CreateVSwitchCidrReservationResponseBody) SetRequestId(v string) *CreateVSwitchCidrReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVSwitchCidrReservationResponseBody) SetVSwitchCidrReservationId(v string) *CreateVSwitchCidrReservationResponseBody {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *CreateVSwitchCidrReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
