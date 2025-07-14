// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserGroupMembersResponseBody) *DeleteUserGroupMembersResponse
	GetBody() *DeleteUserGroupMembersResponseBody
}

type DeleteUserGroupMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserGroupMembersResponse) GetBody() *DeleteUserGroupMembersResponseBody {
	return s.Body
}

func (s *DeleteUserGroupMembersResponse) SetHeaders(v map[string]*string) *DeleteUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupMembersResponse) SetStatusCode(v int32) *DeleteUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupMembersResponse) SetBody(v *DeleteUserGroupMembersResponseBody) *DeleteUserGroupMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteUserGroupMembersResponse) Validate() error {
	return dara.Validate(s)
}
