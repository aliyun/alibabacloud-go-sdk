// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchCidrReservationAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVSwitchCidrReservationAttributeResponseBody
	GetRequestId() *string
}

type ModifyVSwitchCidrReservationAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVSwitchCidrReservationAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchCidrReservationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchCidrReservationAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVSwitchCidrReservationAttributeResponseBody) SetRequestId(v string) *ModifyVSwitchCidrReservationAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
