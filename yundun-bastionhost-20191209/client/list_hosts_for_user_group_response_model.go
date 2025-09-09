// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostsForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostsForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListHostsForUserGroupResponseBody) *ListHostsForUserGroupResponse
	GetBody() *ListHostsForUserGroupResponseBody
}

type ListHostsForUserGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostsForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostsForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostsForUserGroupResponse) GetBody() *ListHostsForUserGroupResponseBody {
	return s.Body
}

func (s *ListHostsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostsForUserGroupResponse) SetStatusCode(v int32) *ListHostsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsForUserGroupResponse) SetBody(v *ListHostsForUserGroupResponseBody) *ListHostsForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListHostsForUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
