// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComponentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComponentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ComponentTemplate) *UpdateComponentTemplateResponse
	GetBody() *ComponentTemplate
}

type UpdateComponentTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComponentTemplate `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComponentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComponentTemplateResponse) GetBody() *ComponentTemplate {
	return s.Body
}

func (s *UpdateComponentTemplateResponse) SetHeaders(v map[string]*string) *UpdateComponentTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentTemplateResponse) SetStatusCode(v int32) *UpdateComponentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComponentTemplateResponse) SetBody(v *ComponentTemplate) *UpdateComponentTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateComponentTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
