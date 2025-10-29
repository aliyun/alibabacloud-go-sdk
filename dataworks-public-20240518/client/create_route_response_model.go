// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteResponseBody) *CreateRouteResponse
	GetBody() *CreateRouteResponseBody
}

type CreateRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteResponse) GetBody() *CreateRouteResponseBody {
	return s.Body
}

func (s *CreateRouteResponse) SetHeaders(v map[string]*string) *CreateRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteResponse) SetStatusCode(v int32) *CreateRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteResponse) SetBody(v *CreateRouteResponseBody) *CreateRouteResponse {
	s.Body = v
	return s
}

func (s *CreateRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
