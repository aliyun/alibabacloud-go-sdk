// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResolverEndpointResponseBody
	GetRequestId() *string
}

type DeleteResolverEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 35134B4A-CEC0-43C8-AAD4-BA54AE3268B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResolverEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResolverEndpointResponseBody) SetRequestId(v string) *DeleteResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResolverEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
