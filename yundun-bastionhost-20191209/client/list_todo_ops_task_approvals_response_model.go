// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodoOpsTaskApprovalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTodoOpsTaskApprovalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTodoOpsTaskApprovalsResponse
	GetStatusCode() *int32
	SetBody(v *ListTodoOpsTaskApprovalsResponseBody) *ListTodoOpsTaskApprovalsResponse
	GetBody() *ListTodoOpsTaskApprovalsResponseBody
}

type ListTodoOpsTaskApprovalsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTodoOpsTaskApprovalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTodoOpsTaskApprovalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTodoOpsTaskApprovalsResponse) GoString() string {
	return s.String()
}

func (s *ListTodoOpsTaskApprovalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTodoOpsTaskApprovalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTodoOpsTaskApprovalsResponse) GetBody() *ListTodoOpsTaskApprovalsResponseBody {
	return s.Body
}

func (s *ListTodoOpsTaskApprovalsResponse) SetHeaders(v map[string]*string) *ListTodoOpsTaskApprovalsResponse {
	s.Headers = v
	return s
}

func (s *ListTodoOpsTaskApprovalsResponse) SetStatusCode(v int32) *ListTodoOpsTaskApprovalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsResponse) SetBody(v *ListTodoOpsTaskApprovalsResponseBody) *ListTodoOpsTaskApprovalsResponse {
	s.Body = v
	return s
}

func (s *ListTodoOpsTaskApprovalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
