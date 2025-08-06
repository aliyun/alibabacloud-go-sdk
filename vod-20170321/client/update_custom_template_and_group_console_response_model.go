// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTemplateAndGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomTemplateAndGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomTemplateAndGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomTemplateAndGroupConsoleResponseBody) *UpdateCustomTemplateAndGroupConsoleResponse
	GetBody() *UpdateCustomTemplateAndGroupConsoleResponseBody
}

type UpdateCustomTemplateAndGroupConsoleResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomTemplateAndGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomTemplateAndGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTemplateAndGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) GetBody() *UpdateCustomTemplateAndGroupConsoleResponseBody {
	return s.Body
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) SetHeaders(v map[string]*string) *UpdateCustomTemplateAndGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) SetStatusCode(v int32) *UpdateCustomTemplateAndGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) SetBody(v *UpdateCustomTemplateAndGroupConsoleResponseBody) *UpdateCustomTemplateAndGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomTemplateAndGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
