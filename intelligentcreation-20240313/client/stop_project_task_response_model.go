// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProjectTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopProjectTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopProjectTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopProjectTaskResponseBody) *StopProjectTaskResponse
	GetBody() *StopProjectTaskResponseBody
}

type StopProjectTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopProjectTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *StopProjectTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopProjectTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopProjectTaskResponse) GetBody() *StopProjectTaskResponseBody {
	return s.Body
}

func (s *StopProjectTaskResponse) SetHeaders(v map[string]*string) *StopProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *StopProjectTaskResponse) SetStatusCode(v int32) *StopProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopProjectTaskResponse) SetBody(v *StopProjectTaskResponseBody) *StopProjectTaskResponse {
	s.Body = v
	return s
}

func (s *StopProjectTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
