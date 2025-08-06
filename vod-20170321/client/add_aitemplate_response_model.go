// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddAITemplateResponseBody) *AddAITemplateResponse
	GetBody() *AddAITemplateResponseBody
}

type AddAITemplateResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAITemplateResponse) GoString() string {
	return s.String()
}

func (s *AddAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAITemplateResponse) GetBody() *AddAITemplateResponseBody {
	return s.Body
}

func (s *AddAITemplateResponse) SetHeaders(v map[string]*string) *AddAITemplateResponse {
	s.Headers = v
	return s
}

func (s *AddAITemplateResponse) SetStatusCode(v int32) *AddAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAITemplateResponse) SetBody(v *AddAITemplateResponseBody) *AddAITemplateResponse {
	s.Body = v
	return s
}

func (s *AddAITemplateResponse) Validate() error {
	return dara.Validate(s)
}
