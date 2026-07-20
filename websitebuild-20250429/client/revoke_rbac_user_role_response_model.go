// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRbacUserRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeRbacUserRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeRbacUserRoleResponse
	GetStatusCode() *int32
	SetBody(v *RevokeRbacUserRoleResponseBody) *RevokeRbacUserRoleResponse
	GetBody() *RevokeRbacUserRoleResponseBody
}

type RevokeRbacUserRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeRbacUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeRbacUserRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeRbacUserRoleResponse) GoString() string {
	return s.String()
}

func (s *RevokeRbacUserRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeRbacUserRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeRbacUserRoleResponse) GetBody() *RevokeRbacUserRoleResponseBody {
	return s.Body
}

func (s *RevokeRbacUserRoleResponse) SetHeaders(v map[string]*string) *RevokeRbacUserRoleResponse {
	s.Headers = v
	return s
}

func (s *RevokeRbacUserRoleResponse) SetStatusCode(v int32) *RevokeRbacUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRbacUserRoleResponse) SetBody(v *RevokeRbacUserRoleResponseBody) *RevokeRbacUserRoleResponse {
	s.Body = v
	return s
}

func (s *RevokeRbacUserRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
