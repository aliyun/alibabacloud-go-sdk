// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *Repository) *GetRepositoryResponse
	GetBody() *Repository
}

type GetRepositoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Repository        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepositoryResponse) GetBody() *Repository {
	return s.Body
}

func (s *GetRepositoryResponse) SetHeaders(v map[string]*string) *GetRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryResponse) SetStatusCode(v int32) *GetRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryResponse) SetBody(v *Repository) *GetRepositoryResponse {
	s.Body = v
	return s
}

func (s *GetRepositoryResponse) Validate() error {
	return dara.Validate(s)
}
