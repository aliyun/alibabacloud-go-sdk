// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRbacRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRbacRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRbacRoleResponseBody) *DeleteRbacRoleResponse
	GetBody() *DeleteRbacRoleResponseBody
}

type DeleteRbacRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRbacRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRbacRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRbacRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRbacRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRbacRoleResponse) GetBody() *DeleteRbacRoleResponseBody {
	return s.Body
}

func (s *DeleteRbacRoleResponse) SetHeaders(v map[string]*string) *DeleteRbacRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRbacRoleResponse) SetStatusCode(v int32) *DeleteRbacRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRbacRoleResponse) SetBody(v *DeleteRbacRoleResponseBody) *DeleteRbacRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteRbacRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
