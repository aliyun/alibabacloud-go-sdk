// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantDirectoryResponseBody) *ListTenantDirectoryResponse
	GetBody() *ListTenantDirectoryResponseBody
}

type ListTenantDirectoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListTenantDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantDirectoryResponse) GetBody() *ListTenantDirectoryResponseBody {
	return s.Body
}

func (s *ListTenantDirectoryResponse) SetHeaders(v map[string]*string) *ListTenantDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListTenantDirectoryResponse) SetStatusCode(v int32) *ListTenantDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantDirectoryResponse) SetBody(v *ListTenantDirectoryResponseBody) *ListTenantDirectoryResponse {
	s.Body = v
	return s
}

func (s *ListTenantDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
