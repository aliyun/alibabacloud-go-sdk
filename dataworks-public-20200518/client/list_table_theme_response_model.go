// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableThemeResponse
	GetStatusCode() *int32
	SetBody(v *ListTableThemeResponseBody) *ListTableThemeResponse
	GetBody() *ListTableThemeResponseBody
}

type ListTableThemeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableThemeResponse) GoString() string {
	return s.String()
}

func (s *ListTableThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableThemeResponse) GetBody() *ListTableThemeResponseBody {
	return s.Body
}

func (s *ListTableThemeResponse) SetHeaders(v map[string]*string) *ListTableThemeResponse {
	s.Headers = v
	return s
}

func (s *ListTableThemeResponse) SetStatusCode(v int32) *ListTableThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableThemeResponse) SetBody(v *ListTableThemeResponseBody) *ListTableThemeResponse {
	s.Body = v
	return s
}

func (s *ListTableThemeResponse) Validate() error {
	return dara.Validate(s)
}
