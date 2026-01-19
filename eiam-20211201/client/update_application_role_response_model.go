// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationRoleResponseBody) *UpdateApplicationRoleResponse
	GetBody() *UpdateApplicationRoleResponseBody
}

type UpdateApplicationRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationRoleResponse) GetBody() *UpdateApplicationRoleResponseBody {
	return s.Body
}

func (s *UpdateApplicationRoleResponse) SetHeaders(v map[string]*string) *UpdateApplicationRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationRoleResponse) SetStatusCode(v int32) *UpdateApplicationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationRoleResponse) SetBody(v *UpdateApplicationRoleResponseBody) *UpdateApplicationRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
