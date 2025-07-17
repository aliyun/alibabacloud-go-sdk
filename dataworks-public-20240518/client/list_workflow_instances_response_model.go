// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowInstancesResponseBody) *ListWorkflowInstancesResponse
	GetBody() *ListWorkflowInstancesResponseBody
}

type ListWorkflowInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowInstancesResponse) GetBody() *ListWorkflowInstancesResponseBody {
	return s.Body
}

func (s *ListWorkflowInstancesResponse) SetHeaders(v map[string]*string) *ListWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowInstancesResponse) SetStatusCode(v int32) *ListWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowInstancesResponse) SetBody(v *ListWorkflowInstancesResponseBody) *ListWorkflowInstancesResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowInstancesResponse) Validate() error {
	return dara.Validate(s)
}
