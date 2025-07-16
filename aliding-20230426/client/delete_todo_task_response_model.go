// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse
	GetBody() *DeleteTodoTaskResponseBody
}

type DeleteTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTodoTaskResponse) GetBody() *DeleteTodoTaskResponseBody {
	return s.Body
}

func (s *DeleteTodoTaskResponse) SetHeaders(v map[string]*string) *DeleteTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTodoTaskResponse) SetStatusCode(v int32) *DeleteTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTodoTaskResponse) SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteTodoTaskResponse) Validate() error {
	return dara.Validate(s)
}
