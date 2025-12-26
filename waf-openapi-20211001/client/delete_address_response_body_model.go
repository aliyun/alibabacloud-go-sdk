// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAddressResponseBody
	GetRequestId() *string
}

type DeleteAddressResponseBody struct {
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A8****6A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAddressResponseBody) SetRequestId(v string) *DeleteAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
