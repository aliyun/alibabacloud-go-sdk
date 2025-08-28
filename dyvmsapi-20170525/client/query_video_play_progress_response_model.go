// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoPlayProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVideoPlayProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVideoPlayProgressResponse
	GetStatusCode() *int32
	SetBody(v *QueryVideoPlayProgressResponseBody) *QueryVideoPlayProgressResponse
	GetBody() *QueryVideoPlayProgressResponseBody
}

type QueryVideoPlayProgressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVideoPlayProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVideoPlayProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoPlayProgressResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVideoPlayProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVideoPlayProgressResponse) GetBody() *QueryVideoPlayProgressResponseBody {
	return s.Body
}

func (s *QueryVideoPlayProgressResponse) SetHeaders(v map[string]*string) *QueryVideoPlayProgressResponse {
	s.Headers = v
	return s
}

func (s *QueryVideoPlayProgressResponse) SetStatusCode(v int32) *QueryVideoPlayProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVideoPlayProgressResponse) SetBody(v *QueryVideoPlayProgressResponseBody) *QueryVideoPlayProgressResponse {
	s.Body = v
	return s
}

func (s *QueryVideoPlayProgressResponse) Validate() error {
	return dara.Validate(s)
}
