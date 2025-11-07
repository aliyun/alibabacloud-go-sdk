// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGitRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGitRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGitRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateGitRepositoryResponseBody) *CreateGitRepositoryResponse
	GetBody() *CreateGitRepositoryResponseBody
}

type CreateGitRepositoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGitRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGitRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGitRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateGitRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGitRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGitRepositoryResponse) GetBody() *CreateGitRepositoryResponseBody {
	return s.Body
}

func (s *CreateGitRepositoryResponse) SetHeaders(v map[string]*string) *CreateGitRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateGitRepositoryResponse) SetStatusCode(v int32) *CreateGitRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGitRepositoryResponse) SetBody(v *CreateGitRepositoryResponseBody) *CreateGitRepositoryResponse {
	s.Body = v
	return s
}

func (s *CreateGitRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
