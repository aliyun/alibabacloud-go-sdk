// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteResponseBody) *UpdateRouteResponse
	GetBody() *UpdateRouteResponseBody
}

type UpdateRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteResponse) GetBody() *UpdateRouteResponseBody {
	return s.Body
}

func (s *UpdateRouteResponse) SetHeaders(v map[string]*string) *UpdateRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteResponse) SetStatusCode(v int32) *UpdateRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteResponse) SetBody(v *UpdateRouteResponseBody) *UpdateRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
