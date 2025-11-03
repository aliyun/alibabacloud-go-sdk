// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse
	GetBody() *CreateRepositoryResponseBody
}

type CreateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepositoryResponse) GetBody() *CreateRepositoryResponseBody {
	return s.Body
}

func (s *CreateRepositoryResponse) SetHeaders(v map[string]*string) *CreateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryResponse) SetStatusCode(v int32) *CreateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse {
	s.Body = v
	return s
}

func (s *CreateRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
