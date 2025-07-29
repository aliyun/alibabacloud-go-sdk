// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowInstanceResponseBody) *ListWorkflowInstanceResponse
	GetBody() *ListWorkflowInstanceResponseBody
}

type ListWorkflowInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowInstanceResponse) GetBody() *ListWorkflowInstanceResponseBody {
	return s.Body
}

func (s *ListWorkflowInstanceResponse) SetHeaders(v map[string]*string) *ListWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowInstanceResponse) SetStatusCode(v int32) *ListWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowInstanceResponse) SetBody(v *ListWorkflowInstanceResponseBody) *ListWorkflowInstanceResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowInstanceResponse) Validate() error {
	return dara.Validate(s)
}
