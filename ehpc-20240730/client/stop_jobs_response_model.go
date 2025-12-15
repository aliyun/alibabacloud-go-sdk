// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopJobsResponse
	GetStatusCode() *int32
	SetBody(v *StopJobsResponseBody) *StopJobsResponse
	GetBody() *StopJobsResponseBody
}

type StopJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopJobsResponse) GetBody() *StopJobsResponseBody {
	return s.Body
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetStatusCode(v int32) *StopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

func (s *StopJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
