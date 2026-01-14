// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *CreateBasicEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetRequestId(v string) *CreateBasicEndpointGroupResponseBody
	GetRequestId() *string
}

type CreateBasicEndpointGroupResponseBody struct {
	// The endpoint group ID.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateBasicEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicEndpointGroupResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointGroupResponseBody) SetRequestId(v string) *CreateBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
