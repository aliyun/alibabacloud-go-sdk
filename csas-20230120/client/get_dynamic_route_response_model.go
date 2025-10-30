// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *GetDynamicRouteResponseBody) *GetDynamicRouteResponse
	GetBody() *GetDynamicRouteResponseBody
}

type GetDynamicRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDynamicRouteResponse) GetBody() *GetDynamicRouteResponseBody {
	return s.Body
}

func (s *GetDynamicRouteResponse) SetHeaders(v map[string]*string) *GetDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *GetDynamicRouteResponse) SetStatusCode(v int32) *GetDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDynamicRouteResponse) SetBody(v *GetDynamicRouteResponseBody) *GetDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *GetDynamicRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
