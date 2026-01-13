// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopTrafficControlTaskResponseBody) *StopTrafficControlTaskResponse
	GetBody() *StopTrafficControlTaskResponseBody
}

type StopTrafficControlTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTrafficControlTaskResponse) GetBody() *StopTrafficControlTaskResponseBody {
	return s.Body
}

func (s *StopTrafficControlTaskResponse) SetHeaders(v map[string]*string) *StopTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTrafficControlTaskResponse) SetStatusCode(v int32) *StopTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrafficControlTaskResponse) SetBody(v *StopTrafficControlTaskResponseBody) *StopTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *StopTrafficControlTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
