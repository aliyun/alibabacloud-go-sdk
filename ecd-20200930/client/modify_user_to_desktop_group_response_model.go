// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserToDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserToDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserToDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserToDesktopGroupResponseBody) *ModifyUserToDesktopGroupResponse
	GetBody() *ModifyUserToDesktopGroupResponseBody
}

type ModifyUserToDesktopGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserToDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserToDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserToDesktopGroupResponse) GetBody() *ModifyUserToDesktopGroupResponseBody {
	return s.Body
}

func (s *ModifyUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserToDesktopGroupResponse) SetStatusCode(v int32) *ModifyUserToDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserToDesktopGroupResponse) SetBody(v *ModifyUserToDesktopGroupResponseBody) *ModifyUserToDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyUserToDesktopGroupResponse) Validate() error {
	return dara.Validate(s)
}
