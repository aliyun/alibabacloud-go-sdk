// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPromptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPromptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetPromptTemplateResponseBody) *GetPromptTemplateResponse
	GetBody() *GetPromptTemplateResponseBody
}

type GetPromptTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPromptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPromptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetPromptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPromptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPromptTemplateResponse) GetBody() *GetPromptTemplateResponseBody {
	return s.Body
}

func (s *GetPromptTemplateResponse) SetHeaders(v map[string]*string) *GetPromptTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetPromptTemplateResponse) SetStatusCode(v int32) *GetPromptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptTemplateResponse) SetBody(v *GetPromptTemplateResponseBody) *GetPromptTemplateResponse {
	s.Body = v
	return s
}

func (s *GetPromptTemplateResponse) Validate() error {
	return dara.Validate(s)
}
