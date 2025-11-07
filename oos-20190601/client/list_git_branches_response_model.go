// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGitBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGitBranchesResponse
	GetStatusCode() *int32
	SetBody(v *ListGitBranchesResponseBody) *ListGitBranchesResponse
	GetBody() *ListGitBranchesResponseBody
}

type ListGitBranchesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGitBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListGitBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGitBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGitBranchesResponse) GetBody() *ListGitBranchesResponseBody {
	return s.Body
}

func (s *ListGitBranchesResponse) SetHeaders(v map[string]*string) *ListGitBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListGitBranchesResponse) SetStatusCode(v int32) *ListGitBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitBranchesResponse) SetBody(v *ListGitBranchesResponseBody) *ListGitBranchesResponse {
	s.Body = v
	return s
}

func (s *ListGitBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
