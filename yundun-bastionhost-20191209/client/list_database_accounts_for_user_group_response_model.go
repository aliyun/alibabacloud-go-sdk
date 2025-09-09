// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseAccountsForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseAccountsForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseAccountsForUserGroupResponseBody) *ListDatabaseAccountsForUserGroupResponse
	GetBody() *ListDatabaseAccountsForUserGroupResponseBody
}

type ListDatabaseAccountsForUserGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseAccountsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseAccountsForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseAccountsForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseAccountsForUserGroupResponse) GetBody() *ListDatabaseAccountsForUserGroupResponseBody {
	return s.Body
}

func (s *ListDatabaseAccountsForUserGroupResponse) SetHeaders(v map[string]*string) *ListDatabaseAccountsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponse) SetStatusCode(v int32) *ListDatabaseAccountsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponse) SetBody(v *ListDatabaseAccountsForUserGroupResponseBody) *ListDatabaseAccountsForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
