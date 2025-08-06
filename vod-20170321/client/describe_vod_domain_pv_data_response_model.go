// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainPvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainPvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainPvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainPvDataResponseBody) *DescribeVodDomainPvDataResponse
	GetBody() *DescribeVodDomainPvDataResponseBody
}

type DescribeVodDomainPvDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainPvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainPvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainPvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainPvDataResponse) GetBody() *DescribeVodDomainPvDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainPvDataResponse) SetStatusCode(v int32) *DescribeVodDomainPvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainPvDataResponse) SetBody(v *DescribeVodDomainPvDataResponseBody) *DescribeVodDomainPvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainPvDataResponse) Validate() error {
	return dara.Validate(s)
}
