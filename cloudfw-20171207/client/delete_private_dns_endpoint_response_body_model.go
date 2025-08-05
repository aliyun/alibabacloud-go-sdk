// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateDnsEndpointResponseBody
	GetRequestId() *string
}

type DeletePrivateDnsEndpointResponseBody struct {
	// example:
	//
	// 822B9125-6E1A-551C-8EAF-6E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateDnsEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateDnsEndpointResponseBody) SetRequestId(v string) *DeletePrivateDnsEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateDnsEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
