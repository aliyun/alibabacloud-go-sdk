// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomRoutingEndpointGroupsResponseBody) *CreateCustomRoutingEndpointGroupsResponse
	GetBody() *CreateCustomRoutingEndpointGroupsResponseBody
}

type CreateCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomRoutingEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomRoutingEndpointGroupsResponse) GetBody() *CreateCustomRoutingEndpointGroupsResponseBody {
	return s.Body
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetBody(v *CreateCustomRoutingEndpointGroupsResponseBody) *CreateCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
