// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserGroupMemberResponseBody) *RemoveUserGroupMemberResponse
	GetBody() *RemoveUserGroupMemberResponseBody
}

type RemoveUserGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserGroupMemberResponse) GetBody() *RemoveUserGroupMemberResponseBody {
	return s.Body
}

func (s *RemoveUserGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserGroupMemberResponse) SetStatusCode(v int32) *RemoveUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserGroupMemberResponse) SetBody(v *RemoveUserGroupMemberResponseBody) *RemoveUserGroupMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveUserGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
