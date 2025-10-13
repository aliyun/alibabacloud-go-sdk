// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGreyTagRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGreyTagRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGreyTagRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGreyTagRouteResponseBody) *UpdateGreyTagRouteResponse
	GetBody() *UpdateGreyTagRouteResponseBody
}

type UpdateGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGreyTagRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGreyTagRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGreyTagRouteResponse) GetBody() *UpdateGreyTagRouteResponseBody {
	return s.Body
}

func (s *UpdateGreyTagRouteResponse) SetHeaders(v map[string]*string) *UpdateGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGreyTagRouteResponse) SetStatusCode(v int32) *UpdateGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGreyTagRouteResponse) SetBody(v *UpdateGreyTagRouteResponseBody) *UpdateGreyTagRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateGreyTagRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
