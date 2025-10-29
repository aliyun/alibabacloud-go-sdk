// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRouteResponse
	GetStatusCode() *int32
	SetBody(v *GetRouteResponseBody) *GetRouteResponse
	GetBody() *GetRouteResponseBody
}

type GetRouteResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRouteResponse) GoString() string {
	return s.String()
}

func (s *GetRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRouteResponse) GetBody() *GetRouteResponseBody {
	return s.Body
}

func (s *GetRouteResponse) SetHeaders(v map[string]*string) *GetRouteResponse {
	s.Headers = v
	return s
}

func (s *GetRouteResponse) SetStatusCode(v int32) *GetRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRouteResponse) SetBody(v *GetRouteResponseBody) *GetRouteResponse {
	s.Body = v
	return s
}

func (s *GetRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
