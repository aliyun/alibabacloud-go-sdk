// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepositoryExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckGitRepositoryExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckGitRepositoryExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckGitRepositoryExistsResponseBody) *CheckGitRepositoryExistsResponse
	GetBody() *CheckGitRepositoryExistsResponseBody
}

type CheckGitRepositoryExistsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckGitRepositoryExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckGitRepositoryExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepositoryExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckGitRepositoryExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckGitRepositoryExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckGitRepositoryExistsResponse) GetBody() *CheckGitRepositoryExistsResponseBody {
	return s.Body
}

func (s *CheckGitRepositoryExistsResponse) SetHeaders(v map[string]*string) *CheckGitRepositoryExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckGitRepositoryExistsResponse) SetStatusCode(v int32) *CheckGitRepositoryExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGitRepositoryExistsResponse) SetBody(v *CheckGitRepositoryExistsResponseBody) *CheckGitRepositoryExistsResponse {
	s.Body = v
	return s
}

func (s *CheckGitRepositoryExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
