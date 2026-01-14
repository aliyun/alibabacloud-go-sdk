// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationMonitorResponseBody) *ListApplicationMonitorResponse
	GetBody() *ListApplicationMonitorResponseBody
}

type ListApplicationMonitorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationMonitorResponse) GetBody() *ListApplicationMonitorResponseBody {
	return s.Body
}

func (s *ListApplicationMonitorResponse) SetHeaders(v map[string]*string) *ListApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationMonitorResponse) SetStatusCode(v int32) *ListApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorResponse) SetBody(v *ListApplicationMonitorResponseBody) *ListApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *ListApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
