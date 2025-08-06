// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateGroupConsoleResponseBody) *ListTemplateGroupConsoleResponse
	GetBody() *ListTemplateGroupConsoleResponseBody
}

type ListTemplateGroupConsoleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateGroupConsoleResponse) GetBody() *ListTemplateGroupConsoleResponseBody {
	return s.Body
}

func (s *ListTemplateGroupConsoleResponse) SetHeaders(v map[string]*string) *ListTemplateGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateGroupConsoleResponse) SetStatusCode(v int32) *ListTemplateGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateGroupConsoleResponse) SetBody(v *ListTemplateGroupConsoleResponseBody) *ListTemplateGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *ListTemplateGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
