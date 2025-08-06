// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomTemplateAndGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomTemplateAndGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomTemplateAndGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomTemplateAndGroupConsoleResponseBody) *AddCustomTemplateAndGroupConsoleResponse
	GetBody() *AddCustomTemplateAndGroupConsoleResponseBody
}

type AddCustomTemplateAndGroupConsoleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomTemplateAndGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomTemplateAndGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomTemplateAndGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *AddCustomTemplateAndGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomTemplateAndGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomTemplateAndGroupConsoleResponse) GetBody() *AddCustomTemplateAndGroupConsoleResponseBody {
	return s.Body
}

func (s *AddCustomTemplateAndGroupConsoleResponse) SetHeaders(v map[string]*string) *AddCustomTemplateAndGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponse) SetStatusCode(v int32) *AddCustomTemplateAndGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponse) SetBody(v *AddCustomTemplateAndGroupConsoleResponseBody) *AddCustomTemplateAndGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *AddCustomTemplateAndGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
