// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaypointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWaypointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWaypointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWaypointResponseBody) *UpdateWaypointResponse
	GetBody() *UpdateWaypointResponseBody
}

type UpdateWaypointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaypointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaypointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaypointResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaypointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWaypointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWaypointResponse) GetBody() *UpdateWaypointResponseBody {
	return s.Body
}

func (s *UpdateWaypointResponse) SetHeaders(v map[string]*string) *UpdateWaypointResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaypointResponse) SetStatusCode(v int32) *UpdateWaypointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaypointResponse) SetBody(v *UpdateWaypointResponseBody) *UpdateWaypointResponse {
	s.Body = v
	return s
}

func (s *UpdateWaypointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
