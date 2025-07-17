// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWorkflowInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopWorkflowInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopWorkflowInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StopWorkflowInstancesResponseBody) *StopWorkflowInstancesResponse
	GetBody() *StopWorkflowInstancesResponseBody
}

type StopWorkflowInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopWorkflowInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopWorkflowInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopWorkflowInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopWorkflowInstancesResponse) GetBody() *StopWorkflowInstancesResponseBody {
	return s.Body
}

func (s *StopWorkflowInstancesResponse) SetHeaders(v map[string]*string) *StopWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopWorkflowInstancesResponse) SetStatusCode(v int32) *StopWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopWorkflowInstancesResponse) SetBody(v *StopWorkflowInstancesResponseBody) *StopWorkflowInstancesResponse {
	s.Body = v
	return s
}

func (s *StopWorkflowInstancesResponse) Validate() error {
	return dara.Validate(s)
}
