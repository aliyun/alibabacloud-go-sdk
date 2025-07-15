// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteTableResponseBody
	GetRequestId() *string
}

type DeleteRouteTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC668356-BCB4-42FD-9BC3-FA2B2E04B634
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteTableResponseBody) SetRequestId(v string) *DeleteRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
