// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleAssignmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoleAssignmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoleAssignmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListRoleAssignmentsResponseBody) *ListRoleAssignmentsResponse
	GetBody() *ListRoleAssignmentsResponseBody
}

type ListRoleAssignmentsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoleAssignmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoleAssignmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoleAssignmentsResponse) GoString() string {
	return s.String()
}

func (s *ListRoleAssignmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoleAssignmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoleAssignmentsResponse) GetBody() *ListRoleAssignmentsResponseBody {
	return s.Body
}

func (s *ListRoleAssignmentsResponse) SetHeaders(v map[string]*string) *ListRoleAssignmentsResponse {
	s.Headers = v
	return s
}

func (s *ListRoleAssignmentsResponse) SetStatusCode(v int32) *ListRoleAssignmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleAssignmentsResponse) SetBody(v *ListRoleAssignmentsResponseBody) *ListRoleAssignmentsResponse {
	s.Body = v
	return s
}

func (s *ListRoleAssignmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
