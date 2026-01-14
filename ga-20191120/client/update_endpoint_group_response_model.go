// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEndpointGroupResponseBody) *UpdateEndpointGroupResponse
	GetBody() *UpdateEndpointGroupResponseBody
}

type UpdateEndpointGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEndpointGroupResponse) GetBody() *UpdateEndpointGroupResponseBody {
	return s.Body
}

func (s *UpdateEndpointGroupResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupResponse) SetStatusCode(v int32) *UpdateEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupResponse) SetBody(v *UpdateEndpointGroupResponseBody) *UpdateEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
