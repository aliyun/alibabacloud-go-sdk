// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowTasksResponseBody) *ListWorkflowTasksResponse
	GetBody() *ListWorkflowTasksResponseBody
}

type ListWorkflowTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowTasksResponse) GetBody() *ListWorkflowTasksResponseBody {
	return s.Body
}

func (s *ListWorkflowTasksResponse) SetHeaders(v map[string]*string) *ListWorkflowTasksResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowTasksResponse) SetStatusCode(v int32) *ListWorkflowTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowTasksResponse) SetBody(v *ListWorkflowTasksResponseBody) *ListWorkflowTasksResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowTasksResponse) Validate() error {
	return dara.Validate(s)
}
