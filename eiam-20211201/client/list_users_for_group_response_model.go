// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUsersForGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUsersForGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse
	GetBody() *ListUsersForGroupResponseBody
}

type ListUsersForGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUsersForGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUsersForGroupResponse) GetBody() *ListUsersForGroupResponseBody {
	return s.Body
}

func (s *ListUsersForGroupResponse) SetHeaders(v map[string]*string) *ListUsersForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForGroupResponse) SetStatusCode(v int32) *ListUsersForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForGroupResponse) SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse {
	s.Body = v
	return s
}

func (s *ListUsersForGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
