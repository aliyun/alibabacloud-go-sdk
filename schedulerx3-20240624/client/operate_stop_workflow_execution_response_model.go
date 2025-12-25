// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopWorkflowExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateStopWorkflowExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateStopWorkflowExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateStopWorkflowExecutionResponseBody) *OperateStopWorkflowExecutionResponse
	GetBody() *OperateStopWorkflowExecutionResponseBody
}

type OperateStopWorkflowExecutionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateStopWorkflowExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateStopWorkflowExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateStopWorkflowExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateStopWorkflowExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateStopWorkflowExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateStopWorkflowExecutionResponse) GetBody() *OperateStopWorkflowExecutionResponseBody {
	return s.Body
}

func (s *OperateStopWorkflowExecutionResponse) SetHeaders(v map[string]*string) *OperateStopWorkflowExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateStopWorkflowExecutionResponse) SetStatusCode(v int32) *OperateStopWorkflowExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateStopWorkflowExecutionResponse) SetBody(v *OperateStopWorkflowExecutionResponseBody) *OperateStopWorkflowExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateStopWorkflowExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
