// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLiveMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLiveMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopLiveMPUTaskResponseBody) *StopLiveMPUTaskResponse
	GetBody() *StopLiveMPUTaskResponseBody
}

type StopLiveMPUTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLiveMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLiveMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLiveMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StopLiveMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLiveMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLiveMPUTaskResponse) GetBody() *StopLiveMPUTaskResponseBody {
	return s.Body
}

func (s *StopLiveMPUTaskResponse) SetHeaders(v map[string]*string) *StopLiveMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StopLiveMPUTaskResponse) SetStatusCode(v int32) *StopLiveMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveMPUTaskResponse) SetBody(v *StopLiveMPUTaskResponseBody) *StopLiveMPUTaskResponse {
	s.Body = v
	return s
}

func (s *StopLiveMPUTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
