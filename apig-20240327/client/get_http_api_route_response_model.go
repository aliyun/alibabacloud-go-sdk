// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpApiRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpApiRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpApiRouteResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpApiRouteResponseBody) *GetHttpApiRouteResponse
	GetBody() *GetHttpApiRouteResponseBody
}

type GetHttpApiRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpApiRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpApiRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpApiRouteResponse) GoString() string {
	return s.String()
}

func (s *GetHttpApiRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpApiRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpApiRouteResponse) GetBody() *GetHttpApiRouteResponseBody {
	return s.Body
}

func (s *GetHttpApiRouteResponse) SetHeaders(v map[string]*string) *GetHttpApiRouteResponse {
	s.Headers = v
	return s
}

func (s *GetHttpApiRouteResponse) SetStatusCode(v int32) *GetHttpApiRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpApiRouteResponse) SetBody(v *GetHttpApiRouteResponseBody) *GetHttpApiRouteResponse {
	s.Body = v
	return s
}

func (s *GetHttpApiRouteResponse) Validate() error {
	return dara.Validate(s)
}
