// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchRepositoryResponseBody) *ListSearchRepositoryResponse
	GetBody() *ListSearchRepositoryResponseBody
}

type ListSearchRepositoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryResponse) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchRepositoryResponse) GetBody() *ListSearchRepositoryResponseBody {
	return s.Body
}

func (s *ListSearchRepositoryResponse) SetHeaders(v map[string]*string) *ListSearchRepositoryResponse {
	s.Headers = v
	return s
}

func (s *ListSearchRepositoryResponse) SetStatusCode(v int32) *ListSearchRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchRepositoryResponse) SetBody(v *ListSearchRepositoryResponseBody) *ListSearchRepositoryResponse {
	s.Body = v
	return s
}

func (s *ListSearchRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
