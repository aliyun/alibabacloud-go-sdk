// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCancelCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskCancelCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskCancelCallResponse
	GetStatusCode() *int32
	SetBody(v *TaskCancelCallResponseBody) *TaskCancelCallResponse
	GetBody() *TaskCancelCallResponseBody
}

type TaskCancelCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCancelCallResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskCancelCallResponse) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskCancelCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskCancelCallResponse) GetBody() *TaskCancelCallResponseBody {
	return s.Body
}

func (s *TaskCancelCallResponse) SetHeaders(v map[string]*string) *TaskCancelCallResponse {
	s.Headers = v
	return s
}

func (s *TaskCancelCallResponse) SetStatusCode(v int32) *TaskCancelCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCancelCallResponse) SetBody(v *TaskCancelCallResponseBody) *TaskCancelCallResponse {
	s.Body = v
	return s
}

func (s *TaskCancelCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
