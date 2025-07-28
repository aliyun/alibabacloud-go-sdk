// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *AddResolverEndpointResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *AddResolverEndpointResponseBody
	GetRequestId() *string
}

type AddResolverEndpointResponseBody struct {
	// The endpoint ID.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32436208-E1AF-4DAB-B3B8-24F5F25B0950
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddResolverEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *AddResolverEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddResolverEndpointResponseBody) SetEndpointId(v string) *AddResolverEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *AddResolverEndpointResponseBody) SetRequestId(v string) *AddResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddResolverEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
