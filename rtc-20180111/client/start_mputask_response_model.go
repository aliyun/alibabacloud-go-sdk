// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartMPUTaskResponseBody) *StartMPUTaskResponse
	GetBody() *StartMPUTaskResponseBody
}

type StartMPUTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StartMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMPUTaskResponse) GetBody() *StartMPUTaskResponseBody {
	return s.Body
}

func (s *StartMPUTaskResponse) SetHeaders(v map[string]*string) *StartMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StartMPUTaskResponse) SetStatusCode(v int32) *StartMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMPUTaskResponse) SetBody(v *StartMPUTaskResponseBody) *StartMPUTaskResponse {
	s.Body = v
	return s
}

func (s *StartMPUTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
