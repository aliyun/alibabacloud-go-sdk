// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAiCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAiCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopAiCallTaskResponseBody) *StopAiCallTaskResponse
	GetBody() *StopAiCallTaskResponseBody
}

type StopAiCallTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAiCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAiCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAiCallTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAiCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAiCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAiCallTaskResponse) GetBody() *StopAiCallTaskResponseBody {
	return s.Body
}

func (s *StopAiCallTaskResponse) SetHeaders(v map[string]*string) *StopAiCallTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAiCallTaskResponse) SetStatusCode(v int32) *StopAiCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAiCallTaskResponse) SetBody(v *StopAiCallTaskResponseBody) *StopAiCallTaskResponse {
	s.Body = v
	return s
}

func (s *StopAiCallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
