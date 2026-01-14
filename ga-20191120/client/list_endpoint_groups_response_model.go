// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListEndpointGroupsResponseBody) *ListEndpointGroupsResponse
	GetBody() *ListEndpointGroupsResponseBody
}

type ListEndpointGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEndpointGroupsResponse) GetBody() *ListEndpointGroupsResponseBody {
	return s.Body
}

func (s *ListEndpointGroupsResponse) SetHeaders(v map[string]*string) *ListEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointGroupsResponse) SetStatusCode(v int32) *ListEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointGroupsResponse) SetBody(v *ListEndpointGroupsResponseBody) *ListEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *ListEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
