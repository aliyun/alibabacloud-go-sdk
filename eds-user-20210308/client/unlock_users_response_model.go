// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockUsersResponse
	GetStatusCode() *int32
	SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse
	GetBody() *UnlockUsersResponseBody
}

type UnlockUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponse) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockUsersResponse) GetBody() *UnlockUsersResponseBody {
	return s.Body
}

func (s *UnlockUsersResponse) SetHeaders(v map[string]*string) *UnlockUsersResponse {
	s.Headers = v
	return s
}

func (s *UnlockUsersResponse) SetStatusCode(v int32) *UnlockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUsersResponse) SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse {
	s.Body = v
	return s
}

func (s *UnlockUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
