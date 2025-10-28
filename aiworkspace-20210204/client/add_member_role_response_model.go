// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMemberRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMemberRoleResponse
	GetStatusCode() *int32
	SetBody(v *AddMemberRoleResponseBody) *AddMemberRoleResponse
	GetBody() *AddMemberRoleResponseBody
}

type AddMemberRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *AddMemberRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMemberRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMemberRoleResponse) GetBody() *AddMemberRoleResponseBody {
	return s.Body
}

func (s *AddMemberRoleResponse) SetHeaders(v map[string]*string) *AddMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *AddMemberRoleResponse) SetStatusCode(v int32) *AddMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberRoleResponse) SetBody(v *AddMemberRoleResponseBody) *AddMemberRoleResponse {
	s.Body = v
	return s
}

func (s *AddMemberRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
