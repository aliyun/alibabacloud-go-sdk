// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopOversoldGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewDesktopOversoldGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewDesktopOversoldGroupResponse
	GetStatusCode() *int32
	SetBody(v *RenewDesktopOversoldGroupResponseBody) *RenewDesktopOversoldGroupResponse
	GetBody() *RenewDesktopOversoldGroupResponseBody
}

type RenewDesktopOversoldGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewDesktopOversoldGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewDesktopOversoldGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopOversoldGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopOversoldGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewDesktopOversoldGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewDesktopOversoldGroupResponse) GetBody() *RenewDesktopOversoldGroupResponseBody {
	return s.Body
}

func (s *RenewDesktopOversoldGroupResponse) SetHeaders(v map[string]*string) *RenewDesktopOversoldGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewDesktopOversoldGroupResponse) SetStatusCode(v int32) *RenewDesktopOversoldGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewDesktopOversoldGroupResponse) SetBody(v *RenewDesktopOversoldGroupResponseBody) *RenewDesktopOversoldGroupResponse {
	s.Body = v
	return s
}

func (s *RenewDesktopOversoldGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
