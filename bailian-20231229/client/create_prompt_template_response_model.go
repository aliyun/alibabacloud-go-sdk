// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePromptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePromptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreatePromptTemplateResponseBody) *CreatePromptTemplateResponse
	GetBody() *CreatePromptTemplateResponseBody
}

type CreatePromptTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePromptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePromptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatePromptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePromptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePromptTemplateResponse) GetBody() *CreatePromptTemplateResponseBody {
	return s.Body
}

func (s *CreatePromptTemplateResponse) SetHeaders(v map[string]*string) *CreatePromptTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatePromptTemplateResponse) SetStatusCode(v int32) *CreatePromptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePromptTemplateResponse) SetBody(v *CreatePromptTemplateResponseBody) *CreatePromptTemplateResponse {
	s.Body = v
	return s
}

func (s *CreatePromptTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
