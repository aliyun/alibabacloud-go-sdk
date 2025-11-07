// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitOrganizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGitOrganizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGitOrganizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListGitOrganizationsResponseBody) *ListGitOrganizationsResponse
	GetBody() *ListGitOrganizationsResponseBody
}

type ListGitOrganizationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitOrganizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGitOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListGitOrganizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGitOrganizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGitOrganizationsResponse) GetBody() *ListGitOrganizationsResponseBody {
	return s.Body
}

func (s *ListGitOrganizationsResponse) SetHeaders(v map[string]*string) *ListGitOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *ListGitOrganizationsResponse) SetStatusCode(v int32) *ListGitOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitOrganizationsResponse) SetBody(v *ListGitOrganizationsResponseBody) *ListGitOrganizationsResponse {
	s.Body = v
	return s
}

func (s *ListGitOrganizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
