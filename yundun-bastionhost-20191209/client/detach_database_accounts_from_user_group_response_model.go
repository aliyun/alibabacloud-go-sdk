// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDatabaseAccountsFromUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDatabaseAccountsFromUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DetachDatabaseAccountsFromUserGroupResponseBody) *DetachDatabaseAccountsFromUserGroupResponse
	GetBody() *DetachDatabaseAccountsFromUserGroupResponseBody
}

type DetachDatabaseAccountsFromUserGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDatabaseAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDatabaseAccountsFromUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) GetBody() *DetachDatabaseAccountsFromUserGroupResponseBody {
	return s.Body
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) SetHeaders(v map[string]*string) *DetachDatabaseAccountsFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) SetStatusCode(v int32) *DetachDatabaseAccountsFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) SetBody(v *DetachDatabaseAccountsFromUserGroupResponseBody) *DetachDatabaseAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
