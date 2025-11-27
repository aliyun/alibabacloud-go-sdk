// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitTenantAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitTenantAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitTenantAliasResponse
	GetStatusCode() *int32
	SetBody(v *InitTenantAliasResponseBody) *InitTenantAliasResponse
	GetBody() *InitTenantAliasResponseBody
}

type InitTenantAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitTenantAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitTenantAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s InitTenantAliasResponse) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitTenantAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitTenantAliasResponse) GetBody() *InitTenantAliasResponseBody {
	return s.Body
}

func (s *InitTenantAliasResponse) SetHeaders(v map[string]*string) *InitTenantAliasResponse {
	s.Headers = v
	return s
}

func (s *InitTenantAliasResponse) SetStatusCode(v int32) *InitTenantAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *InitTenantAliasResponse) SetBody(v *InitTenantAliasResponseBody) *InitTenantAliasResponse {
	s.Body = v
	return s
}

func (s *InitTenantAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
