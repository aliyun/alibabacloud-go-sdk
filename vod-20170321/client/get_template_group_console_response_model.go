// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateGroupConsoleResponseBody) *GetTemplateGroupConsoleResponse
	GetBody() *GetTemplateGroupConsoleResponseBody
}

type GetTemplateGroupConsoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateGroupConsoleResponse) GetBody() *GetTemplateGroupConsoleResponseBody {
	return s.Body
}

func (s *GetTemplateGroupConsoleResponse) SetHeaders(v map[string]*string) *GetTemplateGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateGroupConsoleResponse) SetStatusCode(v int32) *GetTemplateGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateGroupConsoleResponse) SetBody(v *GetTemplateGroupConsoleResponseBody) *GetTemplateGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *GetTemplateGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
