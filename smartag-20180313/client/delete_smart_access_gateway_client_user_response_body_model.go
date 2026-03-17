// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSmartAccessGatewayClientUserResponseBody
	GetRequestId() *string
}

type DeleteSmartAccessGatewayClientUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 09AD82DC-FE26-4B66-B526-2FA6BE82A4D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmartAccessGatewayClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmartAccessGatewayClientUserResponseBody) SetRequestId(v string) *DeleteSmartAccessGatewayClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
