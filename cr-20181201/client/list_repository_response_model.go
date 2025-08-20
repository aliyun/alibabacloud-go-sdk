// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryResponseBody) *ListRepositoryResponse
	GetBody() *ListRepositoryResponseBody
}

type ListRepositoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryResponse) GetBody() *ListRepositoryResponseBody {
	return s.Body
}

func (s *ListRepositoryResponse) SetHeaders(v map[string]*string) *ListRepositoryResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryResponse) SetStatusCode(v int32) *ListRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryResponse) SetBody(v *ListRepositoryResponseBody) *ListRepositoryResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryResponse) Validate() error {
	return dara.Validate(s)
}
