// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsForUserResponseBody) *ListUserGroupsForUserResponse
	GetBody() *ListUserGroupsForUserResponseBody
}

type ListUserGroupsForUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsForUserResponse) GetBody() *ListUserGroupsForUserResponseBody {
	return s.Body
}

func (s *ListUserGroupsForUserResponse) SetHeaders(v map[string]*string) *ListUserGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsForUserResponse) SetStatusCode(v int32) *ListUserGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsForUserResponse) SetBody(v *ListUserGroupsForUserResponseBody) *ListUserGroupsForUserResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
