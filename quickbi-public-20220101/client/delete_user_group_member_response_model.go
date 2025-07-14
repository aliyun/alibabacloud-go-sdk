// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserGroupMemberResponseBody) *DeleteUserGroupMemberResponse
	GetBody() *DeleteUserGroupMemberResponseBody
}

type DeleteUserGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserGroupMemberResponse) GetBody() *DeleteUserGroupMemberResponseBody {
	return s.Body
}

func (s *DeleteUserGroupMemberResponse) SetHeaders(v map[string]*string) *DeleteUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupMemberResponse) SetStatusCode(v int32) *DeleteUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupMemberResponse) SetBody(v *DeleteUserGroupMemberResponseBody) *DeleteUserGroupMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteUserGroupMemberResponse) Validate() error {
	return dara.Validate(s)
}
