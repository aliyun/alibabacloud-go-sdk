// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLiveMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLiveMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartLiveMPUTaskResponseBody) *StartLiveMPUTaskResponse
	GetBody() *StartLiveMPUTaskResponseBody
}

type StartLiveMPUTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLiveMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLiveMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLiveMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLiveMPUTaskResponse) GetBody() *StartLiveMPUTaskResponseBody {
	return s.Body
}

func (s *StartLiveMPUTaskResponse) SetHeaders(v map[string]*string) *StartLiveMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StartLiveMPUTaskResponse) SetStatusCode(v int32) *StartLiveMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLiveMPUTaskResponse) SetBody(v *StartLiveMPUTaskResponseBody) *StartLiveMPUTaskResponse {
	s.Body = v
	return s
}

func (s *StartLiveMPUTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
