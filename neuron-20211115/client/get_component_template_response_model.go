// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComponentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComponentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComponentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ComponentTemplate) *GetComponentTemplateResponse
	GetBody() *ComponentTemplate
}

type GetComponentTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComponentTemplate `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComponentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComponentTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetComponentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComponentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComponentTemplateResponse) GetBody() *ComponentTemplate {
	return s.Body
}

func (s *GetComponentTemplateResponse) SetHeaders(v map[string]*string) *GetComponentTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetComponentTemplateResponse) SetStatusCode(v int32) *GetComponentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComponentTemplateResponse) SetBody(v *ComponentTemplate) *GetComponentTemplateResponse {
	s.Body = v
	return s
}

func (s *GetComponentTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
