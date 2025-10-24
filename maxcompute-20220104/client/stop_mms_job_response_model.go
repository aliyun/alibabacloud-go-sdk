// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *StopMmsJobResponseBody) *StopMmsJobResponse
	GetBody() *StopMmsJobResponseBody
}

type StopMmsJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMmsJobResponse) GoString() string {
	return s.String()
}

func (s *StopMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMmsJobResponse) GetBody() *StopMmsJobResponseBody {
	return s.Body
}

func (s *StopMmsJobResponse) SetHeaders(v map[string]*string) *StopMmsJobResponse {
	s.Headers = v
	return s
}

func (s *StopMmsJobResponse) SetStatusCode(v int32) *StopMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMmsJobResponse) SetBody(v *StopMmsJobResponseBody) *StopMmsJobResponse {
	s.Body = v
	return s
}

func (s *StopMmsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
