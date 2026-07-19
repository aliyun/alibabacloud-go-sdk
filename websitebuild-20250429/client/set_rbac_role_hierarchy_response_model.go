// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRbacRoleHierarchyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRbacRoleHierarchyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRbacRoleHierarchyResponse
	GetStatusCode() *int32
	SetBody(v *SetRbacRoleHierarchyResponseBody) *SetRbacRoleHierarchyResponse
	GetBody() *SetRbacRoleHierarchyResponseBody
}

type SetRbacRoleHierarchyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRbacRoleHierarchyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRbacRoleHierarchyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRbacRoleHierarchyResponse) GoString() string {
	return s.String()
}

func (s *SetRbacRoleHierarchyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRbacRoleHierarchyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRbacRoleHierarchyResponse) GetBody() *SetRbacRoleHierarchyResponseBody {
	return s.Body
}

func (s *SetRbacRoleHierarchyResponse) SetHeaders(v map[string]*string) *SetRbacRoleHierarchyResponse {
	s.Headers = v
	return s
}

func (s *SetRbacRoleHierarchyResponse) SetStatusCode(v int32) *SetRbacRoleHierarchyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRbacRoleHierarchyResponse) SetBody(v *SetRbacRoleHierarchyResponseBody) *SetRbacRoleHierarchyResponse {
	s.Body = v
	return s
}

func (s *SetRbacRoleHierarchyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
