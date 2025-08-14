// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNADBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDNADBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDNADBResponse
	GetStatusCode() *int32
	SetBody(v *ListDNADBResponseBody) *ListDNADBResponse
	GetBody() *ListDNADBResponseBody
}

type ListDNADBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDNADBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDNADBResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBResponse) GoString() string {
	return s.String()
}

func (s *ListDNADBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDNADBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDNADBResponse) GetBody() *ListDNADBResponseBody {
	return s.Body
}

func (s *ListDNADBResponse) SetHeaders(v map[string]*string) *ListDNADBResponse {
	s.Headers = v
	return s
}

func (s *ListDNADBResponse) SetStatusCode(v int32) *ListDNADBResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDNADBResponse) SetBody(v *ListDNADBResponseBody) *ListDNADBResponse {
	s.Body = v
	return s
}

func (s *ListDNADBResponse) Validate() error {
	return dara.Validate(s)
}
