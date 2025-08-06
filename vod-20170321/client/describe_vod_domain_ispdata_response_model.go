// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainISPDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainISPDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainISPDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainISPDataResponseBody) *DescribeVodDomainISPDataResponse
	GetBody() *DescribeVodDomainISPDataResponseBody
}

type DescribeVodDomainISPDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainISPDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainISPDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainISPDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainISPDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainISPDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainISPDataResponse) GetBody() *DescribeVodDomainISPDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainISPDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainISPDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainISPDataResponse) SetStatusCode(v int32) *DescribeVodDomainISPDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainISPDataResponse) SetBody(v *DescribeVodDomainISPDataResponseBody) *DescribeVodDomainISPDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainISPDataResponse) Validate() error {
	return dara.Validate(s)
}
