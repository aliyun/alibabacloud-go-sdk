// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEndpointGroupResponseBody) *DeleteEndpointGroupResponse
	GetBody() *DeleteEndpointGroupResponseBody
}

type DeleteEndpointGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEndpointGroupResponse) GetBody() *DeleteEndpointGroupResponseBody {
	return s.Body
}

func (s *DeleteEndpointGroupResponse) SetHeaders(v map[string]*string) *DeleteEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointGroupResponse) SetStatusCode(v int32) *DeleteEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointGroupResponse) SetBody(v *DeleteEndpointGroupResponseBody) *DeleteEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
