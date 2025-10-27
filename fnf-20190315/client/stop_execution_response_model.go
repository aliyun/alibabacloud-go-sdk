// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StopExecutionResponseBody) *StopExecutionResponse
	GetBody() *StopExecutionResponseBody
}

type StopExecutionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopExecutionResponse) GetBody() *StopExecutionResponseBody {
	return s.Body
}

func (s *StopExecutionResponse) SetHeaders(v map[string]*string) *StopExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopExecutionResponse) SetStatusCode(v int32) *StopExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExecutionResponse) SetBody(v *StopExecutionResponseBody) *StopExecutionResponse {
	s.Body = v
	return s
}

func (s *StopExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
