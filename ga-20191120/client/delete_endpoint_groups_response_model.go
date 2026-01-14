// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEndpointGroupsResponseBody) *DeleteEndpointGroupsResponse
	GetBody() *DeleteEndpointGroupsResponseBody
}

type DeleteEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEndpointGroupsResponse) GetBody() *DeleteEndpointGroupsResponseBody {
	return s.Body
}

func (s *DeleteEndpointGroupsResponse) SetHeaders(v map[string]*string) *DeleteEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointGroupsResponse) SetStatusCode(v int32) *DeleteEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointGroupsResponse) SetBody(v *DeleteEndpointGroupsResponseBody) *DeleteEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *DeleteEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
