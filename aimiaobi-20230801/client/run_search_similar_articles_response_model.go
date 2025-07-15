// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchSimilarArticlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSearchSimilarArticlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSearchSimilarArticlesResponse
	GetStatusCode() *int32
	SetBody(v *RunSearchSimilarArticlesResponseBody) *RunSearchSimilarArticlesResponse
	GetBody() *RunSearchSimilarArticlesResponseBody
}

type RunSearchSimilarArticlesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchSimilarArticlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchSimilarArticlesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponse) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSearchSimilarArticlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSearchSimilarArticlesResponse) GetBody() *RunSearchSimilarArticlesResponseBody {
	return s.Body
}

func (s *RunSearchSimilarArticlesResponse) SetHeaders(v map[string]*string) *RunSearchSimilarArticlesResponse {
	s.Headers = v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetStatusCode(v int32) *RunSearchSimilarArticlesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchSimilarArticlesResponse) SetBody(v *RunSearchSimilarArticlesResponseBody) *RunSearchSimilarArticlesResponse {
	s.Body = v
	return s
}

func (s *RunSearchSimilarArticlesResponse) Validate() error {
	return dara.Validate(s)
}
