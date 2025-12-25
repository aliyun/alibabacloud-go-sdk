// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateEnableWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateEnableWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *OperateEnableWorkflowsResponseBody) *OperateEnableWorkflowsResponse
	GetBody() *OperateEnableWorkflowsResponseBody
}

type OperateEnableWorkflowsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateEnableWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateEnableWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *OperateEnableWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateEnableWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateEnableWorkflowsResponse) GetBody() *OperateEnableWorkflowsResponseBody {
	return s.Body
}

func (s *OperateEnableWorkflowsResponse) SetHeaders(v map[string]*string) *OperateEnableWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *OperateEnableWorkflowsResponse) SetStatusCode(v int32) *OperateEnableWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateEnableWorkflowsResponse) SetBody(v *OperateEnableWorkflowsResponseBody) *OperateEnableWorkflowsResponse {
	s.Body = v
	return s
}

func (s *OperateEnableWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
