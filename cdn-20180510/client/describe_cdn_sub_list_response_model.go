// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSubListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnSubListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnSubListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnSubListResponseBody) *DescribeCdnSubListResponse
	GetBody() *DescribeCdnSubListResponseBody
}

type DescribeCdnSubListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnSubListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnSubListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSubListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnSubListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnSubListResponse) GetBody() *DescribeCdnSubListResponseBody {
	return s.Body
}

func (s *DescribeCdnSubListResponse) SetHeaders(v map[string]*string) *DescribeCdnSubListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSubListResponse) SetStatusCode(v int32) *DescribeCdnSubListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnSubListResponse) SetBody(v *DescribeCdnSubListResponseBody) *DescribeCdnSubListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnSubListResponse) Validate() error {
	return dara.Validate(s)
}
