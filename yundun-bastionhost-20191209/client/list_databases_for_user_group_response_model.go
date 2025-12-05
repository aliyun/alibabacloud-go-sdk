// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabasesForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabasesForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabasesForUserGroupResponseBody) *ListDatabasesForUserGroupResponse
	GetBody() *ListDatabasesForUserGroupResponseBody
}

type ListDatabasesForUserGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabasesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabasesForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabasesForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabasesForUserGroupResponse) GetBody() *ListDatabasesForUserGroupResponseBody {
	return s.Body
}

func (s *ListDatabasesForUserGroupResponse) SetHeaders(v map[string]*string) *ListDatabasesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDatabasesForUserGroupResponse) SetStatusCode(v int32) *ListDatabasesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabasesForUserGroupResponse) SetBody(v *ListDatabasesForUserGroupResponseBody) *ListDatabasesForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListDatabasesForUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
