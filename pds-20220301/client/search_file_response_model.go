// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFileResponse
	GetStatusCode() *int32
	SetBody(v *SearchFileResponseBody) *SearchFileResponse
	GetBody() *SearchFileResponseBody
}

type SearchFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFileResponse) GoString() string {
	return s.String()
}

func (s *SearchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFileResponse) GetBody() *SearchFileResponseBody {
	return s.Body
}

func (s *SearchFileResponse) SetHeaders(v map[string]*string) *SearchFileResponse {
	s.Headers = v
	return s
}

func (s *SearchFileResponse) SetStatusCode(v int32) *SearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFileResponse) SetBody(v *SearchFileResponseBody) *SearchFileResponse {
	s.Body = v
	return s
}

func (s *SearchFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
