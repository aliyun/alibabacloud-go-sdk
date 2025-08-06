// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomTemplateConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomTemplateConsoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomTemplateConsoleResponseBody) *DeleteCustomTemplateConsoleResponse
	GetBody() *DeleteCustomTemplateConsoleResponseBody
}

type DeleteCustomTemplateConsoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTemplateConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTemplateConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateConsoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomTemplateConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomTemplateConsoleResponse) GetBody() *DeleteCustomTemplateConsoleResponseBody {
	return s.Body
}

func (s *DeleteCustomTemplateConsoleResponse) SetHeaders(v map[string]*string) *DeleteCustomTemplateConsoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTemplateConsoleResponse) SetStatusCode(v int32) *DeleteCustomTemplateConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTemplateConsoleResponse) SetBody(v *DeleteCustomTemplateConsoleResponseBody) *DeleteCustomTemplateConsoleResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomTemplateConsoleResponse) Validate() error {
	return dara.Validate(s)
}
