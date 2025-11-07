// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeGitRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeGitRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeGitRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeGitRepositoryResponseBody) *AnalyzeGitRepositoryResponse
	GetBody() *AnalyzeGitRepositoryResponseBody
}

type AnalyzeGitRepositoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeGitRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeGitRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeGitRepositoryResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeGitRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeGitRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeGitRepositoryResponse) GetBody() *AnalyzeGitRepositoryResponseBody {
	return s.Body
}

func (s *AnalyzeGitRepositoryResponse) SetHeaders(v map[string]*string) *AnalyzeGitRepositoryResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeGitRepositoryResponse) SetStatusCode(v int32) *AnalyzeGitRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeGitRepositoryResponse) SetBody(v *AnalyzeGitRepositoryResponseBody) *AnalyzeGitRepositoryResponse {
	s.Body = v
	return s
}

func (s *AnalyzeGitRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
