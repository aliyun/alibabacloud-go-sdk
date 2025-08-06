// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeQpsDataResponseBody) *DescribeVodDomainRealTimeQpsDataResponse
	GetBody() *DescribeVodDomainRealTimeQpsDataResponseBody
}

type DescribeVodDomainRealTimeQpsDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) GetBody() *DescribeVodDomainRealTimeQpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) SetBody(v *DescribeVodDomainRealTimeQpsDataResponseBody) *DescribeVodDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
