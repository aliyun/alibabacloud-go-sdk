// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOffZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TurnOffZoneResponseBody
	GetRequestId() *string
}

type TurnOffZoneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TurnOffZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TurnOffZoneResponseBody) GoString() string {
	return s.String()
}

func (s *TurnOffZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TurnOffZoneResponseBody) SetRequestId(v string) *TurnOffZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *TurnOffZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
