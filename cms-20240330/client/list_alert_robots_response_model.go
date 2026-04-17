// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRobotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertRobotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertRobotsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertRobotsResponseBody) *ListAlertRobotsResponse
	GetBody() *ListAlertRobotsResponseBody
}

type ListAlertRobotsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertRobotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertRobotsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRobotsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertRobotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertRobotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertRobotsResponse) GetBody() *ListAlertRobotsResponseBody {
	return s.Body
}

func (s *ListAlertRobotsResponse) SetHeaders(v map[string]*string) *ListAlertRobotsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertRobotsResponse) SetStatusCode(v int32) *ListAlertRobotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertRobotsResponse) SetBody(v *ListAlertRobotsResponseBody) *ListAlertRobotsResponse {
	s.Body = v
	return s
}

func (s *ListAlertRobotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
