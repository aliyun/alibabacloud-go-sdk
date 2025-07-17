// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableColumnsResponse
	GetStatusCode() *int32
	SetBody(v *ListTableColumnsResponseBody) *ListTableColumnsResponse
	GetBody() *ListTableColumnsResponseBody
}

type ListTableColumnsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListTableColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableColumnsResponse) GetBody() *ListTableColumnsResponseBody {
	return s.Body
}

func (s *ListTableColumnsResponse) SetHeaders(v map[string]*string) *ListTableColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListTableColumnsResponse) SetStatusCode(v int32) *ListTableColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableColumnsResponse) SetBody(v *ListTableColumnsResponseBody) *ListTableColumnsResponse {
	s.Body = v
	return s
}

func (s *ListTableColumnsResponse) Validate() error {
	return dara.Validate(s)
}
