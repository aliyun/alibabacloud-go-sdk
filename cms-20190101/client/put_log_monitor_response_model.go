// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLogMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutLogMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutLogMonitorResponse
	GetStatusCode() *int32
	SetBody(v *PutLogMonitorResponseBody) *PutLogMonitorResponse
	GetBody() *PutLogMonitorResponseBody
}

type PutLogMonitorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutLogMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutLogMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorResponse) GoString() string {
	return s.String()
}

func (s *PutLogMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutLogMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutLogMonitorResponse) GetBody() *PutLogMonitorResponseBody {
	return s.Body
}

func (s *PutLogMonitorResponse) SetHeaders(v map[string]*string) *PutLogMonitorResponse {
	s.Headers = v
	return s
}

func (s *PutLogMonitorResponse) SetStatusCode(v int32) *PutLogMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *PutLogMonitorResponse) SetBody(v *PutLogMonitorResponseBody) *PutLogMonitorResponse {
	s.Body = v
	return s
}

func (s *PutLogMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
