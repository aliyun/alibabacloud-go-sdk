// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAITemplateResponseBody) *UpdateAITemplateResponse
	GetBody() *UpdateAITemplateResponseBody
}

type UpdateAITemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAITemplateResponse) GetBody() *UpdateAITemplateResponseBody {
	return s.Body
}

func (s *UpdateAITemplateResponse) SetHeaders(v map[string]*string) *UpdateAITemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAITemplateResponse) SetStatusCode(v int32) *UpdateAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAITemplateResponse) SetBody(v *UpdateAITemplateResponseBody) *UpdateAITemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateAITemplateResponse) Validate() error {
	return dara.Validate(s)
}
