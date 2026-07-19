// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRbacRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRbacRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRbacRoleResponseBody) *CreateRbacRoleResponse
	GetBody() *CreateRbacRoleResponseBody
}

type CreateRbacRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRbacRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRbacRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRbacRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRbacRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRbacRoleResponse) GetBody() *CreateRbacRoleResponseBody {
	return s.Body
}

func (s *CreateRbacRoleResponse) SetHeaders(v map[string]*string) *CreateRbacRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRbacRoleResponse) SetStatusCode(v int32) *CreateRbacRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRbacRoleResponse) SetBody(v *CreateRbacRoleResponseBody) *CreateRbacRoleResponse {
	s.Body = v
	return s
}

func (s *CreateRbacRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
