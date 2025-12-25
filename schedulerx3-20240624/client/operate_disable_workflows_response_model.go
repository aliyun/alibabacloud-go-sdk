// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateDisableWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateDisableWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *OperateDisableWorkflowsResponseBody) *OperateDisableWorkflowsResponse
	GetBody() *OperateDisableWorkflowsResponseBody
}

type OperateDisableWorkflowsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateDisableWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateDisableWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *OperateDisableWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateDisableWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateDisableWorkflowsResponse) GetBody() *OperateDisableWorkflowsResponseBody {
	return s.Body
}

func (s *OperateDisableWorkflowsResponse) SetHeaders(v map[string]*string) *OperateDisableWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *OperateDisableWorkflowsResponse) SetStatusCode(v int32) *OperateDisableWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateDisableWorkflowsResponse) SetBody(v *OperateDisableWorkflowsResponseBody) *OperateDisableWorkflowsResponse {
	s.Body = v
	return s
}

func (s *OperateDisableWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
