// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeReqHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeReqHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeReqHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeReqHitRateDataResponseBody) *DescribeVodDomainRealTimeReqHitRateDataResponse
	GetBody() *DescribeVodDomainRealTimeReqHitRateDataResponseBody
}

type DescribeVodDomainRealTimeReqHitRateDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) GetBody() *DescribeVodDomainRealTimeReqHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeVodDomainRealTimeReqHitRateDataResponseBody) *DescribeVodDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
