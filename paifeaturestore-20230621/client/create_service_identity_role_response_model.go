// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceIdentityRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceIdentityRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceIdentityRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceIdentityRoleResponseBody) *CreateServiceIdentityRoleResponse
	GetBody() *CreateServiceIdentityRoleResponseBody
}

type CreateServiceIdentityRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceIdentityRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceIdentityRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceIdentityRoleResponse) GetBody() *CreateServiceIdentityRoleResponseBody {
	return s.Body
}

func (s *CreateServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *CreateServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetStatusCode(v int32) *CreateServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetBody(v *CreateServiceIdentityRoleResponseBody) *CreateServiceIdentityRoleResponse {
	s.Body = v
	return s
}

func (s *CreateServiceIdentityRoleResponse) Validate() error {
	return dara.Validate(s)
}
