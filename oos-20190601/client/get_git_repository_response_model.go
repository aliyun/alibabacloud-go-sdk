// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGitRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGitRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *GetGitRepositoryResponseBody) *GetGitRepositoryResponse
	GetBody() *GetGitRepositoryResponseBody
}

type GetGitRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGitRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGitRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGitRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetGitRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGitRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGitRepositoryResponse) GetBody() *GetGitRepositoryResponseBody {
	return s.Body
}

func (s *GetGitRepositoryResponse) SetHeaders(v map[string]*string) *GetGitRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetGitRepositoryResponse) SetStatusCode(v int32) *GetGitRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGitRepositoryResponse) SetBody(v *GetGitRepositoryResponseBody) *GetGitRepositoryResponse {
	s.Body = v
	return s
}

func (s *GetGitRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
