// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *StopDtsJobResponseBody) *StopDtsJobResponse
	GetBody() *StopDtsJobResponseBody
}

type StopDtsJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobResponse) GoString() string {
	return s.String()
}

func (s *StopDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDtsJobResponse) GetBody() *StopDtsJobResponseBody {
	return s.Body
}

func (s *StopDtsJobResponse) SetHeaders(v map[string]*string) *StopDtsJobResponse {
	s.Headers = v
	return s
}

func (s *StopDtsJobResponse) SetStatusCode(v int32) *StopDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDtsJobResponse) SetBody(v *StopDtsJobResponseBody) *StopDtsJobResponse {
	s.Body = v
	return s
}

func (s *StopDtsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
