// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserVipsByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserVipsByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserVipsByDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserVipsByDomainResponseBody) *DescribeVodUserVipsByDomainResponse
	GetBody() *DescribeVodUserVipsByDomainResponseBody
}

type DescribeVodUserVipsByDomainResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserVipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserVipsByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserVipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserVipsByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserVipsByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserVipsByDomainResponse) GetBody() *DescribeVodUserVipsByDomainResponseBody {
	return s.Body
}

func (s *DescribeVodUserVipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeVodUserVipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserVipsByDomainResponse) SetStatusCode(v int32) *DescribeVodUserVipsByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserVipsByDomainResponse) SetBody(v *DescribeVodUserVipsByDomainResponseBody) *DescribeVodUserVipsByDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserVipsByDomainResponse) Validate() error {
	return dara.Validate(s)
}
