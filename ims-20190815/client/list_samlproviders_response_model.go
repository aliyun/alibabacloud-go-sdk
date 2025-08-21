// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSAMLProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSAMLProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSAMLProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListSAMLProvidersResponseBody) *ListSAMLProvidersResponse
	GetBody() *ListSAMLProvidersResponseBody
}

type ListSAMLProvidersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSAMLProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSAMLProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSAMLProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSAMLProvidersResponse) GetBody() *ListSAMLProvidersResponseBody {
	return s.Body
}

func (s *ListSAMLProvidersResponse) SetHeaders(v map[string]*string) *ListSAMLProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListSAMLProvidersResponse) SetStatusCode(v int32) *ListSAMLProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSAMLProvidersResponse) SetBody(v *ListSAMLProvidersResponseBody) *ListSAMLProvidersResponse {
	s.Body = v
	return s
}

func (s *ListSAMLProvidersResponse) Validate() error {
	return dara.Validate(s)
}
