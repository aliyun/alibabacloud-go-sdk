// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchNewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchNewsResponse
	GetStatusCode() *int32
	SetBody(v *SearchNewsResponseBody) *SearchNewsResponse
	GetBody() *SearchNewsResponseBody
}

type SearchNewsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchNewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchNewsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchNewsResponse) GoString() string {
	return s.String()
}

func (s *SearchNewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchNewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchNewsResponse) GetBody() *SearchNewsResponseBody {
	return s.Body
}

func (s *SearchNewsResponse) SetHeaders(v map[string]*string) *SearchNewsResponse {
	s.Headers = v
	return s
}

func (s *SearchNewsResponse) SetStatusCode(v int32) *SearchNewsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchNewsResponse) SetBody(v *SearchNewsResponseBody) *SearchNewsResponse {
	s.Body = v
	return s
}

func (s *SearchNewsResponse) Validate() error {
	return dara.Validate(s)
}
