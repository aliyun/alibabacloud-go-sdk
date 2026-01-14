// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *CreateEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetRequestId(v string) *CreateEndpointGroupResponseBody
	GetRequestId() *string
}

type CreateEndpointGroupResponseBody struct {
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEndpointGroupResponseBody) SetEndpointGroupId(v string) *CreateEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateEndpointGroupResponseBody) SetRequestId(v string) *CreateEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
