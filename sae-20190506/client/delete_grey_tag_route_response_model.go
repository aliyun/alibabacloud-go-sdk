// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGreyTagRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGreyTagRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGreyTagRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGreyTagRouteResponseBody) *DeleteGreyTagRouteResponse
	GetBody() *DeleteGreyTagRouteResponseBody
}

type DeleteGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGreyTagRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGreyTagRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGreyTagRouteResponse) GetBody() *DeleteGreyTagRouteResponseBody {
	return s.Body
}

func (s *DeleteGreyTagRouteResponse) SetHeaders(v map[string]*string) *DeleteGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteGreyTagRouteResponse) SetStatusCode(v int32) *DeleteGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGreyTagRouteResponse) SetBody(v *DeleteGreyTagRouteResponseBody) *DeleteGreyTagRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteGreyTagRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
