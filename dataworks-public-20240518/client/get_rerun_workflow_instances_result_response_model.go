// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRerunWorkflowInstancesResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRerunWorkflowInstancesResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRerunWorkflowInstancesResultResponse
	GetStatusCode() *int32
	SetBody(v *GetRerunWorkflowInstancesResultResponseBody) *GetRerunWorkflowInstancesResultResponse
	GetBody() *GetRerunWorkflowInstancesResultResponseBody
}

type GetRerunWorkflowInstancesResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRerunWorkflowInstancesResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRerunWorkflowInstancesResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRerunWorkflowInstancesResultResponse) GoString() string {
	return s.String()
}

func (s *GetRerunWorkflowInstancesResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRerunWorkflowInstancesResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRerunWorkflowInstancesResultResponse) GetBody() *GetRerunWorkflowInstancesResultResponseBody {
	return s.Body
}

func (s *GetRerunWorkflowInstancesResultResponse) SetHeaders(v map[string]*string) *GetRerunWorkflowInstancesResultResponse {
	s.Headers = v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponse) SetStatusCode(v int32) *GetRerunWorkflowInstancesResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponse) SetBody(v *GetRerunWorkflowInstancesResultResponseBody) *GetRerunWorkflowInstancesResultResponse {
	s.Body = v
	return s
}

func (s *GetRerunWorkflowInstancesResultResponse) Validate() error {
	return dara.Validate(s)
}
