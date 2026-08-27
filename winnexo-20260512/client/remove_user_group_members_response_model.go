// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserGroupMembersResponseBody) *RemoveUserGroupMembersResponse
	GetBody() *RemoveUserGroupMembersResponseBody
}

type RemoveUserGroupMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserGroupMembersResponse) GetBody() *RemoveUserGroupMembersResponseBody {
	return s.Body
}

func (s *RemoveUserGroupMembersResponse) SetHeaders(v map[string]*string) *RemoveUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserGroupMembersResponse) SetStatusCode(v int32) *RemoveUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserGroupMembersResponse) SetBody(v *RemoveUserGroupMembersResponseBody) *RemoveUserGroupMembersResponse {
	s.Body = v
	return s
}

func (s *RemoveUserGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
