// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *RenewDesktopGroupResponseBody) *RenewDesktopGroupResponse
	GetBody() *RenewDesktopGroupResponseBody
}

type RenewDesktopGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewDesktopGroupResponse) GetBody() *RenewDesktopGroupResponseBody {
	return s.Body
}

func (s *RenewDesktopGroupResponse) SetHeaders(v map[string]*string) *RenewDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewDesktopGroupResponse) SetStatusCode(v int32) *RenewDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewDesktopGroupResponse) SetBody(v *RenewDesktopGroupResponseBody) *RenewDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *RenewDesktopGroupResponse) Validate() error {
	return dara.Validate(s)
}
