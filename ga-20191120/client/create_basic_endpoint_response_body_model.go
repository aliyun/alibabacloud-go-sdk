// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *CreateBasicEndpointResponseBody
	GetEndpointGroupId() *string
	SetEndpointId(v string) *CreateBasicEndpointResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *CreateBasicEndpointResponseBody
	GetRequestId() *string
}

type CreateBasicEndpointResponseBody struct {
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
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

func (s CreateBasicEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateBasicEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateBasicEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicEndpointResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointResponseBody) SetEndpointId(v string) *CreateBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicEndpointResponseBody) SetRequestId(v string) *CreateBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
