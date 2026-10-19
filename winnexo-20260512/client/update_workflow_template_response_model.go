// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkflowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkflowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkflowTemplateResponseBody) *UpdateWorkflowTemplateResponse
	GetBody() *UpdateWorkflowTemplateResponseBody
}

type UpdateWorkflowTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkflowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkflowTemplateResponse) GetBody() *UpdateWorkflowTemplateResponseBody {
	return s.Body
}

func (s *UpdateWorkflowTemplateResponse) SetHeaders(v map[string]*string) *UpdateWorkflowTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowTemplateResponse) SetStatusCode(v int32) *UpdateWorkflowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowTemplateResponse) SetBody(v *UpdateWorkflowTemplateResponseBody) *UpdateWorkflowTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkflowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
