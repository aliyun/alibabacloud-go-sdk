// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceLinkedRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse
	GetBody() *CreateServiceLinkedRoleResponseBody
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceLinkedRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceLinkedRoleResponse) GetBody() *CreateServiceLinkedRoleResponseBody {
	return s.Body
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) Validate() error {
	return dara.Validate(s)
}
