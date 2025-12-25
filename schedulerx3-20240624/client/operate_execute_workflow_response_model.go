// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateExecuteWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateExecuteWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *OperateExecuteWorkflowResponseBody) *OperateExecuteWorkflowResponse
	GetBody() *OperateExecuteWorkflowResponseBody
}

type OperateExecuteWorkflowResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateExecuteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateExecuteWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteWorkflowResponse) GoString() string {
	return s.String()
}

func (s *OperateExecuteWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateExecuteWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateExecuteWorkflowResponse) GetBody() *OperateExecuteWorkflowResponseBody {
	return s.Body
}

func (s *OperateExecuteWorkflowResponse) SetHeaders(v map[string]*string) *OperateExecuteWorkflowResponse {
	s.Headers = v
	return s
}

func (s *OperateExecuteWorkflowResponse) SetStatusCode(v int32) *OperateExecuteWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateExecuteWorkflowResponse) SetBody(v *OperateExecuteWorkflowResponseBody) *OperateExecuteWorkflowResponse {
	s.Body = v
	return s
}

func (s *OperateExecuteWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
