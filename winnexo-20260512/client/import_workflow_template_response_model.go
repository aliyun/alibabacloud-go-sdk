// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportWorkflowTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportWorkflowTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ImportWorkflowTemplateResponseBody) *ImportWorkflowTemplateResponse
	GetBody() *ImportWorkflowTemplateResponseBody
}

type ImportWorkflowTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportWorkflowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportWorkflowTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateResponse) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportWorkflowTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportWorkflowTemplateResponse) GetBody() *ImportWorkflowTemplateResponseBody {
	return s.Body
}

func (s *ImportWorkflowTemplateResponse) SetHeaders(v map[string]*string) *ImportWorkflowTemplateResponse {
	s.Headers = v
	return s
}

func (s *ImportWorkflowTemplateResponse) SetStatusCode(v int32) *ImportWorkflowTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportWorkflowTemplateResponse) SetBody(v *ImportWorkflowTemplateResponseBody) *ImportWorkflowTemplateResponse {
	s.Body = v
	return s
}

func (s *ImportWorkflowTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
