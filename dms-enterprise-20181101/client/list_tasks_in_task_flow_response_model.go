// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksInTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTasksInTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTasksInTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListTasksInTaskFlowResponseBody) *ListTasksInTaskFlowResponse
	GetBody() *ListTasksInTaskFlowResponseBody
}

type ListTasksInTaskFlowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksInTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksInTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTasksInTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *ListTasksInTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTasksInTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTasksInTaskFlowResponse) GetBody() *ListTasksInTaskFlowResponseBody {
	return s.Body
}

func (s *ListTasksInTaskFlowResponse) SetHeaders(v map[string]*string) *ListTasksInTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *ListTasksInTaskFlowResponse) SetStatusCode(v int32) *ListTasksInTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksInTaskFlowResponse) SetBody(v *ListTasksInTaskFlowResponseBody) *ListTasksInTaskFlowResponse {
	s.Body = v
	return s
}

func (s *ListTasksInTaskFlowResponse) Validate() error {
	return dara.Validate(s)
}
