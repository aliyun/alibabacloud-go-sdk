// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDatabaseAccountsToUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDatabaseAccountsToUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AttachDatabaseAccountsToUserGroupResponseBody) *AttachDatabaseAccountsToUserGroupResponse
	GetBody() *AttachDatabaseAccountsToUserGroupResponseBody
}

type AttachDatabaseAccountsToUserGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDatabaseAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDatabaseAccountsToUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDatabaseAccountsToUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDatabaseAccountsToUserGroupResponse) GetBody() *AttachDatabaseAccountsToUserGroupResponseBody {
	return s.Body
}

func (s *AttachDatabaseAccountsToUserGroupResponse) SetHeaders(v map[string]*string) *AttachDatabaseAccountsToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponse) SetStatusCode(v int32) *AttachDatabaseAccountsToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponse) SetBody(v *AttachDatabaseAccountsToUserGroupResponseBody) *AttachDatabaseAccountsToUserGroupResponse {
	s.Body = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
