// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupIds(v []*string) *CreateCustomRoutingEndpointGroupsResponseBody
	GetEndpointGroupIds() []*string
	SetRequestId(v string) *CreateCustomRoutingEndpointGroupsResponseBody
	GetRequestId() *string
}

type CreateCustomRoutingEndpointGroupsResponseBody struct {
	// The IDs of the endpoint groups.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *CreateCustomRoutingEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
