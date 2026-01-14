// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationMonitorDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationMonitorDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationMonitorDetectResultResponseBody) *ListApplicationMonitorDetectResultResponse
	GetBody() *ListApplicationMonitorDetectResultResponseBody
}

type ListApplicationMonitorDetectResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationMonitorDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationMonitorDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationMonitorDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationMonitorDetectResultResponse) GetBody() *ListApplicationMonitorDetectResultResponseBody {
	return s.Body
}

func (s *ListApplicationMonitorDetectResultResponse) SetHeaders(v map[string]*string) *ListApplicationMonitorDetectResultResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationMonitorDetectResultResponse) SetStatusCode(v int32) *ListApplicationMonitorDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponse) SetBody(v *ListApplicationMonitorDetectResultResponseBody) *ListApplicationMonitorDetectResultResponse {
	s.Body = v
	return s
}

func (s *ListApplicationMonitorDetectResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
