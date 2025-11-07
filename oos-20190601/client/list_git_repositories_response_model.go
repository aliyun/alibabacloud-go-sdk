// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGitRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGitRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListGitRepositoriesResponseBody) *ListGitRepositoriesResponse
	GetBody() *ListGitRepositoriesResponseBody
}

type ListGitRepositoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGitRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGitRepositoriesResponse) GetBody() *ListGitRepositoriesResponseBody {
	return s.Body
}

func (s *ListGitRepositoriesResponse) SetHeaders(v map[string]*string) *ListGitRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListGitRepositoriesResponse) SetStatusCode(v int32) *ListGitRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitRepositoriesResponse) SetBody(v *ListGitRepositoriesResponseBody) *ListGitRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListGitRepositoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
