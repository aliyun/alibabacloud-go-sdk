// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostGroupsForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostGroupsForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListHostGroupsForUserGroupResponseBody) *ListHostGroupsForUserGroupResponse
	GetBody() *ListHostGroupsForUserGroupResponseBody
}

type ListHostGroupsForUserGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostGroupsForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostGroupsForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostGroupsForUserGroupResponse) GetBody() *ListHostGroupsForUserGroupResponseBody {
	return s.Body
}

func (s *ListHostGroupsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostGroupsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsForUserGroupResponse) SetStatusCode(v int32) *ListHostGroupsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponse) SetBody(v *ListHostGroupsForUserGroupResponseBody) *ListHostGroupsForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListHostGroupsForUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
