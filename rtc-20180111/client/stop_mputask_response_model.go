// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopMPUTaskResponseBody) *StopMPUTaskResponse
	GetBody() *StopMPUTaskResponseBody
}

type StopMPUTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StopMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMPUTaskResponse) GetBody() *StopMPUTaskResponseBody {
	return s.Body
}

func (s *StopMPUTaskResponse) SetHeaders(v map[string]*string) *StopMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StopMPUTaskResponse) SetStatusCode(v int32) *StopMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMPUTaskResponse) SetBody(v *StopMPUTaskResponseBody) *StopMPUTaskResponse {
	s.Body = v
	return s
}

func (s *StopMPUTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
