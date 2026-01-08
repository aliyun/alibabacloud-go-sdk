// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstagramPostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstagramPostsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstagramPostsResponseBody) *ListInstagramPostsResponse
	GetBody() *ListInstagramPostsResponseBody
}

type ListInstagramPostsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstagramPostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstagramPostsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPostsResponse) GoString() string {
	return s.String()
}

func (s *ListInstagramPostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstagramPostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstagramPostsResponse) GetBody() *ListInstagramPostsResponseBody {
	return s.Body
}

func (s *ListInstagramPostsResponse) SetHeaders(v map[string]*string) *ListInstagramPostsResponse {
	s.Headers = v
	return s
}

func (s *ListInstagramPostsResponse) SetStatusCode(v int32) *ListInstagramPostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstagramPostsResponse) SetBody(v *ListInstagramPostsResponseBody) *ListInstagramPostsResponse {
	s.Body = v
	return s
}

func (s *ListInstagramPostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
