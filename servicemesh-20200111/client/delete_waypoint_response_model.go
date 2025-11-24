// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaypointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWaypointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWaypointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWaypointResponseBody) *DeleteWaypointResponse
	GetBody() *DeleteWaypointResponseBody
}

type DeleteWaypointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaypointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaypointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaypointResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaypointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWaypointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWaypointResponse) GetBody() *DeleteWaypointResponseBody {
	return s.Body
}

func (s *DeleteWaypointResponse) SetHeaders(v map[string]*string) *DeleteWaypointResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaypointResponse) SetStatusCode(v int32) *DeleteWaypointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaypointResponse) SetBody(v *DeleteWaypointResponseBody) *DeleteWaypointResponse {
	s.Body = v
	return s
}

func (s *DeleteWaypointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
