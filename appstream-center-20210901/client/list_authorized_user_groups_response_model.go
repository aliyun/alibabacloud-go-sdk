// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUserGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedUserGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedUserGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedUserGroupsResponseBody) *ListAuthorizedUserGroupsResponse
	GetBody() *ListAuthorizedUserGroupsResponseBody
}

type ListAuthorizedUserGroupsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedUserGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUserGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedUserGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedUserGroupsResponse) GetBody() *ListAuthorizedUserGroupsResponseBody {
	return s.Body
}

func (s *ListAuthorizedUserGroupsResponse) SetHeaders(v map[string]*string) *ListAuthorizedUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedUserGroupsResponse) SetStatusCode(v int32) *ListAuthorizedUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedUserGroupsResponse) SetBody(v *ListAuthorizedUserGroupsResponseBody) *ListAuthorizedUserGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedUserGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
