// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse
	GetBody() *UpdateTemplateResponseBody
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTemplateResponse) GetBody() *UpdateTemplateResponseBody {
	return s.Body
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateTemplateResponse) Validate() error {
	return dara.Validate(s)
}
