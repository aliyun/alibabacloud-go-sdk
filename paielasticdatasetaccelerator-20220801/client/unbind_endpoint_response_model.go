// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UnbindEndpointResponseBody) *UnbindEndpointResponse
	GetBody() *UnbindEndpointResponseBody
}

type UnbindEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindEndpointResponse) GoString() string {
	return s.String()
}

func (s *UnbindEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindEndpointResponse) GetBody() *UnbindEndpointResponseBody {
	return s.Body
}

func (s *UnbindEndpointResponse) SetHeaders(v map[string]*string) *UnbindEndpointResponse {
	s.Headers = v
	return s
}

func (s *UnbindEndpointResponse) SetStatusCode(v int32) *UnbindEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEndpointResponse) SetBody(v *UnbindEndpointResponseBody) *UnbindEndpointResponse {
	s.Body = v
	return s
}

func (s *UnbindEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
