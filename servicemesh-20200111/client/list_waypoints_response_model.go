// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaypointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWaypointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWaypointsResponse
	GetStatusCode() *int32
	SetBody(v *ListWaypointsResponseBody) *ListWaypointsResponse
	GetBody() *ListWaypointsResponseBody
}

type ListWaypointsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaypointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaypointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWaypointsResponse) GoString() string {
	return s.String()
}

func (s *ListWaypointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWaypointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWaypointsResponse) GetBody() *ListWaypointsResponseBody {
	return s.Body
}

func (s *ListWaypointsResponse) SetHeaders(v map[string]*string) *ListWaypointsResponse {
	s.Headers = v
	return s
}

func (s *ListWaypointsResponse) SetStatusCode(v int32) *ListWaypointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaypointsResponse) SetBody(v *ListWaypointsResponseBody) *ListWaypointsResponse {
	s.Body = v
	return s
}

func (s *ListWaypointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
