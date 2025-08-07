// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBEndpointAddressResponseBody
	GetRequestId() *string
}

type ModifyDBEndpointAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBEndpointAddressResponseBody) SetRequestId(v string) *ModifyDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
