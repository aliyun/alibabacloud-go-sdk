// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupIds(v []*string) *CreateEndpointGroupsResponseBody
	GetEndpointGroupIds() []*string
	SetRequestId(v string) *CreateEndpointGroupsResponseBody
	GetRequestId() *string
}

type CreateEndpointGroupsResponseBody struct {
	// The IDs of the endpoint groups.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsResponseBody) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *CreateEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *CreateEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *CreateEndpointGroupsResponseBody) SetRequestId(v string) *CreateEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
