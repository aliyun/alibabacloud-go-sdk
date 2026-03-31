// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoliciesForRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoliciesForRoleResponse
	GetStatusCode() *int32
	SetBody(v *ListPoliciesForRoleResponseBody) *ListPoliciesForRoleResponse
	GetBody() *ListPoliciesForRoleResponseBody
}

type ListPoliciesForRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForRoleResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoliciesForRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoliciesForRoleResponse) GetBody() *ListPoliciesForRoleResponseBody {
	return s.Body
}

func (s *ListPoliciesForRoleResponse) SetHeaders(v map[string]*string) *ListPoliciesForRoleResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForRoleResponse) SetStatusCode(v int32) *ListPoliciesForRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForRoleResponse) SetBody(v *ListPoliciesForRoleResponseBody) *ListPoliciesForRoleResponse {
	s.Body = v
	return s
}

func (s *ListPoliciesForRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
