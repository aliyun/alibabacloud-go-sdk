// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationMembersResponseBody) *ListOrganizationMembersResponse
	GetBody() *ListOrganizationMembersResponseBody
}

type ListOrganizationMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationMembersResponse) GetBody() *ListOrganizationMembersResponseBody {
	return s.Body
}

func (s *ListOrganizationMembersResponse) SetHeaders(v map[string]*string) *ListOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationMembersResponse) SetStatusCode(v int32) *ListOrganizationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationMembersResponse) SetBody(v *ListOrganizationMembersResponseBody) *ListOrganizationMembersResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
