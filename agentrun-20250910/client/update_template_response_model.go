// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTemplateResponse
	GetStatusCode() *int32
	SetBody(v *TemplateResult) *UpdateTemplateResponse
	GetBody() *TemplateResult
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTemplateResponse) GetBody() *TemplateResult {
	return s.Body
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *TemplateResult) *UpdateTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
