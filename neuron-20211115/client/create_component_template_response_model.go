// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComponentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComponentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ComponentTemplate) *CreateComponentTemplateResponse
	GetBody() *ComponentTemplate
}

type CreateComponentTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComponentTemplate `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateComponentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComponentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComponentTemplateResponse) GetBody() *ComponentTemplate {
	return s.Body
}

func (s *CreateComponentTemplateResponse) SetHeaders(v map[string]*string) *CreateComponentTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateComponentTemplateResponse) SetStatusCode(v int32) *CreateComponentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComponentTemplateResponse) SetBody(v *ComponentTemplate) *CreateComponentTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateComponentTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
