// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchContentResponse
	GetStatusCode() *int32
	SetBody(v *SearchContentResponseBody) *SearchContentResponse
	GetBody() *SearchContentResponseBody
}

type SearchContentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchContentResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchContentResponse) GoString() string {
	return s.String()
}

func (s *SearchContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchContentResponse) GetBody() *SearchContentResponseBody {
	return s.Body
}

func (s *SearchContentResponse) SetHeaders(v map[string]*string) *SearchContentResponse {
	s.Headers = v
	return s
}

func (s *SearchContentResponse) SetStatusCode(v int32) *SearchContentResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchContentResponse) SetBody(v *SearchContentResponseBody) *SearchContentResponse {
	s.Body = v
	return s
}

func (s *SearchContentResponse) Validate() error {
	return dara.Validate(s)
}
