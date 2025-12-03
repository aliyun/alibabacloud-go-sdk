// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryTreeResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryTreeResponseBody) *ListRepositoryTreeResponse
	GetBody() *ListRepositoryTreeResponseBody
}

type ListRepositoryTreeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTreeResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryTreeResponse) GetBody() *ListRepositoryTreeResponseBody {
	return s.Body
}

func (s *ListRepositoryTreeResponse) SetHeaders(v map[string]*string) *ListRepositoryTreeResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryTreeResponse) SetStatusCode(v int32) *ListRepositoryTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryTreeResponse) SetBody(v *ListRepositoryTreeResponseBody) *ListRepositoryTreeResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
