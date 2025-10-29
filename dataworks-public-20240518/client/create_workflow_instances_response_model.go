// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowInstancesResponseBody) *CreateWorkflowInstancesResponse
	GetBody() *CreateWorkflowInstancesResponseBody
}

type CreateWorkflowInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowInstancesResponse) GetBody() *CreateWorkflowInstancesResponseBody {
	return s.Body
}

func (s *CreateWorkflowInstancesResponse) SetHeaders(v map[string]*string) *CreateWorkflowInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowInstancesResponse) SetStatusCode(v int32) *CreateWorkflowInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowInstancesResponse) SetBody(v *CreateWorkflowInstancesResponseBody) *CreateWorkflowInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
