// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParseProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParseProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetParseProgressResponseBody) *GetParseProgressResponse
	GetBody() *GetParseProgressResponseBody
}

type GetParseProgressResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParseProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParseProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParseProgressResponse) GoString() string {
	return s.String()
}

func (s *GetParseProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParseProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParseProgressResponse) GetBody() *GetParseProgressResponseBody {
	return s.Body
}

func (s *GetParseProgressResponse) SetHeaders(v map[string]*string) *GetParseProgressResponse {
	s.Headers = v
	return s
}

func (s *GetParseProgressResponse) SetStatusCode(v int32) *GetParseProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParseProgressResponse) SetBody(v *GetParseProgressResponseBody) *GetParseProgressResponse {
	s.Body = v
	return s
}

func (s *GetParseProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
