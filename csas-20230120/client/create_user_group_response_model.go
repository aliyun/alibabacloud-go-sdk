// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse
	GetBody() *CreateUserGroupResponseBody
}

type CreateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserGroupResponse) GetBody() *CreateUserGroupResponseBody {
	return s.Body
}

func (s *CreateUserGroupResponse) SetHeaders(v map[string]*string) *CreateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupResponse) SetStatusCode(v int32) *CreateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

func (s *CreateUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
