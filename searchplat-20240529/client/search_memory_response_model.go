// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMemoryResponse
	GetStatusCode() *int32
	SetBody(v *SearchMemoryResponseBody) *SearchMemoryResponse
	GetBody() *SearchMemoryResponseBody
}

type SearchMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponse) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMemoryResponse) GetBody() *SearchMemoryResponseBody {
	return s.Body
}

func (s *SearchMemoryResponse) SetHeaders(v map[string]*string) *SearchMemoryResponse {
	s.Headers = v
	return s
}

func (s *SearchMemoryResponse) SetStatusCode(v int32) *SearchMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMemoryResponse) SetBody(v *SearchMemoryResponseBody) *SearchMemoryResponse {
	s.Body = v
	return s
}

func (s *SearchMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
