// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpApiRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpApiRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpApiRouteResponseBody) *CreateHttpApiRouteResponse
	GetBody() *CreateHttpApiRouteResponseBody
}

type CreateHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpApiRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpApiRouteResponse) GetBody() *CreateHttpApiRouteResponseBody {
	return s.Body
}

func (s *CreateHttpApiRouteResponse) SetHeaders(v map[string]*string) *CreateHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiRouteResponse) SetStatusCode(v int32) *CreateHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiRouteResponse) SetBody(v *CreateHttpApiRouteResponseBody) *CreateHttpApiRouteResponse {
	s.Body = v
	return s
}

func (s *CreateHttpApiRouteResponse) Validate() error {
	return dara.Validate(s)
}
