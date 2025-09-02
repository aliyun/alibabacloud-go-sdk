// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTableThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTableThemeResponse
	GetStatusCode() *int32
	SetBody(v *CreateTableThemeResponseBody) *CreateTableThemeResponse
	GetBody() *CreateTableThemeResponseBody
}

type CreateTableThemeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTableThemeResponse) GoString() string {
	return s.String()
}

func (s *CreateTableThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTableThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTableThemeResponse) GetBody() *CreateTableThemeResponseBody {
	return s.Body
}

func (s *CreateTableThemeResponse) SetHeaders(v map[string]*string) *CreateTableThemeResponse {
	s.Headers = v
	return s
}

func (s *CreateTableThemeResponse) SetStatusCode(v int32) *CreateTableThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableThemeResponse) SetBody(v *CreateTableThemeResponseBody) *CreateTableThemeResponse {
	s.Body = v
	return s
}

func (s *CreateTableThemeResponse) Validate() error {
	return dara.Validate(s)
}
