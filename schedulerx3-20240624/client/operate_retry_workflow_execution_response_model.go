// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryWorkflowExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateRetryWorkflowExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateRetryWorkflowExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateRetryWorkflowExecutionResponseBody) *OperateRetryWorkflowExecutionResponse
	GetBody() *OperateRetryWorkflowExecutionResponseBody
}

type OperateRetryWorkflowExecutionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateRetryWorkflowExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateRetryWorkflowExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryWorkflowExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateRetryWorkflowExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateRetryWorkflowExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateRetryWorkflowExecutionResponse) GetBody() *OperateRetryWorkflowExecutionResponseBody {
	return s.Body
}

func (s *OperateRetryWorkflowExecutionResponse) SetHeaders(v map[string]*string) *OperateRetryWorkflowExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateRetryWorkflowExecutionResponse) SetStatusCode(v int32) *OperateRetryWorkflowExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateRetryWorkflowExecutionResponse) SetBody(v *OperateRetryWorkflowExecutionResponseBody) *OperateRetryWorkflowExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateRetryWorkflowExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
