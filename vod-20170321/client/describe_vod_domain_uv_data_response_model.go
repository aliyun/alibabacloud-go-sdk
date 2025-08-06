// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainUvDataResponseBody) *DescribeVodDomainUvDataResponse
	GetBody() *DescribeVodDomainUvDataResponseBody
}

type DescribeVodDomainUvDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainUvDataResponse) GetBody() *DescribeVodDomainUvDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainUvDataResponse) SetStatusCode(v int32) *DescribeVodDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainUvDataResponse) SetBody(v *DescribeVodDomainUvDataResponseBody) *DescribeVodDomainUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainUvDataResponse) Validate() error {
	return dara.Validate(s)
}
