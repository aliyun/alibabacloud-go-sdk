// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByHybridResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaByHybridResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaByHybridResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaByHybridResponseBody) *SearchMediaByHybridResponse
	GetBody() *SearchMediaByHybridResponseBody
}

type SearchMediaByHybridResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaByHybridResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaByHybridResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaByHybridResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaByHybridResponse) GetBody() *SearchMediaByHybridResponseBody {
	return s.Body
}

func (s *SearchMediaByHybridResponse) SetHeaders(v map[string]*string) *SearchMediaByHybridResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaByHybridResponse) SetStatusCode(v int32) *SearchMediaByHybridResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaByHybridResponse) SetBody(v *SearchMediaByHybridResponseBody) *SearchMediaByHybridResponse {
	s.Body = v
	return s
}

func (s *SearchMediaByHybridResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
