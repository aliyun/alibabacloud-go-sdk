// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockUsersResponse
	GetStatusCode() *int32
	SetBody(v *LockUsersResponseBody) *LockUsersResponse
	GetBody() *LockUsersResponseBody
}

type LockUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponse) GoString() string {
	return s.String()
}

func (s *LockUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockUsersResponse) GetBody() *LockUsersResponseBody {
	return s.Body
}

func (s *LockUsersResponse) SetHeaders(v map[string]*string) *LockUsersResponse {
	s.Headers = v
	return s
}

func (s *LockUsersResponse) SetStatusCode(v int32) *LockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *LockUsersResponse) SetBody(v *LockUsersResponseBody) *LockUsersResponse {
	s.Body = v
	return s
}

func (s *LockUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
