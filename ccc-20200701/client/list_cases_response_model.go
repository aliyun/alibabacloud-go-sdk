// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCasesResponse
	GetStatusCode() *int32
	SetBody(v *ListCasesResponseBody) *ListCasesResponse
	GetBody() *ListCasesResponseBody
}

type ListCasesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponse) GoString() string {
	return s.String()
}

func (s *ListCasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCasesResponse) GetBody() *ListCasesResponseBody {
	return s.Body
}

func (s *ListCasesResponse) SetHeaders(v map[string]*string) *ListCasesResponse {
	s.Headers = v
	return s
}

func (s *ListCasesResponse) SetStatusCode(v int32) *ListCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCasesResponse) SetBody(v *ListCasesResponseBody) *ListCasesResponse {
	s.Body = v
	return s
}

func (s *ListCasesResponse) Validate() error {
	return dara.Validate(s)
}
