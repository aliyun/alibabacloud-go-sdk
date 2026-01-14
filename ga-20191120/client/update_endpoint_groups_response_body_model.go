// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupIds(v []*string) *UpdateEndpointGroupsResponseBody
	GetEndpointGroupIds() []*string
	SetRequestId(v string) *UpdateEndpointGroupsResponseBody
	GetRequestId() *string
}

type UpdateEndpointGroupsResponseBody struct {
	// The IDs of the endpoint groups.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsResponseBody) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *UpdateEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *UpdateEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *UpdateEndpointGroupsResponseBody) SetRequestId(v string) *UpdateEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEndpointGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
