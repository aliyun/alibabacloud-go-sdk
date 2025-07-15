// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignUsersResponse
	GetStatusCode() *int32
	SetBody(v *AssignUsersResponseBody) *AssignUsersResponse
	GetBody() *AssignUsersResponseBody
}

type AssignUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignUsersResponse) GoString() string {
	return s.String()
}

func (s *AssignUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignUsersResponse) GetBody() *AssignUsersResponseBody {
	return s.Body
}

func (s *AssignUsersResponse) SetHeaders(v map[string]*string) *AssignUsersResponse {
	s.Headers = v
	return s
}

func (s *AssignUsersResponse) SetStatusCode(v int32) *AssignUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignUsersResponse) SetBody(v *AssignUsersResponseBody) *AssignUsersResponse {
	s.Body = v
	return s
}

func (s *AssignUsersResponse) Validate() error {
	return dara.Validate(s)
}
