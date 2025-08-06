// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainConfigsResponseBody) *DescribeVodDomainConfigsResponse
	GetBody() *DescribeVodDomainConfigsResponseBody
}

type DescribeVodDomainConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainConfigsResponse) GetBody() *DescribeVodDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeVodDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeVodDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainConfigsResponse) SetStatusCode(v int32) *DescribeVodDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainConfigsResponse) SetBody(v *DescribeVodDomainConfigsResponseBody) *DescribeVodDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
