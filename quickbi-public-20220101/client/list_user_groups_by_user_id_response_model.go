// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsByUserIdResponseBody) *ListUserGroupsByUserIdResponse
	GetBody() *ListUserGroupsByUserIdResponseBody
}

type ListUserGroupsByUserIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsByUserIdResponse) GetBody() *ListUserGroupsByUserIdResponseBody {
	return s.Body
}

func (s *ListUserGroupsByUserIdResponse) SetHeaders(v map[string]*string) *ListUserGroupsByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsByUserIdResponse) SetStatusCode(v int32) *ListUserGroupsByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsByUserIdResponse) SetBody(v *ListUserGroupsByUserIdResponseBody) *ListUserGroupsByUserIdResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsByUserIdResponse) Validate() error {
	return dara.Validate(s)
}
