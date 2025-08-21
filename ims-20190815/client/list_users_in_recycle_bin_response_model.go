// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUsersInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUsersInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *ListUsersInRecycleBinResponseBody) *ListUsersInRecycleBinResponse
	GetBody() *ListUsersInRecycleBinResponseBody
}

type ListUsersInRecycleBinResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUsersInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUsersInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUsersInRecycleBinResponse) GetBody() *ListUsersInRecycleBinResponseBody {
	return s.Body
}

func (s *ListUsersInRecycleBinResponse) SetHeaders(v map[string]*string) *ListUsersInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *ListUsersInRecycleBinResponse) SetStatusCode(v int32) *ListUsersInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersInRecycleBinResponse) SetBody(v *ListUsersInRecycleBinResponseBody) *ListUsersInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *ListUsersInRecycleBinResponse) Validate() error {
	return dara.Validate(s)
}
