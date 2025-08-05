// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoleAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoleAclResponse
	GetStatusCode() *int32
	SetBody(v *GetRoleAclResponseBody) *GetRoleAclResponse
	GetBody() *GetRoleAclResponseBody
}

type GetRoleAclResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleAclResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoleAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoleAclResponse) GetBody() *GetRoleAclResponseBody {
	return s.Body
}

func (s *GetRoleAclResponse) SetHeaders(v map[string]*string) *GetRoleAclResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclResponse) SetStatusCode(v int32) *GetRoleAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclResponse) SetBody(v *GetRoleAclResponseBody) *GetRoleAclResponse {
	s.Body = v
	return s
}

func (s *GetRoleAclResponse) Validate() error {
	return dara.Validate(s)
}
