// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePromptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePromptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePromptTemplateResponseBody) *UpdatePromptTemplateResponse
	GetBody() *UpdatePromptTemplateResponseBody
}

type UpdatePromptTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePromptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePromptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdatePromptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePromptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePromptTemplateResponse) GetBody() *UpdatePromptTemplateResponseBody {
	return s.Body
}

func (s *UpdatePromptTemplateResponse) SetHeaders(v map[string]*string) *UpdatePromptTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdatePromptTemplateResponse) SetStatusCode(v int32) *UpdatePromptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePromptTemplateResponse) SetBody(v *UpdatePromptTemplateResponseBody) *UpdatePromptTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdatePromptTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
