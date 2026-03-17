// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type DeleteSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E26DBAAE-A796-4A48-98B4-B45AFCD1F299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmartAccessGatewayResponseBody) SetRequestId(v string) *DeleteSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
