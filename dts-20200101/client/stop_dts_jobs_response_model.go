// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDtsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDtsJobsResponse
	GetStatusCode() *int32
	SetBody(v *StopDtsJobsResponseBody) *StopDtsJobsResponse
	GetBody() *StopDtsJobsResponseBody
}

type StopDtsJobsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDtsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *StopDtsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDtsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDtsJobsResponse) GetBody() *StopDtsJobsResponseBody {
	return s.Body
}

func (s *StopDtsJobsResponse) SetHeaders(v map[string]*string) *StopDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *StopDtsJobsResponse) SetStatusCode(v int32) *StopDtsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDtsJobsResponse) SetBody(v *StopDtsJobsResponseBody) *StopDtsJobsResponse {
	s.Body = v
	return s
}

func (s *StopDtsJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
