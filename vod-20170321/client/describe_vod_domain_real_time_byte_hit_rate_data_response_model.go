// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeByteHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeByteHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeByteHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeByteHitRateDataResponseBody) *DescribeVodDomainRealTimeByteHitRateDataResponse
	GetBody() *DescribeVodDomainRealTimeByteHitRateDataResponseBody
}

type DescribeVodDomainRealTimeByteHitRateDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) GetBody() *DescribeVodDomainRealTimeByteHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeByteHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeVodDomainRealTimeByteHitRateDataResponseBody) *DescribeVodDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeByteHitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
