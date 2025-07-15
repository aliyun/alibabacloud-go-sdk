// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopOversoldUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopOversoldUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopOversoldUserGroupResponseBody) *ModifyDesktopOversoldUserGroupResponse
	GetBody() *ModifyDesktopOversoldUserGroupResponseBody
}

type ModifyDesktopOversoldUserGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopOversoldUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopOversoldUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopOversoldUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopOversoldUserGroupResponse) GetBody() *ModifyDesktopOversoldUserGroupResponseBody {
	return s.Body
}

func (s *ModifyDesktopOversoldUserGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopOversoldUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponse) SetStatusCode(v int32) *ModifyDesktopOversoldUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponse) SetBody(v *ModifyDesktopOversoldUserGroupResponseBody) *ModifyDesktopOversoldUserGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
