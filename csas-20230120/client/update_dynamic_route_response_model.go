// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDynamicRouteResponseBody) *UpdateDynamicRouteResponse
	GetBody() *UpdateDynamicRouteResponseBody
}

type UpdateDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDynamicRouteResponse) GetBody() *UpdateDynamicRouteResponseBody {
	return s.Body
}

func (s *UpdateDynamicRouteResponse) SetHeaders(v map[string]*string) *UpdateDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateDynamicRouteResponse) SetStatusCode(v int32) *UpdateDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDynamicRouteResponse) SetBody(v *UpdateDynamicRouteResponseBody) *UpdateDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateDynamicRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
