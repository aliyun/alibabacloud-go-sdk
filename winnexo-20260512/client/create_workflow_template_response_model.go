// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowTemplateResponseBody) *CreateWorkflowTemplateResponse
	GetBody() *CreateWorkflowTemplateResponseBody
}

type CreateWorkflowTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowTemplateResponse) GetBody() *CreateWorkflowTemplateResponseBody {
	return s.Body
}

func (s *CreateWorkflowTemplateResponse) SetHeaders(v map[string]*string) *CreateWorkflowTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowTemplateResponse) SetStatusCode(v int32) *CreateWorkflowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowTemplateResponse) SetBody(v *CreateWorkflowTemplateResponseBody) *CreateWorkflowTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
