// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointIds(v []*string) *CreateCustomRoutingEndpointsResponseBody
	GetEndpointIds() []*string
	SetRequestId(v string) *CreateCustomRoutingEndpointsResponseBody
	GetRequestId() *string
}

type CreateCustomRoutingEndpointsResponseBody struct {
	// The IDs of the endpoints.
	EndpointIds []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsResponseBody) GetEndpointIds() []*string {
	return s.EndpointIds
}

func (s *CreateCustomRoutingEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomRoutingEndpointsResponseBody) SetEndpointIds(v []*string) *CreateCustomRoutingEndpointsResponseBody {
	s.EndpointIds = v
	return s
}

func (s *CreateCustomRoutingEndpointsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomRoutingEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}
