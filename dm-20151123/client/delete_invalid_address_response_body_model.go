// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvalidAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInvalidAddressResponseBody
	GetRequestId() *string
}

type DeleteInvalidAddressResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 2D086F6-xxxx-xxxx-xxxx-006DED011A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInvalidAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvalidAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInvalidAddressResponseBody) SetRequestId(v string) *DeleteInvalidAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInvalidAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
