// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApiAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApiAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApiAuthoritiesResponseBody) *ListDataServiceApiAuthoritiesResponse
	GetBody() *ListDataServiceApiAuthoritiesResponseBody
}

type ListDataServiceApiAuthoritiesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApiAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApiAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApiAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApiAuthoritiesResponse) GetBody() *ListDataServiceApiAuthoritiesResponseBody {
	return s.Body
}

func (s *ListDataServiceApiAuthoritiesResponse) SetHeaders(v map[string]*string) *ListDataServiceApiAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponse) SetStatusCode(v int32) *ListDataServiceApiAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponse) SetBody(v *ListDataServiceApiAuthoritiesResponseBody) *ListDataServiceApiAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponse) Validate() error {
	return dara.Validate(s)
}
