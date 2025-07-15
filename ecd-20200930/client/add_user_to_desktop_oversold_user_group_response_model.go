// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopOversoldUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToDesktopOversoldUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToDesktopOversoldUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToDesktopOversoldUserGroupResponseBody) *AddUserToDesktopOversoldUserGroupResponse
	GetBody() *AddUserToDesktopOversoldUserGroupResponseBody
}

type AddUserToDesktopOversoldUserGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToDesktopOversoldUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToDesktopOversoldUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopOversoldUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopOversoldUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToDesktopOversoldUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToDesktopOversoldUserGroupResponse) GetBody() *AddUserToDesktopOversoldUserGroupResponseBody {
	return s.Body
}

func (s *AddUserToDesktopOversoldUserGroupResponse) SetHeaders(v map[string]*string) *AddUserToDesktopOversoldUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupResponse) SetStatusCode(v int32) *AddUserToDesktopOversoldUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupResponse) SetBody(v *AddUserToDesktopOversoldUserGroupResponseBody) *AddUserToDesktopOversoldUserGroupResponse {
	s.Body = v
	return s
}

func (s *AddUserToDesktopOversoldUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
