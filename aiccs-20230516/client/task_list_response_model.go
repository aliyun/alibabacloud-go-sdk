// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskListResponse
	GetStatusCode() *int32
	SetBody(v *TaskListResponseBody) *TaskListResponse
	GetBody() *TaskListResponseBody
}

type TaskListResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskListResponse) GoString() string {
	return s.String()
}

func (s *TaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskListResponse) GetBody() *TaskListResponseBody {
	return s.Body
}

func (s *TaskListResponse) SetHeaders(v map[string]*string) *TaskListResponse {
	s.Headers = v
	return s
}

func (s *TaskListResponse) SetStatusCode(v int32) *TaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskListResponse) SetBody(v *TaskListResponseBody) *TaskListResponse {
	s.Body = v
	return s
}

func (s *TaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
