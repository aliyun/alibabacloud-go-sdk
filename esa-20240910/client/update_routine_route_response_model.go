// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRoutineRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRoutineRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRoutineRouteResponseBody) *UpdateRoutineRouteResponse
	GetBody() *UpdateRoutineRouteResponseBody
}

type UpdateRoutineRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRoutineRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRoutineRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoutineRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRoutineRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRoutineRouteResponse) GetBody() *UpdateRoutineRouteResponseBody {
	return s.Body
}

func (s *UpdateRoutineRouteResponse) SetHeaders(v map[string]*string) *UpdateRoutineRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoutineRouteResponse) SetStatusCode(v int32) *UpdateRoutineRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoutineRouteResponse) SetBody(v *UpdateRoutineRouteResponseBody) *UpdateRoutineRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateRoutineRouteResponse) Validate() error {
	return dara.Validate(s)
}
