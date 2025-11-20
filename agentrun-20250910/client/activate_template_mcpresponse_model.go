// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateTemplateMCPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateTemplateMCPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateTemplateMCPResponse
	GetStatusCode() *int32
	SetBody(v *TemplateResult) *ActivateTemplateMCPResponse
	GetBody() *TemplateResult
}

type ActivateTemplateMCPResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateTemplateMCPResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateTemplateMCPResponse) GoString() string {
	return s.String()
}

func (s *ActivateTemplateMCPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateTemplateMCPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateTemplateMCPResponse) GetBody() *TemplateResult {
	return s.Body
}

func (s *ActivateTemplateMCPResponse) SetHeaders(v map[string]*string) *ActivateTemplateMCPResponse {
	s.Headers = v
	return s
}

func (s *ActivateTemplateMCPResponse) SetStatusCode(v int32) *ActivateTemplateMCPResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateTemplateMCPResponse) SetBody(v *TemplateResult) *ActivateTemplateMCPResponse {
	s.Body = v
	return s
}

func (s *ActivateTemplateMCPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
