// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTableResponse
	GetStatusCode() *int32
	SetBody(v *SearchTableResponseBody) *SearchTableResponse
	GetBody() *SearchTableResponseBody
}

type SearchTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTableResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTableResponse) GoString() string {
	return s.String()
}

func (s *SearchTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTableResponse) GetBody() *SearchTableResponseBody {
	return s.Body
}

func (s *SearchTableResponse) SetHeaders(v map[string]*string) *SearchTableResponse {
	s.Headers = v
	return s
}

func (s *SearchTableResponse) SetStatusCode(v int32) *SearchTableResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTableResponse) SetBody(v *SearchTableResponseBody) *SearchTableResponse {
	s.Body = v
	return s
}

func (s *SearchTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
