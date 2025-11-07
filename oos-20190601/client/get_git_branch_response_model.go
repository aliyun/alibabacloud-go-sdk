// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGitBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGitBranchResponse
	GetStatusCode() *int32
	SetBody(v *GetGitBranchResponseBody) *GetGitBranchResponse
	GetBody() *GetGitBranchResponseBody
}

type GetGitBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGitBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGitBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGitBranchResponse) GoString() string {
	return s.String()
}

func (s *GetGitBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGitBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGitBranchResponse) GetBody() *GetGitBranchResponseBody {
	return s.Body
}

func (s *GetGitBranchResponse) SetHeaders(v map[string]*string) *GetGitBranchResponse {
	s.Headers = v
	return s
}

func (s *GetGitBranchResponse) SetStatusCode(v int32) *GetGitBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGitBranchResponse) SetBody(v *GetGitBranchResponseBody) *GetGitBranchResponse {
	s.Body = v
	return s
}

func (s *GetGitBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
