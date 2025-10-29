// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOnZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TurnOnZoneResponseBody
	GetRequestId() *string
}

type TurnOnZoneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TurnOnZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TurnOnZoneResponseBody) GoString() string {
	return s.String()
}

func (s *TurnOnZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TurnOnZoneResponseBody) SetRequestId(v string) *TurnOnZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *TurnOnZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
