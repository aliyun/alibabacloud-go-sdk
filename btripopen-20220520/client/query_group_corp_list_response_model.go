// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupCorpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGroupCorpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGroupCorpListResponse
	GetStatusCode() *int32
	SetBody(v *QueryGroupCorpListResponseBody) *QueryGroupCorpListResponse
	GetBody() *QueryGroupCorpListResponseBody
}

type QueryGroupCorpListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupCorpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupCorpListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupCorpListResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupCorpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGroupCorpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGroupCorpListResponse) GetBody() *QueryGroupCorpListResponseBody {
	return s.Body
}

func (s *QueryGroupCorpListResponse) SetHeaders(v map[string]*string) *QueryGroupCorpListResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupCorpListResponse) SetStatusCode(v int32) *QueryGroupCorpListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupCorpListResponse) SetBody(v *QueryGroupCorpListResponseBody) *QueryGroupCorpListResponse {
	s.Body = v
	return s
}

func (s *QueryGroupCorpListResponse) Validate() error {
	return dara.Validate(s)
}
