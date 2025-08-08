// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTaskResponse
	GetStatusCode() *int32
	SetBody(v *Task) *StartTaskResponse
	GetBody() *Task
}

type StartTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTaskResponse) GetBody() *Task {
	return s.Body
}

func (s *StartTaskResponse) SetHeaders(v map[string]*string) *StartTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTaskResponse) SetStatusCode(v int32) *StartTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTaskResponse) SetBody(v *Task) *StartTaskResponse {
	s.Body = v
	return s
}

func (s *StartTaskResponse) Validate() error {
	return dara.Validate(s)
}
