// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAccessGatewayUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSmartAccessGatewayUserResponseBody
	GetRequestId() *string
}

type DisableSmartAccessGatewayUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 05E86199-6CF6-4F4E-A9CE-9BFC5B020B72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSmartAccessGatewayUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAccessGatewayUserResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSmartAccessGatewayUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSmartAccessGatewayUserResponseBody) SetRequestId(v string) *DisableSmartAccessGatewayUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSmartAccessGatewayUserResponseBody) Validate() error {
	return dara.Validate(s)
}
