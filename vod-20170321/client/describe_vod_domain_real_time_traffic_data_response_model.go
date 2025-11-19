// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeTrafficDataResponseBody) *DescribeVodDomainRealTimeTrafficDataResponse
	GetBody() *DescribeVodDomainRealTimeTrafficDataResponseBody
}

type DescribeVodDomainRealTimeTrafficDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) GetBody() *DescribeVodDomainRealTimeTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) SetBody(v *DescribeVodDomainRealTimeTrafficDataResponseBody) *DescribeVodDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
