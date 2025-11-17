// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddUserGroupMemberResponseBody) *AddUserGroupMemberResponse
	GetBody() *AddUserGroupMemberResponseBody
}

type AddUserGroupMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserGroupMemberResponse) GetBody() *AddUserGroupMemberResponseBody {
	return s.Body
}

func (s *AddUserGroupMemberResponse) SetHeaders(v map[string]*string) *AddUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddUserGroupMemberResponse) SetStatusCode(v int32) *AddUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserGroupMemberResponse) SetBody(v *AddUserGroupMemberResponseBody) *AddUserGroupMemberResponse {
	s.Body = v
	return s
}

func (s *AddUserGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
