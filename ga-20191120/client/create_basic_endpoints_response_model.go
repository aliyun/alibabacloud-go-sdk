// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicEndpointsResponseBody) *CreateBasicEndpointsResponse
	GetBody() *CreateBasicEndpointsResponseBody
}

type CreateBasicEndpointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointsResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicEndpointsResponse) GetBody() *CreateBasicEndpointsResponseBody {
	return s.Body
}

func (s *CreateBasicEndpointsResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointsResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointsResponse) SetStatusCode(v int32) *CreateBasicEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointsResponse) SetBody(v *CreateBasicEndpointsResponseBody) *CreateBasicEndpointsResponse {
	s.Body = v
	return s
}

func (s *CreateBasicEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
