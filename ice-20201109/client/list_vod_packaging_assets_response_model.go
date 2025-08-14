// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodPackagingAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodPackagingAssetsResponse
	GetStatusCode() *int32
	SetBody(v *ListVodPackagingAssetsResponseBody) *ListVodPackagingAssetsResponse
	GetBody() *ListVodPackagingAssetsResponseBody
}

type ListVodPackagingAssetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodPackagingAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodPackagingAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListVodPackagingAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodPackagingAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodPackagingAssetsResponse) GetBody() *ListVodPackagingAssetsResponseBody {
	return s.Body
}

func (s *ListVodPackagingAssetsResponse) SetHeaders(v map[string]*string) *ListVodPackagingAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListVodPackagingAssetsResponse) SetStatusCode(v int32) *ListVodPackagingAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodPackagingAssetsResponse) SetBody(v *ListVodPackagingAssetsResponseBody) *ListVodPackagingAssetsResponse {
	s.Body = v
	return s
}

func (s *ListVodPackagingAssetsResponse) Validate() error {
	return dara.Validate(s)
}
