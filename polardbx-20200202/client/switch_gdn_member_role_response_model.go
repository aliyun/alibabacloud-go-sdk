// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGdnMemberRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchGdnMemberRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchGdnMemberRoleResponse
	GetStatusCode() *int32
	SetBody(v *SwitchGdnMemberRoleResponseBody) *SwitchGdnMemberRoleResponse
	GetBody() *SwitchGdnMemberRoleResponseBody
}

type SwitchGdnMemberRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchGdnMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchGdnMemberRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchGdnMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchGdnMemberRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchGdnMemberRoleResponse) GetBody() *SwitchGdnMemberRoleResponseBody {
	return s.Body
}

func (s *SwitchGdnMemberRoleResponse) SetHeaders(v map[string]*string) *SwitchGdnMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *SwitchGdnMemberRoleResponse) SetStatusCode(v int32) *SwitchGdnMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGdnMemberRoleResponse) SetBody(v *SwitchGdnMemberRoleResponseBody) *SwitchGdnMemberRoleResponse {
	s.Body = v
	return s
}

func (s *SwitchGdnMemberRoleResponse) Validate() error {
	return dara.Validate(s)
}
