// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostAccountsForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostAccountsForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListHostAccountsForUserGroupResponseBody) *ListHostAccountsForUserGroupResponse
	GetBody() *ListHostAccountsForUserGroupResponseBody
}

type ListHostAccountsForUserGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostAccountsForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostAccountsForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostAccountsForUserGroupResponse) GetBody() *ListHostAccountsForUserGroupResponseBody {
	return s.Body
}

func (s *ListHostAccountsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostAccountsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForUserGroupResponse) SetStatusCode(v int32) *ListHostAccountsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponse) SetBody(v *ListHostAccountsForUserGroupResponseBody) *ListHostAccountsForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListHostAccountsForUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
