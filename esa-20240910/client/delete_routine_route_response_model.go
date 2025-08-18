// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineRouteResponseBody) *DeleteRoutineRouteResponse
	GetBody() *DeleteRoutineRouteResponseBody
}

type DeleteRoutineRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineRouteResponse) GetBody() *DeleteRoutineRouteResponseBody {
	return s.Body
}

func (s *DeleteRoutineRouteResponse) SetHeaders(v map[string]*string) *DeleteRoutineRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineRouteResponse) SetStatusCode(v int32) *DeleteRoutineRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineRouteResponse) SetBody(v *DeleteRoutineRouteResponseBody) *DeleteRoutineRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineRouteResponse) Validate() error {
	return dara.Validate(s)
}
