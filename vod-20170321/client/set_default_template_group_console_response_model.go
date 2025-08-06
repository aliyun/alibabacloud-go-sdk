// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTemplateGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultTemplateGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultTemplateGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultTemplateGroupConsoleResponseBody) *SetDefaultTemplateGroupConsoleResponse
	GetBody() *SetDefaultTemplateGroupConsoleResponseBody
}

type SetDefaultTemplateGroupConsoleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultTemplateGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultTemplateGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTemplateGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultTemplateGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultTemplateGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultTemplateGroupConsoleResponse) GetBody() *SetDefaultTemplateGroupConsoleResponseBody {
	return s.Body
}

func (s *SetDefaultTemplateGroupConsoleResponse) SetHeaders(v map[string]*string) *SetDefaultTemplateGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultTemplateGroupConsoleResponse) SetStatusCode(v int32) *SetDefaultTemplateGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleResponse) SetBody(v *SetDefaultTemplateGroupConsoleResponseBody) *SetDefaultTemplateGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *SetDefaultTemplateGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
