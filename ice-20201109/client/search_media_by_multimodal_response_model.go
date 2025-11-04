// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByMultimodalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaByMultimodalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaByMultimodalResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaByMultimodalResponseBody) *SearchMediaByMultimodalResponse
	GetBody() *SearchMediaByMultimodalResponseBody
}

type SearchMediaByMultimodalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaByMultimodalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaByMultimodalResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByMultimodalResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaByMultimodalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaByMultimodalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaByMultimodalResponse) GetBody() *SearchMediaByMultimodalResponseBody {
	return s.Body
}

func (s *SearchMediaByMultimodalResponse) SetHeaders(v map[string]*string) *SearchMediaByMultimodalResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaByMultimodalResponse) SetStatusCode(v int32) *SearchMediaByMultimodalResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaByMultimodalResponse) SetBody(v *SearchMediaByMultimodalResponseBody) *SearchMediaByMultimodalResponse {
	s.Body = v
	return s
}

func (s *SearchMediaByMultimodalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
