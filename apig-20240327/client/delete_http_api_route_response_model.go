// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpApiRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpApiRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpApiRouteResponseBody) *DeleteHttpApiRouteResponse
	GetBody() *DeleteHttpApiRouteResponseBody
}

type DeleteHttpApiRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpApiRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpApiRouteResponse) GetBody() *DeleteHttpApiRouteResponseBody {
	return s.Body
}

func (s *DeleteHttpApiRouteResponse) SetHeaders(v map[string]*string) *DeleteHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiRouteResponse) SetStatusCode(v int32) *DeleteHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiRouteResponse) SetBody(v *DeleteHttpApiRouteResponseBody) *DeleteHttpApiRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpApiRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
