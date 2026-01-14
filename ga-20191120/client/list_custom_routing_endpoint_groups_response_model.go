// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingEndpointGroupsResponseBody) *ListCustomRoutingEndpointGroupsResponse
	GetBody() *ListCustomRoutingEndpointGroupsResponseBody
}

type ListCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingEndpointGroupsResponse) GetBody() *ListCustomRoutingEndpointGroupsResponseBody {
	return s.Body
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetBody(v *ListCustomRoutingEndpointGroupsResponseBody) *ListCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
