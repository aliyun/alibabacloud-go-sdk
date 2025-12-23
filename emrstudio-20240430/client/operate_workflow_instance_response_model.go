// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWorkflowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateWorkflowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateWorkflowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *OperateWorkflowInstanceResponseBody) *OperateWorkflowInstanceResponse
	GetBody() *OperateWorkflowInstanceResponseBody
}

type OperateWorkflowInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateWorkflowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *OperateWorkflowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateWorkflowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateWorkflowInstanceResponse) GetBody() *OperateWorkflowInstanceResponseBody {
	return s.Body
}

func (s *OperateWorkflowInstanceResponse) SetHeaders(v map[string]*string) *OperateWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *OperateWorkflowInstanceResponse) SetStatusCode(v int32) *OperateWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateWorkflowInstanceResponse) SetBody(v *OperateWorkflowInstanceResponseBody) *OperateWorkflowInstanceResponse {
	s.Body = v
	return s
}

func (s *OperateWorkflowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
