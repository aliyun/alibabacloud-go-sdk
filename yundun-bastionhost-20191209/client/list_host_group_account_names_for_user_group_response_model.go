// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostGroupAccountNamesForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListHostGroupAccountNamesForUserGroupResponseBody) *ListHostGroupAccountNamesForUserGroupResponse
	GetBody() *ListHostGroupAccountNamesForUserGroupResponseBody
}

type ListHostGroupAccountNamesForUserGroupResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupAccountNamesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) GetBody() *ListHostGroupAccountNamesForUserGroupResponseBody {
	return s.Body
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetStatusCode(v int32) *ListHostGroupAccountNamesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetBody(v *ListHostGroupAccountNamesForUserGroupResponseBody) *ListHostGroupAccountNamesForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
