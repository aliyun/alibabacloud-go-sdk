// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateRegistryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateRegistryListResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateRegistryListResponseBody) *ListPrivateRegistryListResponse
	GetBody() *ListPrivateRegistryListResponseBody
}

type ListPrivateRegistryListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateRegistryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateRegistryListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryListResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateRegistryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateRegistryListResponse) GetBody() *ListPrivateRegistryListResponseBody {
	return s.Body
}

func (s *ListPrivateRegistryListResponse) SetHeaders(v map[string]*string) *ListPrivateRegistryListResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateRegistryListResponse) SetStatusCode(v int32) *ListPrivateRegistryListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateRegistryListResponse) SetBody(v *ListPrivateRegistryListResponseBody) *ListPrivateRegistryListResponse {
	s.Body = v
	return s
}

func (s *ListPrivateRegistryListResponse) Validate() error {
	return dara.Validate(s)
}
