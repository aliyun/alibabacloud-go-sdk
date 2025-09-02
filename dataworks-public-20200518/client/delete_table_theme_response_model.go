// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTableThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTableThemeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTableThemeResponseBody) *DeleteTableThemeResponse
	GetBody() *DeleteTableThemeResponseBody
}

type DeleteTableThemeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableThemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTableThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTableThemeResponse) GetBody() *DeleteTableThemeResponseBody {
	return s.Body
}

func (s *DeleteTableThemeResponse) SetHeaders(v map[string]*string) *DeleteTableThemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableThemeResponse) SetStatusCode(v int32) *DeleteTableThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableThemeResponse) SetBody(v *DeleteTableThemeResponseBody) *DeleteTableThemeResponse {
	s.Body = v
	return s
}

func (s *DeleteTableThemeResponse) Validate() error {
	return dara.Validate(s)
}
