// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowResponseBody) *CreateWorkflowResponse
	GetBody() *CreateWorkflowResponseBody
}

type CreateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowResponse) GetBody() *CreateWorkflowResponseBody {
	return s.Body
}

func (s *CreateWorkflowResponse) SetHeaders(v map[string]*string) *CreateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowResponse) SetStatusCode(v int32) *CreateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowResponse) SetBody(v *CreateWorkflowResponseBody) *CreateWorkflowResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
