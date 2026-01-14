// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicEndpointGroupResponseBody) *UpdateBasicEndpointGroupResponse
	GetBody() *UpdateBasicEndpointGroupResponseBody
}

type UpdateBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicEndpointGroupResponse) GetBody() *UpdateBasicEndpointGroupResponseBody {
	return s.Body
}

func (s *UpdateBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *UpdateBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicEndpointGroupResponse) SetStatusCode(v int32) *UpdateBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicEndpointGroupResponse) SetBody(v *UpdateBasicEndpointGroupResponseBody) *UpdateBasicEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
