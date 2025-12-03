// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryCommitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryCommitsResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryCommitsResponseBody) *ListRepositoryCommitsResponse
	GetBody() *ListRepositoryCommitsResponseBody
}

type ListRepositoryCommitsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryCommitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryCommitsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryCommitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryCommitsResponse) GetBody() *ListRepositoryCommitsResponseBody {
	return s.Body
}

func (s *ListRepositoryCommitsResponse) SetHeaders(v map[string]*string) *ListRepositoryCommitsResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCommitsResponse) SetStatusCode(v int32) *ListRepositoryCommitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCommitsResponse) SetBody(v *ListRepositoryCommitsResponseBody) *ListRepositoryCommitsResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryCommitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
