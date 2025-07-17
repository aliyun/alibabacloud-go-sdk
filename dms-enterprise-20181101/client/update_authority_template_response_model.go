// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorityTemplateResponseBody) *UpdateAuthorityTemplateResponse
	GetBody() *UpdateAuthorityTemplateResponseBody
}

type UpdateAuthorityTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorityTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorityTemplateResponse) GetBody() *UpdateAuthorityTemplateResponseBody {
	return s.Body
}

func (s *UpdateAuthorityTemplateResponse) SetHeaders(v map[string]*string) *UpdateAuthorityTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorityTemplateResponse) SetStatusCode(v int32) *UpdateAuthorityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorityTemplateResponse) SetBody(v *UpdateAuthorityTemplateResponseBody) *UpdateAuthorityTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorityTemplateResponse) Validate() error {
	return dara.Validate(s)
}
