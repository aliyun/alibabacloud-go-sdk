// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessWorkflowExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateMarkSuccessWorkflowExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateMarkSuccessWorkflowExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateMarkSuccessWorkflowExecutionResponseBody) *OperateMarkSuccessWorkflowExecutionResponse
	GetBody() *OperateMarkSuccessWorkflowExecutionResponseBody
}

type OperateMarkSuccessWorkflowExecutionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateMarkSuccessWorkflowExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateMarkSuccessWorkflowExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessWorkflowExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) GetBody() *OperateMarkSuccessWorkflowExecutionResponseBody {
	return s.Body
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) SetHeaders(v map[string]*string) *OperateMarkSuccessWorkflowExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) SetStatusCode(v int32) *OperateMarkSuccessWorkflowExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) SetBody(v *OperateMarkSuccessWorkflowExecutionResponseBody) *OperateMarkSuccessWorkflowExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
