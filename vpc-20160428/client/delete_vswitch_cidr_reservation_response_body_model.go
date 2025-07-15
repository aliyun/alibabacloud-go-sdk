// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchCidrReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVSwitchCidrReservationResponseBody
	GetRequestId() *string
}

type DeleteVSwitchCidrReservationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVSwitchCidrReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchCidrReservationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchCidrReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVSwitchCidrReservationResponseBody) SetRequestId(v string) *DeleteVSwitchCidrReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVSwitchCidrReservationResponseBody) Validate() error {
	return dara.Validate(s)
}
