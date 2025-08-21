// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainDetailResponseBody) *DescribeVsDomainDetailResponse
	GetBody() *DescribeVsDomainDetailResponseBody
}

type DescribeVsDomainDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainDetailResponse) GetBody() *DescribeVsDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeVsDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeVsDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainDetailResponse) SetStatusCode(v int32) *DescribeVsDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainDetailResponse) SetBody(v *DescribeVsDomainDetailResponseBody) *DescribeVsDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
