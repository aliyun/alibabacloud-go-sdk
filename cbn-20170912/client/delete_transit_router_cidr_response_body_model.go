// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterCidrResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterCidrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterCidrResponseBody) SetRequestId(v string) *DeleteTransitRouterCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
