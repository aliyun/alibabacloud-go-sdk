// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrgTodoTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrgTodoTasksResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrgTodoTasksResponseBody) *QueryOrgTodoTasksResponse
	GetBody() *QueryOrgTodoTasksResponseBody
}

type QueryOrgTodoTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgTodoTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrgTodoTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrgTodoTasksResponse) GetBody() *QueryOrgTodoTasksResponseBody {
	return s.Body
}

func (s *QueryOrgTodoTasksResponse) SetHeaders(v map[string]*string) *QueryOrgTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetStatusCode(v int32) *QueryOrgTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetBody(v *QueryOrgTodoTasksResponseBody) *QueryOrgTodoTasksResponse {
	s.Body = v
	return s
}

func (s *QueryOrgTodoTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
