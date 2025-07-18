// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse
	GetBody() *ListUserGroupsResponseBody
}

type ListUserGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserGroupsResponse) GetBody() *ListUserGroupsResponseBody {
	return s.Body
}

func (s *ListUserGroupsResponse) SetHeaders(v map[string]*string) *ListUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsResponse) SetStatusCode(v int32) *ListUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

func (s *ListUserGroupsResponse) Validate() error {
	return dara.Validate(s)
}
