// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEndpointGroupsResponseBody) *UpdateEndpointGroupsResponse
	GetBody() *UpdateEndpointGroupsResponseBody
}

type UpdateEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEndpointGroupsResponse) GetBody() *UpdateEndpointGroupsResponseBody {
	return s.Body
}

func (s *UpdateEndpointGroupsResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupsResponse) SetStatusCode(v int32) *UpdateEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupsResponse) SetBody(v *UpdateEndpointGroupsResponseBody) *UpdateEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *UpdateEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
