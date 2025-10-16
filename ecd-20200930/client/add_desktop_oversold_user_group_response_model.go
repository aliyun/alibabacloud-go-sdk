// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesktopOversoldUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDesktopOversoldUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDesktopOversoldUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddDesktopOversoldUserGroupResponseBody) *AddDesktopOversoldUserGroupResponse
	GetBody() *AddDesktopOversoldUserGroupResponseBody
}

type AddDesktopOversoldUserGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDesktopOversoldUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDesktopOversoldUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDesktopOversoldUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDesktopOversoldUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDesktopOversoldUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDesktopOversoldUserGroupResponse) GetBody() *AddDesktopOversoldUserGroupResponseBody {
	return s.Body
}

func (s *AddDesktopOversoldUserGroupResponse) SetHeaders(v map[string]*string) *AddDesktopOversoldUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDesktopOversoldUserGroupResponse) SetStatusCode(v int32) *AddDesktopOversoldUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDesktopOversoldUserGroupResponse) SetBody(v *AddDesktopOversoldUserGroupResponseBody) *AddDesktopOversoldUserGroupResponse {
	s.Body = v
	return s
}

func (s *AddDesktopOversoldUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
