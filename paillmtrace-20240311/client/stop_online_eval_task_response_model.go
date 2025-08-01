// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOnlineEvalTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopOnlineEvalTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopOnlineEvalTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopOnlineEvalTaskResponseBody) *StopOnlineEvalTaskResponse
	GetBody() *StopOnlineEvalTaskResponseBody
}

type StopOnlineEvalTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopOnlineEvalTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *StopOnlineEvalTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopOnlineEvalTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopOnlineEvalTaskResponse) GetBody() *StopOnlineEvalTaskResponseBody {
	return s.Body
}

func (s *StopOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *StopOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *StopOnlineEvalTaskResponse) SetStatusCode(v int32) *StopOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopOnlineEvalTaskResponse) SetBody(v *StopOnlineEvalTaskResponseBody) *StopOnlineEvalTaskResponse {
	s.Body = v
	return s
}

func (s *StopOnlineEvalTaskResponse) Validate() error {
	return dara.Validate(s)
}
