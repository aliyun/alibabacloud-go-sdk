// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryBranchesResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryBranchesResponseBody) *ListRepositoryBranchesResponse
	GetBody() *ListRepositoryBranchesResponseBody
}

type ListRepositoryBranchesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryBranchesResponse) GetBody() *ListRepositoryBranchesResponseBody {
	return s.Body
}

func (s *ListRepositoryBranchesResponse) SetHeaders(v map[string]*string) *ListRepositoryBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryBranchesResponse) SetStatusCode(v int32) *ListRepositoryBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryBranchesResponse) SetBody(v *ListRepositoryBranchesResponseBody) *ListRepositoryBranchesResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
