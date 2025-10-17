// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDocResponse
	GetStatusCode() *int32
	SetBody(v *SearchDocResponseBody) *SearchDocResponse
	GetBody() *SearchDocResponseBody
}

type SearchDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDocResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDocResponse) GoString() string {
	return s.String()
}

func (s *SearchDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDocResponse) GetBody() *SearchDocResponseBody {
	return s.Body
}

func (s *SearchDocResponse) SetHeaders(v map[string]*string) *SearchDocResponse {
	s.Headers = v
	return s
}

func (s *SearchDocResponse) SetStatusCode(v int32) *SearchDocResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDocResponse) SetBody(v *SearchDocResponseBody) *SearchDocResponse {
	s.Body = v
	return s
}

func (s *SearchDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
