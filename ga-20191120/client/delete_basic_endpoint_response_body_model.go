// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DeleteBasicEndpointResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *DeleteBasicEndpointResponseBody
	GetRequestId() *string
}

type DeleteBasicEndpointResponseBody struct {
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteBasicEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicEndpointResponseBody) SetEndpointId(v string) *DeleteBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicEndpointResponseBody) SetRequestId(v string) *DeleteBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
