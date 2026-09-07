// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantAppResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantAppResponseBody) *ListTenantAppResponse
	GetBody() *ListTenantAppResponseBody
}

type ListTenantAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAppResponse) GoString() string {
	return s.String()
}

func (s *ListTenantAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantAppResponse) GetBody() *ListTenantAppResponseBody {
	return s.Body
}

func (s *ListTenantAppResponse) SetHeaders(v map[string]*string) *ListTenantAppResponse {
	s.Headers = v
	return s
}

func (s *ListTenantAppResponse) SetStatusCode(v int32) *ListTenantAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantAppResponse) SetBody(v *ListTenantAppResponseBody) *ListTenantAppResponse {
	s.Body = v
	return s
}

func (s *ListTenantAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
