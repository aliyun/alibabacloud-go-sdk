// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSslCertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSslCertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSslCertsResponse
	GetStatusCode() *int32
	SetBody(v *ListSslCertsResponseBody) *ListSslCertsResponse
	GetBody() *ListSslCertsResponseBody
}

type ListSslCertsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSslCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSslCertsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsResponse) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSslCertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSslCertsResponse) GetBody() *ListSslCertsResponseBody {
	return s.Body
}

func (s *ListSslCertsResponse) SetHeaders(v map[string]*string) *ListSslCertsResponse {
	s.Headers = v
	return s
}

func (s *ListSslCertsResponse) SetStatusCode(v int32) *ListSslCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSslCertsResponse) SetBody(v *ListSslCertsResponseBody) *ListSslCertsResponse {
	s.Body = v
	return s
}

func (s *ListSslCertsResponse) Validate() error {
	return dara.Validate(s)
}
