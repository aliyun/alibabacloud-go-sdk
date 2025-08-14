// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPublicMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchPublicMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchPublicMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *SearchPublicMediaInfoResponseBody) *SearchPublicMediaInfoResponse
	GetBody() *SearchPublicMediaInfoResponseBody
}

type SearchPublicMediaInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPublicMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPublicMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchPublicMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *SearchPublicMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchPublicMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchPublicMediaInfoResponse) GetBody() *SearchPublicMediaInfoResponseBody {
	return s.Body
}

func (s *SearchPublicMediaInfoResponse) SetHeaders(v map[string]*string) *SearchPublicMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *SearchPublicMediaInfoResponse) SetStatusCode(v int32) *SearchPublicMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPublicMediaInfoResponse) SetBody(v *SearchPublicMediaInfoResponseBody) *SearchPublicMediaInfoResponse {
	s.Body = v
	return s
}

func (s *SearchPublicMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
