// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderDepartmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityProviderDepartmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityProviderDepartmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentityProviderDepartmentsResponseBody) *ListIdentityProviderDepartmentsResponse
	GetBody() *ListIdentityProviderDepartmentsResponseBody
}

type ListIdentityProviderDepartmentsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityProviderDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityProviderDepartmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityProviderDepartmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityProviderDepartmentsResponse) GetBody() *ListIdentityProviderDepartmentsResponseBody {
	return s.Body
}

func (s *ListIdentityProviderDepartmentsResponse) SetHeaders(v map[string]*string) *ListIdentityProviderDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityProviderDepartmentsResponse) SetStatusCode(v int32) *ListIdentityProviderDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponse) SetBody(v *ListIdentityProviderDepartmentsResponseBody) *ListIdentityProviderDepartmentsResponse {
	s.Body = v
	return s
}

func (s *ListIdentityProviderDepartmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
