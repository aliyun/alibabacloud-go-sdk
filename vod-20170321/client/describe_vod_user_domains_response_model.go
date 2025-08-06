// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserDomainsResponseBody) *DescribeVodUserDomainsResponse
	GetBody() *DescribeVodUserDomainsResponseBody
}

type DescribeVodUserDomainsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserDomainsResponse) GetBody() *DescribeVodUserDomainsResponseBody {
	return s.Body
}

func (s *DescribeVodUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeVodUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserDomainsResponse) SetStatusCode(v int32) *DescribeVodUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserDomainsResponse) SetBody(v *DescribeVodUserDomainsResponseBody) *DescribeVodUserDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserDomainsResponse) Validate() error {
	return dara.Validate(s)
}
