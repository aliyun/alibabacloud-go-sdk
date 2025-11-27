// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStoriesResponse
	GetStatusCode() *int32
	SetBody(v *QueryStoriesResponseBody) *QueryStoriesResponse
	GetBody() *QueryStoriesResponseBody
}

type QueryStoriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStoriesResponse) GoString() string {
	return s.String()
}

func (s *QueryStoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStoriesResponse) GetBody() *QueryStoriesResponseBody {
	return s.Body
}

func (s *QueryStoriesResponse) SetHeaders(v map[string]*string) *QueryStoriesResponse {
	s.Headers = v
	return s
}

func (s *QueryStoriesResponse) SetStatusCode(v int32) *QueryStoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStoriesResponse) SetBody(v *QueryStoriesResponseBody) *QueryStoriesResponse {
	s.Body = v
	return s
}

func (s *QueryStoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
