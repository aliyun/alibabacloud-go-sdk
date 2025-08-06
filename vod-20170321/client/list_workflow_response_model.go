// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkflowResponseBody) *ListWorkflowResponse
	GetBody() *ListWorkflowResponseBody
}

type ListWorkflowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkflowResponse) GetBody() *ListWorkflowResponseBody {
	return s.Body
}

func (s *ListWorkflowResponse) SetHeaders(v map[string]*string) *ListWorkflowResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowResponse) SetStatusCode(v int32) *ListWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowResponse) SetBody(v *ListWorkflowResponseBody) *ListWorkflowResponse {
	s.Body = v
	return s
}

func (s *ListWorkflowResponse) Validate() error {
	return dara.Validate(s)
}
