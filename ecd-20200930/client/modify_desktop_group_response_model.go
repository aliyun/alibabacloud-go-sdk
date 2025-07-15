// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopGroupResponseBody) *ModifyDesktopGroupResponse
	GetBody() *ModifyDesktopGroupResponseBody
}

type ModifyDesktopGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopGroupResponse) GetBody() *ModifyDesktopGroupResponseBody {
	return s.Body
}

func (s *ModifyDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopGroupResponse) SetStatusCode(v int32) *ModifyDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopGroupResponse) SetBody(v *ModifyDesktopGroupResponseBody) *ModifyDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopGroupResponse) Validate() error {
	return dara.Validate(s)
}
