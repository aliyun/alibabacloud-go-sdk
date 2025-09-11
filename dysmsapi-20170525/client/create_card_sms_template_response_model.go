// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCardSmsTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCardSmsTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCardSmsTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCardSmsTemplateResponseBody) *CreateCardSmsTemplateResponse
	GetBody() *CreateCardSmsTemplateResponseBody
}

type CreateCardSmsTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCardSmsTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCardSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCardSmsTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCardSmsTemplateResponse) GetBody() *CreateCardSmsTemplateResponseBody {
	return s.Body
}

func (s *CreateCardSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateCardSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCardSmsTemplateResponse) SetStatusCode(v int32) *CreateCardSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCardSmsTemplateResponse) SetBody(v *CreateCardSmsTemplateResponseBody) *CreateCardSmsTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCardSmsTemplateResponse) Validate() error {
	return dara.Validate(s)
}
