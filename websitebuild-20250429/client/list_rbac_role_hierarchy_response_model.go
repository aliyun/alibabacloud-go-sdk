// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRoleHierarchyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRbacRoleHierarchyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRbacRoleHierarchyResponse
	GetStatusCode() *int32
	SetBody(v *ListRbacRoleHierarchyResponseBody) *ListRbacRoleHierarchyResponse
	GetBody() *ListRbacRoleHierarchyResponseBody
}

type ListRbacRoleHierarchyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRbacRoleHierarchyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRbacRoleHierarchyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyResponse) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRbacRoleHierarchyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRbacRoleHierarchyResponse) GetBody() *ListRbacRoleHierarchyResponseBody {
	return s.Body
}

func (s *ListRbacRoleHierarchyResponse) SetHeaders(v map[string]*string) *ListRbacRoleHierarchyResponse {
	s.Headers = v
	return s
}

func (s *ListRbacRoleHierarchyResponse) SetStatusCode(v int32) *ListRbacRoleHierarchyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRbacRoleHierarchyResponse) SetBody(v *ListRbacRoleHierarchyResponseBody) *ListRbacRoleHierarchyResponse {
	s.Body = v
	return s
}

func (s *ListRbacRoleHierarchyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
