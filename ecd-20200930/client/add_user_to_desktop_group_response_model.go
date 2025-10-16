// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToDesktopGroupResponseBody) *AddUserToDesktopGroupResponse
	GetBody() *AddUserToDesktopGroupResponseBody
}

type AddUserToDesktopGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToDesktopGroupResponse) GetBody() *AddUserToDesktopGroupResponseBody {
	return s.Body
}

func (s *AddUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *AddUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToDesktopGroupResponse) SetStatusCode(v int32) *AddUserToDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToDesktopGroupResponse) SetBody(v *AddUserToDesktopGroupResponseBody) *AddUserToDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *AddUserToDesktopGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
