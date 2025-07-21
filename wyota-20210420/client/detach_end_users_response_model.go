// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachEndUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachEndUsersResponse
	GetStatusCode() *int32
	SetBody(v *DetachEndUsersResponseBody) *DetachEndUsersResponse
	GetBody() *DetachEndUsersResponseBody
}

type DetachEndUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEndUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEndUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUsersResponse) GoString() string {
	return s.String()
}

func (s *DetachEndUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachEndUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachEndUsersResponse) GetBody() *DetachEndUsersResponseBody {
	return s.Body
}

func (s *DetachEndUsersResponse) SetHeaders(v map[string]*string) *DetachEndUsersResponse {
	s.Headers = v
	return s
}

func (s *DetachEndUsersResponse) SetStatusCode(v int32) *DetachEndUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEndUsersResponse) SetBody(v *DetachEndUsersResponseBody) *DetachEndUsersResponse {
	s.Body = v
	return s
}

func (s *DetachEndUsersResponse) Validate() error {
	return dara.Validate(s)
}
