// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineRouteResponseBody) *CreateRoutineRouteResponse
	GetBody() *CreateRoutineRouteResponseBody
}

type CreateRoutineRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineRouteResponse) GetBody() *CreateRoutineRouteResponseBody {
	return s.Body
}

func (s *CreateRoutineRouteResponse) SetHeaders(v map[string]*string) *CreateRoutineRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineRouteResponse) SetStatusCode(v int32) *CreateRoutineRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineRouteResponse) SetBody(v *CreateRoutineRouteResponseBody) *CreateRoutineRouteResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
