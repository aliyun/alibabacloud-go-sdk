// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBEndpointAddressResponseBody
	GetRequestId() *string
}

type DeleteDBEndpointAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBEndpointAddressResponseBody) SetRequestId(v string) *DeleteDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
