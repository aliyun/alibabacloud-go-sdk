// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRobotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertRobotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertRobotsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertRobotsResponseBody) *DeleteAlertRobotsResponse
	GetBody() *DeleteAlertRobotsResponseBody
}

type DeleteAlertRobotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertRobotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertRobotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRobotsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRobotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertRobotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertRobotsResponse) GetBody() *DeleteAlertRobotsResponseBody {
	return s.Body
}

func (s *DeleteAlertRobotsResponse) SetHeaders(v map[string]*string) *DeleteAlertRobotsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRobotsResponse) SetStatusCode(v int32) *DeleteAlertRobotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRobotsResponse) SetBody(v *DeleteAlertRobotsResponseBody) *DeleteAlertRobotsResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertRobotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
