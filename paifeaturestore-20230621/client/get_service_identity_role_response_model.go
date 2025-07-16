// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceIdentityRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceIdentityRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceIdentityRoleResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceIdentityRoleResponseBody) *GetServiceIdentityRoleResponse
	GetBody() *GetServiceIdentityRoleResponseBody
}

type GetServiceIdentityRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceIdentityRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceIdentityRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceIdentityRoleResponse) GetBody() *GetServiceIdentityRoleResponseBody {
	return s.Body
}

func (s *GetServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *GetServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetStatusCode(v int32) *GetServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetBody(v *GetServiceIdentityRoleResponseBody) *GetServiceIdentityRoleResponse {
	s.Body = v
	return s
}

func (s *GetServiceIdentityRoleResponse) Validate() error {
	return dara.Validate(s)
}
