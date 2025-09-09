// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse
	GetBody() *DeleteUserGroupResponseBody
}

type DeleteUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserGroupResponse) GetBody() *DeleteUserGroupResponseBody {
	return s.Body
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetStatusCode(v int32) *DeleteUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
