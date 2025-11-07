// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoryContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGitRepositoryContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGitRepositoryContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListGitRepositoryContentsResponseBody) *ListGitRepositoryContentsResponse
	GetBody() *ListGitRepositoryContentsResponseBody
}

type ListGitRepositoryContentsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitRepositoryContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitRepositoryContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoryContentsResponse) GoString() string {
	return s.String()
}

func (s *ListGitRepositoryContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGitRepositoryContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGitRepositoryContentsResponse) GetBody() *ListGitRepositoryContentsResponseBody {
	return s.Body
}

func (s *ListGitRepositoryContentsResponse) SetHeaders(v map[string]*string) *ListGitRepositoryContentsResponse {
	s.Headers = v
	return s
}

func (s *ListGitRepositoryContentsResponse) SetStatusCode(v int32) *ListGitRepositoryContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitRepositoryContentsResponse) SetBody(v *ListGitRepositoryContentsResponseBody) *ListGitRepositoryContentsResponse {
	s.Body = v
	return s
}

func (s *ListGitRepositoryContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
