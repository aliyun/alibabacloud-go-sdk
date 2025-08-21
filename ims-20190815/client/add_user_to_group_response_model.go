// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToGroupResponseBody) *AddUserToGroupResponse
	GetBody() *AddUserToGroupResponseBody
}

type AddUserToGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToGroupResponse) GetBody() *AddUserToGroupResponseBody {
	return s.Body
}

func (s *AddUserToGroupResponse) SetHeaders(v map[string]*string) *AddUserToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToGroupResponse) SetStatusCode(v int32) *AddUserToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToGroupResponse) SetBody(v *AddUserToGroupResponseBody) *AddUserToGroupResponse {
	s.Body = v
	return s
}

func (s *AddUserToGroupResponse) Validate() error {
	return dara.Validate(s)
}
