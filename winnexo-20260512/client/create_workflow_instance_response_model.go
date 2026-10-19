// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowInstanceResponseBody) *CreateWorkflowInstanceResponse
	GetBody() *CreateWorkflowInstanceResponseBody
}

type CreateWorkflowInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowInstanceResponse) GetBody() *CreateWorkflowInstanceResponseBody {
	return s.Body
}

func (s *CreateWorkflowInstanceResponse) SetHeaders(v map[string]*string) *CreateWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowInstanceResponse) SetStatusCode(v int32) *CreateWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowInstanceResponse) SetBody(v *CreateWorkflowInstanceResponseBody) *CreateWorkflowInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
