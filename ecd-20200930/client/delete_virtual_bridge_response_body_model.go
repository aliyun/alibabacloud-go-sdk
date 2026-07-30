// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBridgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVirtualBridgeResponseBody
	GetRequestId() *string
}

type DeleteVirtualBridgeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualBridgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBridgeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBridgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirtualBridgeResponseBody) SetRequestId(v string) *DeleteVirtualBridgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualBridgeResponseBody) Validate() error {
	return dara.Validate(s)
}
