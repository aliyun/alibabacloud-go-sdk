// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopTaskResponseBody) *StopTaskResponse
	GetBody() *StopTaskResponseBody
}

type StopTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTaskResponse) GetBody() *StopTaskResponseBody {
	return s.Body
}

func (s *StopTaskResponse) SetHeaders(v map[string]*string) *StopTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTaskResponse) SetStatusCode(v int32) *StopTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskResponse) SetBody(v *StopTaskResponseBody) *StopTaskResponse {
	s.Body = v
	return s
}

func (s *StopTaskResponse) Validate() error {
	return dara.Validate(s)
}
