// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStorageGatewayResponseBody
	GetRequestId() *string
}

type DeleteStorageGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 853D6E71-E087-1557-B65C-32BFBEE5CD97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStorageGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStorageGatewayResponseBody) SetRequestId(v string) *DeleteStorageGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
