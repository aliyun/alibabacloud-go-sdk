// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateDraftVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkflowTemplateDraftVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkflowTemplateDraftVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkflowTemplateDraftVersionResponseBody) *CreateWorkflowTemplateDraftVersionResponse
	GetBody() *CreateWorkflowTemplateDraftVersionResponseBody
}

type CreateWorkflowTemplateDraftVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowTemplateDraftVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowTemplateDraftVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateDraftVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateDraftVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkflowTemplateDraftVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkflowTemplateDraftVersionResponse) GetBody() *CreateWorkflowTemplateDraftVersionResponseBody {
	return s.Body
}

func (s *CreateWorkflowTemplateDraftVersionResponse) SetHeaders(v map[string]*string) *CreateWorkflowTemplateDraftVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponse) SetStatusCode(v int32) *CreateWorkflowTemplateDraftVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponse) SetBody(v *CreateWorkflowTemplateDraftVersionResponseBody) *CreateWorkflowTemplateDraftVersionResponse {
	s.Body = v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
