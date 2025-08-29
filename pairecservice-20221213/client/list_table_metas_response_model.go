// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTableMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTableMetasResponse
	GetStatusCode() *int32
	SetBody(v *ListTableMetasResponseBody) *ListTableMetasResponse
	GetBody() *ListTableMetasResponseBody
}

type ListTableMetasResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTableMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTableMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTableMetasResponse) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTableMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTableMetasResponse) GetBody() *ListTableMetasResponseBody {
	return s.Body
}

func (s *ListTableMetasResponse) SetHeaders(v map[string]*string) *ListTableMetasResponse {
	s.Headers = v
	return s
}

func (s *ListTableMetasResponse) SetStatusCode(v int32) *ListTableMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableMetasResponse) SetBody(v *ListTableMetasResponseBody) *ListTableMetasResponse {
	s.Body = v
	return s
}

func (s *ListTableMetasResponse) Validate() error {
	return dara.Validate(s)
}
