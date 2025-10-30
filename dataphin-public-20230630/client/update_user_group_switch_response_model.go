// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserGroupSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserGroupSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserGroupSwitchResponseBody) *UpdateUserGroupSwitchResponse
	GetBody() *UpdateUserGroupSwitchResponseBody
}

type UpdateUserGroupSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserGroupSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserGroupSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserGroupSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserGroupSwitchResponse) GetBody() *UpdateUserGroupSwitchResponseBody {
	return s.Body
}

func (s *UpdateUserGroupSwitchResponse) SetHeaders(v map[string]*string) *UpdateUserGroupSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupSwitchResponse) SetStatusCode(v int32) *UpdateUserGroupSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGroupSwitchResponse) SetBody(v *UpdateUserGroupSwitchResponseBody) *UpdateUserGroupSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateUserGroupSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
