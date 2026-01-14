// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *UpdateBasicEndpointResponseBody
	GetEndpointGroupId() *string
	SetEndpointId(v string) *UpdateBasicEndpointResponseBody
	GetEndpointId() *string
	SetName(v string) *UpdateBasicEndpointResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateBasicEndpointResponseBody
	GetRequestId() *string
}

type UpdateBasicEndpointResponseBody struct {
	// The ID of the endpoint group to which the endpoints belong.
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
	// The name of the endpoint.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateBasicEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateBasicEndpointResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateBasicEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicEndpointResponseBody) SetEndpointGroupId(v string) *UpdateBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetEndpointId(v string) *UpdateBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetName(v string) *UpdateBasicEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetRequestId(v string) *UpdateBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
