// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTemplateMCPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTemplateMCPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTemplateMCPResponse
	GetStatusCode() *int32
	SetBody(v *TemplateResult) *StopTemplateMCPResponse
	GetBody() *TemplateResult
}

type StopTemplateMCPResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTemplateMCPResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTemplateMCPResponse) GoString() string {
	return s.String()
}

func (s *StopTemplateMCPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTemplateMCPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTemplateMCPResponse) GetBody() *TemplateResult {
	return s.Body
}

func (s *StopTemplateMCPResponse) SetHeaders(v map[string]*string) *StopTemplateMCPResponse {
	s.Headers = v
	return s
}

func (s *StopTemplateMCPResponse) SetStatusCode(v int32) *StopTemplateMCPResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTemplateMCPResponse) SetBody(v *TemplateResult) *StopTemplateMCPResponse {
	s.Body = v
	return s
}

func (s *StopTemplateMCPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
