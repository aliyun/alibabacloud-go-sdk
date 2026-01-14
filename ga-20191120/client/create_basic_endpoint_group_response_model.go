// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicEndpointGroupResponseBody) *CreateBasicEndpointGroupResponse
	GetBody() *CreateBasicEndpointGroupResponseBody
}

type CreateBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicEndpointGroupResponse) GetBody() *CreateBasicEndpointGroupResponseBody {
	return s.Body
}

func (s *CreateBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointGroupResponse) SetStatusCode(v int32) *CreateBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointGroupResponse) SetBody(v *CreateBasicEndpointGroupResponseBody) *CreateBasicEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *CreateBasicEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
