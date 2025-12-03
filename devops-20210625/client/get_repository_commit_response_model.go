// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryCommitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepositoryCommitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepositoryCommitResponse
	GetStatusCode() *int32
	SetBody(v *GetRepositoryCommitResponseBody) *GetRepositoryCommitResponse
	GetBody() *GetRepositoryCommitResponseBody
}

type GetRepositoryCommitResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepositoryCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepositoryCommitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepositoryCommitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepositoryCommitResponse) GetBody() *GetRepositoryCommitResponseBody {
	return s.Body
}

func (s *GetRepositoryCommitResponse) SetHeaders(v map[string]*string) *GetRepositoryCommitResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryCommitResponse) SetStatusCode(v int32) *GetRepositoryCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryCommitResponse) SetBody(v *GetRepositoryCommitResponseBody) *GetRepositoryCommitResponse {
	s.Body = v
	return s
}

func (s *GetRepositoryCommitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
