// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepoFileExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckGitRepoFileExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckGitRepoFileExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckGitRepoFileExistsResponseBody) *CheckGitRepoFileExistsResponse
	GetBody() *CheckGitRepoFileExistsResponseBody
}

type CheckGitRepoFileExistsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckGitRepoFileExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckGitRepoFileExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepoFileExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckGitRepoFileExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckGitRepoFileExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckGitRepoFileExistsResponse) GetBody() *CheckGitRepoFileExistsResponseBody {
	return s.Body
}

func (s *CheckGitRepoFileExistsResponse) SetHeaders(v map[string]*string) *CheckGitRepoFileExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckGitRepoFileExistsResponse) SetStatusCode(v int32) *CheckGitRepoFileExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGitRepoFileExistsResponse) SetBody(v *CheckGitRepoFileExistsResponseBody) *CheckGitRepoFileExistsResponse {
	s.Body = v
	return s
}

func (s *CheckGitRepoFileExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
