// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DetectApplicationMonitorResponseBody) *DetectApplicationMonitorResponse
	GetBody() *DetectApplicationMonitorResponseBody
}

type DetectApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectApplicationMonitorResponse) GetBody() *DetectApplicationMonitorResponseBody {
	return s.Body
}

func (s *DetectApplicationMonitorResponse) SetHeaders(v map[string]*string) *DetectApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DetectApplicationMonitorResponse) SetStatusCode(v int32) *DetectApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectApplicationMonitorResponse) SetBody(v *DetectApplicationMonitorResponseBody) *DetectApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *DetectApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
