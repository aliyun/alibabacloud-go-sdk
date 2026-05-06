// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchSampleResponse
	GetStatusCode() *int32
	SetBody(v *SearchSampleResponseBody) *SearchSampleResponse
	GetBody() *SearchSampleResponseBody
}

type SearchSampleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchSampleResponse) GoString() string {
	return s.String()
}

func (s *SearchSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchSampleResponse) GetBody() *SearchSampleResponseBody {
	return s.Body
}

func (s *SearchSampleResponse) SetHeaders(v map[string]*string) *SearchSampleResponse {
	s.Headers = v
	return s
}

func (s *SearchSampleResponse) SetStatusCode(v int32) *SearchSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSampleResponse) SetBody(v *SearchSampleResponseBody) *SearchSampleResponse {
	s.Body = v
	return s
}

func (s *SearchSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
