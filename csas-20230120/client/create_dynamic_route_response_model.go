// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateDynamicRouteResponseBody) *CreateDynamicRouteResponse
	GetBody() *CreateDynamicRouteResponseBody
}

type CreateDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDynamicRouteResponse) GetBody() *CreateDynamicRouteResponseBody {
	return s.Body
}

func (s *CreateDynamicRouteResponse) SetHeaders(v map[string]*string) *CreateDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateDynamicRouteResponse) SetStatusCode(v int32) *CreateDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDynamicRouteResponse) SetBody(v *CreateDynamicRouteResponseBody) *CreateDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *CreateDynamicRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
