// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomRoutingEndpointGroupsResponseBody) *DeleteCustomRoutingEndpointGroupsResponse
	GetBody() *DeleteCustomRoutingEndpointGroupsResponseBody
}

type DeleteCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) GetBody() *DeleteCustomRoutingEndpointGroupsResponseBody {
	return s.Body
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetBody(v *DeleteCustomRoutingEndpointGroupsResponseBody) *DeleteCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
