// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateGroupConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateGroupConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateGroupConsoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateGroupConsoleResponseBody) *DeleteTemplateGroupConsoleResponse
	GetBody() *DeleteTemplateGroupConsoleResponseBody
}

type DeleteTemplateGroupConsoleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateGroupConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateGroupConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateGroupConsoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateGroupConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateGroupConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateGroupConsoleResponse) GetBody() *DeleteTemplateGroupConsoleResponseBody {
	return s.Body
}

func (s *DeleteTemplateGroupConsoleResponse) SetHeaders(v map[string]*string) *DeleteTemplateGroupConsoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateGroupConsoleResponse) SetStatusCode(v int32) *DeleteTemplateGroupConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateGroupConsoleResponse) SetBody(v *DeleteTemplateGroupConsoleResponseBody) *DeleteTemplateGroupConsoleResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateGroupConsoleResponse) Validate() error {
	return dara.Validate(s)
}
