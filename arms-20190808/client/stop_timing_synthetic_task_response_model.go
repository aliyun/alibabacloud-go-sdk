// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTimingSyntheticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTimingSyntheticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTimingSyntheticTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopTimingSyntheticTaskResponseBody) *StopTimingSyntheticTaskResponse
	GetBody() *StopTimingSyntheticTaskResponseBody
}

type StopTimingSyntheticTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTimingSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTimingSyntheticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTimingSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTimingSyntheticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTimingSyntheticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTimingSyntheticTaskResponse) GetBody() *StopTimingSyntheticTaskResponseBody {
	return s.Body
}

func (s *StopTimingSyntheticTaskResponse) SetHeaders(v map[string]*string) *StopTimingSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTimingSyntheticTaskResponse) SetStatusCode(v int32) *StopTimingSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTimingSyntheticTaskResponse) SetBody(v *StopTimingSyntheticTaskResponseBody) *StopTimingSyntheticTaskResponse {
	s.Body = v
	return s
}

func (s *StopTimingSyntheticTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
