// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopOversoldUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromDesktopOversoldUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromDesktopOversoldUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserFromDesktopOversoldUserGroupResponseBody) *RemoveUserFromDesktopOversoldUserGroupResponse
	GetBody() *RemoveUserFromDesktopOversoldUserGroupResponseBody
}

type RemoveUserFromDesktopOversoldUserGroupResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromDesktopOversoldUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromDesktopOversoldUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopOversoldUserGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) GetBody() *RemoveUserFromDesktopOversoldUserGroupResponseBody {
	return s.Body
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromDesktopOversoldUserGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) SetStatusCode(v int32) *RemoveUserFromDesktopOversoldUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) SetBody(v *RemoveUserFromDesktopOversoldUserGroupResponseBody) *RemoveUserFromDesktopOversoldUserGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveUserFromDesktopOversoldUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
