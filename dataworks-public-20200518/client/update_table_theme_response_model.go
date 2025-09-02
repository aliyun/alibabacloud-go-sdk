// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableThemeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableThemeResponseBody) *UpdateTableThemeResponse
	GetBody() *UpdateTableThemeResponseBody
}

type UpdateTableThemeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableThemeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableThemeResponse) GetBody() *UpdateTableThemeResponseBody {
	return s.Body
}

func (s *UpdateTableThemeResponse) SetHeaders(v map[string]*string) *UpdateTableThemeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableThemeResponse) SetStatusCode(v int32) *UpdateTableThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableThemeResponse) SetBody(v *UpdateTableThemeResponseBody) *UpdateTableThemeResponse {
	s.Body = v
	return s
}

func (s *UpdateTableThemeResponse) Validate() error {
	return dara.Validate(s)
}
