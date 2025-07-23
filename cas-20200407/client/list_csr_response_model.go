// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCsrResponse
	GetStatusCode() *int32
	SetBody(v *ListCsrResponseBody) *ListCsrResponse
	GetBody() *ListCsrResponseBody
}

type ListCsrResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCsrResponse) GoString() string {
	return s.String()
}

func (s *ListCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCsrResponse) GetBody() *ListCsrResponseBody {
	return s.Body
}

func (s *ListCsrResponse) SetHeaders(v map[string]*string) *ListCsrResponse {
	s.Headers = v
	return s
}

func (s *ListCsrResponse) SetStatusCode(v int32) *ListCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCsrResponse) SetBody(v *ListCsrResponseBody) *ListCsrResponse {
	s.Body = v
	return s
}

func (s *ListCsrResponse) Validate() error {
	return dara.Validate(s)
}
