// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InviteUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InviteUsersResponse
	GetStatusCode() *int32
	SetBody(v *InviteUsersResponseBody) *InviteUsersResponse
	GetBody() *InviteUsersResponseBody
}

type InviteUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersResponse) GoString() string {
	return s.String()
}

func (s *InviteUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InviteUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InviteUsersResponse) GetBody() *InviteUsersResponseBody {
	return s.Body
}

func (s *InviteUsersResponse) SetHeaders(v map[string]*string) *InviteUsersResponse {
	s.Headers = v
	return s
}

func (s *InviteUsersResponse) SetStatusCode(v int32) *InviteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteUsersResponse) SetBody(v *InviteUsersResponseBody) *InviteUsersResponse {
	s.Body = v
	return s
}

func (s *InviteUsersResponse) Validate() error {
	return dara.Validate(s)
}
