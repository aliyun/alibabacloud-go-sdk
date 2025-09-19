// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateRegistryTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateRegistryTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateRegistryTypeResponseBody) *ListPrivateRegistryTypeResponse
	GetBody() *ListPrivateRegistryTypeResponseBody
}

type ListPrivateRegistryTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateRegistryTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateRegistryTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryTypeResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateRegistryTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateRegistryTypeResponse) GetBody() *ListPrivateRegistryTypeResponseBody {
	return s.Body
}

func (s *ListPrivateRegistryTypeResponse) SetHeaders(v map[string]*string) *ListPrivateRegistryTypeResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateRegistryTypeResponse) SetStatusCode(v int32) *ListPrivateRegistryTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateRegistryTypeResponse) SetBody(v *ListPrivateRegistryTypeResponseBody) *ListPrivateRegistryTypeResponse {
	s.Body = v
	return s
}

func (s *ListPrivateRegistryTypeResponse) Validate() error {
	return dara.Validate(s)
}
