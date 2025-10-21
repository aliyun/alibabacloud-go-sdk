// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopJobResponse
	GetStatusCode() *int32
	SetBody(v *StopJobResponseBody) *StopJobResponse
	GetBody() *StopJobResponseBody
}

type StopJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopJobResponse) GetBody() *StopJobResponseBody {
	return s.Body
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
	return s
}

func (s *StopJobResponse) SetStatusCode(v int32) *StopJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

func (s *StopJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
