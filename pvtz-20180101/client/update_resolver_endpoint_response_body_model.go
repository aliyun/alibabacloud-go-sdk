// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResolverEndpointResponseBody
	GetRequestId() *string
}

type UpdateResolverEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC0BDA3A-A92A-4AC8-B351-322A9C79D5C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResolverEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResolverEndpointResponseBody) SetRequestId(v string) *UpdateResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResolverEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
