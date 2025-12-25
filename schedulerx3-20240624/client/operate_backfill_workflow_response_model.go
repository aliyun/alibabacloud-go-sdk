// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBackfillWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateBackfillWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateBackfillWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *OperateBackfillWorkflowResponseBody) *OperateBackfillWorkflowResponse
	GetBody() *OperateBackfillWorkflowResponseBody
}

type OperateBackfillWorkflowResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateBackfillWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateBackfillWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateBackfillWorkflowResponse) GoString() string {
	return s.String()
}

func (s *OperateBackfillWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateBackfillWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateBackfillWorkflowResponse) GetBody() *OperateBackfillWorkflowResponseBody {
	return s.Body
}

func (s *OperateBackfillWorkflowResponse) SetHeaders(v map[string]*string) *OperateBackfillWorkflowResponse {
	s.Headers = v
	return s
}

func (s *OperateBackfillWorkflowResponse) SetStatusCode(v int32) *OperateBackfillWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateBackfillWorkflowResponse) SetBody(v *OperateBackfillWorkflowResponseBody) *OperateBackfillWorkflowResponse {
	s.Body = v
	return s
}

func (s *OperateBackfillWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
