// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomRoleResponseBody) *DeleteCustomRoleResponse
	GetBody() *DeleteCustomRoleResponseBody
}

type DeleteCustomRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomRoleResponse) GetBody() *DeleteCustomRoleResponseBody {
	return s.Body
}

func (s *DeleteCustomRoleResponse) SetHeaders(v map[string]*string) *DeleteCustomRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoleResponse) SetStatusCode(v int32) *DeleteCustomRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoleResponse) SetBody(v *DeleteCustomRoleResponseBody) *DeleteCustomRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomRoleResponse) Validate() error {
	return dara.Validate(s)
}
