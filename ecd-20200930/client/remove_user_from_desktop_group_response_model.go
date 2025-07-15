// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserFromDesktopGroupResponseBody) *RemoveUserFromDesktopGroupResponse
	GetBody() *RemoveUserFromDesktopGroupResponseBody
}

type RemoveUserFromDesktopGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromDesktopGroupResponse) GetBody() *RemoveUserFromDesktopGroupResponseBody {
	return s.Body
}

func (s *RemoveUserFromDesktopGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromDesktopGroupResponse) SetStatusCode(v int32) *RemoveUserFromDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromDesktopGroupResponse) SetBody(v *RemoveUserFromDesktopGroupResponseBody) *RemoveUserFromDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveUserFromDesktopGroupResponse) Validate() error {
	return dara.Validate(s)
}
