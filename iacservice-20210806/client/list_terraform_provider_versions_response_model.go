// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerraformProviderVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTerraformProviderVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTerraformProviderVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTerraformProviderVersionsResponseBody) *ListTerraformProviderVersionsResponse
	GetBody() *ListTerraformProviderVersionsResponseBody
}

type ListTerraformProviderVersionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerraformProviderVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerraformProviderVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTerraformProviderVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTerraformProviderVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTerraformProviderVersionsResponse) GetBody() *ListTerraformProviderVersionsResponseBody {
	return s.Body
}

func (s *ListTerraformProviderVersionsResponse) SetHeaders(v map[string]*string) *ListTerraformProviderVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTerraformProviderVersionsResponse) SetStatusCode(v int32) *ListTerraformProviderVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerraformProviderVersionsResponse) SetBody(v *ListTerraformProviderVersionsResponseBody) *ListTerraformProviderVersionsResponse {
	s.Body = v
	return s
}

func (s *ListTerraformProviderVersionsResponse) Validate() error {
	return dara.Validate(s)
}
