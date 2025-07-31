// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse
	GetBody() *UpdateUserGroupResponseBody
}

type UpdateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserGroupResponse) GetBody() *UpdateUserGroupResponseBody {
	return s.Body
}

func (s *UpdateUserGroupResponse) SetHeaders(v map[string]*string) *UpdateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupResponse) SetStatusCode(v int32) *UpdateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGroupResponse) SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
