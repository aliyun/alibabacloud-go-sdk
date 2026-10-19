// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWorkflowTemplateDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateWorkflowTemplateDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateWorkflowTemplateDraftResponse
	GetStatusCode() *int32
	SetBody(v *GenerateWorkflowTemplateDraftResponseBody) *GenerateWorkflowTemplateDraftResponse
	GetBody() *GenerateWorkflowTemplateDraftResponseBody
}

type GenerateWorkflowTemplateDraftResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateWorkflowTemplateDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateWorkflowTemplateDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftResponse) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateWorkflowTemplateDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateWorkflowTemplateDraftResponse) GetBody() *GenerateWorkflowTemplateDraftResponseBody {
	return s.Body
}

func (s *GenerateWorkflowTemplateDraftResponse) SetHeaders(v map[string]*string) *GenerateWorkflowTemplateDraftResponse {
	s.Headers = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponse) SetStatusCode(v int32) *GenerateWorkflowTemplateDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponse) SetBody(v *GenerateWorkflowTemplateDraftResponseBody) *GenerateWorkflowTemplateDraftResponse {
	s.Body = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
