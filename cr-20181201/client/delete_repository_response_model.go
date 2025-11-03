// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse
	GetBody() *DeleteRepositoryResponseBody
}

type DeleteRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepositoryResponse) GetBody() *DeleteRepositoryResponseBody {
	return s.Body
}

func (s *DeleteRepositoryResponse) SetHeaders(v map[string]*string) *DeleteRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryResponse) SetStatusCode(v int32) *DeleteRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryResponse) SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse {
	s.Body = v
	return s
}

func (s *DeleteRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
