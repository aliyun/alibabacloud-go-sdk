// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRbacRoleHierarchyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRbacRoleHierarchyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRbacRoleHierarchyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRbacRoleHierarchyResponseBody) *RemoveRbacRoleHierarchyResponse
	GetBody() *RemoveRbacRoleHierarchyResponseBody
}

type RemoveRbacRoleHierarchyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRbacRoleHierarchyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRbacRoleHierarchyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRbacRoleHierarchyResponse) GoString() string {
	return s.String()
}

func (s *RemoveRbacRoleHierarchyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRbacRoleHierarchyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRbacRoleHierarchyResponse) GetBody() *RemoveRbacRoleHierarchyResponseBody {
	return s.Body
}

func (s *RemoveRbacRoleHierarchyResponse) SetHeaders(v map[string]*string) *RemoveRbacRoleHierarchyResponse {
	s.Headers = v
	return s
}

func (s *RemoveRbacRoleHierarchyResponse) SetStatusCode(v int32) *RemoveRbacRoleHierarchyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRbacRoleHierarchyResponse) SetBody(v *RemoveRbacRoleHierarchyResponseBody) *RemoveRbacRoleHierarchyResponse {
	s.Body = v
	return s
}

func (s *RemoveRbacRoleHierarchyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
