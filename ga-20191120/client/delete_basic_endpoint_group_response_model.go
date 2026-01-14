// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicEndpointGroupResponseBody) *DeleteBasicEndpointGroupResponse
	GetBody() *DeleteBasicEndpointGroupResponseBody
}

type DeleteBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicEndpointGroupResponse) GetBody() *DeleteBasicEndpointGroupResponseBody {
	return s.Body
}

func (s *DeleteBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *DeleteBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicEndpointGroupResponse) SetStatusCode(v int32) *DeleteBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicEndpointGroupResponse) SetBody(v *DeleteBasicEndpointGroupResponseBody) *DeleteBasicEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
