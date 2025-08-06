// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainHttpCodeDataResponseBody) *DescribeVodDomainHttpCodeDataResponse
	GetBody() *DescribeVodDomainHttpCodeDataResponseBody
}

type DescribeVodDomainHttpCodeDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainHttpCodeDataResponse) GetBody() *DescribeVodDomainHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponse) SetStatusCode(v int32) *DescribeVodDomainHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponse) SetBody(v *DescribeVodDomainHttpCodeDataResponseBody) *DescribeVodDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
