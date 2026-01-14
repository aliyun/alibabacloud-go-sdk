// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointIds(v []*string) *UpdateCustomRoutingEndpointsResponseBody
	GetEndpointIds() []*string
	SetRequestId(v string) *UpdateCustomRoutingEndpointsResponseBody
	GetRequestId() *string
}

type UpdateCustomRoutingEndpointsResponseBody struct {
	// The IDs of the endpoints.
	EndpointIds []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsResponseBody) GetEndpointIds() []*string {
	return s.EndpointIds
}

func (s *UpdateCustomRoutingEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomRoutingEndpointsResponseBody) SetEndpointIds(v []*string) *UpdateCustomRoutingEndpointsResponseBody {
	s.EndpointIds = v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}
