// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddUserGroupMembersResponseBody) *AddUserGroupMembersResponse
	GetBody() *AddUserGroupMembersResponseBody
}

type AddUserGroupMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserGroupMembersResponse) GetBody() *AddUserGroupMembersResponseBody {
	return s.Body
}

func (s *AddUserGroupMembersResponse) SetHeaders(v map[string]*string) *AddUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *AddUserGroupMembersResponse) SetStatusCode(v int32) *AddUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserGroupMembersResponse) SetBody(v *AddUserGroupMembersResponseBody) *AddUserGroupMembersResponse {
	s.Body = v
	return s
}

func (s *AddUserGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
