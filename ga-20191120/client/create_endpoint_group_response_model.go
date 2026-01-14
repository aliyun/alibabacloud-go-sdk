// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateEndpointGroupResponseBody) *CreateEndpointGroupResponse
	GetBody() *CreateEndpointGroupResponseBody
}

type CreateEndpointGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEndpointGroupResponse) GetBody() *CreateEndpointGroupResponseBody {
	return s.Body
}

func (s *CreateEndpointGroupResponse) SetHeaders(v map[string]*string) *CreateEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointGroupResponse) SetStatusCode(v int32) *CreateEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointGroupResponse) SetBody(v *CreateEndpointGroupResponseBody) *CreateEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *CreateEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
