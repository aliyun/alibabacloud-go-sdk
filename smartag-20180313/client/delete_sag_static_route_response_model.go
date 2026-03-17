// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagStaticRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSagStaticRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSagStaticRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSagStaticRouteResponseBody) *DeleteSagStaticRouteResponse
	GetBody() *DeleteSagStaticRouteResponseBody
}

type DeleteSagStaticRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSagStaticRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSagStaticRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagStaticRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteSagStaticRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSagStaticRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSagStaticRouteResponse) GetBody() *DeleteSagStaticRouteResponseBody {
	return s.Body
}

func (s *DeleteSagStaticRouteResponse) SetHeaders(v map[string]*string) *DeleteSagStaticRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteSagStaticRouteResponse) SetStatusCode(v int32) *DeleteSagStaticRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSagStaticRouteResponse) SetBody(v *DeleteSagStaticRouteResponseBody) *DeleteSagStaticRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteSagStaticRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
