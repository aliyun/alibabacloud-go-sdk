// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakeCacheSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakeCacheSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakeCacheSizeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakeCacheSizeResponseBody) *DescribeLakeCacheSizeResponse
	GetBody() *DescribeLakeCacheSizeResponseBody
}

type DescribeLakeCacheSizeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakeCacheSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakeCacheSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakeCacheSizeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakeCacheSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakeCacheSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakeCacheSizeResponse) GetBody() *DescribeLakeCacheSizeResponseBody {
	return s.Body
}

func (s *DescribeLakeCacheSizeResponse) SetHeaders(v map[string]*string) *DescribeLakeCacheSizeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakeCacheSizeResponse) SetStatusCode(v int32) *DescribeLakeCacheSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakeCacheSizeResponse) SetBody(v *DescribeLakeCacheSizeResponseBody) *DescribeLakeCacheSizeResponse {
	s.Body = v
	return s
}

func (s *DescribeLakeCacheSizeResponse) Validate() error {
	return dara.Validate(s)
}
