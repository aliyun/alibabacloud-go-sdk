// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineRouteResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineRouteResponseBody) *GetRoutineRouteResponse
	GetBody() *GetRoutineRouteResponseBody
}

type GetRoutineRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineRouteResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineRouteResponse) GetBody() *GetRoutineRouteResponseBody {
	return s.Body
}

func (s *GetRoutineRouteResponse) SetHeaders(v map[string]*string) *GetRoutineRouteResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineRouteResponse) SetStatusCode(v int32) *GetRoutineRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineRouteResponse) SetBody(v *GetRoutineRouteResponseBody) *GetRoutineRouteResponse {
	s.Body = v
	return s
}

func (s *GetRoutineRouteResponse) Validate() error {
	return dara.Validate(s)
}
