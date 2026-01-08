// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacebookPostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFacebookPostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFacebookPostsResponse
	GetStatusCode() *int32
	SetBody(v *ListFacebookPostsResponseBody) *ListFacebookPostsResponse
	GetBody() *ListFacebookPostsResponseBody
}

type ListFacebookPostsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFacebookPostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFacebookPostsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFacebookPostsResponse) GoString() string {
	return s.String()
}

func (s *ListFacebookPostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFacebookPostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFacebookPostsResponse) GetBody() *ListFacebookPostsResponseBody {
	return s.Body
}

func (s *ListFacebookPostsResponse) SetHeaders(v map[string]*string) *ListFacebookPostsResponse {
	s.Headers = v
	return s
}

func (s *ListFacebookPostsResponse) SetStatusCode(v int32) *ListFacebookPostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFacebookPostsResponse) SetBody(v *ListFacebookPostsResponseBody) *ListFacebookPostsResponse {
	s.Body = v
	return s
}

func (s *ListFacebookPostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
