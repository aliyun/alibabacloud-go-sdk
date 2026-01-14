// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTaskResponse
	GetStatusCode() *int32
	SetBody(v *AddTaskResponseBody) *AddTaskResponse
	GetBody() *AddTaskResponseBody
}

type AddTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTaskResponse) GoString() string {
	return s.String()
}

func (s *AddTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTaskResponse) GetBody() *AddTaskResponseBody {
	return s.Body
}

func (s *AddTaskResponse) SetHeaders(v map[string]*string) *AddTaskResponse {
	s.Headers = v
	return s
}

func (s *AddTaskResponse) SetStatusCode(v int32) *AddTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTaskResponse) SetBody(v *AddTaskResponseBody) *AddTaskResponse {
	s.Body = v
	return s
}

func (s *AddTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
