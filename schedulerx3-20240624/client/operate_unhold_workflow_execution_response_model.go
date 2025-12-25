// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdWorkflowExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateUnholdWorkflowExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateUnholdWorkflowExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateUnholdWorkflowExecutionResponseBody) *OperateUnholdWorkflowExecutionResponse
	GetBody() *OperateUnholdWorkflowExecutionResponseBody
}

type OperateUnholdWorkflowExecutionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateUnholdWorkflowExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateUnholdWorkflowExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdWorkflowExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateUnholdWorkflowExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateUnholdWorkflowExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateUnholdWorkflowExecutionResponse) GetBody() *OperateUnholdWorkflowExecutionResponseBody {
	return s.Body
}

func (s *OperateUnholdWorkflowExecutionResponse) SetHeaders(v map[string]*string) *OperateUnholdWorkflowExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponse) SetStatusCode(v int32) *OperateUnholdWorkflowExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponse) SetBody(v *OperateUnholdWorkflowExecutionResponseBody) *OperateUnholdWorkflowExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateUnholdWorkflowExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
