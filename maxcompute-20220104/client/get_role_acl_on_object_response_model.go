// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleAclOnObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoleAclOnObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoleAclOnObjectResponse
	GetStatusCode() *int32
	SetBody(v *GetRoleAclOnObjectResponseBody) *GetRoleAclOnObjectResponse
	GetBody() *GetRoleAclOnObjectResponseBody
}

type GetRoleAclOnObjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleAclOnObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleAclOnObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclOnObjectResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoleAclOnObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoleAclOnObjectResponse) GetBody() *GetRoleAclOnObjectResponseBody {
	return s.Body
}

func (s *GetRoleAclOnObjectResponse) SetHeaders(v map[string]*string) *GetRoleAclOnObjectResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetStatusCode(v int32) *GetRoleAclOnObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetBody(v *GetRoleAclOnObjectResponseBody) *GetRoleAclOnObjectResponse {
	s.Body = v
	return s
}

func (s *GetRoleAclOnObjectResponse) Validate() error {
	return dara.Validate(s)
}
