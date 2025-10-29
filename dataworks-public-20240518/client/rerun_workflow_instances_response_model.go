// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerunWorkflowInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerunWorkflowInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RerunWorkflowInstancesResponseBody) *RerunWorkflowInstancesResponse
	GetBody() *RerunWorkflowInstancesResponseBody
}

type RerunWorkflowInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunWorkflowInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *RerunWorkflowInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerunWorkflowInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerunWorkflowInstancesResponse) GetBody() *RerunWorkflowInstancesResponseBody {
	return s.Body
}

func (s *RerunWorkflowInstancesResponse) SetHeaders(v map[string]*string) *RerunWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *RerunWorkflowInstancesResponse) SetStatusCode(v int32) *RerunWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunWorkflowInstancesResponse) SetBody(v *RerunWorkflowInstancesResponseBody) *RerunWorkflowInstancesResponse {
	s.Body = v
	return s
}

func (s *RerunWorkflowInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
