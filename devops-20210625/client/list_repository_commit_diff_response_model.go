// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitDiffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryCommitDiffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryCommitDiffResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryCommitDiffResponseBody) *ListRepositoryCommitDiffResponse
	GetBody() *ListRepositoryCommitDiffResponseBody
}

type ListRepositoryCommitDiffResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryCommitDiffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryCommitDiffResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitDiffResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryCommitDiffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryCommitDiffResponse) GetBody() *ListRepositoryCommitDiffResponseBody {
	return s.Body
}

func (s *ListRepositoryCommitDiffResponse) SetHeaders(v map[string]*string) *ListRepositoryCommitDiffResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetStatusCode(v int32) *ListRepositoryCommitDiffResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetBody(v *ListRepositoryCommitDiffResponseBody) *ListRepositoryCommitDiffResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryCommitDiffResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
