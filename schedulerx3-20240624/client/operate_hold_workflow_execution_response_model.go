// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldWorkflowExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateHoldWorkflowExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateHoldWorkflowExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateHoldWorkflowExecutionResponseBody) *OperateHoldWorkflowExecutionResponse
	GetBody() *OperateHoldWorkflowExecutionResponseBody
}

type OperateHoldWorkflowExecutionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateHoldWorkflowExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateHoldWorkflowExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldWorkflowExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateHoldWorkflowExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateHoldWorkflowExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateHoldWorkflowExecutionResponse) GetBody() *OperateHoldWorkflowExecutionResponseBody {
	return s.Body
}

func (s *OperateHoldWorkflowExecutionResponse) SetHeaders(v map[string]*string) *OperateHoldWorkflowExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateHoldWorkflowExecutionResponse) SetStatusCode(v int32) *OperateHoldWorkflowExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateHoldWorkflowExecutionResponse) SetBody(v *OperateHoldWorkflowExecutionResponseBody) *OperateHoldWorkflowExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateHoldWorkflowExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
