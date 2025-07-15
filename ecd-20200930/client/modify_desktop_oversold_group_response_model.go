// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopOversoldGroupResponseBody) *ModifyDesktopOversoldGroupResponse
	GetBody() *ModifyDesktopOversoldGroupResponseBody
}

type ModifyDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopOversoldGroupResponse) GetBody() *ModifyDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *ModifyDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopOversoldGroupResponse) SetStatusCode(v int32) *ModifyDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopOversoldGroupResponse) SetBody(v *ModifyDesktopOversoldGroupResponseBody) *ModifyDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopOversoldGroupResponse) Validate() error {
	return dara.Validate(s)
}
