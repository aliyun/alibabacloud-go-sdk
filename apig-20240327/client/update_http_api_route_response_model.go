// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpApiRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpApiRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpApiRouteResponseBody) *UpdateHttpApiRouteResponse
	GetBody() *UpdateHttpApiRouteResponseBody
}

type UpdateHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpApiRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpApiRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpApiRouteResponse) GetBody() *UpdateHttpApiRouteResponseBody {
	return s.Body
}

func (s *UpdateHttpApiRouteResponse) SetHeaders(v map[string]*string) *UpdateHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpApiRouteResponse) SetStatusCode(v int32) *UpdateHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpApiRouteResponse) SetBody(v *UpdateHttpApiRouteResponseBody) *UpdateHttpApiRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpApiRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
