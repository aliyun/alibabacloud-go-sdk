// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse
	GetBody() *CreateTemplateResponseBody
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTemplateResponse) GetBody() *CreateTemplateResponseBody {
	return s.Body
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateTemplateResponse) Validate() error {
	return dara.Validate(s)
}
