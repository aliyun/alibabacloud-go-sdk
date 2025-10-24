// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopBenchmarkTaskResponseBody) *StopBenchmarkTaskResponse
	GetBody() *StopBenchmarkTaskResponseBody
}

type StopBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopBenchmarkTaskResponse) GetBody() *StopBenchmarkTaskResponseBody {
	return s.Body
}

func (s *StopBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StopBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StopBenchmarkTaskResponse) SetStatusCode(v int32) *StopBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBenchmarkTaskResponse) SetBody(v *StopBenchmarkTaskResponseBody) *StopBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *StopBenchmarkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
