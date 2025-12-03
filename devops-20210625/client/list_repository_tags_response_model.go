// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryTagsResponseBody) *ListRepositoryTagsResponse
	GetBody() *ListRepositoryTagsResponseBody
}

type ListRepositoryTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryTagsResponse) GetBody() *ListRepositoryTagsResponseBody {
	return s.Body
}

func (s *ListRepositoryTagsResponse) SetHeaders(v map[string]*string) *ListRepositoryTagsResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryTagsResponse) SetStatusCode(v int32) *ListRepositoryTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryTagsResponse) SetBody(v *ListRepositoryTagsResponseBody) *ListRepositoryTagsResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
